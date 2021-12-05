package json

import (
	"testing"
)

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

var (
	strJSON string
	user    *User
	m       map[string]interface{}
)

func init() {
	user = &User{
		ID:   "1",
		Name: "pibigstar",
	}
	m = make(map[string]interface{})
	strJSON = `{"id":"1","name":"pibigstar"}`
}

func TestStructToJson(t *testing.T) {
	strJSON = StructToJSON(user)
	t.Log(strJSON)
}

func TestJsonToStruct(t *testing.T) {
	var user User
	err := JSONToStruct(strJSON, &user)
	if err != nil {
		t.Error(err)
	}
	t.Logf("%+v", user)
}

func TestMapToJson(t *testing.T) {
	m := make(map[string]interface{})
	m["name"] = "pibigstar"
	m["id"] = "1"
	strJSON = MapToJSON(m)
	t.Log(strJSON)
}

func TestJsonToMap(t *testing.T) {
	m := make(map[string]interface{})
	err := JSONToMap(strJSON, m)
	if err != nil {
		t.Error(err)
	}
	t.Logf("%+v", m)
}

func TestMapToStruct(t *testing.T) {
	m := make(map[string]interface{})
	m["name"] = "pibigstar"
	m["id"] = "1"
	var user User
	err := MapToStruct(m, &user)
	if err != nil {
		t.Error(err)
	}
	t.Logf("%+v", user)
}

func TestStructToMap(t *testing.T) {
	user := &User{
		ID:   "1",
		Name: "pibigstar",
	}
	m := make(map[string]interface{})
	err := StructToMap(user, m)
	if err != nil {
		t.Error(err)
	}
	t.Logf("%+v", m)
}

func TestMarshInterface(t *testing.T) {
	jsonStream := `{"name":"pibigstar", "num": 123456789}`
	err := MarshInterface(jsonStream)
	if err != nil {
		t.Error(err)
	}
}

func TestFormatMarshal(t *testing.T) {
	u := User{
		ID:   "123",
		Name: "派大星",
	}
	s := FormatMarshal(u)
	t.Log(s)
}

func TestMarshalUnEscapeHTML(t *testing.T) {
	u := User{
		ID:   "123",
		Name: "&&&&",
	}
	s := FormatMarshal(u)
	t.Log(s)
	s = MarshalUnEscapeHTML(u)
	t.Log(s)
}
