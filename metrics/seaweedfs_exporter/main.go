package main

import (
	"flag"
	"fmt"
	"net/http"
	collector "seaweedfs_exporter/collector"
	"strings"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	log "github.com/sirupsen/logrus"
)

var (
	opsadmin      string = "xuang@moqi.ai"
	version       string = "v1.0"
	help          bool
	disable       string   //命令行传入的需要关闭的指标
	disables      []string //处理命令行传入的根据,分割为一个切片做处理
	listenaddress = flag.String("listenAddress", ":9088", "Address on which to expose metrics and web interface.")
)

func init() {
	flag.BoolVar(&help, "h", false, "help")
	flag.StringVar(&disable, "disable", "", "关闭的指标收集器")
}

func main() {
	flag.Parse()
	if help {
		flag.Usage()
		return
	}
	disables = strings.Split(disable, ",")
	//手动开关
	//通过用户输入的我们做关闭
	for scraper, _ := range collector.Scrapers {
		for _, v := range disables {
			if v == scraper.Name() {
				collector.Scrapers[scraper] = false
				break
			}
		}
	}

	//访问/的时候返回一些基础提示
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`<html>
             <head><title>` + collector.Name() + "_exporter" + `</title></head>
             <body>
             <h1><a style="text-decoration:none" href=''>` + "SeaWeedFs_exporter" + `</a></h1>
             <p><a href='` + "metrics" + `'>Metrics</a></p>
             <h2>Build</h2>
             <pre>` + version + `</pre>
             </body>
             </html>`))
	})
	//根据开关来判断指标的是否需要收集  这里只有代码里面的判断  用户手动开关还未做
	enabledScrapers := []collector.Scraper{}
	for scraper, enabled := range collector.Scrapers {
		if enabled {
			log.Info("Scraper enabled ", scraper.Name())
			enabledScrapers = append(enabledScrapers, scraper)
		}
	}

	//注册自身采集器
	exporter := collector.New(collector.NewMetrics(), enabledScrapers)
	prometheus.MustRegister(exporter)

	http.Handle("/metrics", promhttp.Handler())
	//监听端口
	fmt.Println("Beginning to serve on port", *listenaddress)
	log.Fatal(http.ListenAndServe(*listenaddress, nil))
	/*
		if err := http.ListenAndServe(*listenaddress, nil); err != nil {
			log.Printf("Error occur when start server %v", err)
		}
	*/
}
