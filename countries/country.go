package countries

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type Country struct {
	Name          string // official name in english
	Cca2          string // country code on 2 digits
	Cca3          string // country code on 3 digits
	Region        string
	Subregion     string
	Currency      string //Note countries can have multiple currencies, using only first one to abide by struct provided
	CurrencySymbo string
}

type RawCountry struct {
	Name       Name       `json:"name"`
	Cca2       string     `json:"cca2"`
	Cca3       string     `json:"cca3"`
	Currencies Currencies `json:"currencies"`
	Region     string     `json:"region"`
	Subregion  string     `json:"subregion"`
}

type Name struct {
	Common     string      `json:"common"`
	Official   string      `json:"official"`
	NativeName NativeNames `json:"nativeName"`
}

type NativeNames map[string]NativeName

type NativeName struct {
	Official string `json:"official"`
	Common   string `json:"common"`
}

type Currencies map[string]Currency

type Currency struct {
	Name   string `json:"name"`
	Symbol string `json:"symbol"`
}

var baseUrl = "https://restcountries.com/v3.1/"
var filter = "fields=name,cca2,cca3,region,subregion,currencies"

func FetchByRegion(region string) ([]Country, error) {

	endpoint := baseUrl + "region/" + url.PathEscape(region) + "?" + filter
	result, err := fetch(endpoint)
	if err != nil {
		return nil, err
	}

	countries := parseCountries(result)

	return countries, nil

}

func FetchBySubRegion(subregion string) ([]Country, error) {

	endpoint := baseUrl + "subregion/" + url.PathEscape(subregion) + "?" + filter
	result, err := fetch(endpoint)
	if err != nil {
		return nil, err
	}

	countries := parseCountries(result)

	return countries, nil

}

func fetch(endpoint string) ([]RawCountry, error) {
	resp, err := http.Get(endpoint)

	if err != nil {
		return nil, err
	}

	responseData, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var rawCountries []RawCountry

	err = json.Unmarshal([]byte(responseData), &rawCountries)
	if err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}

	return rawCountries, nil

}

func parseCountries(rawCountries []RawCountry) []Country {

	countries := []Country{}

	for _, rawCountry := range rawCountries {
		var currencyName, currencySymbol string
		for _, currency := range rawCountry.Currencies {
			currencyName = currency.Name
			currencySymbol = currency.Symbol
			break
		}

		country := Country{
			rawCountry.Name.Official,
			rawCountry.Cca2,
			rawCountry.Cca3,
			rawCountry.Region,
			rawCountry.Subregion,
			currencyName,
			currencySymbol,
		}

		countries = append(countries, country)
	}

	return countries

}
