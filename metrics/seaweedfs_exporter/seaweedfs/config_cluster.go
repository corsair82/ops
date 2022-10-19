package seaweedfs

type ClusterInfo struct {
	IsLeader    bool     `json:"IsLeader"`
	Leader      string   `json:"Leader"`
	Peers       []string `json:"Peers"`
	MaxVolumeId float64  `json:"MaxVolumeId"`
}
