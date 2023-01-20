zcutils-go
=====

golang常用工具类库

# protobuf
提供protobuf相关工具，例如:
- protoreflect 提供获取目标proto消息的字段信息的相关函数。

# zctoken
提供支持国密算法以及国际主流密码学算法的token生成与校验函数:
- `SM2-SM3` : 国密算法，使用SM2签名，使用SM3散列
- `ECDSA-SHA256` : 使用ecdsa签名，使用SHA256散列
- `ED25519-SHA256` : 使用ed25519签名，使用SHA256散列
- `HMAC-SM3` : 采用国密散列算法SM3的HMAC认证码算法
- `HMAC-SHA256` : 采用散列算法SHA256的HMAC认证码算法

# zctime
提供time相关常量定义。

# zcpath
提供本地目录与文件的相关处理。
