package prom

import (
	"fmt"
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

//ConvertToBytes is
func ConvertToBytes(val float64, uom string) float64 {
	var result float64
	// unit of measure
	// Note: Not using 1 Kb = 1024 bytes to avoid skewing the values returned by the Ping API
	switch uom {
	case "GB":
		result = val * (1000 * 1000 * 1000) // convert GB to Bytes
	case "MB":
		result = val * (1000 * 1000) // convert MB to Bytes
	default:
		result = val // no uom? return original value
	}
	log.Debugf("Converted string value %v to %v ", val, result)
	return result
}

func strToFloat64(arg string) (float64, float64, error) {
	log.Debugf("Converting %s to float64", arg)

	// Matches 1234, 123.45 GB, 123 MB, 123.99
	regex := regexp.MustCompile(`(\d*[.,]?\d+)\s?(\w+)?`)
	result := regex.FindAllStringSubmatch(arg, -1)

	// Sometimes the api returns values like N/A. If this happens, return an error and do not render the metric from the result
	if len(result) == 0 {
		log.Errorf("Unexpected return value from API: %s", arg)
		return 0, 0, fmt.Errorf("Unexpected return value from API: %s", arg)
	}

	value, err := strconv.ParseFloat(result[0][1], 64)
	if err != nil {
		log.Errorf("Unable to convert value %s to float64. %v", arg, err)
		return 0, 0, fmt.Errorf("Unable to convert value %s to float64. %v", arg, err)
	}

	// Result[0][2] is the 2nd capture group which is the unit of measure (e.g 123.4 MB)
	if result[0][2] != "" {
		value = ConvertToBytes(value, result[0][2])
	}

	return value, 1, nil
}

func (c *pingCollector) Collect(ch chan<- prometheus.Metric) {
	var up float64 = 0
	var value float64 = 0

	heartbeat, err := c.client.GetPingAccessHearthbeat()
	if err != nil {
		up = 0
		log.Errorf("Received invalid status code from RPC call: %v\n", err)
	}

	for _, hb := range heartbeat.Items {
		mux.Lock()
		defer mux.Unlock()

		// ResponseStatisticsWindowSeconds
		if value, up, err = strToFloat64(hb.ResponseStatisticsWindowSeconds); err == nil {
			ch <- prometheus.MustNewConstMetric(
				MetricDescriptions.ResponseStatisticsWindowSecondsDesc,
				prometheus.GaugeValue,
				value,
				hb.Hostname,
			)
		}

		// ResponseStatisticsCount
		if value, up, err = strToFloat64(hb.ResponseStatisticsCount); err == nil {
			ch <- prometheus.MustNewConstMetric(
				MetricDescriptions.ResponseStatisticsCountDesc,
				prometheus.GaugeValue,
				value,
				hb.Hostname,
			)
		}

		// ResponseTimeStatistics90Percentile
		if value, up, err = strToFloat64(hb.ResponseTimeStatistics90Percentile); err == nil {
			ch <- prometheus.MustNewConstMetric(
				MetricDescriptions.ResponseTimeStatistics90PercentileDesc,
				prometheus.GaugeValue,
				value,
				hb.Hostname,
			)
		}

		// ResponseTimeStatisticsMean
		if value, up, err = strToFloat64(hb.ResponseTimeStatisticsMean); err == nil {
			ch <- prometheus.MustNewConstMetric(
				MetricDescriptions.ResponseTimeStatisticsMeanDesc,
				prometheus.GaugeValue,
				value,
				hb.Hostname,
			)
		}

		// ResponseTimeStatisticsMax
		if value, up, err = strToFloat64(hb.ResponseTimeStatisticsMax); err == nil {
			ch <- prometheus.MustNewConstMetric(
				MetricDescriptions.ResponseTimeStatisticsMaxDesc,
				prometheus.GaugeValue,
				value,
				hb.Hostname,
			)
		}

		// ResponseTimeStatisticsMin
		if value, up, err = strToFloat64(hb.ResponseTimeStatisticsMin); err == nil {
			ch <- prometheus.MustNewConstMetric(
				MetricDescriptions.ResponseTimeStatisticsMinDesc,
				prometheus.GaugeValue,
				value,
				hb.Hostname,
			)
		}

		// ResponseConcurrencyStatistics90Percentile
		if value, up, err = strToFloat64(hb.ResponseConcurrencyStatistics90Percentile); err == nil {
			ch <- prometheus.MustNewConstMetric(
				MetricDescriptions.ResponseConcurrencyStatistics90PercentileDesc,
				prometheus.GaugeValue,
				value,
				hb.Hostname,
			)
		}

		// ResponseConcurrencyStatisticsMean
		if value, up, err = strToFloat64(hb.ResponseConcurrencyStatisticsMean); err == nil {
			ch <- prometheus.MustNewConstMetric(
				MetricDescriptions.ResponseConcurrencyStatisticsMeanDesc,
				prometheus.GaugeValue,
				value,
				hb.Hostname,
			)
		}

		// ResponseConcurrencyStatisticsMax
		if value, up, err = strToFloat64(hb.ResponseConcurrencyStatisticsMax); err == nil {
			ch <- prometheus.MustNewConstMetric(
				MetricDescriptions.ResponseConcurrencyStatisticsMaxDesc,
				prometheus.GaugeValue,
				value,
				hb.Hostname,
			)
		}

		// ResponseConcurrencyStatisticsMin
		if value, up, err = strToFloat64(hb.ResponseConcurrencyStatisticsMin); err == nil {
			ch <- prometheus.MustNewConstMetric(
				MetricDescriptions.ResponseConcurrencyStatisticsMinDesc,
				prometheus.GaugeValue,
				value,
				hb.Hostname,
			)
		}

		// CPULoad
		if value, up, err = strToFloat64(hb.CPULoad); err == nil {
			ch <- prometheus.MustNewConstMetric(
				MetricDescriptions.CPULoadDesc,
				prometheus.GaugeValue,
				value,
				hb.Hostname,
			)
		}

		// TotalJvmMemory
		if value, up, err = strToFloat64(hb.TotalJvmMemory); err == nil {
			ch <- prometheus.MustNewConstMetric(
				MetricDescriptions.TotalJvmMemoryDesc,
				prometheus.GaugeValue,
				value,
				hb.Hostname,
			)
		}

		// FreeJvmMemory
		if value, up, err = strToFloat64(hb.FreeJvmMemory); err == nil {
			ch <- prometheus.MustNewConstMetric(
				MetricDescriptions.FreeJvmMemoryDesc,
				prometheus.GaugeValue,
				value,
				hb.Hostname,
			)
		}

		// UsedJvmMemory
		if value, up, err = strToFloat64(hb.UsedJvmMemory); err == nil {
			ch <- prometheus.MustNewConstMetric(
				MetricDescriptions.UsedJvmMemoryDesc,
				prometheus.GaugeValue,
				value,
				hb.Hostname,
			)
		}

		// TotalPhysicalSystemMemory
		if value, up, err = strToFloat64(hb.TotalPhysicalSystemMemory); err == nil {
			ch <- prometheus.MustNewConstMetric(
				MetricDescriptions.TotalPhysicalSystemMemoryDesc,
				prometheus.GaugeValue,
				value,
				hb.Hostname,
			)
		}

		// TotalFreePhysicalSystemMemory
		if value, up, err = strToFloat64(hb.TotalFreePhysicalSystemMemory); err == nil {
			ch <- prometheus.MustNewConstMetric(
				MetricDescriptions.TotalFreePhysicalSystemMemoryDesc,
				prometheus.GaugeValue,
				value,
				hb.Hostname,
			)
		}

		// TotalUsedPhysicalSystemMemory
		if value, up, err = strToFloat64(hb.TotalUsedPhysicalSystemMemory); err == nil {
			ch <- prometheus.MustNewConstMetric(
				MetricDescriptions.TotalUsedPhysicalSystemMemoryDesc,
				prometheus.GaugeValue,
				value,
				hb.Hostname,
			)
		}

		// NumberOfCpus
		if value, up, err = strToFloat64(hb.NumberOfCpus); err == nil {
			ch <- prometheus.MustNewConstMetric(
				MetricDescriptions.NumberOfCpusDesc,
				prometheus.GaugeValue,
				value,
				hb.Hostname,
			)
		}

		// Hostname
		ch <- prometheus.MustNewConstMetric(
			MetricDescriptions.HostnameDesc,
			prometheus.GaugeValue,
			1,
			hb.Hostname,
		)

		// OpenClientConnections
		if value, up, err = strToFloat64(hb.OpenClientConnections); err == nil {
			ch <- prometheus.MustNewConstMetric(
				MetricDescriptions.OpenClientConnectionsDesc,
				prometheus.GaugeValue,
				value,
				hb.Hostname,
			)
		}

		// NumberOfApplications
		if value, up, err = strToFloat64(hb.NumberOfApplications); err == nil {
			ch <- prometheus.MustNewConstMetric(
				MetricDescriptions.NumberOfApplicationsDesc,
				prometheus.GaugeValue,
				value,
				hb.Hostname,
			)
		}

		// NumberOfVirtualHosts
		if value, up, err = strToFloat64(hb.NumberOfVirtualHosts); err == nil {
			ch <- prometheus.MustNewConstMetric(
				MetricDescriptions.NumberOfVirtualHostsDesc,
				prometheus.GaugeValue,
				value,
				hb.Hostname,
			)
		}

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
