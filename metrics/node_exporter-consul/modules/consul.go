package modules

import (
	"fmt"
	"log"
	"net"
	_ "net/http/pprof"
	"strconv"
	"strings"

	consulapi "github.com/hashicorp/consul/api"
)

func Consulregister(address, nodeport, moqitags *string) {
	portStr := strings.Split(*nodeport, ":")
	port, _ := strconv.Atoi(portStr[1])
	config := consulapi.DefaultConfig()
	config.Address = *address
	regtags := []string{"moqi-node"}
	client, err := consulapi.NewClient(config)
	if err != nil {
		log.Fatal("consul client error : ", err)
	}

	registration := new(consulapi.AgentServiceRegistration)
	registration.ID = localIP() + "-node"          // 服务节点的名称
	registration.Name = "node-exporter"            // 服务名称
	registration.Port = port                       // 服务端口
	registration.Tags = append(regtags, *moqitags) // tag，可以为空
	registration.Address = localIP()               // 服务 IP

	checkPort := port
	registration.Check = &consulapi.AgentServiceCheck{ // 健康检查
		HTTP:     fmt.Sprintf("http://%s:%d%s", registration.Address, checkPort, "/metrics"),
		Timeout:  "3s",
		Interval: "5s", // 健康检查间隔
		//	DeregisterCriticalServiceAfter: "30s", //check失败后30秒删除本服务，注销时间，相当于过期时间
		//  GRPC:     fmt.Sprintf("%v:%v/%v", IP, r.Port, r.Service),// grpc 支持，执行健康检查的地址，service 会传到 Health.Check 函数中
	}

	err = client.Agent().ServiceRegister(registration)
	if err != nil {
		log.Fatal("register server error : ", err)
	}
}

func localIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}
	for _, address := range addrs {
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return ""
}
