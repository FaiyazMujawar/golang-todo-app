package s3

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"mime/multipart"
	"time"

	"github.com/FaiyazMujawar/golang-todo-app/src/config"
	"github.com/FaiyazMujawar/golang-todo-app/src/utils"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
)

func UploadFiles(files ...*multipart.FileHeader) []string {
	urls := utils.Map(files, uploadFile)
	return utils.Filter(urls, filterEmptyString)
}

func DeleteFiles(keys ...string) error {
	fmt.Println("keys", keys)
	client := GetS3Client()
	s3Config := config.GetS3Config()
	objectKeys := utils.Map(keys, stringToObjectIdentifier)
	_, err := client.DeleteObjects(context.TODO(), &s3.DeleteObjectsInput{
		Bucket: &s3Config.Bucket,
		Delete: &types.Delete{Objects: objectKeys},
	})
	return err
}

// Helper functions

func uploadFile(fileHeader *multipart.FileHeader) string {
	fileExtension := extractExtensionFromFilename(fileHeader.Filename)
	file, err := fileHeader.Open()
	if err != nil {
		return ""
	}
	filename := fmt.Sprintf("%d.%s", time.Now().UnixNano(), fileExtension)
	err = upload(filename, file)
	if err != nil {
		log.Println("error uploading file: ", err)
		return ""
	}
	return constructFileUrl(filename)
}

func upload(filename string, file multipart.File) error {
	defer file.Close()
	client := GetS3Client()
	s3Config := config.GetS3Config()
	var filedata []byte
	_, err := file.Read(filedata)
	if err != nil {
		return fmt.Errorf("error reading file: %s", err.Error())
	}

	_, err = client.PutObject(context.TODO(), &s3.PutObjectInput{
		Key:    &filename,
		Bucket: &s3Config.Bucket,
		Body:   bytes.NewReader(filedata),
	})
	return err
}

func constructFileUrl(filename string) string {
	s3Config := config.GetS3Config()
	return fmt.Sprintf("%s/%s/%s", s3Config.BaseEndpoint, s3Config.Bucket, filename)
}
