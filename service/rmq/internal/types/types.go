// Code generated by goctl. DO NOT EDIT.
package types

type EbDownMsgReq struct {
	SubBizType     string       `json:"subBizType"`     //业务子类型，枚举值：SEND_MESSAGE-发送消息
	BizType        string       `json:"bizType"`        //业务类型，枚举值：IM-消息
	Payload        *PayloadItem `json:"payload"`        //业务结构
	PlatformShopId string       `json:"platformShopId"` //平台门店ID
}

type PayloadItem struct {
	SenderId    string   `json:"senderId"`    //消息发送方ID；格式：角色+随机数字串；角色：10(用户)、20(骑手)、30(商家)、32（连锁账号登录）；示例：20235760123
	ReceiverIds []string `json:"receiverIds"` //接收方ID；格式：角色+随机数字串；角色：10(用户)、20(骑手)、30(商家)、32（连锁账号登录）
	CreateTime  int64    `json:"createTime"`  //时间戳
	GroupId     string   `json:"groupId"`     //会话ID
	MsgId       string   `json:"msgId"`       //消息ID
	ContentType string   `json:"contentType"` //消息类型，枚举值：1-普通文本
	Content     string   `json:"content"`     //消息内容，格式：json
}

type ContentItem struct {
	Text string `json:"text"` //消息
}
