package main

import (
	"bitbucket.org/mozillazg/go-cos"
	"context"
	"fmt"
	"os"
	"time"
)

func main() {
	b, _ := cos.ParseBucketFromDomain("testhuanan-1253846586.cn-south.myqcloud.com")
	c := cos.NewClient(os.Getenv("COS_SECRETID"), os.Getenv("COS_SECRETKEY"), b, nil)
	startTime := time.Now()
	endTime := startTime.Add(time.Hour)
	v, _, err := c.Bucket.GetLifecycle(context.Background(), cos.NewAuthTime(
		startTime, endTime,
		startTime, endTime))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%#v", v)
}