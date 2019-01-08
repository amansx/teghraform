package main

import "fmt"

var index map[string]func() bool = make(map[string]func() bool)

func BuildIndex() {

	index["aws s3 bucket doesn't exist"] = func() bool {
		fmt.Println("Executing::AWS.S3.BucketExists")
		return true
	}

	index["create aws s3 bucket"] = func() bool {
		fmt.Println("Executing::AWS.S3.CreateBucket")
		return true
	}

}
