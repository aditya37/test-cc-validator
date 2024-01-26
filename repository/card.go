package repository

import (
	"errors"
	"sync"

	_interface "github.com/aditya37/test-cc-validator/interface"
	"github.com/aditya37/test-cc-validator/model"
)

// dummy database
type cardStore struct {
	// dummy db connection
	db map[string]model.MSTCard

	// for avoid race condtion on insert data to map
	sync.RWMutex
}

func NewCardReaderWriter() _interface.CCWriterReader {
	return &cardStore{
		db:      make(map[string]model.MSTCard),
		RWMutex: sync.RWMutex{},
	}
}

var ErrDuplicate = errors.New("duplicate number")

func (cs *cardStore) InsertCardNumber(value, network string) error {
	if _, ok := cs.db[value]; ok {
		return ErrDuplicate
	}

	cs.RWMutex.Lock()
	defer cs.RWMutex.Unlock()
	cs.db[value] = model.MSTCard{
		Number:      value,
		CardNetwork: network,
	}
	return nil
}
