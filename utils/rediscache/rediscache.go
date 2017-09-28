package rediscache

import (
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/garyburd/redigo/redis"
	"github.com/gregjones/httpcache"
	"github.com/pkg/errors"
)

// cache is an implementation of httpcache.Cache that caches responses in a
// redis server.
type cache struct {
	p *redis.Pool
	l log.FieldLogger
}

// cacheKey modifies an httpcache key for use in redis. Specifically, it
// prefixes keys to avoid collision with other data stored in redis.
func cacheKey(key string) string {
	return "rediscache:" + key
}

// Get returns the response corresponding to key if present.
func (c *cache) Get(key string) (resp []byte, ok bool) {
	conn := c.p.Get()
	defer conn.Close()

	item, err := redis.Bytes(conn.Do("GET", cacheKey(key)))
	if err != nil {
		return nil, false
	}
	return item, true
}

// Set saves a response to the cache as key.
func (c *cache) Set(key string, resp []byte) {
	conn := c.p.Get()
	defer conn.Close()

	ttl := time.Hour * 24
	_, err := conn.Do("SETEX", cacheKey(key), ttl.Seconds(), resp)
	if err != nil {
		c.l.Error(errors.Wrap(err, "failed to cache item"))
	}
}

// Delete removes the response with key from the cache.
func (c *cache) Delete(key string) {
	conn := c.p.Get()
	defer conn.Close()

	_, err := conn.Do("DEL", cacheKey(key))
	if err != nil {
		c.l.Error(errors.Wrap(err, "failed to delete cached item"))
	}
}

// NewWithClient returns a new Cache with the given redis connection.
func NewWithPool(pool *redis.Pool, logger log.FieldLogger) httpcache.Cache {
	return &cache{
		p: pool,
		l: logger.WithField("m", "rediscache"),
	}
}
