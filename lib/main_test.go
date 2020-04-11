package lib_test

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	println("+=+=+ Start Test +=+=+")
	os.Exit(m.Run())
}
