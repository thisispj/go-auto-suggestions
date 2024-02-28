package service

import (
	"encoding/json"

	"github.com/thisispj/autosuggestions/cache"
	"github.com/thisispj/autosuggestions/store"
)

type CountryService struct {
	countryStore *store.CountryStore
	cache        *cache.RedisCache
}

func NewCountryService(cache *cache.RedisCache, countryStore *store.CountryStore) *CountryService {
	return &CountryService{
		cache:        cache,
		countryStore: countryStore,
	}
}

func (s *CountryService) GetAutocompleteSuggestions(query string) (store.CountryData, error) {
	countries, err := s.cache.Get(query)
	if err != nil {
		suggestions := s.countryStore.FindByPrefix(query)
		if len(suggestions.Countries) == 0 {
			if err := s.cache.Set(query, suggestions); err != nil {
				return store.CountryData{}, err
			}
		}
		return suggestions, nil
	}
	var result store.CountryData
	err = json.Unmarshal([]byte(countries), &result)
	if err != nil {
		return store.CountryData{}, err
	}
	return result, nil
}
