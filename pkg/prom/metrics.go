package prom

import (
	"github.com/prometheus/client_golang/prometheus"
)

// Descriptions is
type Descriptions struct {
	upDesc *prometheus.Desc

	// {"items":[{
	// 	"response.statistics.window.seconds": "10",
	// 	"response.statistics.count": "40",
	// 	"response.time.statistics.90.percentile": "1",
	// 	"response.time.statistics.mean": "0.85",
	// 	"response.time.statistics.max": "1",
	// 	"response.time.statistics.min": "0",
	// 	"response.concurrency.statistics.90.percentile": "2",
	// 	"response.concurrency.statistics.mean": "1.55",
	// 	"response.concurrency.statistics.max": "3",
	// 	"response.concurrency.statistics.min": "1",
	// 	"cpu.load": "0.69",
	// 	"total.jvm.memory": "470.286 MB",
	// 	"free.jvm.memory": "260.423 MB",
	// 	"used.jvm.memory": "209.864 MB",
	// 	"total.physical.system.memory": "3.973 GB",
	// 	"total.free.physical.system.memory": "688.079 MB",
	// 	"total.used.physical.system.memory": "3.285 GB",
	// 	"number.of.cpus": "8",
	// 	"hostname": "lv1stgpingaccess03",
	// 	"open.client.connections": "4",
	// 	"number.of.applications": "53",
	// 	"number.of.virtual.hosts": "325",
	// 	"last.refresh.time": "2020-09-16T13:35:00.000Z"
	// }]}

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
		prometheus.BuildFQName(namespace, "", "pingaccess_statistics_window_seconds"),
		"Statistics Window Seconds",
		[]string{"hostname"},
		nil,
	)

	d.ResponseStatisticsCountDesc = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "pingaccess_statistics_count"),
		"Statistics Count",
		[]string{"hostname"},
		nil,
	)

	d.ResponseTimeStatistics90PercentileDesc = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "pingaccess_response_time_statistics_90_percentile"),
		"Response time statistics 90th percentile",
		[]string{"hostname"},
		nil,
	)

	d.ResponseTimeStatisticsMeanDesc = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "pingaccess_response_time_statistics_mean"),
		"Response time statistics mean",
		[]string{"hostname"},
		nil,
	)

	d.ResponseTimeStatisticsMaxDesc = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "pingaccess_response_time_statistics_max"),
		"Response time statistics max",
		[]string{"hostname"},
		nil,
	)

	d.ResponseTimeStatisticsMinDesc = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "pingaccess_response_time_statistics_min"),
		"Response time statistics min",
		[]string{"hostname"},
		nil,
	)

	d.ResponseConcurrencyStatistics90PercentileDesc = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "pingaccess_response_concurrency_statistics_90_percentile"),
		"Response concurrency statistics 90th percentile",
		[]string{"hostname"},
		nil,
	)

	d.ResponseConcurrencyStatisticsMeanDesc = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "pingaccess_response_concurrency_statistics_mean"),
		"Response Concurrency Statistics Mean",
		[]string{"hostname"},
		nil,
	)

	d.ResponseConcurrencyStatisticsMaxDesc = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "pingaccess_response_concurrency_statistics_max"),
		"Response Concurrency Statistics Max",
		[]string{"hostname"},
		nil,
	)

	d.ResponseConcurrencyStatisticsMinDesc = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "pingaccess_response_concurrency_statistics_min"),
		"Response Concurrency Statistics Min",
		[]string{"hostname"},
		nil,
	)

	d.CPULoadDesc = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "pingaccess_cpu_load"),
		"The current CPU utilization. The value returned is a real number from 0 to 1 which represents the CPU utilization percentage.",
		[]string{"hostname"},
		nil,
	)

	d.TotalJvmMemoryDesc = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "pingaccess_total_jvm_memory_bytes"),
		"Total JVM Memory",
		[]string{"hostname"},
		nil,
	)

	d.FreeJvmMemoryDesc = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "pingaccess_free_jvm_memory_bytes"),
		"Free JVM Memory",
		[]string{"hostname"},
		nil,
	)

	d.UsedJvmMemoryDesc = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "pingaccess_used_jvm_memory_bytes"),
		"Used JVM Memory",
		[]string{"hostname"},
		nil,
	)

	d.TotalPhysicalSystemMemoryDesc = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "pingaccess_total_physical_system_memory_bytes"),
		"Total Physical System Memory",
		[]string{"hostname"},
		nil,
	)

	d.TotalFreePhysicalSystemMemoryDesc = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "pingaccess_total_free_physical_system_memory_bytes"),
		"Total Free Physical System Memory",
		[]string{"hostname"},
		nil,
	)

	d.TotalUsedPhysicalSystemMemoryDesc = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "pingaccess_total_used_physical_system_memory_bytes"),
		"Total Used Physical System Memory",
		[]string{"hostname"},
		nil,
	)

	d.NumberOfCpusDesc = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "pingaccess_number_of_cpus"),
		"Number of CPU",
		[]string{"hostname"},
		nil,
	)

	d.HostnameDesc = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "pingaccess_hostname"),
		"The ping access hostname",
		[]string{"hostname"},
		nil,
	)

	d.OpenClientConnectionsDesc = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "pingaccess_open_client_connections"),
		"Open Client Connections",
		[]string{"hostname"},
		nil,
	)

	d.NumberOfApplicationsDesc = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "pingaccess_number_of_applications"),
		"Number of Applications",
		[]string{"hostname"},
		nil,
	)

	d.NumberOfVirtualHostsDesc = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "pingaccess_number_of_virtual_hosts"),
		"Number of Virtual Hosts",
		[]string{"hostname"},
		nil,
	)

	// d.LastRefreshTimeDesc = prometheus.NewDesc(
	// 	prometheus.BuildFQName(namespace, "", "pingaccess_"),
	// 	"",
	// 	[]string{"hostname"},
	// 	nil,
	// )

	return &d
}
