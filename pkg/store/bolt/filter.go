package store

import (
	"encoding/binary"
	"encoding/json"

	bolt "github.com/coreos/bbolt"
	"github.com/ncarlier/feedpushr/autogen/app"
	"github.com/ncarlier/feedpushr/pkg/common"
)

// FILTER_BUCKET bucket name
var FILTER_BUCKET = []byte("FILTER")

// itob returns an 8-byte big endian representation of v.
func itob(v int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}

// GetFilter returns a stored Filter.
func (store *BoltStore) GetFilter(ID int) (*app.Filter, error) {
	var result app.Filter
	err := store.get(FILTER_BUCKET, itob(ID), &result)
	if err != nil {
		if err == bolt.ErrInvalid {
			return nil, common.ErrFilterNotFound
		}
		return nil, err
	}
	return &result, nil
}

// DeleteFilter removes a filter.
func (store *BoltStore) DeleteFilter(ID int) (*app.Filter, error) {
	filter, err := store.GetFilter(ID)
	if err != nil {
		return nil, err
	}

	err = store.delete(FILTER_BUCKET, itob(filter.ID))
	if err != nil {
		return nil, err
	}
	return filter, nil
}

// SaveFilter stores a filter.
func (store *BoltStore) SaveFilter(filter *app.Filter) error {
	if filter.ID == 0 {
		var err error
		id, err := store.nextSequence(FILTER_BUCKET)
		if err != nil {
			return err
		}
		filter.ID = int(id)
	}
	return store.save(FILTER_BUCKET, itob(filter.ID), &filter)
}

// ListFilters returns a paginated list of filters.
func (store *BoltStore) ListFilters(page, limit int) (*app.FilterCollection, error) {
	bufs, err := store.allAsRaw(FILTER_BUCKET, page, limit)
	if err != nil {
		return nil, err
	}

	filters := app.FilterCollection{}
	for _, buf := range bufs {
		var filter *app.Filter
		if err := json.Unmarshal(buf, &filter); err != nil {
			return nil, err
		}
		filters = append(filters, filter)
	}
	return &filters, nil
}
