package base

import (
	"encoding/binary"
	"math"
	"strconv"
)

//int to str
//strconv.Itoa(i)
//strconv.FormatInt(int64(i),10)

//str to int
//strconv.Atoi(s)
//strconv.ParseInt(s,10,0)

//bytes to float64
func bytesToFloat64(bytes []byte) float64 {
	bits := binary.LittleEndian.Uint64(bytes)
	return math.Float64frombits(bits)
}

//float64 to bytes; []uint8
func float64ToBytes(input float64) []byte {
	bits := math.Float64bits(input)
	bytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(bytes, bits)
	return bytes
}

//float to str 支持指定精度
func FloatToStr(num float64, floatPartLen int) string {
	return strconv.FormatFloat(num, 'f', floatPartLen, 64)
}

//strToFloat64 支持指定精度
func strToFloat64(str string, len int) float64 {
	lenstr := "%." + strconv.Itoa(len) + "f"
	value, _ := strconv.ParseFloat(str, 64)
	nstr := fmt.Sprintf(lenstr, value) //指定精度
	val, _ := strconv.ParseFloat(nstr, 64)
	return val
}

//strToFloat64 支持指定精度， 支持四舍五入
func strToFloat64round(str string, prec int, round bool) float64 {
	f, _ := strconv.ParseFloat(str, 64)
	return Precision(f, prec, round)
}

// float指定精度; round为true时, 表示支持四舍五入
func Precision(f float64, prec int, round bool) float64 {
	pow10_n := math.Pow10(prec)
	if round {
		return math.Trunc((f+0.5/pow10_n)*pow10_n) / pow10_n
	}
	return math.Trunc((f)*pow10_n) / pow10_n
}
