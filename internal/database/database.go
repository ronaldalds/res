package database

type Database struct {
	GormStore  *GormStore
	RedisStore *RedisStore
}

var DB Database

func DbLoad() {
	DB = Database{
		GormStore:  newGormStore(),
		RedisStore: newRedisStore(),
	}
}
