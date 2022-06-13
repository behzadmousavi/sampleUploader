package trial

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

//func main() {
//	ctx := context.Background()
//	endpoint := "play.min.io"
//	accessKeyID := "Q3AM3UQ867SPQQA43P2F"
//	secretAccessKey := "zuf+tfteSlswRu7BJ86wekitnifILbZam1KYY3TG"
//	useSSL := true
//
//	// Initialize minio client object.
//	minioClient, err := minio.New(endpoint, &minio.Options{
//		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
//		Secure: useSSL,
//	})
//	if err != nil {
//		log.Fatalln(err)
//	}
//
//	// Make a new bucket called mymusic.
//	bucketName := "images"
//	location := "us-east-1"
//
//	err = minioClient.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{Region: location})
//	if err != nil {
//		// Check to see if we already own this bucket (which happens if you run this twice)
//		exists, errBucketExists := minioClient.BucketExists(ctx, bucketName)
//		if errBucketExists == nil && exists {
//			log.Printf("We already own %s\n", bucketName)
//		} else {
//			log.Fatalln(err)
//		}
//	} else {
//		log.Printf("Successfully created %s\n", bucketName)
//	}
//
//	// Upload the zip file
//	filePath := "/home/behzad/Desktop/1651318903270.JPG"
//	file, _ := os.OpenFile(filePath, os.O_CREATE, 0644)
//	stat, _ := file.Stat()
//	contentType, contentSubType, _ := DetectContentType(file)
//	objectName := uuid.New().String() + "." + contentSubType
//	log.Printf(objectName)
//	log.Printf(contentType)
//	log.Printf(contentSubType)
//
//	// Upload the zip file with FPutObject
//	info, err := minioClient.PutObject(ctx, bucketName, objectName, file, stat.Size(), minio.PutObjectOptions{
//		ContentType: contentType,
//	})
//	if err != nil {
//		log.Fatalln(err)
//	}
//
//	log.Printf("Successfully uploaded %s of size %d\n", objectName, info.Size)
//	fmt.Println(info.Bucket, info)
//
//	_, err = minioClient.GetObject(ctx, bucketName, objectName, minio.GetObjectOptions{})
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//
//	// Set request parameters for content-disposition.
//	reqParams := make(url.Values)
//	//reqParams.Set("response-content-disposition", "attachment; filename=\"your-filename.txt\"")
//
//	// Generates a presigned url which expires in a day.
//	presignedURL, err := minioClient.PresignedGetObject(context.Background(), bucketName, objectName, time.Second * 24 * 60 * 60, reqParams)
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//	fmt.Println("Successfully generated presigned URL", presignedURL)
//
//	//localFile, err := os.Create("image")
//	//if err != nil {
//	//	fmt.Println(err)
//	//	return
//	//}
//	//if _, err = io.Copy(localFile, object); err != nil {
//	//	fmt.Println(err)
//	//	return
//	//}
//}

//type image struct{
//	file *multipart.FileHeader	`json:"file" form:"file" binding:"required"`
//}

func UploadFile(c *gin.Context) {
	//var in image
	//if err := c.ShouldBind(in); err != nil {
	//	log.Printf(err.Error())
	//	return
	//}
	//f, err := in.file.Open()
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("file err : %s", err.Error()))
		return
	}
	defer file.Close()


	//filename := header.Filename
	//out, err := os.Create(filename)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//defer out.Close()
	//_, err = io.Copy(out, file)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//oldLocation := filename
	//newLocation := "/home/behzad/Desktop/" + filename
	//err = os.Rename(oldLocation, newLocation)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//filepath := newLocation

	ctx := context.Background()
	endpoint := "public-public-minio.apps.private.teh-1.snappcloud.io/"
	accessKeyID := "rasoul_user"
	secretAccessKey := "wL*!joZw}4<i.JO"
	useSSL := true

	// Initialize minio client object.
	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		log.Fatalln(err)
	}

	// Make a new bucket called mymusic.
	bucketName := "notifications"
	//location := "us-east-1"

	//err = minioClient.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{Region: location})
	//if err != nil {
	//	// Check to see if we already own this bucket (which happens if you run this twice)
	//	exists, errBucketExists := minioClient.BucketExists(ctx, bucketName)
	//	if errBucketExists == nil && exists {
	//		log.Printf("We already own %s\n", bucketName)
	//	} else {
	//		log.Fatalln(err)
	//	}
	//} else {
	//	log.Printf("Successfully created %s\n", bucketName)
	//}

	size := header.Size
	//file, _ := header.Open()
	//buffer := make([]byte, size)
	//_, err = file.Read(buffer)
	//defer file.Close()
	objectName, err := CreateObjectName(file)
	//reader := bytes.NewReader(buffer)
	if err != nil {
		log.Printf(err.Error())
	}
	//log.Printf(objectName)
	//log.Printf(contentType)
	//log.Printf(contentSubType)

	// Upload the zip file with FPutObject
	info, err := minioClient.PutObject(ctx, bucketName, objectName, file, size, minio.PutObjectOptions{
		//ContentType: contentType,
	})

	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("Successfully uploaded %s of size %d\n", objectName, info.Size)
	fmt.Println(info.Bucket, info)

	_, err = minioClient.GetObject(ctx, bucketName, objectName, minio.GetObjectOptions{})
	if err != nil {
		fmt.Println(err)
		return
	}

	// Set request parameters for content-disposition.
	reqParams := make(url.Values)
	//reqParams.Set("response-content-disposition", "attachment; filename=\"your-filename.txt\"")

	// Generates a presigned url which expires in a day.
	presignedURL, err := minioClient.PresignedGetObject(context.Background(), bucketName, objectName, time.Second * 24 * 60 * 60, reqParams)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Successfully generated presigned URL", presignedURL)


	c.JSON(http.StatusOK, gin.H{"filepath": presignedURL.String()})
}

func DetectContentType(out multipart.File) (string, error) {

	// Only the first 512 bytes are used to sniff the content type.
	buffer := make([]byte, 512)

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