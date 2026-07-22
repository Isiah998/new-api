package relay

import (
	relaycommon "github.com/QingFlow/qing-api/relay/common"
	"github.com/QingFlow/qing-api/types"
)

func qingAPIErrorFromParamOverride(err error) *types.QingAPIError {
	if fixedErr, ok := relaycommon.AsParamOverrideReturnError(err); ok {
		return relaycommon.QingAPIErrorFromParamOverride(fixedErr)
	}
	return types.NewError(err, types.ErrorCodeChannelParamOverrideInvalid, types.ErrOptionWithSkipRetry())
}
