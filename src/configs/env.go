package configs

import (
	"os"

	_ "github.com/joho/godotenv/autoload"
)

type _Env struct {
	DB_URI         string
	DB_NAME        string
	PORT           string
	APP_ENV        string
	JWT_SECRET_KEY string
	JWT_TIME       string
}

var Env _Env

func (cfg *configs) LoadEnvironments() {
	Env.DB_URI = os.Getenv("DB_URI")
	Env.DB_NAME = os.Getenv("DB_NAME")
	Env.PORT = os.Getenv("PORT")
	Env.APP_ENV = os.Getenv("APP_ENV")
	Env.JWT_SECRET_KEY = os.Getenv("JWT_SECRET_KEY")
	Env.JWT_TIME = os.Getenv("JWT_TIME")
}
