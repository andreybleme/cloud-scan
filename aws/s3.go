package aws

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/s3"
)

func CheckBuckets(region string) {
    session, err := session.NewSession(&aws.Config{
        Region: aws.String(region)},
    )

	svc := s3.New(session)
	input := &s3.ListBucketsInput{}

	result, err := svc.ListBuckets(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			fmt.Println(err.Error())
		}
		return
	}

	for _, bucket := range result.Buckets {
		status := "public"
		c := color.New(color.FgRed)

		if isBucketSecure(svc, aws.StringValue(bucket.Name)) {
			c = color.New(color.FgGreen)
			status = "secure"
		}
		c.Print(status)
		fmt.Print(" ", aws.StringValue(bucket.Name))
		fmt.Println()
	}
}

func isBucketSecure(svc *s3.S3, bucketName string) bool {
	bucketInput := &s3.GetPublicAccessBlockInput{Bucket: &bucketName}
	publicBlock, _ := svc.GetPublicAccessBlock(bucketInput)

	if publicBlock.PublicAccessBlockConfiguration == nil {
		return false
	}

	if !aws.BoolValue(publicBlock.PublicAccessBlockConfiguration.BlockPublicAcls) ||
		!aws.BoolValue(publicBlock.PublicAccessBlockConfiguration.BlockPublicPolicy) ||
		!aws.BoolValue(publicBlock.PublicAccessBlockConfiguration.IgnorePublicAcls) ||
		!aws.BoolValue(publicBlock.PublicAccessBlockConfiguration.RestrictPublicBuckets) {
		return false
	}

	return true
}
