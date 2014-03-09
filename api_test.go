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

func TestGet(t *testing.T) {
	bucket := New(accessId, accessKey).Bucket(testBucket)
	data, err := bucket.Get("readme")
	if err != nil {
		t.Error("unable get resource", err)
		return
	}
	log.Println(string(data))
}
