zcutils-go
=====

golang常用工具类库

# protobuf
提供protobuf相关工具，例如:
- protoreflect 提供获取目标proto消息的字段信息的相关函数。

# zcargs
提供对命令行参数的获取与移除函数。

# zcbitmap
提供位图工具包，包括:
- BitSet8 8位的位图
- BitSet16 16位的位图
- BitSet32 32位的位图
- BitSet64 64位的位图

# zcpath
文件路径相关操作包

# zcrandom
随机数相关操作包

# zcslice
切片相关操作包

# zcstr
字符串相关操作包

# zcsync
同步函数执行工具包

# zctime
提供time相关处理

# zctoken
提供支持国密算法以及国际主流密码学算法的token生成与校验函数:
- `SM2-SM3` : 国密算法，使用SM2签名，使用SM3散列
- `ECDSA-SHA256` : 使用ecdsa签名，使用SHA256散列
- `ED25519-SHA256` : 使用ed25519签名，使用SHA256散列
- `HMAC-SM3` : 采用国密散列算法SM3的HMAC认证码算法
- `HMAC-SHA256` : 采用散列算法SHA256的HMAC认证码算法

# zcutil
其他通用处理函数

# JetBrains support
Thanks to JetBrains for supporting open source projects.

<a href="https://jb.gg/OpenSourceSupport" target="_blank">https://jb.gg/OpenSourceSupport.</a>
