zctoken相关工具包
==========

# BuildToken与CheckToken
创建凭证与校验凭证，支持以下算法:
- `SM2-SM3` : 国密算法，使用SM2签名，使用SM3散列
- `ECDSA-SHA256` : 使用ecdsa签名，使用SHA256散列
- `ED25519-SHA256` : 使用ed25519签名，使用SHA256散列

# BuildTokenWithGM与CheckTokenWithGM
使用`gitee.com/zhaochuninhefei/gmgo`的SM2与SM3算法实现国密token的创建与校验。
> 与使用BuildToken与CheckToken采用`SM2-SM3`算法相同。
