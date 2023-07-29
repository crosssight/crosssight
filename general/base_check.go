package general

import (
	"errors"
)

// CheckResult enum
type CheckResult int

const (
	PASSED CheckResult = iota + 1
	FAILED
	UNKNOWN
)

// ResourceCheck is the interface for checks
type ResourceCheck interface {
	GetInspectedKey() string
	GetExpectedValues() []interface{}
	GetExpectedValue() interface{}
	ScanResourceConf(conf map[string]interface{}) (CheckResult, error)
}

// BaseResourceCheck is the struct that contains fields common to all checks
type BaseResourceCheck struct {
	Name               string
	ID                 string
	Categories         []CheckCategory
	SupportedResources []string
	MissingBlockResult CheckResult
}

// GetInspectedKey provides a default implementation
func (b *BaseResourceCheck) GetInspectedKey() string {
	return ""
}

// GetExpectedValues provides a default implementation
func (b *BaseResourceCheck) GetExpectedValues() []interface{} {
	return []interface{}{b.GetExpectedValue()}
}

// GetExpectedValue provides a default implementation
func (b *BaseResourceCheck) GetExpectedValue() interface{} {
	return true
}

// ScanResourceConf should be implemented by each concrete check
func (b *BaseResourceCheck) ScanResourceConf(conf map[string]interface{}) (CheckResult, error) {
	return UNKNOWN, errors.New("ScanResourceConf not implemented")
}
