package utils

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"github.com/go-playground/validator"
	"github.com/go-redis/redis"
	"github.com/segmentio/ksuid"
	"gitlab.yixinonline.org/miniTools/data-service/utils/cache"
	"gitlab.yixinonline.org/pkg/yrdLog"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

var TimeLocation, _ = time.LoadLocation("Asia/Shanghai") // 当地时间
var Layout = "2006-01-02 15:04-05"
// 获取当月第一天零点
func GetMonthZero() time.Time {
	monthStr := time.Now().Format("200601")
	t, _ := time.Parse("200601", monthStr)
	return t
}

// 获取下月第一天零点
func GetNextMonthZero() time.Time {
	t := time.Now()
	y, m, _ := t.Date()
	m += 1
	month := ""
	if m < 10 {
		month = "0" + strconv.Itoa(int(m))
	} else {
		month = strconv.Itoa(int(m))
	}
	t1, _ := time.Parse("200601", strconv.Itoa(y)+month)
	return t1
}

// 生成随机字符串
func GetRandomString(length int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < length; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

func GetNowDateStr(layout string) string {
	return time.Now().In(TimeLocation).Format(layout)
}

func Md5(s string) string {
	c := md5.New()
	c.Write([]byte(s))
	cipherStr := c.Sum(nil)
	return hex.EncodeToString(cipherStr)
}

func RundomString() string {
	return strings.ToUpper(hex.EncodeToString(ksuid.New().Payload()))
}

func randomInt(min, max int) int {
	return min + rand.Intn(max-min)
}

func GenerateCode(len int) string {
	rand.Seed(time.Now().UnixNano())
	bytes := make([]byte, len)
	for i := 0; i < len; i++ {
		bytes[i] = byte(randomInt(48, 57)) // 65 - 90 A-Z
	}
	return string(bytes)
}

// x/y向上取整
func Ceil(x int64, y int64) int64 {
	n := x / y
	m := x % y
	if m != 0 {
		n++
	}
	return n
}
func Substr(str string, start, length int) string {
	rs := []rune(str)
	rl := len(rs)
	end := 0

	if start < 0 {
		start = rl - 1 + start
	}
	end = start + length

	if start > end {
		start, end = end, start
	}

	if start < 0 {
		start = 0
	}
	if start > rl {
		start = rl
	}
	if end < 0 {
		end = 0
	}
	if end > rl {
		end = rl
	}

	return string(rs[start:end])
}

func ValidateParams(st interface{}, options ...Validation) (paramErr validator.FieldError, err error) {
	// 参数验证
	validate := validator.New()
	for i := 0; i < len(options); i++ {
		options[i](validate)
	}
	err = validate.Struct(st)
	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			return paramErr, err
		}
		for _, paramErr = range err.(validator.ValidationErrors) {
			return paramErr, nil
		}
	}
	return
}

type Validation func(*validator.Validate)

func RegValidation(stFunc func(validator.StructLevel), types interface{}) Validation {
	return func(v *validator.Validate) {
		v.RegisterStructValidation(stFunc, types)
	}
}

// 传入的res在json解析时是按照interface去解析的
func GetClusterCache(key string, t time.Duration, res interface{}, f func() (interface{}, error)) (err error) {
	rs, errs := cache.GetCache().Get(key).Result()
	if errs != nil && errs != redis.Nil {
		yrdLog.GetLogger().AlertWithLevel(yrdLog.ALERTWARNING, "redis缓存", "redis get 出错", "key", key, "错误信息", errs.Error())
	}

	if rs != "" {
		errs = json.Unmarshal([]byte(rs), res)
		if errs != nil {
			yrdLog.GetLogger().AlertWithLevel(yrdLog.ALERTWARNING, "redis缓存", "json 错误", "数据", rs, "错误信息", errs.Error())
		} else {
			return
		}
	}
	// 执行f方法获取数据
	res, err = f()
	if err != nil {
		return
	}

	b, errs := json.Marshal(res)
	if errs != nil {
		yrdLog.GetLogger().AlertWithLevel(yrdLog.ALERTWARNING, "redis缓存", "json 错误", "错误信息", errs.Error())
	} else {
		_, errs = cache.GetCache().Set(key, string(b), t).Result()
		if errs != nil {
			yrdLog.GetLogger().AlertWithLevel(yrdLog.ALERTWARNING, "redis缓存", "redis set 出错", "key", key, "data", string(b), "错误信息", errs.Error())
		}
		// yrdLog.GetLogger().Info("json数据>>>>>>", string(b))
	}

	return
}