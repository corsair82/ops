// This file is auto-generated, don't edit it. Thanks.
package module

import (
	"encoding/json"

	cms20190101 "github.com/alibabacloud-go/cms-20190101/v8/client"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

func AllUserDescribeMetricLast(namespace, metricname string) map[string][]Datapoint {
	aliyunusers := LoadUserInfo()
	metricLastInfo := make(map[string][]Datapoint)
	for _, aliyunuserinfo := range aliyunusers {
		username := aliyunuserinfo.Name
		accesskeyid := aliyunuserinfo.AccessKeyId
		accesskeysecret := aliyunuserinfo.AccessKeySecret
		client, _err := CreateCmsClient(tea.String(accesskeyid), tea.String(accesskeysecret), MonitorEndpoint)
		if _err != nil {
			panic(_err)
		}
		userdatapoints, err := DescribeMetricLast(namespace, metricname, client)
		if err != nil {
			panic(_err)
		}
		metricLastInfo[username] = userdatapoints

	}
	return metricLastInfo
}
func DescribeMetricLast(namespace, metricname string, client *cms20190101.Client) (datapoints []Datapoint, _err error) {
	// 工程代码泄露可能会导致AccessKey泄露，并威胁账号下所有资源的安全性。以下代码示例仅供参考，建议使用更安全的 STS 方式，更多鉴权访问方式请参见：https://help.aliyun.com/document_detail/378661.html
	describeMetricLastRequest := &cms20190101.DescribeMetricLastRequest{
		Namespace:  tea.String(namespace),
		MetricName: tea.String(metricname),
	}
	runtime := &util.RuntimeOptions{}
	resp, _err := client.DescribeMetricLastWithOptions(describeMetricLastRequest, runtime)

	if _err != nil {
		return nil, _err
	}
	json.Unmarshal([]byte(*resp.Body.Datapoints), &datapoints)
	return datapoints, _err
}
