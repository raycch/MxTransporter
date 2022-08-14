package eventbridge

import (
	"os"

	"github.com/cam-inc/mxtransporter/config/constant"
)

type Eventbridge struct {
	Eventbus string
	Source   string
	Region   string
}

func EventbridgeConfig() Eventbridge {
	var ebCfg Eventbridge
	ebCfg.Eventbus = os.Getenv(constant.EVENTBRIDGE_EVENTBUS_NAME)
	ebCfg.Source = os.Getenv(constant.EVENTBRIDGE_SOURCE)
	ebCfg.Region = os.Getenv(constant.EVENTBRIDGE_REGION)
	return ebCfg
}
