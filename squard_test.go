package SquadCache

import (
	"reflect"
	"testing"
)

func TestGetter(t *testing.T)  {
	var f Getter = GetterFunc(func(key string) ([]byte, error) {
		return []byte(key), nil
	})
	expect := []byte("key")
	v, _ := f.Get("key")
	if !reflect.DeepEqual(v, expect) {
		t.Errorf("callback failed")
	}
}




