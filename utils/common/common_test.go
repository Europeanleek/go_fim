package common

import (
	"fmt"
	"testing"
)

func TestMD(t *testing.T) {
	md := MD5([]byte("1234"))
	fmt.Println(md)
}

func TestGetFilePrefix(t *testing.T) {
	fmt.Println(GetFilePrefix("name.png"))
	fmt.Println(GetFilePrefix("name.asdf.png"))
	fmt.Println(GetFilePrefix("..name.png"))
}
