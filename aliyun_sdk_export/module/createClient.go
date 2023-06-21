package module

import (
	cms20190101 "github.com/alibabacloud-go/cms-20190101/v8/client"
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	"github.com/alibabacloud-go/tea/tea"
	vpc20160428 "github.com/alibabacloud-go/vpc-20160428/v2/client"
)

func CreateCmsClient(accessKeyId *string, accessKeySecret *string, endpoint string) (_result *cms20190101.Client, _err error) {
	config := &openapi.Config{
		// 必填，您的 AccessKey ID
		AccessKeyId: accessKeyId,
		// 必填，您的 AccessKey Secret
		AccessKeySecret: accessKeySecret,
	}
	// 访问的域名
	config.Endpoint = tea.String(endpoint)
	//	_result = &cms20190101.Client{}
	_result, _err = cms20190101.NewClient(config)
	return _result, _err
}

func CreateVpcClient(accessKeyId *string, accessKeySecret *string, endpoint string) (_result *vpc20160428.Client, _err error) {
	config := &openapi.Config{
		// 必填，您的 AccessKey ID
		AccessKeyId: accessKeyId,
		// 必填，您的 AccessKey Secret
		AccessKeySecret: accessKeySecret,
	}
	// 访问的域名
	config.Endpoint = tea.String(endpoint)
	//	_result = &vpc20160428.Client{}
	_result, _err = vpc20160428.NewClient(config)
	return _result, _err
}
