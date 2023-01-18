package protoreflect

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	protogh "github.com/golang/protobuf/proto"
	"google.golang.org/protobuf/proto"
)

// 自动生成字段名常量定义
//goland:noinspection GoSnakeCaseUsage,GoUnusedConst
const (
	// protoc-gen-go v1.20之前生成的"XXX_"前缀字段
	// 对应`github.com/golang/protobuf`为版本`1.4.2`之前

	GONAME_XXX_PREFIX         = "XXX_"
	GONAME_X_NOUNKEYEDLITERAL = "XXX_NoUnkeyedLiteral"
	GONAME_X_UNRECOGNIZED     = "XXX_unrecognized"
	GONAME_X_SIZECACHE        = "XXX_sizecache"

	// protoc-gen-go v1.20开始生成的特殊字段
	// 对应`github.com/golang/protobuf`为版本`1.4.2`以后
	// 对应`google.golang.org/protobuf`所有版本(`v1.20.0`为目前最早的版本)

	GONAME_N_STATE           = "state"
	GONAME_N_SIZECACHE       = "sizeCache"
	GONAME_N_UNKNOWNFIELDS   = "unknownFields"
	GONAME_N_EXTENSIONFIELDS = "extensionFields"
	GONAME_N_WEAKFIELDS      = "weakFields"
)

// FieldInfo 字段情报
type FieldInfo struct {
	// FieldName 字段名(go结构体中的字段名)
	FieldName string

	// ProtoName proto字段名(protobuf消息中定义的原始字段名)
	ProtoName string

	// jsonName json字段名(protobuf消息中定义的json字段名)
	JsonName string

	// FieldType 字段类型(go结构体中的字段类型)
	FieldType reflect.Type

	// FiledValue 字段类型(go结构体中的字段值)
	FiledValue reflect.Value

	// ExplicitDef 是否显式定义字段
	//  proto3会自动给消息添加state,sizeCache,unknownFields三个通用字段，它们就不是显式定义字段。
	ExplicitDef bool
}

// String 重写String方法
func (fd *FieldInfo) String() string {
	if fd == nil {
		return "nil"
	}
	return fmt.Sprintf("{FieldName: %s, ProtoName: %s, JsonName: %s, FieldType: %v, FiledValue: %v, ExplicitDef: %v}",
		fd.FieldName,
		fd.ProtoName,
		fd.JsonName,
		fd.FieldType,
		fd.FiledValue,
		fd.ExplicitDef)
}

// GetFields 获取目标proto消息的字段信息
//  @param msg protobuf消息
//  @return []*FieldInfo 字段信息列表
//  @return error
func GetFields(msg proto.Message) ([]*FieldInfo, error) {
	// 对msg进行反射获取reflect.Value
	vMsg := reflect.ValueOf(msg)
	// 判断reflect.Value的类型，这里必须是指针
	if vMsg.Kind() != reflect.Ptr {
		return nil, fmt.Errorf("err01 : msg is not a Pointer, but a %s", vMsg.Kind())
	}
	// 非空检查
	if vMsg.IsNil() {
		return nil, fmt.Errorf("err02 : msg is nil")
	}
	// 获取指针指向的数据结构
	vElem := vMsg.Elem()
	// 检查指针指向的数据结构的类型，必须是一个Struct结构体
	if vElem.Kind() != reflect.Struct {
		return nil, fmt.Errorf("err03 : msg's elem is not a Struct, but a %s", vElem.Kind())
	}
	// 获取msg消息对应的类型
	mElem := vElem.Type()
	// 获取该类型的字段个数
	fdNum := mElem.NumField()
	// 检查字段数量
	if fdNum == 0 {
		return nil, fmt.Errorf("err04 : msg has no fields")
	}

	fieldInfos := make([]*FieldInfo, 0, fdNum)

	// 遍历所有字段
	for i := 0; i < fdNum; i++ {
		f := mElem.Field(i)
		v := vElem.Field(i)
		fieldInfo := &FieldInfo{
			FieldName:  f.Name,
			FieldType:  f.Type,
			FiledValue: v,
		}
		//fmt.Printf("字段名: %s, 字段类型: %s, 字段标签: %s\n", f.Name, f.Type.Name(), f.Tag)
		tagProto := f.Tag.Get("protobuf")
		if tagProto != "" {
			fieldInfo.ExplicitDef = true
			//fmt.Printf("protobuf: %s\n", tagProto)
			tmp := strings.Split(tagProto, ",")
			for _, word := range tmp {
				if strings.HasPrefix(word, "name=") {
					fieldInfo.ProtoName = strings.TrimPrefix(word, "name=")
					//fmt.Println("对应proto消息的字段名:" + fieldInfo.ProtoName)
					continue
				}
				if strings.HasPrefix(word, "json=") {
					fieldInfo.JsonName = strings.TrimPrefix(word, "json=")
					//fmt.Println("对应json消息的字段名:" + fieldInfo.JsonName)
					continue
				}
			}
		} else {
			tagOneof := f.Tag.Get("protobuf_oneof")
			if tagOneof != "" {
				fieldInfo.ExplicitDef = true
				fieldInfo.ProtoName = tagOneof
			} else {
				fieldInfo.ExplicitDef = false
			}
		}
		fieldInfos = append(fieldInfos, fieldInfo)
	}
	return fieldInfos, nil
}

// GetFieldsByProperties 根据StructProperties获取proto消息字段信息
//  注意，该函数使用了`github.com/golang/protobuf/proto`的弃用函数`GetProperties`
//  @param msg protobuf消息
//  @return []*FieldInfo 字段信息列表
//  @return error
func GetFieldsByProperties(msg protogh.Message) ([]*FieldInfo, error) {
	// 对msg进行反射获取reflect.Value
	vMsg := reflect.ValueOf(msg)
	// 判断reflect.Value的类型，这里必须是指针
	if vMsg.Kind() != reflect.Ptr {
		return nil, fmt.Errorf("err01 : msg is not a Pointer, but a %s", vMsg.Kind())
	}
	// 非空检查
	if vMsg.IsNil() {
		return nil, fmt.Errorf("err02 : msg is nil")
	}
	// 获取指针指向的数据结构
	vElem := vMsg.Elem()
	// 检查指针指向的数据结构的类型，必须是一个Struct结构体
	if vElem.Kind() != reflect.Struct {
		return nil, fmt.Errorf("err03 : msg's elem is not a Struct, but a %s", vElem.Kind())
	}
	// 获取msg消息对应的类型
	mElem := vElem.Type()
	// 获取该类型的字段个数
	fdNum := mElem.NumField()
	// 检查字段数量
	if fdNum == 0 {
		return nil, fmt.Errorf("err04 : msg has no fields")
	}

	fieldInfos := make([]*FieldInfo, 0, fdNum)

	//goland:noinspection GoDeprecation
	protoProps := protogh.GetProperties(mElem)
	for _, prop := range protoProps.Prop {
		if strings.HasPrefix(prop.Name, "XXX_") {
			continue
		}
		fieldValue := vElem.FieldByName(prop.Name)
		fieldTypeStruct, ok := vElem.Type().FieldByName(prop.Name)
		if !ok {
			return nil, fmt.Errorf("programming error: proto does not have field advertised by proto package : %s", prop.Name)
		}
		fieldType := fieldTypeStruct.Type
		fieldInfo := &FieldInfo{
			FieldName:   prop.Name,
			FieldType:   fieldType,
			FiledValue:  fieldValue,
			ProtoName:   prop.OrigName,
			JsonName:    prop.JSONName,
			ExplicitDef: true,
		}
		fieldInfos = append(fieldInfos, fieldInfo)
	}
	return fieldInfos, nil
}

// IsGeneratedAutoFields 判断字段是否是proto-gen-go自动生成的特殊字段
//  @param fieldName 字段名
//  @return bool 是否是proto-gen-go自动生成的特殊字段
//  @return error
func IsGeneratedAutoFields(fieldName string) (bool, error) {
	if fieldName == "" {
		return false, errors.New("fieldName is empty")
	}
	if strings.HasPrefix(fieldName, GONAME_XXX_PREFIX) {
		return true, nil
	}
	if fieldName == GONAME_N_STATE ||
		fieldName == GONAME_N_SIZECACHE ||
		fieldName == GONAME_N_UNKNOWNFIELDS ||
		fieldName == GONAME_N_EXTENSIONFIELDS ||
		fieldName == GONAME_N_WEAKFIELDS {
		return true, nil
	}
	return false, nil
}
