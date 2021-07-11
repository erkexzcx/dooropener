package dooropener

import (
	"log"
	"net/http"
)

func startHTTPServer() {
	http.HandleFunc("/", handlerRoot)
	log.Fatal(http.ListenAndServe(httpConfig.Bind, nil))
}

func handlerRoot(w http.ResponseWriter, r *http.Request) {
	if r.URL.RequestURI() != httpConfig.SecretURI {
		http.Error(w, "not found", 404)
		return
	}

	if getDoorInProgress() {
		sendTelegram("Already in progress...")
		http.Error(w, "already in progress", 200)
		return
	}

	setDoorInProgress(true)
	sendTelegram("Opening!")
	openDoor()
	setDoorInProgress(false)
	sendTelegram("Done!")
	http.Error(w, "success", 200)
}
