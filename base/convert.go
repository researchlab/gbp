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

//BytesToFloat64 bytes to float64
func BytesToFloat64(bytes []byte) float64 {
	bits := binary.LittleEndian.Uint64(bytes)
	return math.Float64frombits(bits)
}

//Float64ToBytes float64 to bytes; []uint8
func Float64ToBytes(input float64) []byte {
	bits := math.Float64bits(input)
	bytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(bytes, bits)
	return bytes
}

//FloatToStr float to str 支持指定精度
func FloatToStr(num float64, floatPartLen int) string {
	return strconv.FormatFloat(num, 'f', floatPartLen, 64)
}

//StrToFloat64 支持指定精度
func StrToFloat64(str string, len int) float64 {
	lenstr := "%." + strconv.Itoa(len) + "f"
	value, _ := strconv.ParseFloat(str, 64)
	nstr := fmt.Sprintf(lenstr, value) //指定精度
	val, _ := strconv.ParseFloat(nstr, 64)
	return val
}

//StrToFloat64round 支持指定精度， 支持四舍五入
func StrToFloat64round(str string, prec int, round bool) float64 {
	f, _ := strconv.ParseFloat(str, 64)
	return FloatPrecision(f, prec, round)
}

// FloatPrecision float指定精度; round为true时, 表示支持四舍五入
func FloatPrecision(f float64, prec int, round bool) float64 {
	pow10N := math.Pow10(prec)
	if round {
		return math.Trunc((f+0.5/pow10N)*pow10N) / pow10N
	}
	return math.Trunc((f)*pow10N) / pow10N
}

//Struct2Map struct convert to Map
func Struct2Map(obj interface{}) map[string]interface{} {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)

	var data = make(map[string]interface{})
	for i := 0; i < t.NumField(); i++ {
		data[t.Field(i).Name] = v.Field(i).Interface()
	}

	return data
}

//Map2Struct map to struct
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
