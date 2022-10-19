package seaweedfs

type TopologyStatus struct {
	Topology Topology `json:"Topology"`
	Version  string   `json:"Version"`
}

type Topology struct {
	DataCenters []DataCenter `json:"DataCenters"`
	Free        float64      `json:"Free"`
	Layouts     []Layout     `json:"Layouts"`
	Max         float64      `json:"Max"`
}

type DataCenter struct {
	Free  float64 `json:"Free"`
	Id    string  `json:"Id"`
	Max   float64 `json:"Max"`
	Racks []Rack  `json:"Racks"`
}

type Rack struct {
	DataNodes []DataNode `json:"DataNodes"`
	Free      float64    `json:"Free"`
	Id        string     `json:"Id"`
	Max       float64    `json:"Max"`
}

type DataNode struct {
	EcShards  float64 `json:"EcShards"`
	Free      float64 `json:"Free"`
	Max       float64 `json:"Max"`
	PublicUrl string  `json:"PublicUrl"`
	Url       string  `json:"Url"`
	VolumeIds string  `json:"VolumeIds"`
	Volumes   float64 `json:"Volumes"`
}

type Layout struct {
	Collection  string    `json:"collection"`
	Replication string    `json:"replication"`
	Ttl         string    `json:"ttl"`
	Writables   []float64 `json:"writables"`
}
