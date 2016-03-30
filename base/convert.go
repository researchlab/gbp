package base

import (
	"encoding/binary"
	"errors"
	"fmt"
	"math"
	"reflect"
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

func Struct2Map(obj interface{}) map[string]interface{} {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)

	var data = make(map[string]interface{})
	for i := 0; i < t.NumField(); i++ {
		data[t.Field(i).Name] = v.Field(i).Interface()
	}

	return data
}

func Map2Struct(obj map[string]interface{}, data interface{}) (interface{}, error) {
	for k, v := range obj {
		err := setField(data, k, v)
		if err != nil {
			return nil, err
		}
	}
	return data, nil
}

func setField(obj interface{}, name string, value interface{}) error {
	structVal := reflect.ValueOf(obj).Elem()
	structFieldVal := structVal.FieldByName(name)

	if !structFieldVal.IsValid() {
		return fmt.Errorf("No such field: %s in obj", name)
	}

	if !structFieldVal.CanSet() {
		return fmt.Errorf("Cannot set %s field value", name)
	}

	structFieldType := structFieldVal.Type()
	val := reflect.ValueOf(value)
	if structFieldType != val.Type() {
		return errors.New("provided value type didn't match obj field type")
	}
	structFieldVal.Set(val)
	return nil
}
