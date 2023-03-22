package ws

import "fmt"

// GetMtflowerWsUrl 美团鲜花wss链接
func GetMtflowerWsUrl(key, token string) (wsUrl string) {
	wsUrl = fmt.Sprintf("wss://wpush.meituan.com/websocket/%s_WMOPEN/%s", key, token)
	return
}
