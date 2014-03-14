## aliyun SDK for Go language [![GoDoc](http://godoc.org/github.com/PinIdea/oss-aliyun-go?status.png)](http://godoc.org/github.com/PinIdea/oss-aliyun-go/)

### Build and Install 

go get:

```
$ go get github.com/PinIdea/oss-aliyun-go
```

### Example 

```go
package main

import (
	"log"

	"github.com/Pinidea/oss-aliyun-go"
)

func main() {
	var (
		accessId  = "z6HPq5j2aoqzAcbt"
		accessKey = "cunI67S8o5s5fvA9myvWkAayTUXhuI"
	)

	// region list:
	// HangZhou         = "oss-cn-hangzhou"
	// QingDao          = "oss-cn-qingdao"
	// HangZhouInternal = "oss-cn-hangzhou-internal"
	// QingdaoInternal  = "oss-cn-qingdao-internal"
	// DefaultRegion    = "oss"
	o := oss.New(oss.DefaultRegion, accessId, accessKey)

	// get a bucket instance
	bucket := o.Bucket("pinidea-test")

	// upload
	var data = []byte("this is a test")
	err := bucket.Put("test.txt", data, "text/plain", oss.PublicRead)
	if err != nil {
		panic(err)
	}
	log.Println("success upload object")

	// get object
	text, err := bucket.Get("test.txt")
	if err != nil {
		panic(err)
	}

	log.Println("get object content:", string(text))

	// delete object
	err = bucket.Del("test.txt")
	if err != nil {
		panic(err)
	}
	log.Println("success delete object")
}
```

### License

licensed under the LGPL

#### Status: Under Developing
