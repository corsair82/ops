支持seaweedfs 2.16以上版本。
通过seaweedfs的master和volume server API进行数据读取。
go build -o seaweedfs-exporter .
启动示例：
    master单机：10.1.1.111
    ./seaweedfs-exporter -weedmasterip="10.1.1.111:9333"
    master集群：10.1.1.111 10.1.1.112 10.1.1.113
    ./seaweedfs-exporter -weedmasterip="10.1.1.111:9333,10.1.1.112:9333,10.1.1.113:9333"

