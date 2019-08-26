package smspva

// IsoCountryCode defines the ISO 3166 country codes
type IsoCountryCode struct {
	name string
	code string
}

// IsoCountryCodes contains the list of ISO 3166 country codes
var IsoCountryCodes = map[string]IsoCountryCode{
	"MD": {
		name: "Moldova, Republic of",
		code: "MD",
	},
	"PL": {
		name: "Poland",
		code: "PL",
	},
	"RO": {
		name: "Romania",
		code: "RO",
	},
	"UA": {
		name: "Ukraine",
		code: "UA",
	},
	"US": {
		name: "United States",
		code: "US",
	},
}
