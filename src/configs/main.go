package configs

type configInterface interface {
	ConnectDatabase()
	LoadEnvironments()
}

type configs struct{}

func New() configInterface {
	var cfg configInterface = &configs{}
	return cfg
}
