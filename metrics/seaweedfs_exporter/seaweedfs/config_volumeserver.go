package seaweedfs

type VolumeServerInfo struct {
	DiskStatuses []DiskStatus   `json:"DiskStatuses"`
	Version      string         `json:"Version"`
	Volumes      []Volumeserver `json:"Volumes"`
}

type DiskStatus struct {
	Dir          string  `json:"dir"`
	All          float64 `json:"all"`
	Used         float64 `json:"used"`
	Free         float64 `json:"free"`
	Percent_free float64 `json:"percent_free"`
	Percent_used float64 `json:"percent_used"`
}

type Volumeserver struct {
	Id                float64          `json:"Id"`
	Size              float64          `json:"Size"`
	ReplicaPlacement  ReplicaPlacement `json:"ReplicaPlacement"`
	VTtl              VTtl             `json:"Ttl"`
	Collection        string           `json:"Collection"`
	Version           float64          `json:"Version"`
	FileCount         float64          `json:"FileCount"`
	DeleteCount       float64          `json:"DeleteCount"`
	DeletedByteCount  float64          `json:"DeletedByteCount"`
	CompactRevision   float64          `json:"CompactRevision"`
	ModifiedAtSecond  float64          `json:"ModifiedAtSecond"`
	ReadOnly          bool             `json:"ReadOnly"`
	RemoteStorageName string           `json:"RemoteStorageName"`
	RemoteStorageKey  string           `json:"RemoteStorageKey"`
}

type ReplicaPlacement struct {
	SameRackCount       float64 `json:"SameRackCount"`
	DiffRackCount       float64 `json:"DiffRackCount"`
	DiffDataCenterCount float64 `json:"DiffDataCenterCount"`
}

type VTtl struct {
	Count float64 `json:"Count"`
	Unit  float64 `json:"Unit"`
}
