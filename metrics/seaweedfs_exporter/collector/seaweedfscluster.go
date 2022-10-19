package collector

import (
	"flag"
	weed "seaweedfs_exporter/seaweedfs"
	"strconv"
	"strings"

	"github.com/prometheus/client_golang/prometheus"
)

// seaweedfscluster info
type WeedCluster struct{}

func (WeedCluster) Name() string {
	return "SeaWeedFs" + "_cluster"
}

var (
	weedmasterip = flag.String("weedmasterip", "none", "The address to connect the weedmaster api.")
	leaderip     string
)

func (WeedCluster) Scrape(ch chan<- prometheus.Metric) error {
	if strings.Contains(*weedmasterip, ",") {
		for _, v := range strings.Split(*weedmasterip, ",") {
			if weed.GetWeedClusterStatus(v).IsLeader {
				leaderip = v
				break
			}
		}
	} else {
		leaderip = *weedmasterip
	}
	weedclusterinfo := weed.GetWeedClusterStatus(leaderip)
	weedtopologystatus := weed.GetWeedTopologyStatus(leaderip)
	weedcollectionsize, weedvidreadonly := weed.GetCollectioninfo(leaderip)
	weedvolumeserverdiskstatus := weed.GetVolumeServerDiskStatus(leaderip)
	datanodesinfo := make(map[weed.DataNode]string)
	datacenters := weedtopologystatus.Topology.DataCenters
	for _, dcs := range datacenters {
		for _, rs := range dcs.Racks {
			//datanodes = append(datanodes, rs.DataNodes...)
			for _, dns := range rs.DataNodes {
				datanodesinfo[dns] = rs.Id
			}
		}
	}
	var (
		isleader    float64 = 1
		leaderValue float64 = 1
	)
	if weedclusterinfo.IsLeader {
		isleader = 0
	}

	ch <- prometheus.MustNewConstMetric(
		//这里的label是固定标签 我们可以通过
		NewDesc("seaweedfs_cluster", "isleader", "wether is seaweedfs cluster leader", []string{"type"}, prometheus.Labels{"IsLeader": "isleader"}),
		prometheus.GaugeValue,
		isleader,
		//动态标签的值 可以有多个动态标签
		"master",
	)

	ch <- prometheus.MustNewConstMetric(
		NewDesc("seaweedfs_cluster", "leader", "seaweedfs cluster's Leader", []string{"type"}, prometheus.Labels{"Leader": weedclusterinfo.Leader}),
		prometheus.GaugeValue,
		leaderValue,
		"master",
	)

	ch <- prometheus.MustNewConstMetric(
		NewDesc("seaweedfs_cluster", "freeVolumeIds", "seaweedfs cluster free volumeids", nil, nil),
		prometheus.GaugeValue,
		weedtopologystatus.Topology.Free,
	)

	ch <- prometheus.MustNewConstMetric(
		NewDesc("seaweedfs_cluster", "maxVolumeId", "seaweedfs cluster maxVolumeId", nil, nil),
		prometheus.GaugeValue,
		weedclusterinfo.MaxVolumeId,
	)

	for k, rackid := range datanodesinfo {
		ch <- prometheus.MustNewConstMetric(
			NewDesc("seaweedfs_cluster", "datanode_free_volumeids", "seaweedfs cluster's datanode free volumeids", []string{"type"}, prometheus.Labels{"DataNode": k.PublicUrl, "RackId": rackid}),
			prometheus.GaugeValue,
			k.Free,
			"volumeserver",
		)
		ch <- prometheus.MustNewConstMetric(
			NewDesc("seaweedfs_cluster", "datanode_max_volumeids", "seaweedfs cluster's datanode max volumeids", []string{"type"}, prometheus.Labels{"DataNode": k.PublicUrl, "RackId": rackid}),
			prometheus.GaugeValue,
			k.Max,
			"volumeserver",
		)
	}

	for c, s := range weedcollectionsize {
		ch <- prometheus.MustNewConstMetric(
			NewDesc("seaweedfs_cluster", "volumeserver_collection_size", "seaweedfs cluster's volumeserver collection size", []string{"type"}, prometheus.Labels{"Collection": c}),
			prometheus.GaugeValue,
			s,
			"volumeserver",
		)
	}

	for vid, isreadonly := range weedvidreadonly {
		var isreadonlyvalue float64 = 0
		if isreadonly {
			isreadonlyvalue = 1
		}
		ch <- prometheus.MustNewConstMetric(
			NewDesc("seaweedfs_cluster", "volumeserver_vid_readonly", "seaweedfs cluster's volumeId wether is readonly[0(false)|1(true))]", []string{"type"}, prometheus.Labels{"VolumeId": strconv.Itoa(int(vid))}),
			prometheus.GaugeValue,
			isreadonlyvalue,
			"volumeserver",
		)
	}

	for vs, diskstatus := range weedvolumeserverdiskstatus {
		for _, dstatus := range diskstatus {
			ch <- prometheus.MustNewConstMetric(
				NewDesc("seaweedfs_cluster", "volumeserver_disk_size", "seaweedfs cluster's volumeserver disk size(byte)", []string{"type"}, prometheus.Labels{"volumeserver": vs, "diskdir": dstatus.Dir}),
				prometheus.GaugeValue,
				dstatus.All,
				"volumeserver",
			)
			ch <- prometheus.MustNewConstMetric(
				NewDesc("seaweedfs_cluster", "volumeserver_disk_used_percent", "seaweedfs cluster's volumeserver disk used(%)", []string{"type"}, prometheus.Labels{"volumeserver": vs, "diskdir": dstatus.Dir}),
				prometheus.GaugeValue,
				dstatus.Percent_used,
				"volumeserver",
			)
		}
	}

	return nil
}
