// Code is generated by ucloud-model, DO NOT EDIT IT.

package unet

import (
	"github.com/ucloud/ucloud-sdk-go/ucloud/request"
	"github.com/ucloud/ucloud-sdk-go/ucloud/response"
)

// UpdateVIPAttributeRequest is request schema for UpdateVIPAttribute action
type UpdateVIPAttributeRequest struct {
	request.CommonBase

	// [公共参数] 项目ID。不填写为默认项目，子帐号必须填写。 请参考[GetProjectList接口](../summary/get_project_list.html)
	// ProjectId *string `required:"true"`

	// [公共参数] 地域。 参见 [地域和可用区列表](../summary/regionlist.html)
	// Region *string `required:"true"`

	// 内网VIP的名称
	Name *string `required:"false"`

	// 内网VIP的备注
	Remark *string `required:"false"`

	// 内网VIP所属的业务组
	Tag *string `required:"false"`

	// 内网VIP的资源Id
	VIPId *string `required:"true"`
}

// UpdateVIPAttributeResponse is response schema for UpdateVIPAttribute action
type UpdateVIPAttributeResponse struct {
	response.CommonBase
}

// NewUpdateVIPAttributeRequest will create request of UpdateVIPAttribute action.
func (c *UNetClient) NewUpdateVIPAttributeRequest() *UpdateVIPAttributeRequest {
	req := &UpdateVIPAttributeRequest{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

// UpdateVIPAttribute - 更新VIP信息
func (c *UNetClient) UpdateVIPAttribute(req *UpdateVIPAttributeRequest) (*UpdateVIPAttributeResponse, error) {
	var err error
	var res UpdateVIPAttributeResponse

	reqCopier := *req

	err = c.Client.InvokeAction("UpdateVIPAttribute", &reqCopier, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}