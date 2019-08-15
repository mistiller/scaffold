package main

import (
	"log"

	c "stillgrove.com/goexp/pkg/cache"
	o "stillgrove.com/goexp/pkg/object"
	db "stillgrove.com/goexp/pkg/database"
)
  
type Product struct {
	db.Object
	Code string
	Price uint
}

func check(err error){
	if err != nil{
		log.Fatalf("%v", err)
	}
}

func main() {
	pg, err := db.NewPostgresDB(
		"db",
		"gorm",
		"gorm",
		"mypassword",
		5432,
		Product{},
	)
	check(err)

	product := new(Product)
	*product = Product{Code: "L1212", Price: 1000}

	pg.CreateOne(product)
	pg.ReadOne(product, "code = ?", "L1212") 
	//pg.UpdateOne(product, "Price", 2000)

	log.Println(*product)
	pg.DeleteOne(product)
}

func cache() {
	cache, err := c.NewRedisCache(
		"redis", 
		6379, 
		"", 
		0,
		true,
	)
	
	check(err)
	defer cache.Close()
	
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

	log.Println(outObj)
}
