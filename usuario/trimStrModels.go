package usuario

import "strings"

func trimStrUserRequest(params *registerUserRequest) *registerUserRequest {
	params.UserName = strings.TrimSpace(params.UserName)
	params.Email = strings.TrimSpace(params.Email)
	params.Password = strings.TrimSpace(params.Password)
	params.Avatar = strings.TrimSpace(params.Avatar)
	return params
}
