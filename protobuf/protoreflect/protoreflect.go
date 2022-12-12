package protoreflect

import (
	"fmt"
	"reflect"
	"strings"

	"google.golang.org/protobuf/proto"
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
			fieldInfo.ExplicitDef = false
		}
		fieldInfos = append(fieldInfos, fieldInfo)
	}
	return fieldInfos, nil
}
