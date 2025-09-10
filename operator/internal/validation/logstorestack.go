package validation

import (
	"context"
	"fmt"
	"time"

	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/util/validation/field"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"

	logstorev1 "example.com/acme/logstore/operator/api/logstore/v1"
	"example.com/acme/logstore/operator/internal/validation/openshift"
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
func (v *LogstoreStackValidator) ValidateCreate(ctx context.Context, obj runtime.Object) (admission.Warnings, error) {
	return v.validate(ctx, obj)
}

// ValidateUpdate implements admission.CustomValidator.
func (v *LogstoreStackValidator) ValidateUpdate(ctx context.Context, _, newObj runtime.Object) (admission.Warnings, error) {
	return v.validate(ctx, newObj)
}

// ValidateDelete implements admission.CustomValidator.
func (v *LogstoreStackValidator) ValidateDelete(_ context.Context, _ runtime.Object) (admission.Warnings, error) {
	// No validation on delete
	return nil, nil
}

func (v *LogstoreStackValidator) validate(ctx context.Context, obj runtime.Object) (admission.Warnings, error) {
	stack, ok := obj.(*logstorev1.LogstoreStack)
	if !ok {
		return nil, apierrors.NewBadRequest(fmt.Sprintf("object is not of type LogstoreStack: %t", obj))
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

	errors = v.validateReplicationSpec(stack.Spec)
	if len(errors) != 0 {
		allErrs = append(allErrs, errors...)
	}

	errors = v.validateHashRingSpec(stack.Spec)
	if len(errors) != 0 {
		allErrs = append(allErrs, errors...)
	}

	if stack.Spec.Limits != nil {
		if (stack.Spec.Limits.Global != nil && stack.Spec.Limits.Global.OTLP != nil) ||
			len(stack.Spec.Limits.Tenants) > 0 {
			// Only need to validate custom OTLP configuration
			allErrs = append(allErrs, v.validateOTLPConfiguration(&stack.Spec)...)
		}
	}

	if v.ExtendedValidator != nil {
		allErrs = append(allErrs, v.ExtendedValidator(ctx, stack)...)
	}

	if len(allErrs) == 0 {
		return nil, nil
	}

	return nil, apierrors.NewInvalid(
		schema.GroupKind{Group: "logstore.acme.com", Kind: "LogstoreStack"},
		stack.Name,
		allErrs,
	)
}

func (v *LogstoreStackValidator) validateOTLPConfiguration(spec *logstorev1.LogstoreStackSpec) field.ErrorList {
	if spec.Tenants == nil {
		return nil
	}

	if spec.Tenants.Mode == logstorev1.OpenshiftLogging {
		// This tenancy mode always provides stream labels
		return openshift.ValidateOTLPInvalidDrop(spec)
	}

	if spec.Tenants.Mode == logstorev1.OpenshiftNetwork {
		// No validation defined for openshift-network tenancy mode
		// TODO can we define a validation for this mode?
		return nil
	}

	if spec.Limits == nil {
		return nil
	}

	hasGlobalStreamLabels := false
	errList := field.ErrorList{}
	var globalOtlp *logstorev1.OTLPSpec
	if spec.Limits.Global != nil && spec.Limits.Global.OTLP != nil {
		globalOtlp = spec.Limits.Global.OTLP

		hasGlobalStreamLabels = v.hasOTLPStreamLabel(globalOtlp)
		errList = append(errList, v.checkOTLPInvalidDrop(field.NewPath("spec", "limits", "global", "otlp"), globalOtlp, nil)...)
	}

	if !hasGlobalStreamLabels && spec.Limits.Tenants == nil {
		// No tenant config and no global stream labels -> error
		errList = append(errList, field.Invalid(
			field.NewPath("spec", "limits", "global", "otlp", "streamLabels", "resourceAttributes"),
			nil,
			logstorev1.ErrOTLPGlobalNoStreamLabel.Error(),
		))
	}

	errList = append(errList, v.validateOTLPTenantConfiguration(spec, globalOtlp, hasGlobalStreamLabels)...)
	return errList
}

func (v *LogstoreStackValidator) validateOTLPTenantConfiguration(spec *logstorev1.LogstoreStackSpec, globalOtlp *logstorev1.OTLPSpec, hasGlobalStreamLabel bool) (errList field.ErrorList) {
	if spec.Limits == nil || spec.Limits.Tenants == nil {
		return nil
	}

	errList = field.ErrorList{}
	for _, tenant := range spec.Tenants.Authentication {
		tenantName := tenant.TenantName
		tenantLimits, ok := spec.Limits.Tenants[tenantName]
		if !ok || tenantLimits.OTLP == nil {
			if !hasGlobalStreamLabel {
				// No tenant limits defined and no global stream labels -> error
				errList = append(errList, field.Invalid(
					field.NewPath("spec", "limits", "tenants", tenantName, "otlp"),
					nil,
					logstorev1.ErrOTLPTenantMissing.Error(),
				))
			}

			continue
		}

		if !hasGlobalStreamLabel && !v.hasOTLPStreamLabel(tenantLimits.OTLP) {
			errList = append(errList, field.Invalid(
				field.NewPath("spec", "limits", "tenants", tenantName, "otlp", "streamLabels", "resourceAttributes"),
				nil,
				logstorev1.ErrOTLPTenantNoStreamLabel.Error(),
			))
		}

		errList = append(errList, v.checkOTLPInvalidDrop(field.NewPath("spec", "limits", "tenants", tenantName, "otlp"), tenantLimits.OTLP, globalOtlp)...)
	}

	return errList
}

func (v *LogstoreStackValidator) hasOTLPStreamLabel(otlp *logstorev1.OTLPSpec) bool {
	if otlp == nil {
		return false
	}

	if otlp.StreamLabels == nil {
		return false
	}

	return len(otlp.StreamLabels.ResourceAttributes) > 0
}

func (v *LogstoreStackValidator) checkOTLPInvalidDrop(basePath *field.Path, otlp, inheritedOtlp *logstorev1.OTLPSpec) field.ErrorList {
	if otlp.Drop == nil {
		return nil
	}

	errList := field.ErrorList{}
	streamAttributes := [][]logstorev1.OTLPAttributeReference{}
	if streamLabels := otlp.StreamLabels; streamLabels != nil {
		streamAttributes = append(streamAttributes, streamLabels.ResourceAttributes)
	}
	if inheritedOtlp != nil && inheritedOtlp.StreamLabels != nil && len(inheritedOtlp.StreamLabels.ResourceAttributes) > 0 {
		streamAttributes = append(streamAttributes, inheritedOtlp.StreamLabels.ResourceAttributes)
	}
	errList = append(errList, v.checkOTLPInvalidDropReference(
		basePath.Child("drop", "resourceAttributes"),
		otlp.Drop.ResourceAttributes,
		streamAttributes,
	)...)

	return errList
}

func (v *LogstoreStackValidator) checkOTLPInvalidDropReference(basePath *field.Path, dropList []logstorev1.OTLPAttributeReference, keepLists [][]logstorev1.OTLPAttributeReference) field.ErrorList {
	if len(dropList) == 0 {
		return nil
	}

	if len(keepLists) == 0 {
		return nil
	}

	attributeNames := map[string]bool{}
	errList := field.ErrorList{}

	for _, keeps := range keepLists {
		for _, attr := range keeps {
			if attr.Regex {
				// skip regular expressions for this check
				continue
			}
			attributeNames[attr.Name] = true
		}
	}

	for i, attr := range dropList {
		if attr.Regex {
			continue
		}

		if !attributeNames[attr.Name] {
			continue
		}

		errList = append(errList, field.Invalid(
			basePath.Index(i),
			attr.Name,
			logstorev1.ErrOTLPInvalidDrop.Error(),
		))
	}

	return errList
}

func (v *LogstoreStackValidator) validateHashRingSpec(s logstorev1.LogstoreStackSpec) field.ErrorList {
	if s.HashRing == nil {
		return nil
	}

	if s.HashRing.MemberList == nil {
		return nil
	}

	if s.HashRing.MemberList.EnableIPv6 && s.HashRing.MemberList.InstanceAddrType == logstorev1.InstanceAddrDefault {
		return field.ErrorList{
			field.Invalid(
				field.NewPath("spec", "hashRing", "memberlist", "instanceAddrType"),
				s.HashRing.MemberList.InstanceAddrType,
				logstorev1.ErrIPv6InstanceAddrTypeNotAllowed.Error(),
			),
		}
	}

	return nil
}

func (v *LogstoreStackValidator) validateReplicationSpec(stack logstorev1.LogstoreStackSpec) field.ErrorList {
	if stack.Replication == nil {
		return nil
	}

	var allErrs field.ErrorList

	// nolint:staticcheck
	if stack.Replication != nil && stack.ReplicationFactor > 0 {
		allErrs = append(allErrs, field.Invalid(
			field.NewPath("spec", "replicationFactor"),
			stack.ReplicationFactor,
			logstorev1.ErrReplicationSpecConflict.Error(),
		))

		return allErrs
	}

	return nil
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
