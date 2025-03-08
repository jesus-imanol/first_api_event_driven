package utils

import (
	"bytes"
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/s3"
    "github.com/aws/aws-sdk-go/aws/credentials"
    "os"
)

func UploadToS3(file []byte, filename string) (string, error) {
    sess, err := session.NewSession(&aws.Config{
        Region: aws.String("us-east-1"),
        Credentials: credentials.NewStaticCredentials(
            os.Getenv("aws_access_key_id"),
            os.Getenv("aws_secret_access_key"),
            os.Getenv("aws_session_token"),
        ),
    })
    if err != nil {
        return "", err
    }

    svc := s3.New(sess)
    bucket := "spontaneity-2025"
    key := filename

    _, err = svc.PutObject(&s3.PutObjectInput{
        Bucket:        aws.String(bucket),
        Key:           aws.String(key),
        Body:          aws.ReadSeekCloser(bytes.NewReader(file)),
        ContentLength: aws.Int64(int64(len(file))),
        ContentType:   aws.String("image/jpeg"), // Ajustar según el tipo de imagen
    })
    if err != nil {
        return "", err
    }

    return "https://" + bucket + ".s3.amazonaws.com/" + key, nil
}
