package utils

import (
	"reflect"
	"testing"
)

type original struct {
	A int    `yaml:"a,omitempty,test"`
	B string `yaml:"b,omitempty"`
	C bool   `yaml:"c,test,omitempty"`
	D int    `json:"d,omitempty"`
	E string
}

type target struct {
	A int    `yaml:"a,test"`
	B string `yaml:"b"`
	C bool   `yaml:"c,test"`
	D int    `json:"d,omitempty"`
	E string
}

func TestRemoveTagsOmitempty(t *testing.T) {
	o := original{1, "a", true, 2, "b"}
	target := target{1, "a", true, 2, "b"}
	modified, err := RemoveTagsOmitempty(o, "yaml")
	if err != nil {
		t.Error(err)
	}
	tVal := reflect.ValueOf(target)
	mVal := reflect.ValueOf(modified)
	tType := tVal.Type()
	mType := mVal.Type()
	for i := 0; i < tVal.NumField(); i++ {
		wantField := tVal.Field(i).Interface()
		gotField := mVal.Field(i).Interface()
		if !reflect.DeepEqual(wantField, gotField) {
			t.Errorf("Fields are not equal: want: %v got: %v", wantField, gotField)
		}
		wantT := tType.Field(i)
		gotT := mType.Field(i)
		if !reflect.DeepEqual(wantT, gotT) {
			t.Errorf("StructFields are not equal: want: %v got: %v", wantT, gotT)
		}
	}
}
