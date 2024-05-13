package main

import (
	"encoding/json"

	log "github.com/sirupsen/logrus"

	client "github.com/synology-community/synology-api/pkg"
	"github.com/synology-community/synology-api/pkg/util/form"
)

func setupLog() {
	log.SetFormatter(&log.JSONFormatter{})
}

func main() {

	log.SetFormatter(&log.JSONFormatter{})

	log.Info("Starting")
	// tst()

	// return

	host := "https://appkins.synology.me:5001" // os.Getenv("SYNOLOGY_HOST")
	user := "terraform"                        // os.Getenv("SYNOLOGY_USER")
	password := "ach2vzw*dnx5BPV9njr"          // os.Getenv("SYNOLOGY_PASSWORD")

	client, err := client.New(host, true)

	if err != nil {
		panic(err)
	}

	r, err := client.Login(user, password)

	if err != nil {
		panic(err)
	}

	rBytes, _ := json.Marshal(r)
	log.Infoln(string(rBytes))
	log.Infoln(string(rBytes))

	log.Infoln("Login successful")
	log.Infof("[INFO] Session: %s\nDeviceID: %s", r.SessionID, r.DeviceID)

	if _, err := client.FileStationAPI().Download("/projects/WireGuard-v1000-1.0.20220627.spk", "download"); err != nil {
		log.Error(err)

		panic(err)
	}

	if _, err := client.FileStationAPI().Upload("/data/foo/bar", &form.File{Name: "main.go", Content: "package main"}, true, true); err != nil {
		panic(err)
	}

	listGuestResp, err := client.VirtualizationAPI().ListGuests()

	if err != nil {
		panic(err)
	}

	listGuestRespBytes, _ := json.Marshal(listGuestResp)

	log.Infoln(string(listGuestRespBytes))

	for _, guest := range listGuestResp.Guests {
		println(guest.Name)
	}

	createFolder(client)
}

func createFolder(client client.SynologyClient) {
	resp, err := client.FileStationAPI().CreateFolder([]string{"/data/foo"}, []string{"bar"}, true)

	if err != nil {
		panic(err)
	}

	for _, folder := range resp.Folders {
		println(folder.Path)
		println(folder.Name)
		println(folder.IsDir)
	}
}
