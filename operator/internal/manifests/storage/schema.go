package storage

import (
	"sort"
	"time"

	"github.com/ViaQ/logerr/v2/kverrors"

	logstorev1 "example.com/acme/logstore/operator/api/logstore/v1"
	"example.com/acme/logstore/operator/internal/validation"
)

// BuildSchemaConfig creates a list of schemas to be used to configure
// the storage schemas for the cluster. This method assumes that the following
// validation has been done to the statuses and specs:
//
// 1. All EffectiveDate fields are able to be parsed
// 2. All EffectiveDate fields are unique in their respective list
func BuildSchemaConfig(
	utcTime time.Time,
	spec logstorev1.ObjectStorageSpec,
	status logstorev1.LogstoreStackStorageStatus,
) ([]logstorev1.ObjectStorageSchema, error) {
	if len(spec.Schemas) == 0 {
		return nil, kverrors.New("spec does not contain any schemas")
	}

	errors := validation.ValidateSchemas(&spec, utcTime, status)
	if len(errors) != 0 {
		return nil, kverrors.Wrap(errors[0], "spec contains invalid schema entry")
	}

	schemas := buildSchemas(spec.Schemas)

	return schemas, nil
}

// buildSchemas creates a sorted and reduced list of schemaConfigs
func buildSchemas(schemas []logstorev1.ObjectStorageSchema) []logstorev1.ObjectStorageSchema {
	sortedSchemas := make([]logstorev1.ObjectStorageSchema, len(schemas))
	copy(sortedSchemas, schemas)

	sort.SliceStable(sortedSchemas, func(i, j int) bool {
		iDate, _ := sortedSchemas[i].EffectiveDate.UTCTime()
		jDate, _ := sortedSchemas[j].EffectiveDate.UTCTime()

		return iDate.Before(jDate)
	})

	return reduceSortedSchemas(sortedSchemas)
}

// reduceSortedSchemas returns a list of schemas that have removed redundant entries.
func reduceSortedSchemas(schemas []logstorev1.ObjectStorageSchema) []logstorev1.ObjectStorageSchema {
	version := ""
	reduced := []logstorev1.ObjectStorageSchema{}

	for _, schema := range schemas {
		strSchemaVersion := string(schema.Version)

		if version != strSchemaVersion {
			version = strSchemaVersion
			reduced = append(reduced, schema)
		}
	}

	return reduced
}
