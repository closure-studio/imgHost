package storage

import (
	"bytes"
	"context"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/closure-studio/imgHost/utils/env"
)

type S3Config struct {
	Client *s3.Client
	Bucket string
}

var S3Instance *S3Config

func S3ClientInit() (*S3Config, error) {

	accessKey := env.Instance.S3_ACCESS_KEY
	secretKey := env.Instance.S3_SECRET_KEY
	endpoint := env.Instance.S3_ENDPOINT
	bucket := env.Instance.S3_BUCKET

	if accessKey == "" || secretKey == "" || endpoint == "" || bucket == "" {
		fmt.Println("S3 配置缺失，请检查环境变量")
		// os exit
		os.Exit(1)
	}

	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(accessKey, secretKey, "")),
		config.WithRegion("auto"),
	)
	if err != nil {
		fmt.Println("无法加载 AWS 配置:", err)
		os.Exit(1)
	}

	client := s3.NewFromConfig(cfg, func(o *s3.Options) {
		o.BaseEndpoint = aws.String(endpoint)
		o.UsePathStyle = true
	})
    S3Instance = &S3Config{
        Client: client,
        Bucket: bucket,
    }
    return S3Instance, nil
}

// UploadFile 上传文件到 S3
func (r *S3Config) UploadFile(ctx context.Context, key string, data []byte, contentType string) error {
	_, err := r.Client.PutObject(ctx, &s3.PutObjectInput{
		Bucket:      &r.Bucket,
		Key:         &key,
		Body:        bytes.NewReader(data),
		ContentType: &contentType,
	})
	if err != nil {
		return fmt.Errorf("上传文件失败: %w", err)
	}
	fmt.Println("文件上传成功:", key)
	return nil
}
