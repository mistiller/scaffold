package cache

import (
	"testing"
	"github.com/alicebob/miniredis"
	o "stillgrove.com/goexp/pkg/object"
)

func TestRedis(t *testing.T){
	m, err := miniredis.Run()
	if err != nil {
		t.Fatalf("%v", err)
	}
	defer m.Close()

	// Optionally set some keys your code expects:
	//s.Set("foo", "bar")

	cache, err := NewRedisCache(
		m.Addr(), 
		0, 
		"", 
		0,
		true,
	)
	if err != nil {
		t.Fatalf("%v", err)
	}
	defer cache.Close()
	
	var inObj = o.Object{
		Field: "test",
	}
	
	err = cache.SaveRecord(
		"testObj", 
		inObj.ToMarshal(), 
		0,
	)
	if err != nil {
		t.Fatalf("%v", err)
	}
	
	var outObj o.Object
	payload, err := cache.LoadRecord("testObj")
	if err != nil {
		t.Fatalf("%v", err)
	}
	
	err = outObj.FromMarshal(payload)
	if err != nil {
		t.Fatalf("%v", err)
	}

	if outObj.Field != inObj.Field {
		t.Fatalf("Object field wasn't reproduced properly -%s", outObj)
	}
}