package cache


import (
	"testing"
	o "stillgrove.com/goexp/pkg/object"
)

func TestBolt(t *testing.T){
	cache, err := NewBoltCache(
		"bolt",
		"test",
	)
	defer cache.Close()
	if err != nil {
		t.Fatalf("%v", err)
	}
	
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
}
