package pwd

import (
	"fmt"
	"testing"
)

func TestHashPwd(t *testing.T) {
	hash := HashPwd("123456890")
	fmt.Println(hash)
}

func TestCheckHashPwd(t *testing.T) {
	hash := HashPwd("123456")
	status := CheckPwd(hash, "13456")
	fmt.Println(status)
}
