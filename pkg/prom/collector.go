package prom

import (
	"regexp"
	"strconv"
	"sync"

	"github.com/mjavier2k/ping-identity-exporter/pkg/ping"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/common/log"
)

type pingCollector struct {
	client *ping.Client
}

var (
	//MetricDescriptions is
	MetricDescriptions = NewMetricDescriptions("ping")
	mux                sync.Mutex
)

func (c *pingCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- MetricDescriptions.upDesc
	ch <- MetricDescriptions.ResponseStatisticsWindowSecondsDesc
	ch <- MetricDescriptions.ResponseStatisticsCountDesc
	ch <- MetricDescriptions.ResponseTimeStatistics90PercentileDesc
	ch <- MetricDescriptions.ResponseTimeStatisticsMeanDesc
	ch <- MetricDescriptions.ResponseTimeStatisticsMaxDesc
	ch <- MetricDescriptions.ResponseTimeStatisticsMinDesc
	ch <- MetricDescriptions.ResponseConcurrencyStatistics90PercentileDesc
	ch <- MetricDescriptions.ResponseConcurrencyStatisticsMeanDesc
	ch <- MetricDescriptions.ResponseConcurrencyStatisticsMaxDesc
	ch <- MetricDescriptions.ResponseConcurrencyStatisticsMinDesc
	ch <- MetricDescriptions.CPULoadDesc
	ch <- MetricDescriptions.TotalJvmMemoryDesc
	ch <- MetricDescriptions.FreeJvmMemoryDesc
	ch <- MetricDescriptions.UsedJvmMemoryDesc
	ch <- MetricDescriptions.TotalPhysicalSystemMemoryDesc
	ch <- MetricDescriptions.TotalFreePhysicalSystemMemoryDesc
	ch <- MetricDescriptions.TotalUsedPhysicalSystemMemoryDesc
	ch <- MetricDescriptions.NumberOfCpusDesc
	ch <- MetricDescriptions.HostnameDesc
	ch <- MetricDescriptions.OpenClientConnectionsDesc
	ch <- MetricDescriptions.NumberOfApplicationsDesc
	ch <- MetricDescriptions.NumberOfVirtualHostsDesc
	// ch <- MetricDescriptions.LastRefreshTimeDesc
}

func strToFloat64(arg string) float64 {
	// matches 123.45 GB, 123 MB, etc
	regex := regexp.MustCompile(`(\d*[.,]?\d+)\s?(\w+)?`)
	result := regex.FindAllStringSubmatch(arg, -1)

	value, err := strconv.ParseFloat(result[0][1], 64)
	if err != nil {
		log.Errorf("Unable to convert %v to string. Err: %v\n", arg, err)
	}

	// result[0][2] is the 2nd capture group for unit of measure (e.g 123.4 MB)
	// Note: Not using 1 Kb = 1024 bytes to avoid skewing the values returned from the API
	if result[0][2] != "" {
		switch result[0][2] {
		case "GB":
			value = value * (1000 * 1000 * 1000) // convert GB to Bytes
		case "MB":
			value = value * (1000 * 1000) // convert MB to Bytes
		default:
			value = value
		}
		log.Debugf("converted string value %v to %v ", arg, value)
	}

	return value
}

func (c *pingCollector) Collect(ch chan<- prometheus.Metric) {
	var up float64 = 1

	heartbeat, err := c.client.GetPingAccessHearthbeat()
	if err != nil {
		up = 0
		log.Errorf("Received invalid status code from RPC call: %v\n", err)
	}

	for _, hb := range heartbeat.Items {
		mux.Lock()
		defer mux.Unlock()

		// ResponseStatisticsWindowSeconds
		ch <- prometheus.MustNewConstMetric(
			MetricDescriptions.ResponseStatisticsWindowSecondsDesc,
			prometheus.GaugeValue,
			strToFloat64(hb.ResponseStatisticsWindowSeconds),
			hb.Hostname,
		)
		// ResponseStatisticsCount
		ch <- prometheus.MustNewConstMetric(
			MetricDescriptions.ResponseStatisticsCountDesc,
			prometheus.GaugeValue,
			strToFloat64(hb.ResponseStatisticsCount),
			hb.Hostname,
		)
		// ResponseTimeStatistics90Percentile
		ch <- prometheus.MustNewConstMetric(
			MetricDescriptions.ResponseTimeStatistics90PercentileDesc,
			prometheus.GaugeValue,
			strToFloat64(hb.ResponseTimeStatistics90Percentile),
			hb.Hostname,
		)
		// ResponseTimeStatisticsMean
		ch <- prometheus.MustNewConstMetric(
			MetricDescriptions.ResponseTimeStatisticsMeanDesc,
			prometheus.GaugeValue,
			strToFloat64(hb.ResponseTimeStatisticsMean),
			hb.Hostname,
		)
		// ResponseTimeStatisticsMax
		ch <- prometheus.MustNewConstMetric(
			MetricDescriptions.ResponseTimeStatisticsMaxDesc,
			prometheus.GaugeValue,
			strToFloat64(hb.ResponseTimeStatisticsMax),
			hb.Hostname,
		)
		// ResponseTimeStatisticsMin
		ch <- prometheus.MustNewConstMetric(
			MetricDescriptions.ResponseTimeStatisticsMinDesc,
			prometheus.GaugeValue,
			strToFloat64(hb.ResponseTimeStatisticsMin),
			hb.Hostname,
		)

		// ResponseConcurrencyStatistics90Percentile
		ch <- prometheus.MustNewConstMetric(
			MetricDescriptions.ResponseConcurrencyStatistics90PercentileDesc,
			prometheus.GaugeValue,
			strToFloat64(hb.ResponseConcurrencyStatistics90Percentile),
			hb.Hostname,
		)
		// ResponseConcurrencyStatisticsMean
		ch <- prometheus.MustNewConstMetric(
			MetricDescriptions.ResponseConcurrencyStatisticsMeanDesc,
			prometheus.GaugeValue,
			strToFloat64(hb.ResponseConcurrencyStatisticsMean),
			hb.Hostname,
		)
		// ResponseConcurrencyStatisticsMax
		ch <- prometheus.MustNewConstMetric(
			MetricDescriptions.ResponseConcurrencyStatisticsMaxDesc,
			prometheus.GaugeValue,
			strToFloat64(hb.ResponseConcurrencyStatisticsMax),
			hb.Hostname,
		)
		// ResponseConcurrencyStatisticsMin
		ch <- prometheus.MustNewConstMetric(
			MetricDescriptions.ResponseConcurrencyStatisticsMinDesc,
			prometheus.GaugeValue,
			strToFloat64(hb.ResponseConcurrencyStatisticsMin),
			hb.Hostname,
		)
		// CPULoad
		ch <- prometheus.MustNewConstMetric(
			MetricDescriptions.CPULoadDesc,
			prometheus.GaugeValue,
			strToFloat64(hb.CPULoad),
			hb.Hostname,
		)
		// TotalJvmMemory
		ch <- prometheus.MustNewConstMetric(
			MetricDescriptions.TotalJvmMemoryDesc,
			prometheus.GaugeValue,
			strToFloat64(hb.TotalJvmMemory),
			hb.Hostname,
		)
		// FreeJvmMemory
		ch <- prometheus.MustNewConstMetric(
			MetricDescriptions.FreeJvmMemoryDesc,
			prometheus.GaugeValue,
			strToFloat64(hb.FreeJvmMemory),
			hb.Hostname,
		)
		// UsedJvmMemory
		ch <- prometheus.MustNewConstMetric(
			MetricDescriptions.UsedJvmMemoryDesc,
			prometheus.GaugeValue,
			strToFloat64(hb.UsedJvmMemory),
			hb.Hostname,
		)
		// TotalPhysicalSystemMemory
		ch <- prometheus.MustNewConstMetric(
			MetricDescriptions.TotalPhysicalSystemMemoryDesc,
			prometheus.GaugeValue,
			strToFloat64(hb.TotalPhysicalSystemMemory),
			hb.Hostname,
		)
		// TotalFreePhysicalSystemMemory
		ch <- prometheus.MustNewConstMetric(
			MetricDescriptions.TotalFreePhysicalSystemMemoryDesc,
			prometheus.GaugeValue,
			strToFloat64(hb.TotalFreePhysicalSystemMemory),
			hb.Hostname,
		)
		// TotalUsedPhysicalSystemMemory
		ch <- prometheus.MustNewConstMetric(
			MetricDescriptions.TotalUsedPhysicalSystemMemoryDesc,
			prometheus.GaugeValue,
			strToFloat64(hb.TotalUsedPhysicalSystemMemory),
			hb.Hostname,
		)
		// NumberOfCpus
		ch <- prometheus.MustNewConstMetric(
			MetricDescriptions.NumberOfCpusDesc,
			prometheus.GaugeValue,
			strToFloat64(hb.NumberOfCpus),
			hb.Hostname,
		)
		// Hostname
		ch <- prometheus.MustNewConstMetric(
			MetricDescriptions.HostnameDesc,
			prometheus.GaugeValue,
			1,
			hb.Hostname,
		)
		// OpenClientConnections
		ch <- prometheus.MustNewConstMetric(
			MetricDescriptions.OpenClientConnectionsDesc,
			prometheus.GaugeValue,
			strToFloat64(hb.OpenClientConnections),
			hb.Hostname,
		)
		// NumberOfApplications
		ch <- prometheus.MustNewConstMetric(
			MetricDescriptions.NumberOfApplicationsDesc,
			prometheus.GaugeValue,
			strToFloat64(hb.NumberOfApplications),
			hb.Hostname,
		)
		// NumberOfVirtualHosts
		ch <- prometheus.MustNewConstMetric(
			MetricDescriptions.NumberOfVirtualHostsDesc,
			prometheus.GaugeValue,
			strToFloat64(hb.NumberOfVirtualHosts),
			hb.Hostname,
		)
		// LastRefreshTime
	}

	// Set scrape success metric to scrapeSuccess
	ch <- prometheus.MustNewConstMetric(
		MetricDescriptions.upDesc,
		prometheus.GaugeValue,
		up,
		c.client.Hostname,
	)

}

//NewCollector is
func NewCollector(hostname string) (*pingCollector, error) {
	return &pingCollector{
		client: ping.NewPingClient(hostname),
	}, nil
}
