package redis

import (
	"gin_practice/src/practice/gin/global/redis/exception"
	"github.com/gomodule/redigo/redis"
	"os"
)

var connection *redis.Conn = nil

func SetUpRedis() (*redis.Conn, error) {
	host := os.Getenv("REDIS_HOST")
	port := os.Getenv("REDIS_PORT")
	conn, err := redis.Dial("tcp", host+":"+port)
	if err != nil {
		return nil, err
	}
	connection = &conn
	return connection, nil
}

func GetRedisConn() (*redis.Conn, error) {
	if connection == nil {
		return nil, exception.ConnectionNullException()
	}
	return connection, nil
}

func SaveValue(hashName string, key string, value string, ttl int) error {
	if connection == nil {
		return exception.ConnectionNullException()
	}
	conn := *connection
	conn.Do("HMSET", hashName+":1", key, value)
	conn.Do("EXPIRE", key, ttl)
	return nil
}

func CloseConn() error {
	if connection == nil {
		return exception.ConnectionNullException()
	}
	(*connection).Close()
	return nil
}
