package protoreflect

import (
	"fmt"
	"gitee.com/zhaochuninhefei/zcutils-go/testdata/myproto-go/owner"
	"google.golang.org/protobuf/proto"
	"testing"
)

func Test_getFields_success(t *testing.T) {
	owner1 := &owner.Owner{
		OwnerId:   1,
		OwnerName: "owner1",
		OwnerDesc: "just test",
	}
	pb := proto.Message(owner1)

	got, err := GetFields(pb)
	if err != nil {
		fmt.Println(err)
	}
	if got != nil {
		for _, fd := range got {
			fmt.Printf("fd: %s\n", fd)
		}
	}
}
