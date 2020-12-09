package prom

import (
	"github.com/prometheus/client_golang/prometheus"
)

// Descriptions is
type Descriptions struct {
	upDesc                                        *prometheus.Desc
	ResponseStatisticsWindowSecondsDesc           *prometheus.Desc
	ResponseStatisticsCountDesc                   *prometheus.Desc
	ResponseTimeStatistics90PercentileDesc        *prometheus.Desc
	ResponseTimeStatisticsMeanDesc                *prometheus.Desc
	ResponseTimeStatisticsMaxDesc                 *prometheus.Desc
	ResponseTimeStatisticsMinDesc                 *prometheus.Desc
	ResponseConcurrencyStatistics90PercentileDesc *prometheus.Desc
	ResponseConcurrencyStatisticsMeanDesc         *prometheus.Desc
	ResponseConcurrencyStatisticsMaxDesc          *prometheus.Desc
	ResponseConcurrencyStatisticsMinDesc          *prometheus.Desc
	CPULoadDesc                                   *prometheus.Desc
	TotalJvmMemoryDesc                            *prometheus.Desc
	FreeJvmMemoryDesc                             *prometheus.Desc
	UsedJvmMemoryDesc                             *prometheus.Desc
	TotalPhysicalSystemMemoryDesc                 *prometheus.Desc
	TotalFreePhysicalSystemMemoryDesc             *prometheus.Desc
	TotalUsedPhysicalSystemMemoryDesc             *prometheus.Desc
	NumberOfCpusDesc                              *prometheus.Desc
	HostnameDesc                                  *prometheus.Desc
	OpenClientConnectionsDesc                     *prometheus.Desc
	NumberOfApplicationsDesc                      *prometheus.Desc
	NumberOfVirtualHostsDesc                      *prometheus.Desc
	LastRefreshTimeDesc                           *prometheus.Desc
}

// NewMetricDescriptions is
func NewMetricDescriptions(namespace string) *Descriptions {
	var d Descriptions

	d.upDesc = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "up"),
		"Whether last scrape against ping access was successful.",
		[]string{"hostname"},
		nil,
	)

	d.ResponseStatisticsWindowSecondsDesc = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "statistics_window_seconds"),
		"Statistics Window Seconds",
		[]string{"hostname"},
		nil,
	)

	d.ResponseStatisticsCountDesc = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "statistics_count"),
		"Statistics Count",
		[]string{"hostname"},
		nil,
	)

	d.ResponseTimeStatistics90PercentileDesc = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "response_time_statistics_90_percentile"),
		"Response time statistics 90th percentile",
		[]string{"hostname"},
		nil,
	)

	d.ResponseTimeStatisticsMeanDesc = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "response_time_statistics_mean"),
		"Response time statistics mean",
		[]string{"hostname"},
		nil,
	)

	d.ResponseTimeStatisticsMaxDesc = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "response_time_statistics_max"),
		"Response time statistics max",
		[]string{"hostname"},
		nil,
	)

	d.ResponseTimeStatisticsMinDesc = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "response_time_statistics_min"),
		"Response time statistics min",
		[]string{"hostname"},
		nil,
	)

	d.ResponseConcurrencyStatistics90PercentileDesc = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "response_concurrency_statistics_90_percentile"),
		"Response concurrency statistics 90th percentile",
		[]string{"hostname"},
		nil,
	)

	d.ResponseConcurrencyStatisticsMeanDesc = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "response_concurrency_statistics_mean"),
		"Response Concurrency Statistics Mean",
		[]string{"hostname"},
		nil,
	)

	d.ResponseConcurrencyStatisticsMaxDesc = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "response_concurrency_statistics_max"),
		"Response Concurrency Statistics Max",
		[]string{"hostname"},
		nil,
	)

	d.ResponseConcurrencyStatisticsMinDesc = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "response_concurrency_statistics_min"),
		"Response Concurrency Statistics Min",
		[]string{"hostname"},
		nil,
	)

	d.CPULoadDesc = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "cpu_load"),
		"The current CPU utilization. The value returned is a real number from 0 to 1 which represents the CPU utilization percentage.",
		[]string{"hostname"},
		nil,
	)

	d.TotalJvmMemoryDesc = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "total_jvm_memory_bytes"),
		"Total JVM Memory",
		[]string{"hostname"},
		nil,
	)

	d.FreeJvmMemoryDesc = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "free_jvm_memory_bytes"),
		"Free JVM Memory",
		[]string{"hostname"},
		nil,
	)

	d.UsedJvmMemoryDesc = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "used_jvm_memory_bytes"),
		"Used JVM Memory",
		[]string{"hostname"},
		nil,
	)

	d.TotalPhysicalSystemMemoryDesc = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "total_physical_system_memory_bytes"),
		"Total Physical System Memory",
		[]string{"hostname"},
		nil,
	)

	d.TotalFreePhysicalSystemMemoryDesc = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "total_free_physical_system_memory_bytes"),
		"Total Free Physical System Memory",
		[]string{"hostname"},
		nil,
	)

	d.TotalUsedPhysicalSystemMemoryDesc = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "total_used_physical_system_memory_bytes"),
		"Total Used Physical System Memory",
		[]string{"hostname"},
		nil,
	)

	d.NumberOfCpusDesc = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "number_of_cpus"),
		"Number of CPU",
		[]string{"hostname"},
		nil,
	)

	d.HostnameDesc = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "hostname"),
		"The ping access hostname",
		[]string{"hostname"},
		nil,
	)

	d.OpenClientConnectionsDesc = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "open_client_connections"),
		"Open Client Connections",
		[]string{"hostname"},
		nil,
	)

	d.NumberOfApplicationsDesc = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "number_of_applications"),
		"Number of Applications",
		[]string{"hostname"},
		nil,
	)

	d.NumberOfVirtualHostsDesc = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "number_of_virtual_hosts"),
		"Number of Virtual Hosts",
		[]string{"hostname"},
		nil,
	)

	// d.LastRefreshTimeDesc = prometheus.NewDesc(
	// 	prometheus.BuildFQName(namespace, "", ""),
	// 	"",
	// 	[]string{"hostname"},
	// 	nil,
	// )

	return &d
}
