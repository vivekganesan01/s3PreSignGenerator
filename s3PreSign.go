// ./s3Sign module helps to create signed url for the given bucket object and timeframe
package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

var awsAccessID string
var awsSecretKey string

// checkCreds checks env var for aws cred
func checkCreds() {
	ac, acok := os.LookupEnv("ACCESS_ID")
	sk, skok := os.LookupEnv("SECRET_KEY")
	if acok && skok {
		awsAccessID = ac
		awsSecretKey = sk
	} else {
		fmt.Println("Fatal Error: Please make sure to define ACCESS_ID & SECRET_KEY as env variable before execution.")
		os.Exit(1)
	}
}

// awsActiveSession helps to fetch aws sdk active session
func awsActiveSession() (source *session.Session) {
	fmt.Println("* - Getting the active session for aws")
	sourceSession, _ := session.NewSession(&aws.Config{
		Region:      aws.String("eu-west-1"),
		Credentials: credentials.NewStaticCredentials(awsAccessID, awsSecretKey, ""),
	})
	return sourceSession
}

func main() {
	// pre-requisite: First Check for the AWS prod cred
	checkCreds()

	// properties
	bn := flag.String("bucketname", "", "(string) source s3 bucket name")
	bnp := flag.String("bucketprefix", "", "(string) source file name inside the bucket")
	flag.Parse()

	if false {
		fmt.Println("./s3PreSign invalid arg count, [help] s3Sign -h \nUsage:")
		flag.PrintDefaults()
		fmt.Println("Example: [./s3PreSign -bucketname=xyz -bucketprefix=folder/filename.txt]")
		os.Exit(1)
	} else {
		if *bn == "" || *bnp == "" {
			fmt.Printf("Either bucketname and bucketprefix cant be blank, BucketName: %v,Prefix: %v", *bn, *bnp)
			flag.PrintDefaults()
			os.Exit(1)
		}
	}
	svc := s3.New(awsActiveSession())
	req, _ := svc.GetObjectRequest(&s3.GetObjectInput{
		Bucket: aws.String(*bn),
		Key:    aws.String(*bnp),
	})

	// hours, _ := strconv.Atoi(*hs)
	urlStr, err := req.Presign(72 * time.Hour)

	if err != nil {
		log.Println("Failed to sign request", err)
	}

	log.Println("The URL is", urlStr)
}
