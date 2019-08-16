package database

import(
	"testing"
)

type Product struct {
	Object
	Code string
	Price uint
}

func TestPostgres(t *testing.T) {
	pg, err := NewPGTestDB(
		"db",
		"gorm",
		"gorm",
		"mypassword",
		5432,
		Product{},
	)
	if err != nil {
		t.Fatalf("%v", err)
	}
	
	var product = Product{Code: "L1212", Price: 1000}
	
	err = pg.CreateOne(&product)
	if err != nil {
		t.Fatalf("%v", err)
	}
	/*err = pg.ReadOne(&product, "code = ?", "L1212") 
	if err != nil {
		t.Fatalf("%v", err)
	}
	//pg.UpdateOne(product, "Price", 2000)
	
	err = pg.DeleteOne(&product)
	if err != nil {
		t.Fatalf("%v", err)
	}*/
}
