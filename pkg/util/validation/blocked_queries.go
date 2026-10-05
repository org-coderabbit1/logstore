package validation

import "example.com/acme/kit/flagext"

type BlockedQuery struct {
	Pattern string                 `yaml:"pattern"`
	Regex   bool                   `yaml:"regex"`
	Types   flagext.StringSliceCSV `yaml:"types"`
}
