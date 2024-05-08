package s3

import (
	"github.com/FaiyazMujawar/golang-todo-app/src/config"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

var s3Client *s3.Client = nil

func GetS3Client() s3.Client {
	if s3Client == nil {
		s3Config := config.GetS3Config()

		s3Client =
			s3.New(s3.Options{
				Region:       *aws.String(s3Config.Region),
				BaseEndpoint: aws.String(s3Config.BaseEndpoint),
				Credentials:  aws.NewCredentialsCache(credentials.NewStaticCredentialsProvider(s3Config.ClientId, s3Config.ClientSecret, "")),
			})
	}
	return *s3Client
}
