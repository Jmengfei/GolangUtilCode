package redis

import (
	"encoding/json"
	"errors"
	"github.com/garyburd/redigo/redis"
	"reflect"
	"strings"
	"time"
)

/**
 * @Author: mf
 * @email: 18539271635@163.com
 * @Date: 2021/8/2 12:00 下午
 * @Desc:
 */
var (
	ErrDataNil   = errors.New("data is nil")
)

type ToolRedis struct {
	redisClient *redis.Pool
}

func (c *ToolRedis) Connect(address, password string, db, maxIdle, maxActive, idleTimeout int) error {
	c.redisClient = &redis.Pool{
		MaxIdle:     maxIdle,
		MaxActive:   maxActive,
		IdleTimeout: time.Duration(idleTimeout) * time.Second,
		Wait:        true,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", address)
			if err != nil {
				return nil, err
			}
			if password != "" {
				if _, err := c.Do("AUTH", password); err != nil {
					_ = c.Close()
					return nil, err
				}
			}
			_, _ = c.Do("SELECT", db)
			return c, nil
		},
	}
	// 检查连接池是否可用
	conn := c.redisClient.Get()
	defer conn.Close()
	_, err := conn.Do("ping")
	if err != nil {
		return err
	}
	return nil
}

func (c *ToolRedis) Close() error {
	if err := c.redisClient.Close(); err != nil {
		return err
	}
	return nil
}

//得到redis 池
func (c *ToolRedis) GetPool() *redis.Pool {
	return c.redisClient
}

//得到一个客户端
func (c *ToolRedis) GetCli() redis.Conn {
	return c.redisClient.Get()
}

func (c *ToolRedis) Do(commandName string, args ...interface{}) (reply interface{}, err error) {
	cli := c.redisClient.Get()
	defer cli.Close()
	reply, err = cli.Do(commandName, args...)
	if err != nil {
		return nil, err
	}
	return
}

// ZREVRange 返回有序集中指定分数区间内的成员，分数从高到低排序
func (c *ToolRedis) ZREVRange(key string, start int, limit int) (reply interface{}, err error) {
	cli := c.redisClient.Get()
	defer cli.Close()
	reply, err = redis.Strings(cli.Do("ZREVRANGE", key, start, start+limit-1))
	if err != nil {
		return nil, err
	}
	return
}

// ZRange 通过索引区间返回有序集合指定区间内的成员
func (c *ToolRedis) ZRange(key string, start int, limit int) (reply interface{}, err error) {
	cli := c.redisClient.Get()
	defer cli.Close()
	reply, err = redis.Strings(cli.Do("ZRANGE", key, start, start+limit-1))
	if err != nil {
		return nil, err
	}
	return
}

// ZScore 返回有序集合，成员的分数值
func (c *ToolRedis) ZScore(key string, defVal int64) (reply interface{}, err error) {
	cli := c.redisClient.Get()
	defer cli.Close()
	reply, err = cli.Do("ZSCORE", key, defVal)
	if err != nil {
		return nil, err
	}
	return
}

// ZRangeWithScore 通过索引区间返回有序集合指定区间内的成员
func (c *ToolRedis) ZRangeWithScore(key string, start int, limit int) (reply interface{}, err error) {
	cli := c.redisClient.Get()
	defer cli.Close()
	reply, err = redis.Strings(cli.Do("ZRANGE", key, start, start+limit-1, "WITHSCORES"))
	if err != nil {
		return nil, err
	}
	return
}

// ZRem 移除有序集合中的一个或多个元素
func (c *ToolRedis) ZRem(key string, defVal int64) error {
	cli := c.redisClient.Get()
	defer cli.Close()
	_, err := cli.Do("ZREM", key, defVal)
	if err != nil {
		return err
	}
	return nil
}

// ZAdd 向有序集合添加一个或多个成员，或者更新已存在成员的分数
func (c *ToolRedis) ZAdd(key string, score int, defVal int64) error {
	cli := c.redisClient.Get()
	defer cli.Close()
	_, err := cli.Do("ZADD", key, score, defVal)
	if err != nil {
		return err
	}
	return nil
}

// ZRangeByScore 通过分数返回有序集合指定区间内的成员，顺序从小到大排列
// +inf 显示右侧搜有的值
// -inf 显示左侧所有的值
func (c *ToolRedis) ZRangeByScore(key string) ([][]byte, error) {
	cli := c.redisClient.Get()
	defer cli.Close()
	score, err := redis.Values(cli.Do("ZRANGEBYSCORE", key, "+inf", "-inf"))
	if err != nil {
		return nil, err
	}
	items := make([][]byte, 0, len(score))
	for _, v := range score {
		p, ok := v.([]byte)
		if !ok {
			break
		}
		items = append(items, p)
	}
	return items, nil
}

// ZREVRangeByScore 通过分数返回有序集合指定区间内的成员，顺序从大到小排列
// +inf 显示右侧搜有的值
// -inf 显示左侧所有的值
func (c *ToolRedis) ZREVRangeByScore(key string) ([][]byte, error) {
	cli := c.redisClient.Get()
	defer cli.Close()
	score, err := redis.Values(cli.Do("ZREVRANGEBYSCORE", key, "+inf", "-inf"))
	if err != nil {
		return nil, err
	}
	items := make([][]byte, 0, len(score))
	for _, v := range score {
		p, ok := v.([]byte)
		if !ok {
			break
		}
		items = append(items, p)
	}
	return items, nil
}

// Incr 将 key 中储存的数字值增一。如果 key 不存在，那么 key 的值会先被初始化为 0 ，然后再执行 INCR 操作。取值范围 int64
func (c *ToolRedis) Incr(key string) error {
	cli := c.redisClient.Get()
	defer cli.Close()
	_, err := cli.Do("INCR", key)
	if err != nil {
		return err
	}
	return nil
}

// Decr 将 key 中储存的数字值减一。如果 key 不存在，那么 key 的值会先被初始化为 0 ，然后再执行 DECR 操作。取值范围 int64
func (c *ToolRedis) Decr(key string) error {
	cli := c.redisClient.Get()
	defer cli.Close()
	_, err := cli.Do("DECR", key)
	if err != nil {
		return err
	}
	return nil
}

// HIncr 命令用于为哈希表中的字段值加上指定增量值.
// 如果哈希表的 key 不存在，一个新的哈希表被创建并执行 HINCRBY 命令.
// 如果指定的字段不存在，那么在执行命令前，字段的值被初始化为 0 。
// 取值范围 int64
func (c *ToolRedis) HIncr(key string, field string) error {
	cli := c.redisClient.Get()
	defer cli.Close()
	_, err := cli.Do("HINCRBY", key, field, 1)
	if err != nil {
		return err
	}
	return nil
}

// HDecr 命令用于为哈希表中的字段值加上指定增量值.
// 如果哈希表的 key 不存在，一个新的哈希表被创建并执行 HINCRBY 命令.
// 如果指定的字段不存在，那么在执行命令前，字段的值被初始化为 0 。
// 取值范围 int64
func (c *ToolRedis) HDecr(key string, feild string) error {
	cli := c.redisClient.Get()
	defer cli.Close()
	_, err := cli.Do("HINCRBY", key, feild, -1)
	if err != nil {
		return err
	}
	return nil
}

// HIDCount 命令用于为哈希表中的字段值加上指定增量值.
// 如果哈希表的 key 不存在，一个新的哈希表被创建并执行 HINCRBY 命令.
// 如果指定的字段不存在，那么在执行命令前，字段的值被初始化为 0 。
// 取值范围 int64
func (c *ToolRedis) HIDCount(key string, field string, count int) error {
	cli := c.redisClient.Get()
	defer cli.Close()
	_, err := cli.Do("HINCRBY", key, field, count)
	if err != nil {
		return err
	}
	return nil
}

// Set 设置值, 设置expire 可以设置过期时间
func (c *ToolRedis) Set(key string, value interface{}, expire ...int) error {
	cli := c.redisClient.Get()
	defer cli.Close()
	var err error
	_, err = cli.Do("SET", key, value)
	if err == nil && len(expire) > 0 {
		_, err = cli.Do("EXPIRE", key, expire[0])
	}
	if err != nil {
		return err
	}
	return nil
}

// MSet 用于同时设置一个或多个 key-value 对
func (c *ToolRedis) MSet(kvMap map[string]interface{}, expire ...int) error {
	cli := c.redisClient.Get()
	defer cli.Close()
	var err error
	args := make([]interface{}, 0, len(kvMap)*2)
	for k, v := range kvMap {
		args = append(args, k, v)
	}
	_, err = cli.Do("MSET", args...)
	if err == nil && len(expire) > 0 {
		go func() {
			for key := range kvMap {
				_, err = cli.Do("EXPIRE", key, expire[0])
			}
		}()
	}
	if err != nil {
		return err
	}
	return nil
}

// Get 根据key得到值
func (c *ToolRedis) Get(key string) (interface{}, error) {
	cli := c.redisClient.Get()
	defer cli.Close()
	val, err := cli.Do("GET", key)
	return val, err
}
// MGet 用于同时获取一个或多个 key-value 对
func (c *ToolRedis) MGet(keys ...string) ([][]byte, error) {
	cli := c.redisClient.Get()
	defer cli.Close()
	args := make([]interface{}, 0, len(keys))
	for _, key := range keys {
		args = append(args, key)
	}
	values, err := redis.Values(cli.Do("MGET", args...))
	if err != nil {
		return nil, err
	}
	ret := make([][]byte, 0, len(values))
	for _, v := range values {
		p := v.([]byte)
		ret = append(ret, p)
	}
	return ret, nil
}

// AsInt 得到 int 值, 如果出错得到默认值
func (c *ToolRedis) AsInt(key string, defVal int) (int, error) {
	value, err := redis.Int(c.Get(key))
	if err != nil {
		return defVal, err
	}
	return value, nil
}

// AsInt64 得到int64值，如果出错得到默认值
func (c *ToolRedis) AsInt64(key string, defVal int64) (int64, error) {
	value, err := redis.Int64(c.Get(key))
	if err != nil {
		return defVal, err
	}
	return value, nil
}

// AsBytesAndExist 得到 int64 值，如果不存在，会报错 ErrDataNil
func (c *ToolRedis) AsInt64AndExist(key string, defVal int64) (int64, error) {
	bValue, err := c.Get(key)
	if err != nil {
		return defVal, err
	}
	if bValue == nil {
		return defVal, ErrDataNil
	}
	value, err := redis.Int64(bValue, err)
	if err != nil {
		return defVal, err
	}
	return value, nil
}

// AsString 得到 string 值, 如果出错得到默认值
func (c *ToolRedis) AsString(key string, defVal string) (string, error) {
	value, err := redis.String(c.Get(key))
	if err != nil {
		return defVal, err
	}
	return value, nil
}

// AsBytesAndExist 得到 []byte 值，如果不存在，会报错 ErrDataNil
func (c *ToolRedis) AsBytesAndExist(key string, defVal []byte) ([]byte, error) {
	bValue, err := c.Get(key)
	if err != nil {
		return defVal, err
	}
	if bValue == nil {
		return defVal, ErrDataNil
	}
	value, err := redis.Bytes(bValue, err)
	if err != nil {
		return defVal, err
	}
	return value, nil
}

// AsBytes 得到 []byte 值
func (c *ToolRedis) AsBytes(key string, defVal []byte) ([]byte, error) {
	value, err := redis.Bytes(c.Get(key))
	if err != nil {
		return defVal, err
	}
	return value, nil
}

// Del 删除key
func (c *ToolRedis) Del(key string) error {
	cli := c.redisClient.Get()
	defer cli.Close()
	_, err := cli.Do("DEL", key)
	return err
}

// IsExist 判断key是否存在
func (c *ToolRedis) IsExist(key string) (bool, error) {
	cli := c.redisClient.Get()
	defer cli.Close()
	isExist, err := redis.Bool(cli.Do("EXISTS", key))
	if err != nil {
		isExist = false
	}
	return isExist, err
}

// Append 追加字符串， 如果不存在直接创建
func (c *ToolRedis) Append(key string, value string) error {
	if value == "" {
		return nil
	}
	cli := c.redisClient.Get()
	defer cli.Close()
	_, err := cli.Do("APPEND", key, value)
	return err
}

// Expire 给某个key设置过期时间
func (c *ToolRedis) Expire(key string, expire int) error {
	cli := c.redisClient.Get()
	defer cli.Close()
	_, err := redis.Bool(cli.Do("EXPIRE", key, expire))
	return err
}

// SetObj 通过json序列化对象到redis
func (c *ToolRedis) SetObj(key string, obj interface{}, expire ...int) error {
	value, err := json.Marshal(obj)
	if err != nil {
		return err
	}
	return c.Set(key, value, expire...)
}

// AsObj 解析json序列化的数据
func (c *ToolRedis) AsObj(key string, obj interface{}) error {
	value, err := c.AsBytes(key, nil)
	if err != nil {
		return err
	}
	err = json.Unmarshal(value, &obj)
	return err
}

/*
返回数据类型
返回值：
	error (表示出错)
	none (key不存在)
	string (字符串)
	list (列表)
	set (集合)
	zset (有序集)
	hash (哈希表)
*/
func (c *ToolRedis) AsType(key string) string {
	cli := c.redisClient.Get()
	defer cli.Close()
	value, err := redis.String(cli.Do("TYPE", key))
	if err != nil {
		return "error"
	}
	return value
}

// HMSet 命令用于同时将多个 field-value (字段-值)对设置到哈希表中
// 此命令会覆盖哈希表中已存在的字段。
// 如果哈希表不存在，会创建一个空哈希表，并执行 HMSET 操作。
func (c *ToolRedis) HMSet(key string, hMap map[string]interface{}) error {
	cli := c.redisClient.Get()
	defer cli.Close()
	_, err := cli.Do("HMSET", redis.Args{}.Add(key).AddFlat(hMap)...)
	return err
}

// HMSet 命令用于同时将多个 field-value (字段-值)对设置到哈希表中
// 设置结构体到hash
func (c *ToolRedis) HMSetStruct(key string, value interface{}) error {
	cli := c.redisClient.Get()
	defer cli.Close()
	_, err := cli.Do("HMSET", redis.Args{}.Add(key).AddFlat(value)...)
	return err
}

// HSet 在hash中设置值
func (c *ToolRedis) HSet(key, filed string, value interface{}) error {
	cli := c.redisClient.Get()
	defer cli.Close()
	valStr := parseToString(value)
	_, err := cli.Do("HSET", key, filed, valStr)
	if err != nil {
		return err
	}
	return nil
}

// HGet 获取hash值 返回 interface{}
func (c *ToolRedis) HGet(key, field string) (interface{}, error) {
	cli := c.redisClient.Get()
	defer cli.Close()
	val, err := cli.Do("HGET", key, field)
	return val, err
}

// AsHByte 获取hash值， 返回 []byte
func (c *ToolRedis) AsHByte(key, field string) ([]byte, error) {
	val, err := redis.Bytes(c.HGet(key, field))
	return val, err
}

// AsHString 获取hash的 string 值
func (c *ToolRedis) AsHString(key, filed, defVal string) (string, error) {
	val, err := redis.String(c.HGet(key, filed))
	if err != nil {
		return defVal, err
	}
	return val, nil
}

// AsHInt 获取hash的int值
func (c *ToolRedis) AsHInt(key, filed string, defVal int) (int, error) {
	val, err := redis.Int(c.HGet(key, filed))
	if err != nil {
		return defVal, err
	}
	return val, nil
}

// AsHInt64 获取hash的int64值
func (c *ToolRedis) AsHInt64(key, filed string, defVal int64) (int64, error) {
	val, err := redis.Int64(c.HGet(key, filed))
	if err != nil {
		return defVal, err
	}
	return val, nil
}

// AsHFloat 获取hash的浮点值
func (c *ToolRedis) AsHFloat(key, filed string, defVal float64) (float64, error) {
	val, err := redis.Float64(c.HGet(key, filed))
	if err != nil {
		return defVal, err
	}
	return val, nil
}

// AsHStruct 获取hash插入的结构体的数据
func (c *ToolRedis) AsHStruct(key string, out interface{}) error {
	cli := c.redisClient.Get()
	defer cli.Close()
	val, err := redis.Values(cli.Do("HGETALL", key))
	if err != nil {
		return err
	}
	if err = redis.ScanStruct(val, out); err != nil {
		return err
	}
	return nil
}

// HGetAll 获取hash的数据，返回[][]byte
func (c *ToolRedis) HGetAll(key string) ([][]byte, error) {
	cli := c.redisClient.Get()
	defer cli.Close()
	bVal, err := redis.ByteSlices(cli.Do("HGETALL", key))
	if err != nil {
		return nil, err
	}
	if bVal == nil {
		return nil, ErrDataNil
	}
	return bVal, nil
}

// HDel 用于删除哈希表 key 中的一个或多个指定字段，不存在的字段将被忽略
func (c *ToolRedis) HDel(key string, fields ...string) error {
	cli := c.redisClient.Get()
	defer cli.Close()
	_, err := cli.Do("HDEL", redis.Args{}.Add(key).AddFlat(fields)...)
	if err != nil {
		if c.IsNil(err) {
			return nil
		}
		return err
	}
	return nil
}

// SisMember set 集合中判断成员元素是否是集合的成员
func (c *ToolRedis) SisMember(key string, field string) (bool, error) {
	cli := c.redisClient.Get()
	defer cli.Close()
	isExist, err := redis.Int(cli.Do("SISMEMBER", key, field))
	return isExist == 1, err
}

// HExists 查看哈希表的指定字段是否存在
func (c *ToolRedis) HExists(key string, field string) (bool, error) {
	cli := c.redisClient.Get()
	defer cli.Close()
	isExist, err := redis.Bool(cli.Do("HEXISTS", key, field))
	if err != nil {
		isExist = false
	}
	return isExist, err
}

// SMembers 返回集合中的所有的成员。 不存在的集合 key 被视为空集合
func (c *ToolRedis) SMembers(key string) ([][]byte, error) {
	cli := c.redisClient.Get()
	defer cli.Close()
	vals, err := redis.ByteSlices(cli.Do("SMEMBERS", key))
	if err != nil {
		return nil, err
	}
	return vals, nil
}

// SScan 迭代集合中键的元素
func (c *ToolRedis) SScan(key string, cursor uint64, pattern string, count int) (uint64, []string, error) {
	cli := c.redisClient.Get()
	defer cli.Close()
	args := packArgs(key, cursor)
	if pattern != "" {
		args = append(args, "MATCH", pattern)
	}
	if count > 0 {
		args = append(args, "COUNT", count)
	}
	values, err := redis.Values(cli.Do("SSCAN", args...))
	if err != nil {
		return 0, nil, err
	}
	var items []string
	_, err = redis.Scan(values, &cursor, &items)
	if err != nil {
		return 0, nil, err
	}
	return cursor, items, nil
}

// SetLPos list 在集合某个位置添加元素
// 通过索引来设置元素的值
// 当索引参数超出范围，或对一个空列表进行 LSET 时，返回一个错误
func (c *ToolRedis) SetLPos(key string, pos int, value interface{}) error {
	cli := c.redisClient.Get()
	defer cli.Close()
	_, err := cli.Do("LSET", key, pos, value)
	return err
}

// LPush list 在集合头部添加
func (c *ToolRedis) LPush(key string, value interface{}) error {
	cli := c.redisClient.Get()
	defer cli.Close()
	_, err := cli.Do("LPUSH", key, value)
	return err
}

// RPush list 在集合尾部添加
func (c *ToolRedis) RPush(key string, value interface{}) error {
	cli := c.redisClient.Get()
	defer cli.Close()
	_, err := cli.Do("RPUSH", key, value)
	return err
}

// RPushStr 将一个或多个值插入到列表的尾部(最右边)
func (c *ToolRedis) RPushStr(key string, members ...string) (int, error) {
	cli := c.redisClient.Get()
	defer cli.Close()
	return redis.Int(cli.Do("RPUSH", redis.Args{}.Add(key).AddFlat(members)...))
}

// SRandMemberCount 返回集合中的一个随机元素
// 如果 count 为正数，且小于集合基数，那么命令返回一个包含 count 个元素的数组，数组中的元素各不相同。如果 count 大于等于集合基数，那么返回整个集合。
// 如果 count 为负数，那么命令返回一个数组，数组中的元素可能会重复出现多次，而数组的长度为 count 的绝对值。
// 该操作和 SPOP 相似，但 SPOP 将随机元素从集合中移除并返回，而 Srandmember 则仅仅返回随机元素，而不对集合进行任何改动
func (c *ToolRedis) SRandMemberCount(key string, count int) ([][]byte, error) {
	cli := c.redisClient.Get()
	defer cli.Close()
	res, err := redis.ByteSlices(cli.Do("SRANDMEMBER", key, count))
	if err != nil {
		return nil, err
	}
	return res, nil
}

// SRandMember 返回一个随机元素；如果集合为空，返回 nil
func (c *ToolRedis) SRandMember(key string) ([]byte, error) {
	cli := c.redisClient.Get()
	defer cli.Close()
	res, err := redis.Bytes(cli.Do("SRANDMEMBER", key))
	if err != nil {
		return nil, err
	}
	return res, nil
}

// SAdd 将一个或多个成员元素加入到集合中，已经存在于集合的成员元素将被忽略
func (c *ToolRedis) SAdd(key string, members ...string) error {
	cli := c.redisClient.Get()
	defer cli.Close()
	args := packArgs(key, members)
	_, err := cli.Do("SADD", args...)
	return err
}

// SMove 指定成员 member 元素从 source 集合移动到 destination 集合。
// SMove 是原子性操作。删除source，添加到destination集合中
func (c *ToolRedis) SMove(srcKey, destKey string, member string) error {
	cli := c.redisClient.Get()
	defer cli.Close()
	_, err := cli.Do("SMOVE", srcKey, destKey, member)
	return err
}

// LTrim 对一个列表进行修剪(trim)，就是说，让列表只保留指定区间内的元素，不在指定区间之内的元素都将被删除。
func (c *ToolRedis) LTrim(key string, start int, stop int) (bool, error) {
	cli := c.redisClient.Get()
	defer cli.Close()
	if res, err := redis.String(cli.Do("LTRIM", key, start, stop)); err == redis.ErrNil {
		return false, nil
	} else if err != nil || strings.ToLower(res) != "ok" {
		return false, err
	} else {
		return true, nil
	}
}

// SRem 移除集合中的一个或多个成员元素，不存在的成员元素会被忽略
func (c *ToolRedis) SRem(key string, members ...string) error {
	cli := c.redisClient.Get()
	defer cli.Close()
	args := packArgs(key, members)
	_, err := cli.Do("SREM", args...)
	return err
}

// SCard 返回set集合中元素的数量
func (c *ToolRedis) SCard(key string) (int, error) {
	cli := c.redisClient.Get()
	defer cli.Close()
	val, err := redis.Int(cli.Do("SCARD", key))
	if err != nil {
		return 0, err
	}
	return val, nil
}

// LLen 获取list集合长度
func (c *ToolRedis) LLen(key string) (int, error) {
	cli := c.redisClient.Get()
	defer cli.Close()
	num, err := redis.Int(cli.Do("llen", key))
	if err != nil {
		return 0, err
	}
	return num, nil
}

// LDataByIndex 通过索引获取列表中的元素。
// 你也可以使用负数下标，以 -1 表示列表的最后一个元素， -2 表示列表的倒数第二个元素，以此类推。
func (c *ToolRedis) LDataByIndex(key string, index int) ([]byte, error) {
	cli := c.redisClient.Get()
	defer cli.Close()
	value, err := redis.Bytes(cli.Do("lindex", key, index))
	if err != nil {
		return nil, err
	}
	return value, nil
}

// LDataByRang 获取，bg到ed范围的元素
func (c *ToolRedis) LDataByRang(key string, bg, ed int) ([][]byte, error) {
	cli := c.redisClient.Get()
	defer cli.Close()
	values, err := redis.ByteSlices(cli.Do("lrange", key, bg, ed))
	if err != nil {
		if c.IsNil(err) {
			return nil, nil
		}
		return nil, err
	}
	return values, nil
}

// LRem 移除列表中与参数 VALUE 相等的元素
// count > 0 : 从表头开始向表尾搜索，移除与 VALUE 相等的元素，数量为 COUNT
// count < 0 : 从表尾开始向表头搜索，移除与 VALUE 相等的元素，数量为 COUNT 的绝对值
// count = 0 : 移除表中所有与 VALUE 相等的值
func (c *ToolRedis) LRem(key string, value interface{}, count int) error {
	cli := c.redisClient.Get()
	defer cli.Close()
	_, err := cli.Do("LREM", key, count, value)
	return err
}

// LPop 删除头部的元素，并且返回头部的值
func (c *ToolRedis) LPop(key string) ([]byte, error) {
	cli := c.redisClient.Get()
	defer cli.Close()
	value, err := redis.Bytes(cli.Do("LPOP", key))
	return value, err
}

// RPop 删除尾部的元素，并且返回尾部的值
func (c *ToolRedis) RPop(key string) ([]byte, error) {
	cli := c.redisClient.Get()
	defer cli.Close()
	value, err := redis.Bytes(cli.Do("RPOP", key))
	return value, err
}

// TryLock (SET if Not eXists) 在指定的 key 不存在时，为 key 设置指定的值
func (c *ToolRedis) TryLock(key string, val string, expire uint64) (bool, error) {
	cli := c.redisClient.Get()
	defer cli.Close()
	res, err := redis.String(cli.Do("SET", key, val, "PX", expire, "NX"))
	if err != nil {
		return false, err
	} else if strings.ToLower(res) != "ok" {
		return false, nil
	}
	return true, nil
}

// TryUnlock 解除锁
func (c *ToolRedis) TryUnlock(key string) (bool, error) {
	cli := c.redisClient.Get()
	defer cli.Close()
	res, err := redis.Int(cli.Do("DEL", key))
	if err != nil {
		return false, err
	} else if res <= 0 {
		return false, nil
	}
	return true, nil
}

// IsNil 判断是否是空
func (c *ToolRedis) IsNil(err error) bool {
	if err == redis.ErrNil {
		return true
	}
	return false
}

func parseToString(value interface{}) string {
	switch v := value.(type) {
	case string:
		return v
	case []byte:
		return string(v)
	default:
		b, e := json.Marshal(v)
		if e != nil {
			return ""
		}
		return string(b)
	}
}

func packArgs(items ...interface{}) (args []interface{}) {
	for _, item := range items {
		v := reflect.ValueOf(item)
		switch v.Kind() {
		case reflect.Slice:
			if v.IsNil() {
				continue
			}
			for i := 0; i < v.Len(); i++ {
				args = append(args, v.Index(i).Interface())
			}
		case reflect.Map:
			if v.IsNil() {
				continue
			}
			for _, key := range v.MapKeys() {
				args = append(args, key.Interface(), v.MapIndex(key).Interface())
			}
		default:
			args = append(args, v.Interface())
		}
	}
	return args
}