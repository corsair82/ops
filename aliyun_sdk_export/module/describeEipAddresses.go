// This file is auto-generated, don't edit it. Thanks.
package module

import (
	"encoding/json"

	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
	vpc20160428 "github.com/alibabacloud-go/vpc-20160428/v2/client"
)

func DescribeEipAddresses(username string) (eipaddressinfomap map[string]string, _err error) {
	regionIds := []string{"cn-beijing", "us-west-1"}
	aliyunusers := LoadUserInfo()
	for _, aliyunuserinfo := range aliyunusers {
		if username == aliyunuserinfo.Name {
			accesskeyid := aliyunuserinfo.AccessKeyId
			accesskeysecret := aliyunuserinfo.AccessKeySecret
			eipaddressinfomap = make(map[string]string)
			client, _err := CreateVpcClient(tea.String(accesskeyid), tea.String(accesskeysecret), EipEndpoint)
			for _, ri := range regionIds {
				describeEipAddressesRequest := &vpc20160428.DescribeEipAddressesRequest{
					RegionId:               tea.String(ri),
					IncludeReservationData: tea.Bool(false),
					PageSize:               tea.Int32(50),
				}
				runtime := &util.RuntimeOptions{}
				resp, _err := client.DescribeEipAddressesWithOptions(describeEipAddressesRequest, runtime)
				if _err != nil {
					return nil, _err
				}

				var eipaddresslist map[string]interface{}
				marshal, _ := json.Marshal(*resp.Body.EipAddresses)
				json.Unmarshal([]byte(string(marshal)), &eipaddresslist)
				if eipaddresslist["EipAddress"] == nil {
					continue
				}
				eipaddressinfos := eipaddresslist["EipAddress"].([]interface{})
				for v := range eipaddressinfos {
					eipaddressinfo := eipaddressinfos[v].(map[string]interface{})
					ipaddress := string(eipaddressinfo["IpAddress"].(string))
					allocationid := string(eipaddressinfo["AllocationId"].(string))
					eipaddressinfomap[allocationid] = ipaddress + "_" + ri
				}

			}
			return eipaddressinfomap, _err
		}
	}
	return nil, nil

}
