package config

import "os"

type S3Config struct {
	Bucket       string
	Region       string
	BaseEndpoint string
	ClientId     string
	ClientSecret string
}

func GetS3Config() S3Config {
	return S3Config{
		Bucket:       os.Getenv("S3_BUCKET"),
		Region:       os.Getenv("S3_REGION"),
		BaseEndpoint: os.Getenv("S3_BASE_ENDPOINT"),
		ClientId:     os.Getenv("S3_ACCESS_KEY_ID"),
		ClientSecret: os.Getenv("S3_SECRET_ACCESS_KEY"),
	}
}

func Dsn() string {
	return os.Getenv("dsn")
}
