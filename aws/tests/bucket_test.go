package tests

import (
	"crosscheck/aws/checks"
	"crosscheck/general"
	"fmt"
	"io"
	"os"
	"testing"

	crossplanev1beta1 "github.com/crossplane-contrib/provider-aws/apis/s3/v1beta1"

	"github.com/ghodss/yaml"
	"github.com/stretchr/testify/assert"
)

func TestBucketParsing(t *testing.T) {
	file, err := os.Open("demo_bucket.yaml")
	if err != nil {
		t.Fatalf("failed to load demo_bucket.yaml: %v", err)
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		t.Fatalf("failed to read demo_bucket.yaml: %v", err)
	}

	var bucket crossplanev1beta1.Bucket
	err = yaml.Unmarshal(data, &bucket)
	if err != nil {
		t.Fatalf("failed to unmarshal YAML: %v", err)
	}

	check := checks.NewS3RestrictPublicBucketsCheck()
	fmt.Println(check.GetInspectedKey())

	// You can add additional checks here based on what you expect from the parsed bucket
	assert.NotNil(t, bucket.Spec.ForProvider.ACL)
	assert.Equal(t, "example-bucket", bucket.ObjectMeta.Name)
}

func TestScanResourceConf(t *testing.T) {
	// Initialize the check
	check := checks.NewS3RestrictPublicBucketsCheck()

	file, err := os.Open("demo_bucket.yaml")
	if err != nil {
		t.Fatalf("failed to load demo_bucket.yaml: %v", err)
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		t.Fatalf("failed to read demo_bucket.yaml: %v", err)
	}

	var bucket crossplanev1beta1.Bucket
	err = yaml.Unmarshal(data, &bucket)
	if err != nil {
		t.Fatalf("failed to unmarshal YAML: %v", err)
	}

	// Call ScanResourceConf
	result, err := check.ScanResourceConf(&bucket)
	if err != nil {
		t.Fatalf("ScanResourceConf failed: %v", err)
	}

	// Assert the result
	if result != general.PASSED {
		t.Errorf("Expected PASSED, got %v", result)
	}

	assert.Equal(t, general.PASSED, result)
}
