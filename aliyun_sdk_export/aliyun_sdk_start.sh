#!/bin/sh
while true
do
  aliyun_sdk_export_proce=$(ps aux| awk '!/awk/&&/aliyun_sdk_export'/)
  if [[ -z ${aliyun_sdk_export_proce} ]]
  then
	nohup /opt/huatu/gowork/aliyun_sdk_export/aliyun_sdk_export &
  fi
  sleep 10
done
