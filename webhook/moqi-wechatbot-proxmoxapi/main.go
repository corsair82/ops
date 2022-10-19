package main

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"log"
	"math"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"

	"moqi-wechatbot-proxmoxapi/model"

	"moqi-wechatbot-proxmoxapi/proxmox"

	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
)

func ProxmoxLogin() (c *proxmox.Client) {
	tlsconf := &tls.Config{InsecureSkipVerify: true}
	PM_API_URL := "https://pve.moqi.ai/api2/json"
	PM_USER := "root@pam"
	PM_PASS := "liyHgxPsMOsrSj"
	PM_OTP := ""
	c, err := proxmox.NewClient(PM_API_URL, nil, tlsconf, "", 300)
	failError(err)
	if userRequiresAPIToken(PM_USER) {
		c.SetAPIToken(PM_USER, PM_PASS)
		// As test, get the version of the server
		_, err := c.GetVersion()
		if err != nil {
			log.Fatalf("login error: %s", err)
		}
	} else {
		err = c.Login(PM_USER, PM_PASS, PM_OTP)
		failError(err)
	}
	return c
}
func getGxVmConfig() []proxmox.ConfigQemu {
	proxc := ProxmoxLogin()
	vms, err := proxc.GetVmList()
	if err != nil {
		log.Printf("Error listing VMs %+v\n", err)
		os.Exit(1)
	}
	var vmsconfig []proxmox.ConfigQemu
	var vmsid []int
	vmsinfo, isvmsinfo := vms["data"].([]interface{})
	if isvmsinfo {
		for vm := range vmsinfo {
			vminfo := vmsinfo[vm].(map[string]interface{})
			vmstatus := string(vminfo["status"].(string))
			if vmstatus == "running" {
				id := int(math.Ceil(float64(vminfo["vmid"].(float64))))
				vmsid = append(vmsid, id)
			}
		}
	}
	for _, vmid := range vmsid {
		vmr := proxmox.NewVmRef(vmid)
		err := proxc.CheckVmRef(vmr)
		failError(err)
		vmconfig, _ := proxc.GetGxVmConfig(vmr)
		vmsconfig = append(vmsconfig, *vmconfig)
	}
	return vmsconfig
}

func GetNodesRS() string {
	proxc := ProxmoxLogin()
	nodeRemainSource := proxc.GetNodeRemainSource()
	nodesource, err := json.Marshal(nodeRemainSource)
	failError(err)
	return string(nodesource)
}

func WechatApiPost(wechatapiurl, alertdesc string, user []string) {
	//post the context to the wechat api
	postBody, _ := json.Marshal(map[string]interface{}{
		"msgtype": "text",
		"text": map[string]interface{}{
			"content":        alertdesc,
			"mentioned_list": user,
		},
	})
	responseBody := bytes.NewBuffer(postBody)
	resp, err := http.Post(wechatapiurl, "application/json", responseBody)
	if err != nil {
		glog.Error(err)
		return
	}

	defer resp.Body.Close()
}

func failError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

var rxUserRequiresToken = regexp.MustCompile("[a-z0-9]+@[a-z0-9]+![a-z0-9]+")

func userRequiresAPIToken(userID string) bool {
	return rxUserRequiresToken.MatchString(userID)
}

func main() {
	wechatApi := "https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=519dac67-b406-4f21-9cfb-25cb39c7c34cgx"
	router := gin.Default()
	router.GET("/wechatbot", func(c *gin.Context) {
		allowAccessIp := "172.17.0.[2-254]"
		ipRegex, _ := regexp.Compile(allowAccessIp)
		ipMatch := ipRegex.MatchString(c.RemoteIP())
		if !ipMatch {
			c.String(http.StatusForbidden, "No permission to access /wechatbot....")
			return
		}

		//router.SetTrustedProxies([]string{"172.17.0.0/24", "127.0.0.1", "10.10.5.88"})
		c.String(http.StatusOK, "Hello,here is nothing....")
	})
	router.POST("/wechatbot", func(c *gin.Context) {

		var notification model.Notification

		err := c.BindJSON(&notification)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		//router.SetTrustedProxies([]string{"172.17.0.0/24"})
		vmconfig := getGxVmConfig()
		var (
			vmsname, user []string
			vmhostpcis    int
		)
		vmpcis := make(map[string]int)
		for _, v := range vmconfig {
			vmsname = append(vmsname, v.Name)
			if v.Hostpci0 != "" {
				vmpcis[v.Name] += 1
			}
			if v.Hostpci1 != "" {
				vmpcis[v.Name] += 1
			}
			if v.Hostpci2 != "" {
				vmpcis[v.Name] += 1
			}
			if v.Hostpci3 != "" {
				vmpcis[v.Name] += 1
			}
			if v.Hostpci4 != "" {
				vmpcis[v.Name] += 1
			}
			if v.Hostpci5 != "" {
				vmpcis[v.Name] += 1
			}
			if v.Hostpci6 != "" {
				vmpcis[v.Name] += 1
			}
			if v.Hostpci7 != "" {
				vmpcis[v.Name] += 1
			}
		}

		noNotiUserStr := "matching|infra|generic|extraction|Chengdu"
		for _, v := range notification.Alerts {
			vmip := strings.Split(v.Labels["instance"], ":")[0] + "-"

			for _, name := range vmsname {
				if strings.Contains(name, vmip) {
					username := strings.Split(name, "-")[1]
					r, _ := regexp.Compile(noNotiUserStr)
					match := r.MatchString(username)
					if match {
						user = []string{}
					} else {
						user = append(user, username)
						vmhostpcis = vmpcis[name]
					}

					break
				}
			}
			if len(user) == 0 {
				continue
			}
			vmhostpcisstr := strconv.Itoa(vmhostpcis)
			wechatMessage := v.Annotations["description"] + "(GPU: " + vmhostpcisstr + "块)"
			WechatApiPost(wechatApi, wechatMessage, user)
			fmt.Println(v.Annotations["description"], user)
			user = []string{}
		}
		c.JSON(http.StatusOK, gin.H{"message": " successful receive alert notification message!"})

	})
	router.GET("/nodesremainsource", func(c *gin.Context) {
		//c.JSON(http.StatusOK, gin.H{"data": GetNodesRS()})
		c.String(http.StatusOK, GetNodesRS())
	})

	router.GET("/vmconfig", func(c *gin.Context) {
		//c.JSON(http.StatusOK, gin.H{"data": GetNodesRS()})
		vmconfig := getGxVmConfig()

		c.JSON(http.StatusOK, vmconfig)
	})

	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello,here is OPS webhook....You can't do anything!!!\n")
	})
	router.StaticFile("/favicon.ico", "./static/favicon.ico")
	router.Run(":9080")
}
