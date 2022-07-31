package storageProvider

import (
	awsclient "gymn/internal/providers/aws"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
)

func DeleteFile(key *string) error {
	svc := s3.New(awsclient.Session)

	_, err := svc.DeleteObject(&s3.DeleteObjectInput{
		Bucket: aws.String(os.Getenv("AWS_S3_BUCKET")),
		Key:    key,
	})

	if err != nil {
		return err
	}

	return nil
}
