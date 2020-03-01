package datasources

import (
	"time"

	"github.com/go-redis/redis"
	"github.com/rbcervilla/redisstore"
)

var (
	// RedisStore connection
	RedisStore *redisstore.RedisStore
	// RedisClient connection
	RedisClient *redis.Client
)

// NewredisDB creates a new database connection backed by a given redis server.
func NewredisDB(network string, host string, port string, password string, database int, idleTimeout time.Duration, prefix string) (dbStore *redisstore.RedisStore, err error) {

	if len(network) == 0 {
		network = "tcp"
	}
	if len(host) == 0 {
		host = "127.0.0.1"
	}
	if len(port) == 0 {
		port = "6379"
	}
	if idleTimeout == 0 {
		idleTimeout = time.Duration(5 * time.Minute)
	}

	RedisClient = redis.NewClient(&redis.Options{
		Network:     network,
		Addr:        host + ":" + port,
		IdleTimeout: idleTimeout,
		Password:    password,
		DB:          database,
	})

	RedisStore, err := redisstore.NewRedisStore(RedisClient)
	RedisStore.KeyPrefix(prefix)

	return RedisStore, err
}
