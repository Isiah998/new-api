package relay

import (
	relaycommon "github.com/Isiah998/new-api/relay/common"
	"github.com/Isiah998/new-api/types"
)

func qingAPIErrorFromParamOverride(err error) *types.QingAPIError {
	if fixedErr, ok := relaycommon.AsParamOverrideReturnError(err); ok {
		return relaycommon.QingAPIErrorFromParamOverride(fixedErr)
	}
	return types.NewError(err, types.ErrorCodeChannelParamOverrideInvalid, types.ErrOptionWithSkipRetry())
}
