zctoken相关工具包
==========

# BuildTokenWithECC / CheckTokenWithECC
使用椭圆曲线签名算法构建/校验凭证，支持以下算法:
- `SM2-SM3` : 国密算法，使用SM2签名，使用SM3散列
- `ECDSA-SHA256` : 使用ecdsa签名，使用SHA256散列
- `ED25519-SHA256` : 使用ed25519签名，使用SHA256散列

# BuildTokenWithGM / CheckTokenWithGM
使用`gitee.com/zhaochuninhefei/gmgo`的SM2与SM3算法实现国密token的创建与校验。
> 与使用`BuildTokenWithECC/CheckTokenWithECC`时采用`SM2-SM3`算法相同，只是入参出参不同。

# BuildTokenWithHMAC / CheckTokenWithHMAC
使用HMAC算法构建/校验凭证，支持以下算法:
- `HMAC-SM3` : 采用国密散列算法SM3的HMAC认证码算法
- `HMAC-SHA256` : 采用散列算法SHA256的HMAC认证码算法


# 性能测试结果
毫无疑问，采用HMAC算法的凭证构造与校验函数在性能上是有很大优势的。
> 但采用椭圆曲线签名算法来构造和校验凭证的话，在密码学安全性上更有优势。

以下是性能测试结果:
```
GOROOT=/usr/golang/go_1.17.5 #gosetup
GOPATH=/home/zhaochun/work/sources/go_path #gosetup
/usr/golang/go_1.17.5/bin/go test -c -o /tmp/GoLand/___gobench_gitee_com_zhaochuninhefei_zcutils_go_zctoken.test gitee.com/zhaochuninhefei/zcutils-go/zctoken #gosetup
/tmp/GoLand/___gobench_gitee_com_zhaochuninhefei_zcutils_go_zctoken.test -test.v -test.paniconexit0 -test.bench . -test.run ^$
goos: linux
goarch: amd64
pkg: gitee.com/zhaochuninhefei/zcutils-go/zctoken
cpu: 12th Gen Intel(R) Core(TM) i7-12700H
BenchmarkBuildTokenWithSM2SM3
BenchmarkBuildTokenWithSM2SM3-20        	   32943	     37864 ns/op	    7111 B/op	     142 allocs/op
BenchmarkBuildTokenWithECDSA
BenchmarkBuildTokenWithECDSA-20         	   37602	     32474 ns/op	    6257 B/op	     122 allocs/op
BenchmarkBuildTokenWithED25519
BenchmarkBuildTokenWithED25519-20       	   41622	     29124 ns/op	    2473 B/op	      58 allocs/op
BenchmarkBuildTokenWithHMACSM3
BenchmarkBuildTokenWithHMACSM3-20       	  546943	      2132 ns/op	    1793 B/op	      30 allocs/op
BenchmarkBuildTokenWithHMACSHA256
BenchmarkBuildTokenWithHMACSHA256-20    	  586698	      1913 ns/op	    1825 B/op	      28 allocs/op
BenchmarkCheckTokenWithSM2SM3
BenchmarkCheckTokenWithSM2SM3-20        	   20446	     58128 ns/op	    4620 B/op	     101 allocs/op
BenchmarkCheckTokenWithECDSA
BenchmarkCheckTokenWithECDSA-20         	   21111	     57371 ns/op	    4500 B/op	      91 allocs/op
BenchmarkCheckTokenWithED25519
BenchmarkCheckTokenWithED25519-20       	   32475	     36449 ns/op	    2264 B/op	      58 allocs/op
BenchmarkCheckTokenWithHMACSM3
BenchmarkCheckTokenWithHMACSM3-20       	  407744	      3128 ns/op	    1992 B/op	      38 allocs/op
BenchmarkCheckTokenWithHMACSHA256
BenchmarkCheckTokenWithHMACSHA256-20    	  404240	      3069 ns/op	    2008 B/op	      36 allocs/op
PASS

Process finished with the exit code 0
```
