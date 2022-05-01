package handlers

import (
	"austin-go/app/austin-common/enums/channelType"
	"austin-go/app/austin-common/types"
)

type IHandler interface {
	DoHandler(taskInfo types.TaskInfo) (err error)
}

var handlerHolder = map[int]IHandler{
	channelType.Sms:                NewSmsHandler(),
	channelType.Email:              NewEmailHandler(),
	channelType.OfficialAccounts:   NewOfficialAccountHandler(),
	channelType.EnterpriseWeChat:   NewEnterpriseWeChatHandler(),
	channelType.DingDing:           NewDingDingRobotHandler(),
	channelType.DingDingWorkNotice: NewDingDingWorkNoticeHandler(),
}

func GetHandler(sendChannel int) IHandler {
	return handlerHolder[sendChannel]
}
