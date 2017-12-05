package betsy

import (
	"errors"
	"log"
)

// Rule ...
type Rule interface {
	Run(input interface{}) bool
}

// NewRule detects the type of input and creates a rule based on it
func NewRule(input interface{}) (Rule, error) {
	switch input.(type) {
	case string:
		return StringRule(input.(string)), nil
	case int:
		return IntRule(input.(int)), nil
	case bool:
		return BoolRule(input.(bool)), nil
	case []int:
		return IntSliceRule(input.([]int)), nil
	case []string:
		return StringSliceRule(input.([]string)), nil
	}

	return NoopRule{}, errors.New("invalid rule type")
}

// NoopRule ...
type NoopRule struct{}

// Run ...
func (r NoopRule) Run(input interface{}) bool {
	return false
}

// StringRule ...
type StringRule string

// Run a String Rule
func (r StringRule) Run(input interface{}) bool {
	switch input.(type) {
	case string:
		return string(r) == input.(string)
	}

	log.Printf("Error] Unexpected type. Expected string, got `%v`", input)
	return false
}

// StringSliceRule ...
type StringSliceRule []string

// Run ...
func (r StringSliceRule) Run(input interface{}) bool {
	switch input.(type) {
	case []string:
		for _, s := range r {
			if s == input.(string) {
				return true
			}
		}
	}

	log.Printf("[Error] Unexpected type. Expected []string, got `%v`", input)
	return false
}

// BoolRule ...
type BoolRule bool

// Run a Bool Rule
func (r BoolRule) Run(input interface{}) bool {
	switch input.(type) {
	case bool:
		return bool(r) == input.(bool)
	}

	log.Printf("Error] Unexpected type. Expected bool, got `%v`", input)
	return false
}

// IntRule ...
type IntRule int

// Run an IntRule
func (r IntRule) Run(input interface{}) bool {
	switch input.(type) {
	case int:
		return int(r) == input.(int)
	case float64:
		return int(r) == int(input.(float64))
	}

	log.Printf("Error] Unexpected type. Expected int or float64, got `%v`", input)
	return false
}

// IntSliceRule ...
type IntSliceRule []int

// Run an IntSliceRule
func (r IntSliceRule) Run(input interface{}) bool {
	switch input.(type) {
	case int:
		for _, i := range r {
			if int(i) == input.(int) {
				return true
			}
		}
	case float64:
		for _, i := range r {
			if int(i) == int(input.(float64)) {
				return true
			}
		}
	}

	log.Printf("Error] Unexpected type. Expected []int or []float64, got `%v`", input)
	return false
}
