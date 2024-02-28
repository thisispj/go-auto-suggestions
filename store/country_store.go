package store

import (
	"encoding/json"
	"os"
	"strings"
	"sync"
)

type CountryStore struct {
	data CountryData
	sync.RWMutex
}

type CountryData struct {
	Countries []string `json:"countries"`
}

func NewCountryStore(filePath string) (*CountryStore, error) {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var data CountryData
	if err := json.Unmarshal(content, &data); err != nil {
		return nil, err
	}

	return &CountryStore{
		data: data,
	}, nil
}

func (r *CountryStore) FindByPrefix(query string) CountryData {
	r.RLock()
	defer r.RUnlock()

	var countries []string
	for _, country := range r.data.Countries {
		if strings.HasPrefix(strings.ToLower(country), strings.ToLower(query)) {
			countries = append(countries, country)
		}
	}

	return CountryData{Countries: countries}
}
