package checks

import (
	"crosscheck/general"

	crossplanev1beta1 "github.com/crossplane-contrib/provider-aws/apis/s3/v1beta1"
)

type Check interface {
	ScanResourceConf(bucket *crossplanev1beta1.Bucket) (general.CheckResult, error)
	GetInspectedKey() string
}
