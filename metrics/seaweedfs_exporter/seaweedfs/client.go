package seaweedfs

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func Curl(weburl string) []byte {
	response, err := http.Get(weburl)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	return body
}

func GetWeedClusterStatus(masterurl string) ClusterInfo {
	apiurl := "http://" + masterurl + "/cluster/status?pretty=y"
	var clusterinfo ClusterInfo
	json.Unmarshal(Curl(apiurl), &clusterinfo)
	return clusterinfo
}

func GetWeedTopologyStatus(masterurl string) TopologyStatus {
	apiurl := "http://" + masterurl + "/dir/status?pretty=y"
	var topologyStatus TopologyStatus
	json.Unmarshal(Curl(apiurl), &topologyStatus)
	return topologyStatus
}

func GetVolumeServerList(masterurl string) (VolumeServerList [][]string) {
	weedtopologystatus := GetWeedTopologyStatus(masterurl)
	for _, dcenter := range weedtopologystatus.Topology.DataCenters {
		dcvolume := make(map[string]int)
		var dcvolumeservers []string
		for _, rack := range dcenter.Racks {
			for _, vserver := range rack.DataNodes {
				dcvolume[vserver.PublicUrl] = 1
			}
		}
		for k, _ := range dcvolume {
			dcvolumeservers = append(dcvolumeservers, k)
		}
		VolumeServerList = append(VolumeServerList, dcvolumeservers)
	}
	return VolumeServerList
}

func GetWeedVolumeServer(volumeserverurl string) VolumeServerInfo {
	apiurl := "http://" + volumeserverurl + "/status?pretty=y"
	var volumeServerInfo VolumeServerInfo
	json.Unmarshal(Curl(apiurl), &volumeServerInfo)
	return volumeServerInfo
}

func GetCollectioninfo(volumeserverurl string) (map[string]float64, map[float64]bool) {
	collectionsize := make(map[string]float64)
	vidreadonly := make(map[float64]bool)
	volumeserverlist := GetVolumeServerList(volumeserverurl)
	for _, dc := range volumeserverlist {
		for _, vs := range dc {
			weedvolumeserverinfo := GetWeedVolumeServer(vs)
			volumes := weedvolumeserverinfo.Volumes
			for _, v := range volumes {
				collectionsize[v.Collection] += v.Size
				vidreadonly[v.Id] = v.ReadOnly
			}
		}
	}
	return collectionsize, vidreadonly
}

func GetVolumeServerDiskStatus(masterurl string) map[string][]DiskStatus {
	volumeserverdiskstatus := make(map[string][]DiskStatus)
	volumeserverlist := GetVolumeServerList(masterurl)
	for _, dc := range volumeserverlist {
		for _, vs := range dc {
			volumeserverinfo := GetWeedVolumeServer(vs)
			volumeserverdiskstatus[vs] = append(volumeserverdiskstatus[vs], volumeserverinfo.DiskStatuses...)
		}
	}
	return volumeserverdiskstatus
}
