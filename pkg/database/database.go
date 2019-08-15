package database

type DB interface {
	CreateOne(interface{}) error
	CreateMany(objects []interface{}) error
	ReadOne(receiver interface{}, condition, value string) error
	UpdateOne(receiver interface{}, key string, value interface{}) error
	DeleteOne(receiver interface{}) error
	Close()
}