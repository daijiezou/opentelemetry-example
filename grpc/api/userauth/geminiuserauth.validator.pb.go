// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: proto/geminiuserauth/geminiuserauth.proto

package geminiuserauth

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *ReFreshTokenReq) Validate() error {
	return nil
}
func (this *ReFreshTokenReply) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *ReFreshTokenReplyData) Validate() error {
	return nil
}
func (this *UserLoginReq) Validate() error {
	return nil
}
func (this *UserLoginReply) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *UserLoginReplyData) Validate() error {
	return nil
}
func (this *GeminiAuthRequest) Validate() error {
	return nil
}
func (this *GeminiAuthReply) Validate() error {
	return nil
}
func (this *GeminiLoginAuthRequest) Validate() error {
	return nil
}
func (this *GeminiLoginAuthReplay) Validate() error {
	return nil
}
func (this *PullImageAuthRequest) Validate() error {
	return nil
}
func (this *PullImageAuthReplay) Validate() error {
	return nil
}
func (this *PushImageAuthRequest) Validate() error {
	return nil
}
func (this *PushImageAuthReplay) Validate() error {
	return nil
}
func (this *RefreshLicenseRequest) Validate() error {
	return nil
}
func (this *RefreshLicenseReplay) Validate() error {
	return nil
}
func (this *GetUserInfoByIdRequest) Validate() error {
	return nil
}
func (this *UserInfo) Validate() error {
	return nil
}
func (this *GetUserInfoByIdReply) Validate() error {
	if this.UserInfo != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.UserInfo); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("UserInfo", err)
		}
	}
	return nil
}
func (this *CheckKeyRequest) Validate() error {
	return nil
}
func (this *CheckKeyReply) Validate() error {
	return nil
}
func (this *GetKeyRequest) Validate() error {
	return nil
}
func (this *GetKeyReply) Validate() error {
	return nil
}
func (this *GetUsersByRoleIdRequest) Validate() error {
	return nil
}
func (this *GetUsersByRoleIdReply) Validate() error {
	for _, item := range this.UserInfo {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("UserInfo", err)
			}
		}
	}
	return nil
}
func (this *GetRolesInfoByIdRequest) Validate() error {
	return nil
}
func (this *RoleInfo) Validate() error {
	return nil
}
func (this *GetRolesInfoByIdReply) Validate() error {
	for _, item := range this.RoleInfo {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("RoleInfo", err)
			}
		}
	}
	return nil
}
func (this *UserInfoListReq) Validate() error {
	return nil
}
func (this *UserInfoListReply) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *UserInfoListReply_Data) Validate() error {
	for _, item := range this.UserInfoList {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("UserInfoList", err)
			}
		}
	}
	return nil
}
