package fulfilment

import "github.com/TukTuk/core"

var FF *FFClient

type FFClient struct {
	Cfg *core.Config
}

func InitFF(cfg *core.Config) {
	FF = &FFClient{
		Cfg: cfg,
	}
}
