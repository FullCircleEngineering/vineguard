package cache

import (
	"vineguard/internal/ttndb"
)

type repo struct {
	cache *tempCache
}

type tempCache struct {
	// stand-in for our datastore
	registrations map[string][]map[string]string // e.g. {"jake@gmail.com: [{"devId": "foo-device", "phone": "510-444-222"}, ...], ... }
	data          map[string]ttndb.LSB50Msgs
}

func NewRepo() *repo {
	cache := &tempCache{
		registrations: map[string][]map[string]string{},
		data:          map[string]ttndb.LSB50Msgs{},
	}
	return &repo{cache: cache}
}
func (r *repo) AddData(devId string, data ttndb.LSB50Msgs) error {
	// overwrite and update our cache
	r.cache.data[devId] = data
	return nil
}

func (r *repo) AddRegistration(userEmail string, reg map[string]string) error {
	// takes in a registration of device to phone number.  e.g. {"devId": "foo-device", "phone": "510-444-222"}
	if existingRegs, ok := r.cache.registrations[userEmail]; ok {
		r.cache.registrations[userEmail] = append(existingRegs, reg)
	} else {
		r.cache.registrations[userEmail] = []map[string]string{reg}
	}
	return nil
}

func (r *repo) GetData(devId string) (ttndb.LSB50Msgs, error) {
	return r.cache.data[devId], nil
}

func (r *repo) GetUserRegistrations(userEmail string) ([]map[string]string, error) {
	if regs, ok := r.cache.registrations[userEmail]; ok {
		return regs, nil
	} else {
		return []map[string]string{}, nil
	}
}
