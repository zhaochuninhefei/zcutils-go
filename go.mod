module gitee.com/zhaochuninhefei/zcutils-go

go 1.22

require (
	gitee.com/zhaochuninhefei/gmgo v0.1.1
	gitee.com/zhaochuninhefei/zcgolog v0.0.23
	github.com/fsnotify/fsnotify v1.7.0
	// 此处依赖废弃的`github.com/golang/protobuf`包是因为`protoreflect.GetFieldsByProperties`函数使用了`github.com/golang/protobuf/proto`的弃用函数`GetProperties`作为兼容老版本protobuf的功能。
	github.com/golang/protobuf v1.5.4
	github.com/nxadm/tail v1.4.11
	github.com/stretchr/testify v1.9.0
	golang.org/x/crypto v0.22.0
	google.golang.org/grpc v1.63.2
	google.golang.org/protobuf v1.33.0
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/kr/pretty v0.3.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	golang.org/x/net v0.24.0 // indirect
	golang.org/x/sys v0.19.0 // indirect
	golang.org/x/text v0.14.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240415180920-8c6c420018be // indirect
	gopkg.in/check.v1 v1.0.0-20201130134442-10cb98267c6c // indirect
	gopkg.in/tomb.v1 v1.0.0-20141024135613-dd632973f1e7 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
