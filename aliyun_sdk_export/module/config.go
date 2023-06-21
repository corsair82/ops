package module

import (
	"github.com/spf13/viper"
)

const (
	MonitorEndpoint = "metrics.cn-hangzhou.aliyuncs.com"
	EipEndpoint     = "vpc.aliyuncs.com"
)

type AliyunUser struct {
	Name            string
	AccessKeyId     string
	AccessKeySecret string
}

type Datapoint struct {
	Timestamp  string  `json:"timestamp,omitempty"`
	InstanceId string  `json:"instanceId,omitempty"`
	UserId     string  `json:"userId,omitempty"`
	GroupId    string  `json:"groupId,omitempty"`
	Device     string  `json:"device,omitempty"`
	Value      float64 `json:"Value,omitempty"`
	Average    float64 `json:"Average,omitempty"`
	Minimum    float64 `json:"Minimum,omitempty"`
	Maximum    float64 `json:"Maximum,omitempty"`
	Sum        float64 `json:"Sum,omitempty"`
}

func LoadUserInfo() (aliyunuserslice []AliyunUser) {
	v := viper.New()
	// 设置配置文件
	v.SetConfigFile("./config/aliyun-aksk.yml")
	// 设置配置文件类型
	v.SetConfigType("yaml")
	// 加载配置文件内容
	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// 配置文件未找到
			panic("配置文件未找到")
		} else {
			// Config file was found but another error was produced
			panic(err.Error())
		}

	}
	// 读取文件配置项
	err := v.UnmarshalKey("Users", &aliyunuserslice)
	if err != nil {
		panic(err)
	}
	return aliyunuserslice
}
