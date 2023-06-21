package collector

import (
	gxmod "aliyun_sdk_export/module"
	"strings"

	"github.com/prometheus/client_golang/prometheus"
)

// seaweedfscluster info
type AliyunSDK struct{}

func (AliyunSDK) Name() string {
	return "AliyunSDK" + "_INFO"
}

func (AliyunSDK) Scrape(ch chan<- prometheus.Metric) error {
	eipNetRxMetricLast := gxmod.AllUserDescribeMetricLast("acs_vpc_eip", "net_rx.rate")
	eipNetTxMetricLast := gxmod.AllUserDescribeMetricLast("acs_vpc_eip", "net_tx.rate")
	for k, v := range eipNetRxMetricLast {
		eipadressmap, _ := gxmod.DescribeEipAddresses(k)
		for _, d := range v {
			ipaddress := strings.Split(eipadressmap[d.InstanceId], "_")[0]
			regionsid := strings.Split(eipadressmap[d.InstanceId], "_")[1]
			ch <- prometheus.MustNewConstMetric(
				//这里的label是固定标签 我们可以通过
				NewDesc("aliyunsdk_info", "acsvpceip_netrx_rate", "describe eip rx rage (byte/s)", []string{"type"}, prometheus.Labels{"ACSUSER": k, "MetricClass": "EipNetRxRate", "RegionsId": regionsid, "Eipaddress": ipaddress, "InstanceId": d.InstanceId}),
				prometheus.GaugeValue,
				d.Value,
				//动态标签的值 可以有多个动态标签
				"ALIYUNASDK",
			)
		}

	}
	for k, v := range eipNetTxMetricLast {
		eipadressmap, _ := gxmod.DescribeEipAddresses(k)
		for _, d := range v {
			ipaddress := strings.Split(eipadressmap[d.InstanceId], "_")[0]
			regionsid := strings.Split(eipadressmap[d.InstanceId], "_")[1]
			ch <- prometheus.MustNewConstMetric(
				//这里的label是固定标签 我们可以通过
				NewDesc("aliyunsdk_info", "acsvpceip_nettx_rate", "describe eip tx rage (byte/s)", []string{"type"}, prometheus.Labels{"ACSUSER": k, "MetricClass": "EipNetTxRate", "RegionsId": regionsid, "Eipaddress": ipaddress, "InstanceId": d.InstanceId}),
				prometheus.GaugeValue,
				d.Value,
				//动态标签的值 可以有多个动态标签
				"ALIYUNASDK",
			)
		}

	}
	return nil
}
