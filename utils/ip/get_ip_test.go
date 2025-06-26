package ip

import (
	"fmt"
	"testing"
)

func Test_getIP(T *testing.T) {
	addr := GetIP()
	fmt.Println(addr)

}
