package storageProvider

import (
	"fmt"
	"gymn/internal/models"
	awsclient "gymn/internal/services/aws"
	dateUtils "gymn/internal/utils/date"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

func SaveFile(file *models.FileModel, path string) (string, error) {
	uploader := s3manager.NewUploader(awsclient.Session)

	filename := fmt.Sprintf(
		"%d_%s",
		dateUtils.GetCurrentTimestamp(),
		file.Filename,
	)

	key := path + "/" + filename

	_, err := uploader.Upload(&s3manager.UploadInput{
		Bucket:      aws.String(os.Getenv("AWS_S3_BUCKET")),
		Key:         aws.String(key),
		Body:        file.File,
		ContentType: aws.String(file.ContentType),
	})

	if err != nil {
		return "", err
	}

	return key, nil
}
