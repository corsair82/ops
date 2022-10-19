package seaweedfs

type JsonInfo struct {
	Json Jinfo `json:"JSON"`
}

type Jinfo struct {
	Version string   `json:"Version"`
	Volumes []Volume `json:"Volumes"`
}

type Volume struct {
	DataCenters VDataCenter `json:"DataCenters"`
	Free        float64     `json:"Free"`
	Max         float64     `json:"Max"`
}

type VDataCenter struct {
}
