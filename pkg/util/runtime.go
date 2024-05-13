package util

import (
	"context"
	"fmt"
	"path/filepath"
	"regexp"
	"runtime"
	"strings"

	log "github.com/sirupsen/logrus"
)

func IsClientMethod(caller string) (bool, string, error) {

	re1 := regexp.MustCompile(`\.\(\*\w+Client\)\.[\w\.]+$`)

	if !re1.MatchString(caller) {
		return false, "", nil
	}

	re := regexp.MustCompile(`(?:package\.\(\*)(?P<Package>\w+)(?:Client[\)\.]{2})(?P<Method>[\w\.]+)$`)

	matches := re.FindStringSubmatch(caller)

	if len(matches) < 3 {
		return false, "", nil
	}

	pkg := matches[re.SubexpIndex("Package")]
	mtd := matches[re.SubexpIndex("Method")]

	if pkg == "" || mtd == "" {
		return false, "", nil
	}

	return true, fmt.Sprintf("%s.%s", TitleCase(pkg), TitleCase(mtd)), nil
}

func getMethod(skipFrames int) (string, error) {
	// We need the frame at index skipFrames+2, since we never want runtime.Callers and getFrame
	targetFrameIndex := skipFrames + 2

	// Set size to targetFrameIndex+2 to ensure we have room for one more caller than we need
	programCounters := make([]uintptr, targetFrameIndex+2)
	n := runtime.Callers(0, programCounters)

	mtd := ""
	fnList := []string{}
	if n > 0 {
		frames := runtime.CallersFrames(programCounters[:n])
		for more, frameIndex := true, 0; more && frameIndex <= targetFrameIndex; frameIndex++ {
			var frameCandidate runtime.Frame
			frameCandidate, more = frames.Next()
			fnList = append(fnList, frameCandidate.Function)
		}
	}

	for _, f := range fnList {
		// log.Infof("[INFO] Frame %d: %s\n", i, f)
		if ok, v, _ := IsClientMethod(f); ok {
			// log.Infoln(v)
			mtd = v
		}
	}

	return mtd, nil
}

func GetCaller() (caller string, err error) {
	f, err := getMethod(4)
	if err != nil {
		return "", err
	}
	return f, nil

	// log.Infof("[INFO] Caller: %s\n", f)

	// pcs := []uintptr{}
	// runtime.Callers(1, pcs)

	// rxp, err := regexp.Compile(`\.\(*(\w+Client)\)\.`)

	// if err != nil {
	// 	// log.Fatalf("[ERROR] Failed to compile regex: %v", err)
	// 	return "", errors.New("failed to compile regex")
	// }

	// for _, pc := range pcs {
	// 	fn := runtime.FuncForPC(pc).Name()
	// 	if rxp.MatchString(fn) || strings.Contains(fn, "Client") {
	// 		caller = fn // Found the caller
	// 		break       // Exit the loop
	// 	}
	// }
	// caller, err = Sanitize(caller)
	// return
}

func Sanitize(ctx context.Context, caller string) (string, error) {
	// Remove github.com/username/project/ from the caller
	name := filepath.Base(caller)

	// log.Infof("[INFO] Caller Basename: %s\n", name)

	re := regexp.MustCompile(`(?:package\.\(\*)(?P<Package>\w+)(?:Client[\)\.]{2})(?P<Method>[\w\.]+)`)

	if !re.MatchString(name) {
		return caller, nil
	}

	matches := re.FindStringSubmatch(name)

	// log.Infof("[INFO] Subexp: %v\n", re.SubexpNames())

	// log.Infof("[INFO] Matches: %v\n", matches)

	parts := []string{}

	for i, m := range matches {
		log.Info("Match %d: %s\n", i, m)
	}

	for i := range re.SubexpNames() {
		if i > 0 && i <= len(matches) {
			parts = append(parts, TitleCase(matches[i]))
		}
	}

	result := strings.Join(parts, ".")

	// log.Infof("[INFO] Sanitized Caller Name: %s", result)

	return result, nil
}
