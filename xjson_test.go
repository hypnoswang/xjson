package xjson

import (
	//	"reflect"
	"encoding/json"
	"testing"
)

func TestJsonObjectEncodeDecode(t *testing.T) {
	j := `{"Name":"Homer", "Age":45, "Married":true, "height":176.8, "Families":[{"name":"Marge", "relation":"wife"}, {"name":"Bart", "relation":"son"}, {"name":"Lisa", "relation":"daughter"}, {"name":"Maggie", "relation":"son"}, {"name":"Abe", "relation":"father"}, {"name":"Mona", "relation":"mother"}]}`

	jo := NewJsonObject()
	if jo == nil {
		t.Error("NewJsonObject failed")
	}

	err := jo.Decode(j)
	if err != nil {
		t.Error("JsonObject Decode failed: ", err.Error())
	}

	res, err := jo.Encode()
	if err != nil {
		t.Error("JsonObject Encode failed: ", err.Error())
	}

	t.Log(res)
}

func TestJsonArrayEncodeDecode(t *testing.T) {
	j := `[{"name":"Marge", "relation":"wife"}, {"name":"Bart", "relation":"son"}, {"name":"Lisa", "relation":"daughter"}, {"name":"Maggie", "relation":"son"}, {"name":"Abe", "relation":"father"}, {"name":"Mona", "relation":"mother"}]`

	ja := NewJsonArray()
	if ja == nil {
		t.Error("NewJsonArray failed")
	}

	err := ja.Decode(j)
	if err != nil {
		t.Error("JsonArray Decode failed: ", err.Error())
	}

	res, err := ja.Encode()
	if err != nil {
		t.Error("JsonArray Encode failed: ", err.Error())
	}

	t.Log(res)
}

func TestJsonObjectGet(t *testing.T) {
	j := `{"Name":"Homer", "Age":45, "Married":true, "Height":176.8, "Families":[{"name":"Marge", "relation":"wife"}, {"name":"Bart", "relation":"son"}, {"name":"Lisa", "relation":"daughter"}, {"name":"Maggie", "relation":"son"}, {"name":"Abe", "relation":"father"}, {"name":"Mona", "relation":"mother"}], "Car":{"Maker":"Ford","Color":"Pink"}}`

	jo := NewJsonObject()
	if jo == nil {
		t.Error("NewJsonObject failed")
	}

	err := jo.Decode(j)
	if err != nil {
		t.Error("JsonObject Decode failed: ", err.Error())
	}

	if !jo.Has("Name") {
		t.Error("JsonObject Has failed: ")
	}

	if jo.Has("NotExist") {
		t.Error("JsonObject Has failed: ")
	}

	s, err := jo.GetString("Name")
	if err != nil {
		t.Error("JsonObject GetString failed: ", err.Error())
	}

	if s != "Homer" {
		t.Error("JsonObject GetString failed: values not match, ", s)
	}

	i, err := jo.GetInt64("Age")
	if err != nil {
		t.Error("JsonObject GetInt64 failed: ", err.Error())
	}

	if i != 45 {
		t.Error("JsonObject GetString failed: values not match, ", i)
	}

	f, err := jo.GetFloat64("Height")
	if err != nil {
		t.Error("JsonObject GetFloat64 failed: ", err.Error())
	}

	if f != 176.8 {
		t.Error("JsonObject GetHeight failed: values not match, ", f)
	}

	b, err := jo.GetBool("Married")
	if err != nil {
		t.Error("JsonObject GetBool failed: ", err.Error())
	}

	if !b {
		t.Error("JsonObject GetBool failed: values not match")
	}

	o, err := jo.GetObject("Car")
	if err != nil {
		t.Error("JsonObject GetObject failed: ", err.Error())
	}

	ot, err := GetJsonDataType(o)
	if err != nil {
		t.Error("JsonObject GetJsonDataType failed: ", err.Error())
	}

	if ot != "JsonObject" || len(o) <= 0 {
		t.Error("JsonObject GetObject failed: values not match, ", ot, " ", o)
	}

	a, err := jo.GetArray("Families")
	if err != nil {
		t.Error("JsonObject GetArray failed: ", err.Error())
	}

	at, err := GetJsonDataType(a)
	if err != nil {
		t.Error("JsonObject GetJsonDataType failed: ", err.Error())
	}

	if at != "JsonArray" || len(a) <= 0 {
		t.Error("JsonObject GetArray failed: values not match, ", at, " ", a)
	}
}

func TestJsonArrayGetString(t *testing.T) {
	j := `["away", "from", "world"]`

	ja := NewJsonArray()
	if ja == nil {
		t.Error("NewJsonArray failed")
	}

	err := ja.Decode(j)
	if err != nil {
		t.Error("JsonArray Decode failed: ", err.Error())
	}

	if ja.Len() != 3 {
		t.Error("JsonArray Len failed: ")
	}

	s, err := ja.GetString(1)
	if err != nil {
		t.Error("JsonArray GetString failed: ", err.Error())
	}

	if s != "from" {
		t.Error("JsonArray GetString failed: values not match, ", s)
	}

	s, err = ja.GetString(-1)
	if err == nil {
		t.Error("JsonArray GetString failed: index smaller than border")
	}

	s, err = ja.GetString(5)
	if err == nil {
		t.Error("JsonArray GetString failed: index bigger than border")
	}
}

func TestJsonArrayGetInt64(t *testing.T) {
	j := `[2, 4, 6, 8, 10]`

	ja := NewJsonArray()
	if ja == nil {
		t.Error("NewJsonArray failed")
	}

	err := ja.Decode(j)
	if err != nil {
		t.Error("JsonArray Decode failed: ", err.Error())
	}

	i, err := ja.GetInt64(1)
	if err != nil {
		t.Error("JsonArray GetInt64 failed: ", err.Error())
	}

	if i != 4 {
		t.Error("JsonArray GetInt64 failed: values not match, ", i)
	}

}

func TestJsonArrayGetFloat64(t *testing.T) {
	j := `[2.5, 2.6, 2.7, 2.8, 2.9]`

	ja := NewJsonArray()
	if ja == nil {
		t.Error("NewJsonArray failed")
	}

	err := ja.Decode(j)
	if err != nil {
		t.Error("JsonArray Decode failed: ", err.Error())
	}

	f, err := ja.GetFloat64(1)
	if err != nil {
		t.Error("JsonArray GetFloat64 failed: ", err.Error())
	}

	if f != 2.6 {
		t.Error("JsonArray GetFloat64 failed: values not match, ", f)
	}

}

func TestJsonArrayGetBool(t *testing.T) {
	j := `[false, true, false, false, false]`

	ja := NewJsonArray()
	if ja == nil {
		t.Error("NewJsonArray failed")
	}

	err := ja.Decode(j)
	if err != nil {
		t.Error("JsonArray Decode failed: ", err.Error())
	}

	b, err := ja.GetBool(1)
	if err != nil {
		t.Error("JsonArray GetBool failed: ", err.Error())
	}

	if !b {
		t.Error("JsonArray GetBool failed: values not match, ", b)
	}

}

func TestJsonArrayGetArray(t *testing.T) {
	j := `[["a", "b", "c"], ["1", "2", "3"]]`

	ja := NewJsonArray()
	if ja == nil {
		t.Error("NewJsonArray failed")
	}

	err := ja.Decode(j)
	if err != nil {
		t.Error("JsonArray Decode failed: ", err.Error())
	}

	a, err := ja.GetArray(1)
	if err != nil {
		t.Error("JsonArray GetArray failed: ", err.Error())
	}

	at, err := GetJsonDataType(a)
	if err != nil {
		t.Error("JsonArray GetJsonDataType failed: ", err.Error())
	}

	if at != "JsonArray" || len(a) <= 0 {
		t.Error("JsonObject GetObject failed: values not match, ", at, " ", a)
	}

}

func TestJsonArrayGetObject(t *testing.T) {
	j := `[{"name":"Marge", "relation":"wife"}, {"name":"Bart", "relation":"son"}, {"name":"Lisa", "relation":"daughter"}, {"name":"Maggie", "relation":"son"}, {"name":"Abe", "relation":"father"}, {"name":"Mona", "relation":"mother"}]`

	ja := NewJsonArray()
	if ja == nil {
		t.Error("NewJsonArray failed")
	}

	err := ja.Decode(j)
	if err != nil {
		t.Error("JsonArray Decode failed: ", err.Error())
	}

	o, err := ja.GetObject(1)
	if err != nil {
		t.Error("JsonArray GetObject failed: ", err.Error())
	}

	ot, err := GetJsonDataType(o)
	if err != nil {
		t.Error("JsonArray GetJsonDataType failed: ", err.Error())
	}

	if ot != "JsonObject" || len(o) <= 0 {
		t.Error("JsonObject GetObject failed: values not match, ", ot, " ", o)
	}

}

func TestJsonObjectSet(t *testing.T) {
	j := `{"a":3, "b":"ok"}`

	jo := NewJsonObject()
	if jo == nil {
		t.Error("NewJsonObject failed")
	}

	err := jo.Decode(j)
	if err != nil {
		t.Error("JsonObject Decode failed: ", err.Error())
	}

	err = jo.Set("a", 5)
	if err != nil {
		t.Error("JsonObject Set failed: ", err.Error())
	}
	i, err := jo.GetInt64("a")
	if err != nil || i != 5 {
		t.Error("JsonObject Set-Get failed: ", err.Error(), " ", i)
	}

	err = jo.Set("x", 5.09)
	if err != nil {
		t.Error("JsonObject Set failed: ", err.Error())
	}
	f, err := jo.GetFloat64("x")
	if err != nil || f != 5.09 {
		t.Error("JsonObject Set-Get failed: ", err.Error(), " ", i)
	}

	err = jo.Set("c", false)
	if err != nil {
		t.Error("JsonObject Set failed: ", err.Error())
	}
	b, err := jo.GetBool("c")
	if err != nil || b {
		t.Error("JsonObject Set-Get failed: ", err.Error(), " ", b)
	}

	err = jo.Set("b", "nok")
	if err != nil {
		t.Error("JsonObject Set failed: ", err.Error())
	}
	s, err := jo.GetString("b")
	if err != nil || s != "nok" {
		t.Error("JsonObject Set-Get failed: ", err.Error(), " ", s)
	}

	o1 := NewJsonObject()
	o1.Set("n", "Hypnos")
	o1.Set("w", 30)
	err = jo.Set("d", *o1)
	if err != nil {
		t.Error("JsonObject Set failed: ", err.Error())
	}
	o2, err := jo.GetObject("d")
	if err != nil || !o2.Has("n") || !o2.Has("w") {
		t.Error("JsonObject Set-Get failed: ", err.Error(), " ", o2)
	}

	a1 := NewJsonArray()
	a1.Add("Saga")
	a1.Add("Shaka")
	a1.Add("Milo")
	err = jo.Set("e", *a1)
	if err != nil {
		t.Error("JsonObject Set failed: ", err.Error())
	}
	a2, err := jo.GetArray("e")
	if err != nil || a2.Len() != 3 {
		t.Error("JsonObject Set-Get failed: ", err.Error(), " ", o2)
	}
}

func TestJsonObjectDel(t *testing.T) {
	j := `{"Name":"Homer", "Age":45, "Married":true, "Height":176.8, "Families":[{"name":"Marge", "relation":"wife"}, {"name":"Bart", "relation":"son"}, {"name":"Lisa", "relation":"daughter"}, {"name":"Maggie", "relation":"son"}, {"name":"Abe", "relation":"father"}, {"name":"Mona", "relation":"mother"}], "Car":{"Maker":"Ford","Color":"Pink"}}`

	jo := NewJsonObject()
	if jo == nil {
		t.Error("NewJsonObject failed")
	}

	err := jo.Decode(j)
	if err != nil {
		t.Error("JsonObject Decode failed: ", err.Error())
	}

	err = jo.Del("Name")
	if err != nil {
		t.Error("JsonObject Del failed: ", err.Error())
	}
	if jo.Has("Name") {
		t.Error("JsonObject Del string failed: ", err.Error())
	}

	err = jo.Del("Age")
	if err != nil {
		t.Error("JsonObject Del failed: ", err.Error())
	}
	if jo.Has("Age") {
		t.Error("JsonObject Del Int failed: ", err.Error())
	}

	err = jo.Del("Married")
	if err != nil {
		t.Error("JsonObject Del failed: ", err.Error())
	}
	if jo.Has("Married") {
		t.Error("JsonObject Del Bool failed: ", err.Error())
	}

	err = jo.Del("Height")
	if err != nil {
		t.Error("JsonObject Del failed: ", err.Error())
	}
	if jo.Has("Married") {
		t.Error("JsonObject Del float failed: ", err.Error())
	}

	err = jo.Del("Families")
	if err != nil {
		t.Error("JsonObject Del failed: ", err.Error())
	}
	if jo.Has("Families") {
		t.Error("JsonObject Del Array failed: ", err.Error())
	}

	err = jo.Del("Car")
	if err != nil {
		t.Error("JsonObject Del failed: ", err.Error())
	}
	if jo.Has("Car") {
		t.Error("JsonObject Del Object failed: ", err.Error())
	}

	err = jo.Del("NotExist")
	if err != nil {
		t.Error("JsonObject Del failed: ", err.Error())
	}
	if jo.Has("NotExist") {
		t.Error("JsonObject Del Object failed: ", err.Error())
	}
}

func TestJsonObjectClear(t *testing.T) {
	j := `{"Name":"Homer", "Age":45, "Married":true, "Height":176.8, "Families":[{"name":"Marge", "relation":"wife"}, {"name":"Bart", "relation":"son"}, {"name":"Lisa", "relation":"daughter"}, {"name":"Maggie", "relation":"son"}, {"name":"Abe", "relation":"father"}, {"name":"Mona", "relation":"mother"}], "Car":{"Maker":"Ford","Color":"Pink"}}`

	jo := NewJsonObject()
	if jo == nil {
		t.Error("NewJsonObject failed")
	}

	err := jo.Decode(j)
	if err != nil {
		t.Error("JsonObject Decode failed: ", err.Error())
	}

	jo.Clear()

	if jo.Has("Name") || jo.Has("Age") || jo.Has("Married") || jo.Has("Height") || jo.Has("Families") || jo.Has("Car") {
		t.Error("JsonObject Clear failed: ", err.Error())
	}
}

func TestJsonArraySet(t *testing.T) {
	j := `[{"name":"Marge", "relation":"wife"}, {"name":"Bart", "relation":"son"}, {"name":"Lisa", "relation":"daughter"}, {"name":"Maggie", "relation":"son"}, {"name":"Abe", "relation":"father"}, {"name":"Mona", "relation":"mother"}]`

	ja := NewJsonArray()
	if ja == nil {
		t.Error("NewJsonArray failed")
	}

	err := ja.Decode(j)
	if err != nil {
		t.Error("JsonArray Decode failed: ", err.Error())
	}

	err = ja.Set(-3, "NotExist")
	if err == nil {
		t.Error("JsonArray Set failed: low border check failed")
	}

	err = ja.Set(6, "NotExist")
	if err == nil {
		t.Error("JsonArray Set failed: high border check failed")
	}

	err = ja.Set(0, false)
	if err != nil {
		t.Error("JsonArray Set failed: ", err.Error())
	}
	b, err := ja.GetBool(0)
	if err != nil || b {
		t.Error("JsonArray Set-Get failed: ", err.Error(), " ", b)
	}

	err = ja.Set(1, 45)
	if err != nil {
		t.Error("JsonArray Set failed: ", err.Error())
	}
	i, err := ja.GetInt64(1)
	if err != nil || i != 45 {
		t.Error("JsonArray Set-Get failed: ", err.Error(), " ", i)
	}

	err = ja.Set(2, 3.14)
	if err != nil {
		t.Error("JsonArray Set failed: ", err.Error())
	}
	f, err := ja.GetFloat64(2)
	if err != nil || f != 3.14 {
		t.Error("JsonArray Set-Get failed: ", err.Error(), " ", f)
	}

	err = ja.Set(3, "Aolia")
	if err != nil {
		t.Error("JsonArray Set failed: ", err.Error())
	}
	s, err := ja.GetString(3)
	if err != nil || s != "Aolia" {
		t.Error("JsonArray Set-Get failed: ", err.Error(), " ", s)
	}

	o1 := NewJsonObject()
	o1.Set("n", "Wang")
	o1.Set("a", 45)
	err = ja.Set(4, *o1)
	if err != nil {
		t.Error("JsonArray Set failed: ", err.Error())
	}
	o2, err := ja.GetObject(4)
	if err != nil || !o2.Has("n") || !o2.Has("a") {
		t.Error("JsonArray Set-Get failed: ", err.Error(), " ", o2)
	}

	a1 := NewJsonArray()
	a1.Add("Saga")
	a1.Add("Milo")
	a1.Add("Shaka")
	err = ja.Set(5, *a1)
	if err != nil {
		t.Error("JsonArray Set failed: ", err.Error())
	}
	a2, err := ja.GetArray(5)
	if err != nil || a2.Len() != 3 {
		t.Error("JsonArray Set-Get failed: ", err.Error(), " ", o2)
	}
}

func TestJsonArrayAdd(t *testing.T) {
	j := `[{"name":"Marge", "relation":"wife"}, {"name":"Bart", "relation":"son"}, {"name":"Lisa", "relation":"daughter"}, {"name":"Maggie", "relation":"son"}, {"name":"Abe", "relation":"father"}, {"name":"Mona", "relation":"mother"}]`

	ja := NewJsonArray()
	if ja == nil {
		t.Error("NewJsonArray failed")
	}

	err := ja.Decode(j)
	if err != nil {
		t.Error("JsonArray Decode failed: ", err.Error())
	}

	err = ja.Add(45)
	if err != nil {
		t.Error("JsonArray Add failed: ", err.Error())
	}
	i, err := ja.GetInt64(6)
	if err != nil || i != 45 {
		t.Error("JsonArray Add-Get failed: ", err.Error(), " ", i)
	}

	err = ja.Add(3.14)
	if err != nil {
		t.Error("JsonArray Add failed: ", err.Error())
	}
	f, err := ja.GetFloat64(7)
	if err != nil || f != 3.14 {
		t.Error("JsonArray Add-Get failed: ", err.Error(), " ", f)
	}

	err = ja.Add("Aolia")
	if err != nil {
		t.Error("JsonArray Add failed: ", err.Error())
	}
	s, err := ja.GetString(8)
	if err != nil || s != "Aolia" {
		t.Error("JsonArray Add-Get failed: ", err.Error(), " ", s)
	}

	o1 := NewJsonObject()
	o1.Set("n", "Wang")
	o1.Set("a", 45)
	err = ja.Add(*o1)
	if err != nil {
		t.Error("JsonArray Add failed: ", err.Error())
	}
	o2, err := ja.GetObject(9)
	if err != nil || !o2.Has("n") || !o2.Has("a") {
		t.Error("JsonArray Add-Get failed: ", err.Error(), " ", o2)
	}

	a1 := NewJsonArray()
	a1.Add("Saga")
	a1.Add("Milo")
	a1.Add("Shaka")
	err = ja.Add(*a1)
	if err != nil {
		t.Error("JsonArray Add failed: ", err.Error())
	}
	a2, err := ja.GetArray(10)
	if err != nil || a2.Len() != 3 {
		t.Error("JsonArray Add-Get failed: ", err.Error(), " ", o2)
	}

	t.Log("Add resutl: ", *ja)
}
func TestJsonArrayDel(t *testing.T) {
	j := `[{"name":"Marge", "relation":"wife"}, {"name":"Bart", "relation":"son"}, {"name":"Lisa", "relation":"daughter"}, {"name":"Maggie", "relation":"son"}, {"name":"Abe", "relation":"father"}, {"name":"Mona", "relation":"mother"}]`

	ja := NewJsonArray()
	if ja == nil {
		t.Error("NewJsonArray failed")
	}

	err := ja.Decode(j)
	if err != nil {
		t.Error("JsonArray Decode failed: ", err.Error())
	}

	err = ja.Del(-1)
	if err == nil {
		t.Error("JsonArray Del failed: invalid low border failed")
	}

	err = ja.Del(9)
	if err == nil {
		t.Error("JsonArray Del failed: invalid high border failed")
	}

	err = ja.Del(0)
	if err != nil || ja.Len() != 5 {
		t.Error("JsonArray Del failed: low border failed")
	}

	err = ja.Del(4)
	if err != nil || ja.Len() != 4 {
		t.Error("JsonArray Del failed: high border failed")
	}

	err = ja.Del(2)
	if err != nil || ja.Len() != 3 {
		t.Error("JsonArray Del failed")
	}

	t.Log("Del resutl: ", *ja)
}

func TestJsonArrayClear(t *testing.T) {
	j := `[{"name":"Marge", "relation":"wife"}, {"name":"Bart", "relation":"son"}, {"name":"Lisa", "relation":"daughter"}, {"name":"Maggie", "relation":"son"}, {"name":"Abe", "relation":"father"}, {"name":"Mona", "relation":"mother"}]`

	ja := NewJsonArray()
	if ja == nil {
		t.Error("NewJsonArray failed")
	}

	err := ja.Decode(j)
	if err != nil {
		t.Error("JsonArray Decode failed: ", err.Error())
	}

	ja.Clear()
	if ja.Len() > 0 {
		t.Error("JsonArray Clear failed")
	}
}

func TestMapToJsonString(t *testing.T) {
	type testJ struct {
		Name     string `json:name`
		Wife     string `json:wife`
		Daughter string `json:daughter`
		Son      string `json:son`
	}

	m := make(map[string]string)
	m["name"] = "Homer"
	m["wife"] = "Magie"
	m["son"] = "Bart"
	m["daughter"] = "Lisa"

	s := MapToJsonString(m)
	if len(s) <= 0 {
		t.Error("mapToJsonString execution failed")
	}

	tj := &testJ{}
	err := json.Unmarshal([]byte(s), tj)
	if err != nil {
		t.Error("mapToJsonString failed: ", err.Error())
	}

	if tj.Name != "Homer" || tj.Wife != "Magie" || tj.Daughter != "Lisa" || tj.Son != "Bart" {
		t.Log(s)
		t.Error("mapToJsonString failed: json string incorrect")
	}
}

func TestJsonObjectAddEmptyArray(t *testing.T) {
	j := `{"Name":"Homer", "Car":{"Maker":"Ford","Color":"Pink"}}`

	jo := NewJsonObject()
	if jo == nil {
		t.Error("NewJsonObject failed")
	}

	err := jo.Decode(j)
	if err != nil {
		t.Error("JsonObject Decode failed: ", err.Error())
	}

	ja := NewJsonArray()

	jo.Set("ttt", *ja)
	js, _ := jo.Encode()

	t.Log("###########: ", js)

}
