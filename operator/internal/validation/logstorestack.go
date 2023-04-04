package validation

import (
	"context"
	"fmt"
	"time"

	logstorev1 "example.com/acme/logstore/operator/apis/logstore/v1"

	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/util/validation/field"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
)

// objectStorageSchemaMap defines the type for mapping a schema version with a date
type objectStorageSchemaMap map[logstorev1.StorageSchemaEffectiveDate]logstorev1.ObjectStorageSchemaVersion

var _ admission.CustomValidator = &LogstoreStackValidator{}

// LogstoreStackValidator implements a custom validator for LogstoreStack resources.
type LogstoreStackValidator struct {
	ExtendedValidator func(context.Context, *logstorev1.LogstoreStack) field.ErrorList
}

// SetupWebhookWithManager registers the LogstoreStackValidator as a validating webhook
// with the controller-runtime manager or returns an error.
func (v *LogstoreStackValidator) SetupWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).
		For(&logstorev1.LogstoreStack{}).
		WithValidator(v).
		Complete()
}

// ValidateCreate implements admission.CustomValidator.
func (v *LogstoreStackValidator) ValidateCreate(ctx context.Context, obj runtime.Object) error {
	return v.validate(ctx, obj)
}

// ValidateUpdate implements admission.CustomValidator.
func (v *LogstoreStackValidator) ValidateUpdate(ctx context.Context, _, newObj runtime.Object) error {
	return v.validate(ctx, newObj)
}

// ValidateDelete implements admission.CustomValidator.
func (v *LogstoreStackValidator) ValidateDelete(_ context.Context, _ runtime.Object) error {
	// No validation on delete
	return nil
}

func (v *LogstoreStackValidator) validate(ctx context.Context, obj runtime.Object) error {
	stack, ok := obj.(*logstorev1.LogstoreStack)
	if !ok {
		return apierrors.NewBadRequest(fmt.Sprintf("object is not of type LogstoreStack: %t", obj))
	}

	var allErrs field.ErrorList

	storageStatus := logstorev1.LogstoreStackStorageStatus{}
	if stack != nil {
		storageStatus = stack.Status.Storage
	}

	errors := ValidateSchemas(&stack.Spec.Storage, time.Now().UTC(), storageStatus)
	if len(errors) != 0 {
		allErrs = append(allErrs, errors...)
	}

	if v.ExtendedValidator != nil {
		allErrs = append(allErrs, v.ExtendedValidator(ctx, stack)...)
	}

	if len(allErrs) == 0 {
		return nil
	}

	return apierrors.NewInvalid(
		schema.GroupKind{Group: "logstore.acme.com", Kind: "LogstoreStack"},
		stack.Name,
		allErrs,
	)
}

// ValidateSchemas ensures that the schemas are in a valid format
func ValidateSchemas(v *logstorev1.ObjectStorageSpec, utcTime time.Time, status logstorev1.LogstoreStackStorageStatus) field.ErrorList {
	var allErrs field.ErrorList

	appliedSchemasFound := 0
	containsValidStartDate := false
	found := make(map[logstorev1.StorageSchemaEffectiveDate]bool)

	cutoff := utcTime.Add(logstorev1.StorageSchemaUpdateBuffer)
	appliedSchemas := buildAppliedSchemaMap(status.Schemas, cutoff)

	for i, sc := range v.Schemas {
		if found[sc.EffectiveDate] {
			allErrs = append(allErrs, field.Invalid(
				field.NewPath("spec").Child("storage").Child("schemas").Index(i).Child("effectiveDate"),
				sc.EffectiveDate,
				logstorev1.ErrEffectiveDatesNotUnique.Error(),
			))
		}

		found[sc.EffectiveDate] = true

		date, err := sc.EffectiveDate.UTCTime()
		if err != nil {
			allErrs = append(allErrs, field.Invalid(
				field.NewPath("spec").Child("storage").Child("schemas").Index(i).Child("effectiveDate"),
				sc.EffectiveDate,
				logstorev1.ErrParseEffectiveDates.Error(),
			))
		}

		if date.Before(cutoff) {
			containsValidStartDate = true
		}

		// No statuses to compare against or this is a new schema which will be added.
		if len(appliedSchemas) == 0 || date.After(cutoff) {
			continue
		}

		appliedSchemaVersion, ok := appliedSchemas[sc.EffectiveDate]

		if !ok {
			allErrs = append(allErrs, field.Invalid(
				field.NewPath("spec").Child("storage").Child("schemas").Index(i),
				sc,
				logstorev1.ErrSchemaRetroactivelyAdded.Error(),
			))
		}

		if ok && appliedSchemaVersion != sc.Version {
			allErrs = append(allErrs, field.Invalid(
				field.NewPath("spec").Child("storage").Child("schemas").Index(i),
				sc,
				logstorev1.ErrSchemaRetroactivelyChanged.Error(),
			))
		}

		appliedSchemasFound++
	}

	if !containsValidStartDate {
		allErrs = append(allErrs, field.Invalid(
			field.NewPath("spec").Child("storage").Child("schemas"),
			v.Schemas,
			logstorev1.ErrMissingValidStartDate.Error(),
		))
	}

	if appliedSchemasFound != len(appliedSchemas) {
		allErrs = append(allErrs, field.Invalid(
			field.NewPath("spec").Child("storage").Child("schemas"),
			v.Schemas,
			logstorev1.ErrSchemaRetroactivelyRemoved.Error(),
		))
	}

	if len(allErrs) == 0 {
		return nil
	}

	return allErrs
}

// buildAppliedSchemaMap creates a map of schemas which occur before the given time
func buildAppliedSchemaMap(schemas []logstorev1.ObjectStorageSchema, effectiveDate time.Time) objectStorageSchemaMap {
	appliedMap := objectStorageSchemaMap{}

	for _, schema := range schemas {
		date, err := schema.EffectiveDate.UTCTime()

		if err == nil && date.Before(effectiveDate) {
			appliedMap[schema.EffectiveDate] = schema.Version
		}
	}

	return appliedMap
}
