package config

type Global struct {
	logEnable bool
}

func newGlobal(logEnable bool) *Global {
	return &Global{logEnable: logEnable}
}

func OfGlobal() *Global {
	return newGlobal(true)
}

func (g *Global) LogEnable() bool {
	return g.logEnable
}
