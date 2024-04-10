package config

type Global struct {
	logEnable bool
}

func NewGlobal() *Global {
	return &Global{
		logEnable: true,
	}
}

func (g *Global) LogEnable() bool {
	return g.logEnable
}
