// Code generated by goctl. DO NOT EDIT.
package types

type WsReq struct {
	PlatformType int64  `form:"platformType" validate:"oneof=1,2"`
	Token        string `form:"token"`
}