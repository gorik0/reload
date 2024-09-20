package main

import (
	prom "github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
)

func main() {

	http.Handle("/metrics", prom.Handler())
	http.ListenAndServe(":9999", nil)

}
