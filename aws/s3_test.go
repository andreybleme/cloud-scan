package aws

import (
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

const defaultRegion = "us-east-1"

const secureBucketName = "apple-private"
const publicBucketName = "apple-public"

func TestCheckBuckets(t *testing.T) {
	session, _ := session.NewSession(&aws.Config{
		Region: aws.String(defaultRegion)},
	)
	svc := s3.New(session)

	res := isBucketSecure(svc, secureBucketName)
	if !res {
		t.Errorf("Bucket %q should be secure", secureBucketName)
	}

	res = isBucketSecure(svc, publicBucketName)
	if res {
		t.Errorf("Bucket %q should be public", publicBucketName)
	}
}
