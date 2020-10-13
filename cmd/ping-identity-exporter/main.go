package main

import (
	"fmt"
	"net/http"

	"github.com/mjavier2k/ping-identity-exporter/pkg/prom"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/prometheus/common/log"
)

var (
	sha1ver   string // sha1 revision used to build the program
	buildTime string // when the executable was built
)

func main() {
	log.Infof("Version: %v", sha1ver)
	log.Infof("Built: %v", buildTime)

	listenAddr := fmt.Sprintf("127.0.0.1:%v", 9999)
	http.Handle("/metrics", promhttp.Handler())
	log.Infof("Booted and listening on %v/pingaccess\n", listenAddr)

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "UP")
	})

	http.HandleFunc("/pingaccess", func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()

		target := query.Get("target")
		if len(query["target"]) != 1 || target == "" {
			http.Error(w, "'target' parameter must be specified once", 400)
			return
		}

		registry := prometheus.NewRegistry()
		collector, _ := prom.NewCollector(target)
		registry.MustRegister(collector)

		h := promhttp.HandlerFor(registry, promhttp.HandlerOpts{})
		h.ServeHTTP(w, r)

	})

	if err := http.ListenAndServe(listenAddr, nil); err != nil {
		log.Errorln(err)
	}
}
