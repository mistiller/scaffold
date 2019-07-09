package main

import (
	"fmt"
	"os"

	c "stillgrove.com/goexp/pkg/cache"
	o "stillgrove.com/goexp/pkg/object"
)

func check(err error){
	if err != nil{
		fmt.Println(err)
		os.Exit(1)
	}
}

func main() {
	cache, err := c.BuildBoltCache(
		"bolt",
		"test",
	)
	defer cache.Close()
	check(err)

	var inObj = o.Object{
		Field: "test",
	}

	err = cache.SaveRecord(
		"testObj", 
		inObj.ToMarshal(), 
		0,
	)
	check(err)

	var outObj o.Object
	payload, err := cache.LoadRecord("testObj")
	check(err)

	err = outObj.FromMarshal(payload)
	check(err)

	fmt.Println(outObj)
}

func redis() {
	cache, err := c.BuildRedisCache(
		"redis", 
		6379, 
		"", 
		0,
		true,
	)
	check(err)

	var inObj = o.Object{
		Field: "test",
	}

	err = cache.SaveRecord(
		"testObj", 
		inObj.ToMarshal(), 
		0,
	)
	check(err)

	var outObj o.Object
	payload, err := cache.LoadRecord("test")
	check(err)

	err = outObj.FromMarshal(payload)
	check(err)

	fmt.Println(outObj)
}
