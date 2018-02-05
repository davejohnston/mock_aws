package main

import (
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/aws/aws-sdk-go/service/ec2/ec2iface"
	"github.com/davejohnston/aws_sample/client"
)

func getClient(accessKey string, secretAccessKey string, defaultRegion string) ec2iface.EC2API {

	if os.Getenv("MOCK") != "" {
		return &client.MockEC2Client{}
	} else {

		sess, err := session.NewSession(&aws.Config{
			Region: aws.String(defaultRegion),
			Credentials: credentials.NewStaticCredentials(
				*aws.String(accessKey),
				*aws.String(secretAccessKey),
				*aws.String("")),
		})
		if err != nil {
			log.Fatalf("Unable to create new session %s", err)
		}

		return ec2.New(sess)
	}
}

func main() {

	accessKey := os.Getenv("AWS_ACCESS_KEY")
	secretAccessKey := os.Getenv("AWS_SECRET_ACCESS_KEY")
	defaultRegion := os.Getenv("AWS_DEFAULT_REGION")

	ec2Svc := getClient(accessKey, secretAccessKey, defaultRegion)

	instances, err := ec2Svc.DescribeInstances(nil)
	if err != nil {
		log.Println(err)
	}

	for _, v := range instances.Reservations {
		log.Printf("Found Intance %s with key %s", *v.Instances[0].InstanceId, *v.Instances[0].KeyName)
	}
}
