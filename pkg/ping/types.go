package ping

import (
	"net/http"
	"time"
)

var (
	ListenPortFlag        = "listenPort"
	HTTPClientTimeoutFlag = "timeout"
	ListenPortFlagEnv     = "PINGEXPORTER_PORT"
)

//Client is
type Client struct {
	Hostname   string
	Endpoint   string
	HTTPClient *http.Client
}

// PingAccessHBResponse is the response payload from calling http://host:3000/pa/hearthbeat.ping endpoint
type PingAccessHBResponse struct {
	Items []struct {
		ResponseStatisticsWindowSeconds           string    `json:"response.statistics.window.seconds"`
		ResponseStatisticsCount                   string    `json:"response.statistics.count"`
		ResponseTimeStatistics90Percentile        string    `json:"response.time.statistics.90.percentile"`
		ResponseTimeStatisticsMean                string    `json:"response.time.statistics.mean"`
		ResponseTimeStatisticsMax                 string    `json:"response.time.statistics.max"`
		ResponseTimeStatisticsMin                 string    `json:"response.time.statistics.min"`
		ResponseConcurrencyStatistics90Percentile string    `json:"response.concurrency.statistics.90.percentile"`
		ResponseConcurrencyStatisticsMean         string    `json:"response.concurrency.statistics.mean"`
		ResponseConcurrencyStatisticsMax          string    `json:"response.concurrency.statistics.max"`
		ResponseConcurrencyStatisticsMin          string    `json:"response.concurrency.statistics.min"`
		CPULoad                                   string    `json:"cpu.load"`
		TotalJvmMemory                            string    `json:"total.jvm.memory"`
		FreeJvmMemory                             string    `json:"free.jvm.memory"`
		UsedJvmMemory                             string    `json:"used.jvm.memory"`
		TotalPhysicalSystemMemory                 string    `json:"total.physical.system.memory"`
		TotalFreePhysicalSystemMemory             string    `json:"total.free.physical.system.memory"`
		TotalUsedPhysicalSystemMemory             string    `json:"total.used.physical.system.memory"`
		NumberOfCpus                              string    `json:"number.of.cpus"`
		Hostname                                  string    `json:"hostname"`
		OpenClientConnections                     string    `json:"open.client.connections"`
		NumberOfApplications                      string    `json:"number.of.applications"`
		NumberOfVirtualHosts                      string    `json:"number.of.virtual.hosts"`
		LastRefreshTime                           time.Time `json:"last.refresh.time"`
	} `json:"items"`
}
