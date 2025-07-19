package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"io"
	"os"
)

var (
	s3Client *s3.S3
	s3Bucket string
)

func init() {
	sess, err := session.NewSession(
		&aws.Config{
			Region: aws.String("us-east-1"),
			Credentials: credentials.NewStaticCredentials(
				"EXAMPLEACCESSKEY",
				"EXAMPLESECRETKEY",
				"",
			),
		},
	)
	if err != nil {
		panic(err)
	}
	s3Client = s3.New(sess)
	s3Bucket = "goexpert-bucket-exemplo"
}

func main() {
	dir, err := os.Open("./tmp")
	if err != nil {
		panic(err)
	}
	defer dir.Close()
	for {
		files, err := dir.ReadDir(1)
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Sprintf("Error reading directory: %v\n", err)
			continue
		}
		uploadFile(files[0].Name())
	}
}

func uploadFile(fileName string) {
	completeFileName := fmt.Sprintf("./tmp/%s", fileName)
	fmt.Printf("Uploading file %s to bucket %s", fileName, s3Bucket)
	f, err := os.Open(completeFileName)
	if err != nil {
		fmt.Printf("Error opening file: %s\n", completeFileName)
		return
	}
	defer f.Close()
	_, err = s3Client.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(s3Bucket),
		Key:    aws.String(fileName),
		Body:   f,
	})
	if err != nil {
		fmt.Printf("Error uploading file: %s\n", completeFileName)
		return
	}
	fmt.Printf("File uploaded successfully: %s\n", completeFileName)
}
