package querier

import (
	"example.com/acme/kit/flagext"

	"example.com/acme/logstore/pkg/validation"
)

func DefaultLimitsConfig() validation.Limits {
	limits := validation.Limits{}
	flagext.DefaultValues(&limits)
	return limits
}
