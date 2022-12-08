# celt-go
 golang wrapper for celt codec

# supported for libcelt 0.7.1

> 编译 libcelt： https://github.com/scjtqs2/libcelt
> 
# 使用说明参考 example

# 编译说明 demo
> 用到了cgo。因此需要自行安装gcc。
```shell
# 静态编译
CGO_ENABLED="1"  go build -ldflags "-extldflags=-static -s" -o test ./example/encode/
# 非静态编译
CGO_ENABLED="1"  go build  -o test ./example/encode/
```