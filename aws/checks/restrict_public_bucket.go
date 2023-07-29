package checks

import (
	"crosscheck/general"
	"fmt"

	crossplanev1beta1 "github.com/crossplane-contrib/provider-aws/apis/s3/v1beta1"
)

// Ensure S3RestrictPublicBucketsCheck implements the Check interface
var _ Check = (*S3RestrictPublicBucketsCheck)(nil)

// S3RestrictPublicBucketsCheck is a check for S3 bucket restriction
type S3RestrictPublicBucketsCheck struct {
	general.BaseResourceCheck
}

// NewS3RestrictPublicBucketsCheck creates a new S3RestrictPublicBucketsCheck
func NewS3RestrictPublicBucketsCheck() Check {
	return &S3RestrictPublicBucketsCheck{
		BaseResourceCheck: general.BaseResourceCheck{
			Name:               "Ensure S3 bucket has 'restrict_public_bucket' enabled",
			ID:                 "CKV_AWS_56",
			Categories:         []general.CheckCategory{general.GENERAL_SECURITY},
			SupportedResources: []string{"aws_s3_bucket_public_access_block"},
			MissingBlockResult: general.FAILED,
		},
	}
}

// GetInspectedKey returns the key to be inspected
func (s *S3RestrictPublicBucketsCheck) GetInspectedKey() string {
	return "restrict_public_buckets"
}

func (s *S3RestrictPublicBucketsCheck) ScanResourceConf(bucket *crossplanev1beta1.Bucket) (general.CheckResult, error) {
	// Check if "restrict_public_buckets" is set to true
	// fmt.Println("Checking if restrict public buckets is set to true")
	// fmt.Println("The value of RestrictPublicBuckets", bucket.Spec.ForProvider.PublicAccessBlockConfiguration.RestrictPublicBuckets)
	if bucket.Spec.ForProvider.PublicAccessBlockConfiguration.RestrictPublicBuckets != nil && *bucket.Spec.ForProvider.PublicAccessBlockConfiguration.RestrictPublicBuckets {
		// If "restrict_public_buckets" is set to true, the check passes
		// fmt.Println("Restrict public buckets is set to true")
		return general.PASSED, nil
	}

	// If "restrict_public_buckets" is not set to true, the check fails
	return general.FAILED, fmt.Errorf("'restrict_public_buckets' should be set to true")
}
