package msg

//
//import (
//	"fmt"
//
//	"github.com/zeromicro/go-zero/core/jsonx"
//
//	"chatim/service/elmim/api/internal/types"
//)
//
//type (
//	ElmMessage struct {
//		SubBizType string `json:"subBizType"`
//		BizType    string `json:"bizType"`
//		Payload    string `json:"payload"`
//	}
//
//	SendPayload struct {
//		Content     string   `json:"content"`
//		ContentType int      `json:"contentType"`
//		CreateTime  int64    `json:"createTime"`
//		Extensions  string   `json:"extensions"`
//		GroupID     string   `json:"groupId"`
//		MsgID       string   `json:"msgId"`
//		ReceiverIds []string `json:"receiverIds"`
//		SenderID    string   `json:"senderId"`
//	}
//
//	ElmMsgContent struct {
//		Data        string `json:"data,omitempty"`
//		DegradeText string `json:"degradeText,omitempty"`
//		Summary     string `json:"summary,omitempty"`
//		Title       string `json:"title,omitempty"`
//		Type        int    `json:"type,omitempty"`
//		Text        string `json:"text,omitempty"`
//	}
//
//	AtContent struct {
//		Elements []struct {
//			ElementContent string `json:"elementContent"`
//			ElementType    int    `json:"elementType"`
//		} `json:"elements"`
//	}
//	AtExtension struct {
//		Extensions struct {
//		} `json:"extensions"`
//		Text string `json:"text"`
//	}
//
//	TextContent struct {
//		Text string `json:"text"`
//	}
//
//	ImageContent struct {
//		FileType    int    `json:"fileType"`
//		Orientation int    `json:"orientation"`
//		Url         string `json:"url"`
//	}
//
//	ReadPayload struct {
//		Cid   string   `json:"cid"`
//		MsgId []string `json:"msgId"`
//		Uid   string   `json:"uid"`
//	}
//)
//
//const (
//	TypeSendMessage = "SEND_MESSAGE" // 发送消息
//	TypeReadMessage = "READ_MESSAGE" // 已读消息
//
//	ContentTokenError    = -1   // 踢下线消息
//	ContentTypeText      = 1    // 文本消息
//	ContentTypeImage     = 2    // 图片消息
//	ContentTypeVoice     = 3    // 语言消息
//	ContentTypeVideo     = 4    // 视频消息
//	ContentTypeLocation  = 5    // 位置消息
//	ContentTypeAt        = 8    // @消息
//	ContentTypeCustomize = 101  // 自定义消息
//	ContentReadMsg       = 1001 // 已读消息
//
//	PrefixCustomer = "10" // 客户
//	PrefixRider    = "20" // 骑手
//	PrefixShop     = "30" // 商家
//	PrefixPlatform = "40" // 平台
//)
//
//type (
//	DownMessage struct {
//		MsgId     string `json:"msgId"`               // 消息ID
//		GroupId   string `json:"groupId"`             // 会话ID
//		ShopId    string `json:"shopId"`              // 店铺ID
//		ShopName  string `json:"shopName,omitempty"`  // 店铺名称
//		ElmUserId string `json:"elmUserId,omitempty"` // 饿了么UserId
//		SendRole  int64  `json:"sendRole,omitempty"`  // 发消息的角色
//		SendId    int64  `json:"sendId,omitempty"`    // 发消息的admin
//		NickName  string `json:"nickName,omitempty"`  // adminName
//		Avatar    string `json:"avatar,omitempty"`    // admin头像
//		MsgType   int64  `json:"msgType,omitempty"`   // 消息类型
//		Content   string `json:"content,omitempty"`   // 消息内容
//		MsgTime   int64  `json:"msgTime,omitempty"`   // 消息时间
//		IsRead    int64  `json:"isRead"`              // 是否已读
//	}
//
//	UpMessage struct {
//		GroupId  string `json:"groupId"`  // 会话id
//		MsgType  int64  `json:"msgType"`  // 消息类型
//		MsgId    string `json:"msgId"`    // 消息ID
//		Content  string `json:"content"`  // 消息内容
//		SendRole int64  `json:"sendRole"` // 发送者角色
//	}
//)
//
//func BuildSendMsg(p SendPayload) []byte {
//	switch p.ContentType {
//	case ContentTypeImage:
//		c := ImageContent{
//			FileType: 3,
//			Url:      p.Content,
//		}
//		p.Content, _ = jsonx.MarshalToString(c)
//	default:
//	}
//	marshal, _ := jsonx.MarshalToString(p)
//	bytes, _ := jsonx.Marshal(ElmMessage{
//		SubBizType: TypeSendMessage,
//		BizType:    "IM",
//		Payload:    marshal,
//	})
//	return bytes
//}
//
//func BuildReadMsg(msgId string, shopId string) []byte {
//	type read struct {
//		MsgId string `json:"msgId"`
//		Uid   string `json:"uid"`
//	}
//	var r = read{
//		MsgId: msgId,
//		Uid:   fmt.Sprintf("%s%s", PrefixShop, shopId),
//	}
//	marshal, _ := jsonx.MarshalToString(r)
//	bytes, _ := jsonx.Marshal(ElmMessage{
//		SubBizType: TypeReadMessage,
//		BizType:    "IM",
//		Payload:    marshal,
//	})
//	return bytes
//}
//
//func TokenErrorMsg(msg string) []byte {
//	marshal, _ := jsonx.Marshal(types.DownMessage{
//		MsgType: ContentTokenError,
//		Content: msg,
//	})
//	return marshal
//}
