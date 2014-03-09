package oss

import (
	"log"
	"testing"
)

var (
	accessId   = "z6HPq5j2aoqzAcbt"
	accessKey  = "cunI67S8o5s5fvA9myvWkAayTUXhuI"
	testBucket = "pinidea-test"
)

func TestNew(t *testing.T) {
	o := New(accessId, accessKey)
	if o == nil {
		t.Error("Unable new oss")
	}
}

func TestPutBucket(t *testing.T) {
	bucket := New(accessId, accessKey).Bucket(testBucket)
	err := bucket.PutBucket(PublicRead)
	if err != nil {
		t.Error("Unable put bucket:", err)
	}
}

func TestPut(t *testing.T) {
	bucket := New(accessId, accessKey).Bucket(testBucket)
	data := []byte("helloworld")
	err := bucket.Put("readme", data, "text/plain", Private)
	if err != nil {
		t.Error("Unable del object:", err)
	}
}

func TestGet(t *testing.T) {
	bucket := New(accessId, accessKey).Bucket(testBucket)
	data, err := bucket.Get("readme")
	if err != nil {
		t.Error("Unable get object:", err)
		return
	}
	log.Println(string(data))
}

func TestDel(t *testing.T) {
	bucket := New(accessId, accessKey).Bucket(testBucket)
	err := bucket.Del("readme")
	if err != nil {
		t.Error("Unable del object:", err)
	}
}

func TestDelBucket(t *testing.T) {
	bucket := New(accessId, accessKey).Bucket(testBucket)
	err := bucket.DelBucket()
	if err != nil {
		t.Error("Unable del bucket:", err)
	}
}

func TestURL(t *testing.T) {
	bucket := New(accessId, accessKey).Bucket(testBucket)
	url := bucket.URL("readme")
	if url != "http://oss.aliyuncs.com/pinidea-test/readme" {
		t.Error("Unable get correct url:", url)
	}
}
