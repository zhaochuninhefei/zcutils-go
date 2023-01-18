protobuf相关操作包
=====

# 功能列表

- GetFields 获取目标proto消息的字段信息
- GetFieldsByProperties 根据StructProperties获取proto消息字段信息，注意，该函数使用了`github.com/golang/protobuf/proto`的弃用函数`GetProperties`
- IsGeneratedAutoFields 判断字段是否是proto-gen-go自动生成的特殊字段

> GetFields与GetFieldsByProperties功能相同，不同之处在于GetFieldsByProperties使用了已经被弃用的函数`GetProperties`。
> 
> 目前建议使用`protobuf reflection`代替`GetProperties`，但在`google.golang.org/protobuf/proto`的`protobuf reflection`中并未提供与其功能类似的函数，
> 无法获取对应的Go字段名(只能获取到protobuf定义名与json定义名。。。)
> 
> 因此GetFields直接使用golang原生的反射实现了类似`GetProperties`的功能，但性能方面，比`GetProperties`慢大约50%。
