package lib_test

import (
	"cardcaptor/lib"
	"testing"
)

func Test파일이없으면false를반환(t *testing.T) {
	isExist := lib.IsFileExist("./non-existing-file")
	if isExist {
		t.Fail()
	}
}

func Test파일이있으면true를반환(t *testing.T) {
	isExist := lib.IsFileExist("./util.go")
	if !isExist {
		t.Fail()
	}
}

func Test경로가존재하지않으면false를반환(t *testing.T) {
	isExist := lib.IsFileExist("./resources/non-directory/")
	if isExist {
		t.Fail()
	}
}
