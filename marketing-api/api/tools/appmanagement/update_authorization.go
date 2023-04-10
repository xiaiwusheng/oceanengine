package appmanagement

import (
	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/tools/appmanagement"
)

// UpdateAuthorization 更新应用共享关系
// 增加和删除应用资产的共享关系
func UpdateAuthorization(clt *core.SDKClient, accessToken string, req *appmanagement.UpdateAuthorizationRequest) error {
	return clt.Post("2/tools/app_management/update/authorization/", req, nil, accessToken)
}
