package main

import (
	"io"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func readCSVFromS3(bucket, key, region string) (io.ReadCloser, error) {
	// Create AWS Session
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(region),
	},
	)
	if err != nil {
		return nil, err
	}

	// Create S3 Service Client
	svc := s3.New(sess)

	// Download CSV from S3
	req, _ := svc.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	})
	return req.Body, nil
}
