package qjson

import (
	"encoding/json"
	"errors"
	//"fmt"
	"reflect"
)

type JsonObject map[string]interface{}
type JsonArray []interface{}

func NewJsonObject() *JsonObject {
	jo := make(JsonObject)
	return &jo
}

func NewJsonArray() *JsonArray {
	ja := make(JsonArray, 0)
	return &ja
}

func GetJsonDataType(v interface{}) (string, error) {
	vtype := reflect.TypeOf(v).Kind().String()
	rt := vtype

	switch vtype {
	case "slice":
		rt = "JsonArray"
	case "map":
		rt = "JsonObject"
	case "float64":
		fallthrough
	case "string":
		fallthrough
	case "bool":
		rt = vtype
	default:
		return "", errors.New("Invalid json data type")
	}

	return rt, nil
}

func (this *JsonObject) Decode(s string) error {
	if len(s) <= 0 {
		return errors.New("Input empty")
	}

	return json.Unmarshal([]byte(s), this)
}

func (this JsonObject) Encode() (string, error) {
	t, err := json.Marshal(this)
	return string(t), err
}

func (this *JsonArray) Decode(s string) error {
	if len(s) <= 0 {
		return errors.New("Input empty")
	}

	return json.Unmarshal([]byte(s), this)
}

func (this JsonArray) Encode() (string, error) {
	t, err := json.Marshal(this)
	return string(t), err
}

func (this JsonObject) Has(key string) bool {
	if len(key) <= 0 {
		return false
	}

	_, ok := this[key]

	return ok
}

func (this JsonObject) get(key string) (interface{}, error) {
	if len(key) <= 0 {
		return nil, errors.New("Key empty")
	}

	val, ok := this[key]
	if !ok {
		return nil, errors.New("Key not exist")
	}

	//vtype, err := getJsonDataType(val)
	//if err != nil {
	//	return nil, "", err
	//}

	return val, nil
}

func (this JsonObject) GetString(key string) (s string, e error) {
	v, e := this.get(key)
	if e != nil {
		return s, e
	}

	if s, ok := v.(string); ok {
		return s, nil
	}

	return s, errors.New("Value is not of type string")
}

func (this JsonObject) GetStringDef(key string, def string) string {
	v, e := this.GetString(key)
	if e != nil {
		return def
	}

	return v
}

func (this JsonObject) GetInt64(key string) (i int64, e error) {
	v, e := this.get(key)
	if e != nil {
		return i, e
	}

	if ii, ok := v.(float64); ok {
		return int64(ii), nil
	}

	return i, errors.New("Value is not of type int64")
}

func (this JsonObject) GetInt64Def(key string, def int64) int64 {
	v, e := this.GetInt64(key)
	if e != nil {
		return def
	}

	return v
}

func (this JsonObject) GetFloat64(key string) (f float64, e error) {
	v, e := this.get(key)
	if e != nil {
		return f, e
	}

	if f, ok := v.(float64); ok {
		return f, nil
	}

	return f, errors.New("Value is not of type float64")
}

func (this JsonObject) GetFloat64Def(key string, def float64) float64 {
	v, e := this.GetFloat64(key)
	if e != nil {
		return def
	}

	return v
}

func (this JsonObject) GetBool(key string) (b bool, e error) {
	v, e := this.get(key)
	if e != nil {
		return b, e
	}

	if b, ok := v.(bool); ok {
		return b, nil
	}

	return b, errors.New("Value is not of type bool")
}

func (this JsonObject) GetBoolDef(key string, def bool) bool {
	v, e := this.GetBool(key)
	if e != nil {
		return def
	}

	return v
}

func (this JsonObject) GetObject(key string) (o JsonObject, e error) {
	v, e := this.get(key)
	if e != nil {
		return o, e
	}

	if to, ok := v.(map[string]interface{}); ok {
		o = to
		return o, nil
	} else if o, ok = v.(JsonObject); ok {
		return o, nil
	}

	return o, errors.New("Value is not of type JsonObject")
}

func (this JsonObject) GetArray(key string) (a JsonArray, e error) {
	v, e := this.get(key)
	if e != nil {
		return a, e
	}

	if ta, ok := v.([]interface{}); ok {
		a = ta
		return a, nil
	} else if a, ok = v.(JsonArray); ok {
		return a, nil
	}

	return a, errors.New("Value is not of type JsonArray")
}

func (this JsonArray) Len() int {
	return len(this)
}

func (this JsonArray) get(idx int) (interface{}, error) {
	if idx < 0 || idx > (len(this)-1) {
		return nil, errors.New("Idx out of range")
	}

	val := this[idx]

	//vtype, err := getJsonDataType(val)
	//if err != nil {
	//	return nil, "", err
	//}

	return val, nil
}

func (this JsonArray) GetString(idx int) (s string, e error) {
	v, e := this.get(idx)
	if e != nil {
		return s, e
	}

	if s, ok := v.(string); ok {
		return s, nil
	}

	return s, errors.New("Value is not of type string")
}

func (this JsonArray) GetStringDef(idx int, def string) string {
	v, e := this.GetString(idx)
	if e != nil {
		return def
	}

	return v
}

func (this JsonArray) GetInt64(idx int) (i int64, e error) {
	v, e := this.get(idx)
	if e != nil {
		return i, e
	}

	if ii, ok := v.(float64); ok {
		return int64(ii), nil
	}

	return i, errors.New("Value is not of type int64")
}

func (this JsonArray) GetInt64Def(idx int, def int64) int64 {
	v, e := this.GetInt64(idx)
	if e != nil {
		return def
	}

	return v
}

func (this JsonArray) GetFloat64(idx int) (f float64, e error) {
	v, e := this.get(idx)
	if e != nil {
		return f, e
	}

	if f, ok := v.(float64); ok {
		return f, nil
	}

	return f, errors.New("Value is not of type float64")
}

func (this JsonArray) GetFloat64Def(idx int, def float64) float64 {
	v, e := this.GetFloat64(idx)
	if e != nil {
		return def
	}

	return v
}

func (this JsonArray) GetBool(idx int) (b bool, e error) {
	v, e := this.get(idx)
	if e != nil {
		return b, e
	}

	if b, ok := v.(bool); ok {
		return b, nil
	}

	return b, errors.New("Value is not of type bool")
}

func (this JsonArray) GetBoolDef(idx int, def bool) bool {
	v, e := this.GetBool(idx)
	if e != nil {
		return def
	}

	return v
}

func (this JsonArray) GetObject(idx int) (o JsonObject, e error) {
	v, e := this.get(idx)
	if e != nil {
		return o, e
	}

	if to, ok := v.(map[string]interface{}); ok {
		o = to
		return o, nil
	} else if o, ok = v.(JsonObject); ok {
		return o, nil
	}

	return o, errors.New("Value is not of type JsonObject")
}

func (this JsonArray) GetArray(idx int) (a JsonArray, e error) {
	v, e := this.get(idx)
	if e != nil {
		return a, e
	}

	if ta, ok := v.([]interface{}); ok {
		a = ta
		return a, nil
	} else if a, ok = v.(JsonArray); ok {
		return a, nil
	}

	return a, errors.New("Value is not of type JsonArray")
}

func (this *JsonObject) Set(key string, v interface{}) error {
	if len(key) <= 0 {
		return errors.New("Key empty")
	}

	if v == nil {
		return errors.New("Value empty")
	}

	vtype := reflect.TypeOf(v).Kind().String()
	switch vtype {
	case "slice":
		ja, _ := v.(JsonArray)
		(*this)[key] = ja
	case "map":
		jo, _ := v.(JsonObject)
		(*this)[key] = jo
	case "int":
		i, _ := v.(int)
		(*this)[key] = float64(i)
	case "int64":
		i64, _ := v.(int64)
		(*this)[key] = float64(i64)
	case "float32":
		f32, _ := v.(float32)
		(*this)[key] = float64(f32)
	case "float64":
		f64, _ := v.(float64)
		(*this)[key] = f64
	case "string":
		s, _ := v.(string)
		(*this)[key] = s
	case "bool":
		b, _ := v.(bool)
		(*this)[key] = b
	default:
		return errors.New("Unsupported data type")
	}

	return nil
}

func (this *JsonArray) Set(idx int, v interface{}) error {
	if idx < 0 || idx > (len(*this)-1) {
		return errors.New("Idx out of range")
	}

	if v == nil {
		return errors.New("Value empty")
	}

	vtype := reflect.TypeOf(v).Kind().String()
	switch vtype {
	case "slice":
		ja, _ := v.(JsonArray)
		(*this)[idx] = ja
	case "map":
		jo, _ := v.(JsonObject)
		(*this)[idx] = jo
	case "int":
		i, _ := v.(int)
		(*this)[idx] = float64(i)
	case "int64":
		i64, _ := v.(int64)
		(*this)[idx] = float64(i64)
	case "float32":
		f32, _ := v.(float32)
		(*this)[idx] = float64(f32)
	case "float64":
		f64, _ := v.(float64)
		(*this)[idx] = f64
	case "string":
		s, _ := v.(string)
		(*this)[idx] = s
	case "bool":
		b, _ := v.(bool)
		(*this)[idx] = b
	default:
		return errors.New("Unsupported data type")
	}

	return nil
}

func (this *JsonArray) Add(v interface{}) error {
	if v == nil {
		return errors.New("Value empty")
	}

	vtype := reflect.TypeOf(v).Kind().String()
	switch vtype {
	case "slice":
		ja, _ := v.(JsonArray)
		*this = append(*this, ja)
	case "map":
		jo, _ := v.(JsonObject)
		*this = append(*this, jo)
	case "int":
		i, _ := v.(int)
		*this = append(*this, float64(i))
	case "int64":
		i64, _ := v.(int64)
		*this = append(*this, float64(i64))
	case "float32":
		f32, _ := v.(float32)
		*this = append(*this, float64(f32))
	case "float64":
		f64, _ := v.(float64)
		*this = append(*this, float64(f64))
	case "string":
		s, _ := v.(string)
		*this = append(*this, s)
	case "bool":
		b, _ := v.(bool)
		*this = append(*this, b)
	default:
		return errors.New("Unsupported data type")
	}

	return nil
}

func (this *JsonObject) Del(key string) error {
	if len(key) <= 0 {
		return errors.New("Key empty")
	}

	delete(*this, key)

	return nil
}

func (this *JsonArray) Del(idx int) error {
	if idx < 0 || idx > (len(*this)-1) {
		return errors.New("Idx out of range")
	}

	*this = append((*this)[:idx], (*this)[idx+1:]...)

	return nil
}

func (this *JsonObject) Clear() {
	*this = make(JsonObject)
}

func (this *JsonArray) Clear() {
	*this = make(JsonArray, 0)
}

func (this *JsonObject) Iterable() map[string]interface{} {
	return map[string]interface{}(*this)
}

func (this *JsonArray) Iterable() []interface{} {
	return []interface{}(*this)
}

func MapToJsonString(m map[string]string) (s string) {
	msg, err := json.Marshal(m)
	if err != nil {
		return s
	}

	return string(msg)
}
