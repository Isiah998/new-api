package service

import (
	"github.com/Isiah998/new-api/service/relayconvert"
	"github.com/Isiah998/new-api/setting/model_setting"
)

func ShouldChatCompletionsUseResponsesPolicy(policy model_setting.ChatCompletionsToResponsesPolicy, channelID int, channelType int, model string) bool {
	return relayconvert.ShouldChatCompletionsUseResponsesPolicy(policy, channelID, channelType, model)
}

func ShouldChatCompletionsUseResponsesGlobal(channelID int, channelType int, model string) bool {
	return relayconvert.ShouldChatCompletionsUseResponsesGlobal(channelID, channelType, model)
}
