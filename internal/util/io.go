package util

// WriteCounter counts the number of bytes written to it. It implements to the io.Writer interface
// and we can pass this into io.TeeReader() which will report progress on each write cycle.
// type WriteCounter struct {
// 	Total uint64
// }

// func (wc *WriteCounter) Write(p []byte) (int, error) {
// 	n := len(p)
// 	wc.Total += uint64(n)
// 	wc.PrintProgress()
// 	return n, nil
// }

// func (wc WriteCounter) PrintProgress() {
// 	// Clear the line by using a character return to go back to the start and remove
// 	// the remaining characters by filling it with spaces
// 	fmt.Printf("\r%s", strings.Repeat(" ", 35))

// 	// Return again and print current status of download
// 	// We use the humanize package to print the bytes in a meaningful way (e.g. 10 MB)
// 	fmt.Printf("\rDownloading... %s complete", humanize.Bytes(wc.Total))
// }
