package action

import (
	"austin-go/app/austin-common/enums/channelType"
	"austin-go/app/austin-common/enums/idType"
	"austin-go/app/austin-common/taskUtil"
	"austin-go/app/austin-common/types"
	"context"
	"github.com/pkg/errors"
	"regexp"
)

type AfterParamCheckAction struct {
}

func NewAfterParamCheckAction() *AfterParamCheckAction {
	return &AfterParamCheckAction{}
}

func (p AfterParamCheckAction) Process(_ context.Context, data interface{}) error {
	sendTaskModel, ok := data.(*types.SendTaskModel)
	if !ok {
		return errors.Wrapf(sendErr, "AssembleAction 类型错误 err:%v", data)
	}
	// 1. 过滤掉不合法的手机号

	if sendTaskModel.TaskInfo[0].IdType == idType.Phone && sendTaskModel.TaskInfo[0].SendChannel == channelType.Sms {
		var newTask []types.TaskInfo
		for _, item := range sendTaskModel.TaskInfo {
			for _, tel := range item.Receiver {
				matched, _ := regexp.Match(taskUtil.PhoneRegex, []byte(tel))
				if matched {
					newTask = append(newTask, item)
				}
			}
		}
		if len(newTask) <= 0 {
			return errors.Wrapf(sendErr, "AfterParamCheckAction err:%v", data)
		}
		sendTaskModel.TaskInfo = newTask
	}

	return nil
}
