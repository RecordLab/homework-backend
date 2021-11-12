package service

import (
	"mime/multipart"
	"path/filepath"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	uuid "github.com/satori/go.uuid"

	"dailyscoop-backend/config"
)

type AWSService struct {
	cfg config.AWSConfig
}

func NewAWSService(cfg config.AWSConfig) *AWSService {
	return &AWSService{
		cfg: cfg,
	}
}

func (as *AWSService) UploadImage(file multipart.File, fileName string) (string, error) {
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(as.cfg.Region),
		Credentials: credentials.NewStaticCredentials(as.cfg.AccessKey, as.cfg.SecretAccessKey, ""),
	})
	if err != nil {
		return "", err
	}
	uploader := s3manager.NewUploader(sess)
	newFileName := uuid.NewV4().String() + filepath.Ext(fileName)
	_, err = uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(as.cfg.Bucket),
		Key:    aws.String(newFileName),
		Body:   file,
	})
	if err != nil {
		return "", err
	}
	return as.cfg.URL + newFileName, nil
}
