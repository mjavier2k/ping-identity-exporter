package main

import (
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/mjavier2k/ping-identity-exporter/pkg/ping"
	"github.com/mjavier2k/ping-identity-exporter/pkg/prom"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/prometheus/common/log"

	flag "github.com/spf13/pflag"
	"github.com/spf13/viper"
)

var (
	sha1ver   string // sha1 revision used to build the program
	buildTime string // when the executable was built
)

func init() {

	flag.CommandLine.SortFlags = false
	flag.IntP(ping.ListenPortFlag, "p", 9988, fmt.Sprintf("Port for the exporter to listen on."))
	flag.BoolP(ping.InsecureSSLFlag, "i", true, fmt.Sprintf("Whether to disable TLS validation when calling the Ping Identity Application API."))
	flag.Int64P(ping.HTTPClientTimeoutFlag, "t", 30, fmt.Sprintf("HTTP Client timeout (in seconds) per call to Ping Identity Application API."))
	flag.StringP(ping.ConfigFileFlag, "c", "config", fmt.Sprintf("Specify configuration file."))
	flag.StringP(ping.LogLevel, "l", "info", fmt.Sprintf("Specify loging level"))

	flag.Parse()
	viper.BindPFlags(flag.CommandLine)

	viper.SetConfigName(filepath.Base(viper.GetString(ping.ConfigFileFlag)))
	viper.SetConfigType("yaml")
	viper.AddConfigPath(filepath.Dir(viper.GetString(ping.ConfigFileFlag)))
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Infof("No config file found.")
		}
	} else {
		log.Infof("Found configuration file on %v ", viper.GetViper().ConfigFileUsed())
	}

	log.Base().SetLevel(viper.GetString(ping.LogLevel))
}

func main() {
	log.Infof("Version: %v", sha1ver)
	log.Infof("Built: %v", buildTime)

	listenAddr := fmt.Sprintf("127.0.0.1:%v", viper.GetInt(ping.ListenPortFlag))
	http.Handle("/metrics", promhttp.Handler())
	log.Infof("ping-identity-exporter: Started and listening on %v/pingaccess\n", listenAddr)

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
