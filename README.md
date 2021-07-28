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

* ### Snowflake -> 感谢「[snowflake](https://github.com/bwmarrin/snowflake)」

#### ID Format
By default, the ID format follows the original Twitter snowflake format.

* The ID as a whole is a 63 bit integer stored in an int64
* 41 bits are used to store a timestamp with millisecond precision, using a custom epoch.
* 10 bits are used to store a node id - a range from 0 through 1023.
* 12 bits are used to store a sequence number - a range from 0 through 4095.

#### How it Works.
Each time you generate an ID, it works, like this.

* A timestamp with millisecond precision is stored using 41 bits of the ID.
* Then the NodeID is added in subsequent bits.
* Then the Sequence Number is added, starting at 0 and incrementing for each ID generated in the same millisecond. If you generate enough IDs in the same millisecond that the sequence would roll over or overfill then the generate function will pause until the next millisecond.
  
The default Twitter format shown below.
```go
+--------------------------------------------------------------------------+
| 1 Bit Unused | 41 Bit Timestamp |  10 Bit NodeID  |   12 Bit Sequence ID |
+--------------------------------------------------------------------------+
```
Using the default settings, this allows for 4096 unique IDs to be generated every millisecond, per Node ID.

```go
// Create a new Node with a Node number of 1
node, err := snowflake.NewNode(1)
if err != nil {
fmt.Println(err)
return
}

// Generate a snowflake ID.
id := node.Generate()

```

* ### jwt -> 「[jwt-go](https://github.com/dgrijalva/jwt-go)」
```go
func GenerateToken(userId, username interface{}) (string, error) 生成一个token
func ParseToken(strToken string) (*JwtClaims, bool)  解析token
```