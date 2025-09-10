package storage

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	logstorev1 "example.com/acme/logstore/operator/api/logstore/v1"
)

func TestBuildSchemaConfig_NoSchemas(t *testing.T) {
	spec := logstorev1.ObjectStorageSpec{}
	status := logstorev1.LogstoreStackStorageStatus{}

	expected, err := BuildSchemaConfig(time.Now().UTC(), spec, status)

	require.Error(t, err)
	require.Nil(t, expected)
}

func TestBuildSchemaConfig_AddSchema_NoStatuses(t *testing.T) {
	spec := logstorev1.ObjectStorageSpec{
		Schemas: []logstorev1.ObjectStorageSchema{
			{
				Version:       logstorev1.ObjectStorageSchemaV11,
				EffectiveDate: "2020-10-01",
			},
		},
	}
	status := logstorev1.LogstoreStackStorageStatus{}

	actual, err := BuildSchemaConfig(time.Now().UTC(), spec, status)
	expected := []logstorev1.ObjectStorageSchema{
		{
			Version:       logstorev1.ObjectStorageSchemaV11,
			EffectiveDate: "2020-10-01",
		},
	}

	require.NoError(t, err)
	require.Equal(t, expected, actual)
}

func TestBuildSchemaConfig_AddSchema_WithStatuses_WithValidDate(t *testing.T) {
	utcTime := time.Date(2021, 9, 1, 0, 0, 0, 0, time.UTC)
	spec := logstorev1.ObjectStorageSpec{
		Schemas: []logstorev1.ObjectStorageSchema{
			{
				Version:       logstorev1.ObjectStorageSchemaV11,
				EffectiveDate: "2020-10-01",
			},
			{
				Version:       logstorev1.ObjectStorageSchemaV12,
				EffectiveDate: "2021-10-01",
			},
		},
	}
	status := logstorev1.LogstoreStackStorageStatus{
		Schemas: []logstorev1.ObjectStorageSchema{
			{
				Version:       logstorev1.ObjectStorageSchemaV11,
				EffectiveDate: "2020-10-01",
			},
		},
	}

	actual, err := BuildSchemaConfig(utcTime, spec, status)
	expected := []logstorev1.ObjectStorageSchema{
		{
			Version:       logstorev1.ObjectStorageSchemaV11,
			EffectiveDate: "2020-10-01",
		},
		{
			Version:       logstorev1.ObjectStorageSchemaV12,
			EffectiveDate: "2021-10-01",
		},
	}

	require.NoError(t, err)
	require.Equal(t, expected, actual)
}

func TestBuildSchemaConfig_AddSchema_WithStatuses_WithInvalidDate(t *testing.T) {
	utcTime := time.Date(2021, 10, 1, 0, 0, 0, 0, time.UTC)
	updateWindow := utcTime.Add(logstorev1.StorageSchemaUpdateBuffer).Format(logstorev1.StorageSchemaEffectiveDateFormat)
	spec := logstorev1.ObjectStorageSpec{
		Schemas: []logstorev1.ObjectStorageSchema{
			{
				Version:       logstorev1.ObjectStorageSchemaV11,
				EffectiveDate: "2020-10-01",
			},
			{
				Version:       logstorev1.ObjectStorageSchemaV12,
				EffectiveDate: logstorev1.StorageSchemaEffectiveDate(updateWindow),
			},
		},
	}
	status := logstorev1.LogstoreStackStorageStatus{
		Schemas: []logstorev1.ObjectStorageSchema{
			{
				Version:       logstorev1.ObjectStorageSchemaV11,
				EffectiveDate: "2020-10-01",
			},
		},
	}

	expected, err := BuildSchemaConfig(utcTime, spec, status)

	require.Error(t, err)
	require.Nil(t, expected)
}

func TestBuildSchemas(t *testing.T) {
	schemas := []logstorev1.ObjectStorageSchema{
		{
			Version:       logstorev1.ObjectStorageSchemaV12,
			EffectiveDate: "2021-11-01",
		},
		{
			Version:       logstorev1.ObjectStorageSchemaV12,
			EffectiveDate: "2021-06-01",
		},
		{
			Version:       logstorev1.ObjectStorageSchemaV11,
			EffectiveDate: "2020-10-01",
		},
		{
			Version:       logstorev1.ObjectStorageSchemaV12,
			EffectiveDate: "2021-12-01",
		},
		{
			Version:       logstorev1.ObjectStorageSchemaV11,
			EffectiveDate: "2021-10-01",
		},
		{
			Version:       logstorev1.ObjectStorageSchemaV11,
			EffectiveDate: "2021-02-01",
		},
	}

	expected := []logstorev1.ObjectStorageSchema{
		{
			Version:       logstorev1.ObjectStorageSchemaV11,
			EffectiveDate: "2020-10-01",
		},
		{
			Version:       logstorev1.ObjectStorageSchemaV12,
			EffectiveDate: "2021-06-01",
		},
		{
			Version:       logstorev1.ObjectStorageSchemaV11,
			EffectiveDate: "2021-10-01",
		},
		{
			Version:       logstorev1.ObjectStorageSchemaV12,
			EffectiveDate: "2021-11-01",
		},
	}
	actual := buildSchemas(schemas)

	require.Equal(t, expected, actual)
}

func TestReduceSortedSchemas(t *testing.T) {
	schemas := []logstorev1.ObjectStorageSchema{
		{
			Version:       logstorev1.ObjectStorageSchemaV11,
			EffectiveDate: "2020-10-01",
		},
		{
			Version:       logstorev1.ObjectStorageSchemaV11,
			EffectiveDate: "2021-02-01",
		},
		{
			Version:       logstorev1.ObjectStorageSchemaV12,
			EffectiveDate: "2021-06-01",
		},
		{
			Version:       logstorev1.ObjectStorageSchemaV11,
			EffectiveDate: "2021-10-01",
		},
		{
			Version:       logstorev1.ObjectStorageSchemaV12,
			EffectiveDate: "2021-11-01",
		},
		{
			Version:       logstorev1.ObjectStorageSchemaV12,
			EffectiveDate: "2021-12-01",
		},
	}

	expected := []logstorev1.ObjectStorageSchema{
		{
			Version:       logstorev1.ObjectStorageSchemaV11,
			EffectiveDate: "2020-10-01",
		},
		{
			Version:       logstorev1.ObjectStorageSchemaV12,
			EffectiveDate: "2021-06-01",
		},
		{
			Version:       logstorev1.ObjectStorageSchemaV11,
			EffectiveDate: "2021-10-01",
		},
		{
			Version:       logstorev1.ObjectStorageSchemaV12,
			EffectiveDate: "2021-11-01",
		},
	}
	actual := reduceSortedSchemas(schemas)

	require.Equal(t, expected, actual)
}
