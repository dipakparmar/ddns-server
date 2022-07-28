package main

import (
	"log"
	"net/http"

	updater "go.pkg.dipak.io/ddns-server/internal/dns-updater"
)

func main() {
	log.Println("Starting server at 127.0.0.1:80")
	// err := http.ListenAndServeTLS(":443", "/etc/letsencrypt/live/ddns.mydomain.net/fullchain.pem", "/etc/letsencrypt/live/ddns.mydomain.net/privkey.pem", &handler{})
	// err := http.ListenAndServe(":80", &handler{})
	// if err != nil {
	// 	log.Fatal(err)
	// }
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/nic/update", updater.UpdateDNS)
	log.Fatal(http.ListenAndServe(":80", nil))
}

// type rootHandler struct{}

func rootHandler(resp http.ResponseWriter, req *http.Request) {
	// log.Println(req.URL.String())
	auth, found := req.Header["Authorization"]
	if found {
		log.Println("Authorization:", auth)
	}
	agent, found := req.Header["User-Agent"]
	if found {
		log.Println("User-Agent:", agent)
	}

	if req.URL.Path == "/" {
		h := resp.Header()
		h.Set("Content-Type", "text/plain")
		resp.WriteHeader(http.StatusOK)
		resp.Write([]byte("Server is running okay."))
	} else {
		// return a 404
		resp.WriteHeader(http.StatusNotFound)
		resp.Write([]byte("Not found."))
	}

}
