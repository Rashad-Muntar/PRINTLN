package utils

import (
	"bytes"
	"io"
	"log"
	"net/http"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

const (
	AWS_S3_REGION = "println"
	AWS_S3_BUCKET = "af-south-1"
)


// func UploadFile(uploadFileDir string) (string, error) {
// 	session, err := session.NewSession(&aws.Config{Region: aws.String(AWS_S3_REGION)})
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	upFile, err := os.Open(uploadFileDir)
// 	if err != nil {
// 		return "", err
// 	}
// 	defer upFile.Close()

// 	upFileInfo, _ := upFile.Stat()
// 	var fileSize int64 = upFileInfo.Size()
// 	fileBuffer := make([]byte, fileSize)
// 	upFile.Read(fileBuffer)

// 	_, err = s3.New(session).PutObject(&s3.PutObjectInput{
// 		Bucket:               aws.String(AWS_S3_BUCKET),
// 		Key:                  aws.String(uploadFileDir),
// 		ACL:                  aws.String("private"),
// 		Body:                 bytes.NewReader(fileBuffer),
// 		ContentLength:        aws.Int64(fileSize),
// 		ContentType:          aws.String(http.DetectContentType(fileBuffer)),
// 		ContentDisposition:   aws.String("attachment"),
// 		ServerSideEncryption: aws.String("AES256"),
// 	})
// 	if err != nil {
// 		return "", err
// 	}

// 	url := "https://%s.s3.amazonaws.com/%s"

// 	return url, nil
// }

func UploadFile(file io.Reader) (string, error) {
	session, err := session.NewSession(&aws.Config{Region: aws.String(AWS_S3_REGION)})
	if err != nil {
		log.Fatal(err)
	}

	// Read the file contents into a buffer
	fileBuffer := bytes.Buffer{}
	_, err = fileBuffer.ReadFrom(file)
	if err != nil {
		return "", err
	}

	_, err = s3.New(session).PutObject(&s3.PutObjectInput{
		Bucket:               aws.String(AWS_S3_BUCKET),
		Key:                  aws.String("print"),
		ACL:                  aws.String("private"),
		Body:                 bytes.NewReader(fileBuffer.Bytes()),
		ContentLength:        aws.Int64(int64(fileBuffer.Len())),
		ContentType:          aws.String(http.DetectContentType(fileBuffer.Bytes())),
		ContentDisposition:   aws.String("attachment"),
		ServerSideEncryption: aws.String("AES256"),
	})
	if err != nil {
		return "", err
	}

	url := "https://%s.s3.amazonaws.com/%s"

	return url, nil
}