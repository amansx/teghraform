package aws

// import "github.com/aws/aws-sdk-go/service/s3"

type S3 struct {}

func (this *S3Glue) BucketDoesntExistMeta() string{ return "aws s3 bucket doesn't exist" }
func (this *S3Glue) BucketDoesntExist(bucketName string) bool{
	return false
}