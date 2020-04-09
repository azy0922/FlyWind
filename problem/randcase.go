package problem

import (
	"encoding/json"
	"math/rand"
	"time"
)

const (
	StringPool = "abcdefghijklmnopqrstuvwxyz1234567890"
)

type TestCase struct {
	Matcher string
	Target  string
}

type RandCase struct{}

// 把string切片转为interface类型，便于sprintf
func Convert(source []string) []interface{} {
	vals := make([]interface{}, len(source))
	for i, v := range source {
		vals[i] = v
	}
	return vals
}

// 组合返回函数，将若干返回的interface{}组合在一起进行返回
func (r *RandCase) BindReturn(args ...interface{}) []interface{} {
	if len(args) == 0 {
		return nil
	}
	ret := make([]interface{}, len(args))
	for i, v := range args {
		ret[i] = v
	}
	return ret
}

// 把参数列表中的整数切片转换为前端显示友好的字符串
// 形如：[]int{1,2,3,4,5} => "[1,2,3,4,5]"
func ConvertArgs(src []interface{}) []interface{} {
	for i, v := range src {
		switch v.(type) {
		// 只对整数切片进行转换
		case []int:
			src[i] = SliceIntToString(v.([]int))
		}
	}
	return src
}

// 生成[start, end]之间的随机整数用例
func (r *RandCase) GenRandInt(start, end int) interface{} {
	return rand.Intn(end-start+1) + start
}

// 将int数组转换为string类型
func SliceIntToString(n []int) string {
	ba, err := json.Marshal(n)
	if err != nil {
		return ""
	}
	return string(ba)
	//ret = append([]string{}, string(ba))
	//return ret
}

// 生成[start, end]之间的随机整数切片用例
// rand为true则以num为起始长度创建随机长度
func (r *RandCase) GenRandSliceInt(start, end, num int, randNum bool) interface{} {
	if randNum {
		num = rand.Intn(14-num+1) + num
	}
	if num == 0 {
		return nil
	}
	ca := make([]int, 0)
	for i := 0; i < num; i++ {
		ca = append(ca, rand.Intn(end-start+1)+start)
	}
	//return append(ret, ca)
	//fmt.Println(ca)
	return ca
}

// 从stringpool中随机选出pick个字符，组成length长度的字符串
// 并且要求匹配最后match长度的串
func (r *RandCase) GenRandSliceString(pick, length, match int) interface{} {
	var strCase []byte
	var strTarget []byte

	rand.Seed(time.Now().UnixNano())

	// 先从stringPool抽出随机选出8个字符
	for i := 0; i < pick; i++ {
		strCase = append(strCase, StringPool[rand.Intn(len(StringPool))])
	}

	// 生成targetString目标字符串，20字节长

	for i := 0; i < length; i++ {
		strTarget = append(strTarget, strCase[rand.Intn(len(strCase))])
	}
	// 取后6字节为待匹配字符串
	matcher := append([]byte{}, strTarget[len(strTarget)-match+1:]...)

	// 随机插入1个空格
	m := rand.Intn(len(strTarget)-1) + 1
	temp := append([]byte{}, strTarget[m:]...)
	strTarget = append(strTarget[:m], ' ')
	strTarget = append(strTarget, temp...)

	ret := TestCase{string(matcher), string(strTarget)}

	return ret
}
