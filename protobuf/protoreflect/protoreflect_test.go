package protoreflect

import (
	"fmt"
	"testing"

	"gitee.com/zhaochuninhefei/zcutils-go/protobuf/myproto-go/asset"
	"gitee.com/zhaochuninhefei/zcutils-go/protobuf/myproto-go/owner"
	protogh "github.com/golang/protobuf/proto"
	"google.golang.org/protobuf/proto"
)

func BenchmarkGetFieldsByProperties(b *testing.B) {
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
	pb := protogh.Message(asset1)
	for i := 0; i < b.N; i++ {
		_, err := GetFieldsByProperties(pb)
		if err != nil {
			fmt.Println(err)
			break
		}
	}
}

func BenchmarkGetFields(b *testing.B) {
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

func TestGetFieldsOwnerSuccess(t *testing.T) {
	fmt.Println("--- TestGetFieldsOwnerSuccess start ---")
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
			isG, _ := IsGeneratedAutoFields(fd.FieldName)
			fmt.Printf("fd: %s , 是否自动生成: %v\n", fd, isG)
		}
	}
	fmt.Println("--- TestGetFieldsOwnerSuccess end ---")
}

func TestGetFieldsAssetSuccess(t *testing.T) {
	fmt.Println("--- TestGetFieldsAssetSuccess start ---")
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
	fmt.Println("--- TestGetFieldsAssetSuccess end ---")
}

func TestGetFieldsByProperties(t *testing.T) {
	fmt.Println("--- TestGetFieldsByProperties start ---")
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
	pb := protogh.Message(asset1)

	got, err := GetFieldsByProperties(pb)
	if err != nil {
		fmt.Println(err)
	}
	if got != nil {
		for _, fd := range got {
			fmt.Printf("fd: %s\n", fd)
		}
	}
	fmt.Println("--- TestGetFieldsByProperties end ---")
}

func TestGetFieldsByProtoReflect(t *testing.T) {
	fmt.Println("--- TestGetFieldsByProtoReflect start ---")
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

	m := pb.ProtoReflect()
	fds := m.Descriptor().Fields()
	for k := 0; k < fds.Len(); k++ {
		fd := fds.Get(k)
		fv := m.Get(fd)
		fmt.Printf("fieldName: %s, fieldValue: %s, fieldType: %s \n", fd.Name(), fv, fd.Kind())
	}
	fmt.Println("--- TestGetFieldsByProtoReflect end ---")
}
