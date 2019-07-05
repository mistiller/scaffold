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

func cacheObject(c *c.Cache, obj *o.Object, name string) error {
	err := c.SaveRecord(
		name, 
		obj.ToMarshal(), 
		0,
	)
	return err
}

func fetchObject(c *c.Cache, name string) (obj o.Object, err error){
	payload, err := c.LoadRecord(name)
	if err != nil {
		return obj, err
	}
	err = obj.FromMarshal(payload)
	if err != nil {
		return obj, err
	}

	return obj, nil
}

func main() {
	cache, err := c.BuildCache(
		"redis", 
		6379, 
		"", 
		0,
		true,
	)
	check(err)

	var obj = o.Object{
		Field: "test",
	}

	err = cacheObject(&cache, &obj, "testObj")
	check(err)

	payload, err := fetchObject(&cache, "testObj")
	check(err)
	fmt.Println(payload)
}
