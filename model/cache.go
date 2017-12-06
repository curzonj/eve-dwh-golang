package model

import (
	"database/sql"

	"github.com/pkg/errors"
)

func (d *DAO) SetCacheItem(key string, data []byte) error {
	_, err := d.Exec("insert into cache_items (key,content,last_touched_at) values ($1,$2,current_timestamp) ON CONFLICT(key) DO UPDATE SET content = $2, last_touched_at = current_timestamp", key, data)
	if err != nil {
		return errors.Wrap(err, "SetCacheItem")
	}

	return nil
}

func (d *DAO) GetCacheItem(key string) ([]byte, error) {
	var content []byte
	err := d.Get(&content, "select content from cache_items where key = $1 limit 1", key)
	if err != nil {
		if err != sql.ErrNoRows {
			return nil, errors.Wrap(err, "GetCacheItem")
		}

		return nil, nil
	}

	return content, nil
}
