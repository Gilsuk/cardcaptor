package lib_test

import (
	cc "cardcaptor"
	"testing"
)

func Test파일이없으면false를반환(t *testing.T) {
	isExist := cc.IsFileExist("./resources/test/not-existed-conf.json")
	if isExist {
		t.Fail()
	}
}

func Test파일이있으면true를반환(t *testing.T) {
	isExist := cc.IsFileExist("./resources/test/testconf.json")
	if !isExist {
		t.Fail()
	}
}

func Test경로가존재하지않으면false를반환(t *testing.T) {
	isExist := cc.IsFileExist("./resources/non-directory/conf.json")
	if isExist {
		t.Fail()
	}
}
