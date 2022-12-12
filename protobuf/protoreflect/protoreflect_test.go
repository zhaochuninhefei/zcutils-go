package protoreflect

import (
	"fmt"
	"testing"

	"gitee.com/zhaochuninhefei/zcutils-go/protobuf/myproto-go/asset"
	"gitee.com/zhaochuninhefei/zcutils-go/protobuf/myproto-go/owner"
	"google.golang.org/protobuf/proto"
)

func Benchmark(b *testing.B) {
	asset1 := &asset.BasicAsset{
		AssetId:    1,
		AssetName:  "测试资产1",
		AssetPrice: 108,
		AssetOwner: &owner.Owner{
			OwnerId:   0,
			OwnerName: "张三",
		},
		AssetNum:    &asset.BasicAsset_AssetNumInt{AssetNumInt: 123},
		AssetStatus: asset.BasicAsset_CHANGED,
	}
	pb := proto.Message(asset1)
	for i := 0; i < b.N; i++ {
		_, err := GetFields(pb)
		if err != nil {
			fmt.Println(err)
			break
		}
	}
}

func Test_getFields_owner_success(t *testing.T) {
	fmt.Println("--- Test_getFields_owner_success start ---")
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
	fmt.Println("--- Test_getFields_owner_success end ---")
}

func Test_getFields_asset_success(t *testing.T) {
	fmt.Println("--- Test_getFields_asset_success start ---")
	asset1 := &asset.BasicAsset{
		AssetId:    1,
		AssetName:  "测试资产1",
		AssetPrice: 108,
		AssetOwner: &owner.Owner{
			OwnerId:   0,
			OwnerName: "张三",
		},
		AssetNum:    &asset.BasicAsset_AssetNumInt{AssetNumInt: 123},
		AssetStatus: asset.BasicAsset_CHANGED,
	}
	pb := proto.Message(asset1)

	got, err := GetFields(pb)
	if err != nil {
		fmt.Println(err)
	}
	if got != nil {
		for _, fd := range got {
			fmt.Printf("fd: %s\n", fd)
		}
	}
	fmt.Println("--- Test_getFields_asset_success end ---")
}
