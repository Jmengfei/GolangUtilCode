## About

GolangUtilCodeãð¥ æ¯ä¸ä¸ªå¼åä¸­å¸¸ç¨çå·¥å·ç±»ãéæå®æ´çååæµè¯åæ§è½æµè¯ï¼ä¸éè¦å¼ç¨å¤ä½çåºåæä»¶ï¼å¯ä»¥å¼ç®±å³ç¨ã



## APIs

* ### rand -> ã[randææ¡£](https://pkg.go.dev/math/rand@go1.16.6)ã

```go
func GetInt(max int) (int, error)  	çæä¸ä¸ªå å¯çå®å¨çéæºint
func GetIntInsecure(i int) int	ä½¿ç¨å½åç³»ç»æ¶é´çç§å­çæä¸ä¸ªéæºæ´æ°
func String(n int) (string, error)  çæå å¯å®å¨çå­ç¬¦ä¸²
func StringInsecure(n int) (string, error)	çæä¸ä¸ªå¯ç ä¸å®å¨çå­ç¬¦ä¸²
func StringRange(min int, max int) (string, error)	çæç»å®èå´åçå®å¨éæºå­ç¬¦ä¸²
func IntRange(min int, max int) (int, error)    è¿åç»å®èå´ä¹é´çéæºæ´æ°
func Random(n int, charset string, isSecure bool) (string, error) ä»ç»å®çå­ç¬¦éçæéæºæ°æ®
func Bytes(n int) ([]byte, error)   çæä¸ç»å å¯å®å¨çå­è
func Choice(j []string) (string, error)	ä»ä¸æ®µå­ç¬¦ä¸²ä¸­éæºéæ©
func ChoiceInsecure(j []string) string	ä»ä¸æ®µå­ç¬¦ä¸²ä¸­éæºéæ©(ä¸å®å¨)
```

* ### Snowflake -> æè°¢ã[snowflake](https://github.com/bwmarrin/snowflake)ã

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

* ### jwt -> ã[jwt-go](https://github.com/dgrijalva/jwt-go)ã
```go
func GenerateToken(userId, username interface{}) (string, error) çæä¸ä¸ªtoken
func ParseToken(strToken string) (*JwtClaims, bool)  è§£ætoken
```

* ### redis -> ã[redis-go](https://github.com/go-redis/redis)ã
```go

```

* ### compare -> ã[go-version](https://github.com/hashicorp/go-version)ã
```go
// æ ¸å¿ä»£ç 
func versionOrdinal(version string) string {
	// ISO/IEC 14651:2011
	const maxByte = 1<<8 - 1

	vo := make([]byte, 0, len(version)+8)
	j := -1
	for i := 0; i < len(version); i++ {
		b := version[i]
		if '0' > b || b > '9' {
			vo = append(vo, b)
			j = -1
			continue
		}
		if j == -1 {
			vo = append(vo, 0x00)
			j = len(vo) - 1
		}
		if vo[j] == 1 && vo[j+1] == '0' {
			vo[j+1] = b
			continue
		}
		if vo[j]+1 > maxByte {
			panic("VersionOrdinal: invalid version")
		}
		vo = append(vo, b)
		vo[j]++
	}
	return string(vo)
}
```

### ã[isoç¸å³](https://github.com/Jmengfei/GolangUtilCode/blob/main/iso/ISO.md)ã -> ã[ISO](https://github.com/lukes/ISO-3166-Countries-with-Regional-Codes)ã

ã[ISOææ¡£](ISO-3166 Country and Dependent Territories Lists with UN Regional Codes)ã

