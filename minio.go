package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"log"
	"mime/multipart"
	"net/http"
	"net/url"
	"time"
)

func UploadFile(c *gin.Context) {
	// Reading the file from template
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("file err : %s", err.Error()))
		return
	}
	defer file.Close()

	// Declaring minio client credentials
	ctx := context.Background()
	endpoint := "public-public-minio.apps.private.teh-1.snappcloud.io/"
	accessKeyID := "rasoul_user"
	secretAccessKey := "wL*!joZw}4<i.JO"
	useSSL := true
	bucketName := "notifications"

	// Initiating a minio client
	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		log.Fatalln(err)
	}
	// Reading file size
	size := header.Size

	// Creating a new name for the new object
	objectName, err := CreateObjectName(file)
	if err != nil {
		log.Printf(err.Error())
	}

	// Upload the file with PutObject
	info, err := minioClient.PutObject(ctx, bucketName, objectName, file, size, minio.PutObjectOptions{})
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("Successfully uploaded %s of size %d\n", objectName, info.Size)

	// Set request parameters for content-disposition
	reqParams := make(url.Values)

	// Generates a presigned url which expires in a day
	presignedURL, err := minioClient.PresignedGetObject(context.Background(), bucketName, objectName, time.Second * 24 * 60 * 60, reqParams)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Successfully generated presigned URL", presignedURL)

	c.JSON(http.StatusOK, gin.H{"filepath": presignedURL.String()})
}

// DetectContentType Detects file extension
func DetectContentType(out multipart.File) (string, error) {
	// Only the first 512 bytes are used to sniff the content type
	buffer := make([]byte, 512)

	// Reading the file from its stream beginning
	_, err := out.Read(buffer)
	if err != nil {
		return "", err
	}

	// Reset the read pointer if necessary so no part of the stream gets lost
	out.Seek(0, 0)


	// Use the net/http package's handy DectectContentType function. Always returns a valid
	// content-type by returning "application/octet-stream" if no others seemed to match.
	contentType := http.DetectContentType(buffer)

	return contentType, nil
}

var mimeToExt = map[string]string{
	"image/jpeg": ".jpg",
	"image/png":  ".png",
}

// CreateObjectName Creates an object name using uuid and file extension
func CreateObjectName(file multipart.File) (string, error) {
	contentType, err := DetectContentType(file)
	if err != nil {
		return "", err
	}

	contentSubType, ok := mimeToExt[contentType]
	if !ok {
		return "", fmt.Errorf("invalid mime: %s: ", contentType)
	}

	objectName := uuid.New().String() + contentSubType
	return objectName, nil
}

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("template/*")
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "select_file.html", gin.H{})
	})
	router.POST("/upload", UploadFile)
	router.StaticFS("/file", http.Dir("public"))
	router.Run(":8080")
}