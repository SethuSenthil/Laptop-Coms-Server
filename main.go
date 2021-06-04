package main

import (
	"fmt"
	"net/http"
	"strings"

	"sethusenthil.com/main/coms/phoneBackup"
)

func main() {
	port := "3889"

	fmt.Println("Starting Personal Comms Server on port " + port + " ...")
	http.HandleFunc("/", ServerHandler)
	http.ListenAndServe(":"+port, nil)
}

func ServerHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path[1:] == "PingPhoneBackup" && strings.HasPrefix(r.Header.Get("User-Agent"), "BackgroundShortcutRunner") && r.Header.Get("Auth-Token") == "YourToken" {
		fmt.Fprintf(w, phoneBackup.PhoneBackup()) //manage auto backup when connected to MacBook
	}
}
