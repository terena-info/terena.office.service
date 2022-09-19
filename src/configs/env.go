package configs

import (
	"os"

	_ "github.com/joho/godotenv/autoload"
)

type Environments struct {
	APP_ENV             string
	PORT                string
	MONGO_URI           string
	MONGO_DATABASE_NAME string
	SLACK_ERROR_URI     string
	REDIS_URI           string
	REDIS_USER          string
	REDIS_PASSWORD      string
	AWS_SECRET_KEY      string
	AWS_ACCESS_KEY      string
	AWS_BUCKET_NAME     string
	AWS_REGION          string
	JWT_SECRET_KEY      string
	JWT_EXPIRE_TIME     string
}

var Env Environments

func LoadEnv() {
	Env.APP_ENV = os.Getenv("APP_ENV")
	Env.PORT = os.Getenv("PORT")
	Env.MONGO_URI = os.Getenv("MONGO_URI")
	Env.MONGO_DATABASE_NAME = os.Getenv("MONGO_DATABASE_NAME")
	Env.SLACK_ERROR_URI = os.Getenv("SLACK_ERROR_URI")
	Env.REDIS_URI = os.Getenv("REDIS_URI")
	Env.REDIS_USER = os.Getenv("REDIS_USER")
	Env.REDIS_PASSWORD = os.Getenv("REDIS_PASSWORD")
	Env.AWS_SECRET_KEY = os.Getenv("AWS_SECRET_KEY")
	Env.AWS_ACCESS_KEY = os.Getenv("AWS_ACCESS_KEY")
	Env.AWS_BUCKET_NAME = os.Getenv("AWS_BUCKET_NAME")
	Env.AWS_REGION = os.Getenv("AWS_REGION")
	Env.JWT_SECRET_KEY = os.Getenv("JWT_SECRET_KEY")
	Env.JWT_EXPIRE_TIME = os.Getenv("JWT_EXPIRE_TIME")
}
