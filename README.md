## About

GolangUtilCode」🔥 是一个开发中常用的工具类。配有完整的单元测试和性能测试，不需要引用多余的库和文件，可以开箱即用。



## APIs

* ### rand -> 「[rand文档](https://pkg.go.dev/math/rand@go1.16.6)」

```go
func GetInt(max int) (int, error)  	生成一个加密的安全的随机int
func GetIntInsecure(i int) int	使用当前系统时间的种子生成一个随机整数
func String(n int) (string, error)  生成加密安全的字符串
func StringInsecure(n int) (string, error)	生成一个密码不安全的字符串
func StringRange(min int, max int) (string, error)	生成给定范围内的安全随机字符串
func IntRange(min int, max int) (int, error)    返回给定范围之间的随机整数
func Random(n int, charset string, isSecure bool) (string, error) 从给定的字符集生成随机数据
func Bytes(n int) ([]byte, error)   生成一组加密安全的字节
func Choice(j []string) (string, error)	从一段字符串中随机选择
func ChoiceInsecure(j []string) string	从一段字符串中随机选择(不安全)
```

