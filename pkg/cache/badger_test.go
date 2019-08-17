package cache

import (
	"testing"

	o "stillgrove.com/goexp/pkg/object"
)

func TestBadgerCache(t *testing.T) {
	if true {
		t.Skip("skipping testing in short mode")
	}

	c, err := NewBadgerCache("test")
	if err != nil {
		t.Errorf("%v", err)
	}
	//defer cache.Close()

	var inObj = o.Object{
		Field: "test",
	}

	err = c.SaveRecord(
		"testObj",
		inObj.ToMarshal(),
		0,
	)
	if err != nil {
		t.Fatalf("%v", err)
	}

	var outObj o.Object
	payload, err := c.LoadRecord("testObj")
	if err != nil {
		t.Fatalf("%v", err)
	}

	err = outObj.FromMarshal(payload)
	if err != nil {
		t.Fatalf("%v", err)
	}

	if string(outObj.Field) != "test" {
		t.Errorf("Failed to retrieve the expected value from the cache")
	}
}
