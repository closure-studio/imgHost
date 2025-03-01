package env

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// define a struct to store the environment variables
type Config struct {
	// the port to listen on
	PORT               string
	BASE_AUTH_USER     string
	BASE_AUTH_PASS     string
	REDROID_CLIENT_ID  string
	WEBSITES_CLIENT_ID string

	S3_ACCESS_KEY string
	S3_SECRET_KEY string
	S3_ENDPOINT   string
	S3_BUCKET     string
}

// an object of Environment

var Instance = Config{
	PORT:               "3000", // default port
	BASE_AUTH_USER:     "closure",
	BASE_AUTH_PASS:     "s0wDmTt1mcmV04skSmRcYLZEyNtZ4bFT",
	REDROID_CLIENT_ID:  "cdcc7399-6dff-4387-839a-4dfde70b21b7",
	WEBSITES_CLIENT_ID: "35083884-4a3f-476d-bc8b-df34dbe46806",

	S3_ACCESS_KEY: "",
	S3_SECRET_KEY: "",
	S3_ENDPOINT:   "",
	S3_BUCKET:     "",
}

func InstanceInit() {
	err := godotenv.Load()
	if err != nil {
		log.Println("⚠️ 未找到 .env 文件，使用默认值")
	}
	if port := os.Getenv("PORT"); port != "" {
		Instance.PORT = port
	}
	if baseAuthUser := os.Getenv("BASE_AUTH_USER"); baseAuthUser != "" {
		Instance.BASE_AUTH_USER = baseAuthUser
	}
	if baseAuthPass := os.Getenv("BASE_AUTH_PASS"); baseAuthPass != "" {
		Instance.BASE_AUTH_PASS = baseAuthPass
	}
	if redroidClientID := os.Getenv("REDROID_CLIENT_ID"); redroidClientID != "" {
		Instance.REDROID_CLIENT_ID = redroidClientID
	}
	if websitesClientID := os.Getenv("WEBSITES_CLIENT_ID"); websitesClientID != "" {
		Instance.WEBSITES_CLIENT_ID = websitesClientID
	}
	if s3AccessKey := os.Getenv("S3_ACCESS_KEY"); s3AccessKey != "" {
		Instance.S3_ACCESS_KEY = s3AccessKey
	}
	if s3SecretKey := os.Getenv("S3_SECRET_KEY"); s3SecretKey != "" {
		Instance.S3_SECRET_KEY = s3SecretKey
	}
	if s3Endpoint := os.Getenv("S3_ENDPOINT"); s3Endpoint != "" {
		Instance.S3_ENDPOINT = s3Endpoint
	}
	if s3Bucket := os.Getenv("S3_BUCKET"); s3Bucket != "" {
		Instance.S3_BUCKET = s3Bucket
	}
}
