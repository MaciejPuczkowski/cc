package dict

import (
	"encoding/json"
	"testing"
)

func Test_dict(t *testing.T) {
	d := New[string, any]()
	d.Set("a", 1)
	d.Set("b", "bb")
	d.Set("c", true)
	d.Set("d", []int{1, 2, 3})
	d.Set("e", map[string]int{"a": 1, "b": 2})
	if d.Get("a", 0) != 1 {
		t.Errorf("d.Get(\"a\") = %v, want %v", d.Get("a", 0), 1)
	}
	if d.Get("b", "bc") != "bb" {
		t.Errorf("d.Get(\"b\") = %v, want %v", d.Get("b", "bc"), "bb")
	}
	if d.Get("c", false) != true {
		t.Errorf("d.Get(\"c\") = %v, want %v", d.Get("c", false), true)
	}
	if d.Get("x", "aaa") != "aaa" {
		t.Errorf("d.Get(\"x\") = %v, want %v", d.Get("x", "aaa"), "aaa")
	}
	if d.GetF("x", func() any { return "aaa" }) != "aaa" {
		t.Errorf("d.Get(\"x\") = %v, want %v", d.Get("x", "aaa"), "aaa")
	}
	if d.Get("x", nil) != nil {
		t.Errorf("d.Get(\"x\") = %v, want %v", d.Get("x", nil), nil)
	}
	if d.GetF("x", nil) != nil {
		t.Errorf("d.Get(\"x\") = %v, want %v", d.Get("x", nil), nil)
	}
	if d.GetF("x", func() any { return nil }) != nil {
		t.Errorf("d.Get(\"x\") = %v, want %v", d.Get("x", nil), nil)
	}
	d.Delete("a")
	if d.Get("a", 0) != 0 {
		t.Errorf("d.Get(\"a\") = %v, want %v", d.Get("a", 0), 0)
	}
	d.Delete("a")
	d.SetDefault("z", 7)
	if d.Get("z", 0) != 7 {
		t.Errorf("d.Get(\"z\") = %v, want %v", d.Get("z", 0), 7)
	}
	d.SetDefault("b", "zz")
	if d.Get("b", "xx") != "bb" {
		t.Errorf("d.Get(\"b\") = %v, want %v", d.Get("b", "xx"), "bb")
	}

}

func Test_jsonify(t *testing.T) {
	d := New[string, any]()
	d.Set("a", 1)
	d.Set("b", "bb")
	d.Set("c", true)
	d.Set("d", []int{1, 2, 3})
	d.Set("e", map[string]int{"a": 1, "b": 2})
	b, err := json.Marshal(d)
	if err != nil {
		t.Errorf("json.Marshal(d) = %v", err)
	}
	var dd Dict[string, any]
	err = json.Unmarshal(b, &dd)
	if err != nil {
		t.Errorf("json.Unmarshal(b, &dd) = %v", err)
	}
	if dd.Get("a", 0).(float64) != 1 {
		t.Errorf("dd.Get(\"a\") = %v, want %v", dd.Get("a", 0), 1)
	}

}

func Test_interfaces(t *testing.T) {
	var _ Dicter[string, any] = New[string, any]()
	var _ Dicter[string, []string] = Sliced[string, string](New[string, []string]())
}

func Test_sliced(t *testing.T) {
	d := New[string, []string]()
	db := Sliced[string, string](d)
	db.Append("a", "aa")
	if v := d.Get("a", []string{}); len(v) != 1 || v[0] != "aa" {
		t.Errorf("d.Get(\"a\") = %v, want %v", d.Get("a", []string{}), []string{"aa"})
	}
	db.Append("b", "bb")
	if v := d.Get("b", []string{}); len(v) != 1 || v[0] != "bb" {
		t.Errorf("d.Get(\"b\") = %v, want %v", d.Get("b", []string{}), []string{"bb"})
	}
	db.Append("b", "bb")
	if v := d.Get("b", []string{}); len(v) != 2 || v[0] != "bb" || v[1] != "bb" {
		t.Errorf("d.Get(\"b\") = %v, want %v", d.Get("b", []string{}), []string{"bb", "bb"})
	}
}

func Test_setdefault(t *testing.T) {
	d := New[string, []string]()
	d.SetDefault("a", []string{})
	// d.Get("a", []string{}) = append(d.Get("a", []string{}))
}
