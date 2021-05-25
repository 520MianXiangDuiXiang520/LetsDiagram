package utils

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func TestBase64ToFile(t *testing.T) {
	b64, err := ioutil.ReadFile("./b64_test.txt")
	if err != nil {
		t.Errorf("cannot open file b64:%v", err)
	}
	res, err := CompressCover(string(b64))
	if err != nil {
		t.Error(err)
	}
	fmt.Println(res)
}
