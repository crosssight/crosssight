package checks

import (
	"crosscheck/general"
	"fmt"

	crossplanev1beta1 "github.com/crossplane-contrib/provider-aws/apis/s3/v1beta1"
)

// Ensure S3RequireLoggingCheck implements the Check interface
var _ Check = (*S3RequireLoggingCheck)(nil)

// S3RequireLoggingCheck is a check for S3 bucket logging requirement
type S3RequireLoggingCheck struct {
	general.BaseResourceCheck
}

// NewS3RequireLoggingCheck creates a new S3RequireLoggingCheck
func NewS3RequireLoggingCheck() Check {
	return &S3RequireLoggingCheck{
		BaseResourceCheck: general.BaseResourceCheck{
			Name:               "Ensure S3 bucket has logging enabled",
			ID:                 "CKV_AWS_18",
			Categories:         []general.CheckCategory{general.GENERAL_SECURITY},
			SupportedResources: []string{"aws_s3_bucket"},
			MissingBlockResult: general.FAILED,
		},
	}
}

// GetInspectedKey returns the key to be inspected
func (s *S3RequireLoggingCheck) GetInspectedKey() string {
	return "" // Change this key to the appropriate key in the S3 bucket resource
}

func (s *S3RequireLoggingCheck) ScanResourceConf(bucket *crossplanev1beta1.Bucket) (general.CheckResult, error) {
	// Check if "LoggingConfiguration" exists
	if bucket.Spec.ForProvider.LoggingConfiguration != nil {
		// If LoggingConfiguration exists, the check passes
		return general.PASSED, nil
	}
	// If LoggingConfiguration is not set, the check fails
	return general.FAILED, fmt.Errorf("'LoggingConfiguration' should be set in the YAML file")
}
