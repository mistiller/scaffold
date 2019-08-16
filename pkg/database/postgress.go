package database

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	mocket "github.com/selvatico/go-mocket"
)

type PostgresDB struct {
	client *gorm.DB
}

func NewPGTestDB(host, user, dbname, password string, port int, object interface{}) (*PostgresDB,  error) {
	db := PostgresDB{}
	var err error

	mocket.Catcher.Register()
	mocket.Catcher.Logging = true

	connectionString := fmt.Sprintf("host=%s user=%s dbname=%s password=%s port=%d sslmode=disable", host, user, dbname, password, port)
	db.client, err = gorm.Open(mocket.DriverName, connectionString)
	if err != nil {
		return &db, fmt.Errorf("Error setting up test database: - %s",err)
	}

	return &db, nil
}

func NewPostgresDB(host, user, dbname, password string, port int, object interface{}) (*PostgresDB,  error) {
	db := PostgresDB{}
	var err error

	connectionString := fmt.Sprintf("host=%s user=%s dbname=%s password=%s port=%d sslmode=disable", host, user, dbname, password, port)
	db.client, err = gorm.Open("postgres", connectionString)
	if err != nil {
		return &db, fmt.Errorf("Error connecting to database: - %s",err)
	}
	err = db.client.AutoMigrate(object).Error
	if err != nil {
		return &db, err
	}

	return &db, nil
}

func (db PostgresDB) CreateOne(object interface{}) error {
	fmt.Println(object)
	err := db.client.Create(object).Error
	return err
}

func (db PostgresDB) ReadOne(receiver interface{}, condition, value string) error {
	err := db.client.First(receiver, condition, value).Error
	return err
}

func (db PostgresDB) UpdateOne(receiver interface{}, condition, value interface{}) error {
	err := db.client.First(receiver, condition, value).Error
	return err
}

func (db PostgresDB) DeleteOne(receiver interface{}) error {
	err := db.client.Delete(receiver).Error
	return err
}

func (db PostgresDB) Close() {
	db.client.Close()
}

func (db PostgresDB) CreateMany(objects []interface{}) error {
	// Note the use of tx as the database handle once you are within a transaction
	tx := db.client.Begin()
	defer func() {
	  if r := recover(); r != nil {
		tx.Rollback()
	  }
	}()
  
	if err := tx.Error; err != nil {
	  return err
	}
  
	for i := range objects {
		if err := tx.Create(&objects[i]).Error; err != nil {
			tx.Rollback()
			return err
		 }
	}
  
	return tx.Commit().Error
  }