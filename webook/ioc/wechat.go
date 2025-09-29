package ioc

import (
	"webook/internal/service/oauth2/wechat"
	"webook/internal/web"
)

func InitOAuth2WechatService() wechat.Service {
	appId := "123"
	appSecret := "456"
	return wechat.NewService(appId, appSecret)
}

func NewWechatHandler() web.WechatHandlerConfig {
	return web.WechatHandlerConfig{
		Secure: false,
	}

}
