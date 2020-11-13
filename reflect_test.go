package utils

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"
)

type TestReflect struct {
	Foo string
	Map map[string]string
}

func TestAssign(t *testing.T) {
	values := make(map[string]interface{})
	values["Foo"] = "bar"
	values["Map"] = map[string]string{"go": "pher"}
	values["Ja"] = "va"
	test := new(TestReflect)
	Assign(test, values)
	if test.Foo != "bar" {
		t.Errorf("Invalid dume got %s instead of %s", test.Foo, "bar")
	}
	if test.Map == nil {
		t.Error("Missing value for Map")
	}
	if v, ok := test.Map["go"]; ok {
		if v != "pher" {
			t.Errorf("Invalid dume got %s instead of %s", v, "pher")
		}
	} else {
		t.Error("Missing value for map key \"go\"")
	}
	return
}

func TestToInterfaceArray(t *testing.T) {
	ToInterfaceArray([]int{1, 2, 3, 4, 5, 6})
}

func TestAssignStringString(t *testing.T) {
	type TestStruct struct {
		Foo string
	}

	obj := &TestStruct{}
	elem := reflect.ValueOf(obj).Elem().Field(0)
	err := AssignStringString(elem, "bar")
	require.NoError(t, err)
	require.Equal(t, "bar", obj.Foo)
}

func TestAssignBoolString(t *testing.T) {
	type TestStruct struct {
		Foo bool
	}

	obj := &TestStruct{}
	elem := reflect.ValueOf(obj).Elem().Field(0)
	err := AssignBoolString(elem, "true")
	require.NoError(t, err)
	require.True(t, obj.Foo)

	err = AssignBoolString(elem, "false")
	require.NoError(t, err)
	require.False(t, obj.Foo)

	err = AssignBoolString(elem, "blah")
	require.Error(t, err)
}

func TestAssignIntString(t *testing.T) {
	type TestStruct struct {
		Foo int
	}

	obj := &TestStruct{}
	elem := reflect.ValueOf(obj).Elem().Field(0)
	err := AssignIntString(elem, "42")
	require.NoError(t, err)
	require.Equal(t, 42, obj.Foo)
}

func TestAssignUintString(t *testing.T) {
	type TestStruct struct {
		Foo uint
	}

	obj := &TestStruct{}
	elem := reflect.ValueOf(obj).Elem().Field(0)
	err := AssignUintString(elem, "42")
	require.NoError(t, err)
	require.Equal(t, uint(42), obj.Foo)
}

func TestAssignFloatString(t *testing.T) {
	type TestStruct struct {
		Foo float64
	}

	obj := &TestStruct{}
	elem := reflect.ValueOf(obj).Elem().Field(0)
	err := AssignFloatString(elem, "42.26")
	require.NoError(t, err)
	require.Equal(t, 42.26, obj.Foo)
}

func TestAssignSliceValues(t *testing.T) {
	type TestStruct struct {
		Slice []string
	}

	obj := &TestStruct{}
	field := reflect.ValueOf(obj).Elem().Field(0)
	err := AssignJsonSliceString(field, "[\"foo\",\"bar\",\"baz\"]")
	require.NoError(t, err)
	require.EqualValues(t, []string{"foo", "bar", "baz"}, obj.Slice)
}

func TestAssignSliceIntValues(t *testing.T) {
	type TestStruct struct {
		Slice []int
	}

	obj := &TestStruct{}
	field := reflect.ValueOf(obj).Elem().Field(0)
	err := AssignJsonSliceString(field, "[0,1,2]")
	require.NoError(t, err)
	require.EqualValues(t, []int{0, 1, 2}, obj.Slice)
}

func TestAssignMapValuesNilMap(t *testing.T) {
	type TestStruct struct {
		Map map[string]string
	}

	obj := &TestStruct{}
	field := reflect.ValueOf(obj).Elem().Field(0)
	err := AssignJsonMapString(field, "{\"foo\":\"bar\"}")
	require.NoError(t, err)
	require.EqualValues(t, map[string]string{"foo": "bar"}, obj.Map)
}

func TestAssignMapValues(t *testing.T) {
	type TestStruct struct {
		Map map[string]string
	}

	obj := &TestStruct{Map: map[string]string{"foo": "baz", "plip": "plop"}}
	field := reflect.ValueOf(obj).Elem().Field(0)
	err := AssignJsonMapString(field, "{\"foo\":\"bar\"}")
	require.NoError(t, err)
	require.EqualValues(t, map[string]string{"foo": "bar", "plip": "plop"}, obj.Map)
}

func TestAssignStrings(t *testing.T) {
	type TestStruct struct {
		Bool   bool
		String string
		Int    int
		Uint   uint
		Float  float64
		Slice  []interface{}
		Map    map[string]interface{}
		Nope   struct{}
	}

	getter := func(name string) (value string, ok bool) {
		switch name {
		case "Bool":
			return "true", true
		case "String":
			return "string", true
		case "Int":
			return "42", true
		case "Uint":
			return "42", true
		case "Float":
			return "42.26", true
		case "Slice":
			return "[\"foo\",\"bar\"]", true
		case "Map":
			return "{\"foo\":\"bar\"}", true
		default:
			return "", false
		}
	}
	obj := &TestStruct{}
	err := AssignStrings(obj, getter)
	require.NoError(t, err)
	require.True(t, obj.Bool)
	require.Equal(t, "string", obj.String)
	require.Equal(t, 42, obj.Int)
	require.Equal(t, uint(42), obj.Uint)
	require.Equal(t, 42.26, obj.Float)
	require.EqualValues(t, []interface{}{"foo", "bar"}, obj.Slice)
	require.EqualValues(t, map[string]interface{}{"foo": "bar"}, obj.Map)
}
