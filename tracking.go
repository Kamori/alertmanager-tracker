package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/prometheus/alertmanager/template"
)

func main() {

	reqHandler := func(w http.ResponseWriter, req *http.Request) {
		webhook_fire := &template.Data{}
		json.NewDecoder(req.Body).Decode(webhook_fire)

		fmt.Println(webhook_fire.Status)
		io.WriteString(w, webhook_fire.Status)
		for _, alert := range webhook_fire.Alerts {
			fmt.Println(alert.Labels["alertname"])
		}
	}

	http.HandleFunc("/", reqHandler)
	// host.docker.internal on mac
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// POST /api/v1/alerts HTTP/1.1
// Host: localhost
// Content-Type: application/json; charset=utf-8
// Content-Length: 124

// [{"labels":{"alertname":"TylerDTirt","node":"random"},"annotations":{"description":"Sample descript"},"generatorURL":"wtf"}]
