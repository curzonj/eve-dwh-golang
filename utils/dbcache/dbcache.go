package dbcache

import (
	"database/sql"

	log "github.com/Sirupsen/logrus"
	"github.com/gregjones/httpcache"
	"github.com/pkg/errors"
)

type dbType interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
	Get(dest interface{}, query string, args ...interface{}) error
}

// cache is an implementation of httpcache.Cache that caches responses in a
// redis server.
type cache struct {
	db dbType
	l  log.FieldLogger
}

// NewWithClient returns a new Cache with the given redis connection.
func New(db dbType, logger log.FieldLogger) httpcache.Cache {
	return &cache{
		db: db,
		l:  logger.WithField("m", "rediscache"),
	}
}

func (c *cache) Get(key string) (resp []byte, ok bool) {
	var content []byte
	err := c.db.Get(&content, "select content from cache_items where key = $1 limit 1", key)
	if err != nil {
		if err != sql.ErrNoRows {
			c.l.Error(errors.Wrap(err, "failed to get cached item"))
		}

		return nil, false
	}
	return content, true
}

func (c *cache) Set(key string, resp []byte) {
	_, err := c.db.Exec("insert into cache_items (key,content,last_touched_at) values ($1,$2,current_timestamp) ON CONFLICT(key) DO UPDATE SET content = $2, last_touched_at = current_timestamp", key, resp)
	if err != nil {
		c.l.Error(errors.Wrap(err, "failed to cache item"))
	}
}

func (c *cache) Delete(key string) {
	_, err := c.db.Exec("delete from cache_items where key = $1", key)
	if err != nil {
		c.l.Error(errors.Wrap(err, "failed to cache item"))
	}
}
