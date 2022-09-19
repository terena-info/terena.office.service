package configs

type Constants struct {
	PRODUCTION  string
	DEVELOPMENT string
}

var Const Constants

func init() {
	Const.PRODUCTION = "production"
	Const.DEVELOPMENT = "development"
}
