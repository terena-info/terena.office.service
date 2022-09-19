package configs

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/bankonly/goutils/fileupload"
)

var (
	AwsS3    *s3.S3
	S3Sesion *session.Session
)

func InitAwsConfig() {
	AwsS3, S3Sesion = fileupload.NewS3(fileupload.Config{
		SecretAccessKey: Env.AWS_SECRET_KEY,
		AccessKeyID:     Env.AWS_ACCESS_KEY,
	})
}
