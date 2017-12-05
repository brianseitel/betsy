package betsy

import (
	"encoding/json"
	"fmt"
	"log"
)

// Features ...
var Features *FeatureFlags

// FeatureFlags ...
type FeatureFlags struct {
	Features map[string]Rule
}

// NewFeatureFlags ...
func NewFeatureFlags() *FeatureFlags {
	flags := make(map[string]Rule)
	return &FeatureFlags{
		Features: flags,
	}
}

// GetFeatureFlags ...
func GetFeatureFlags() FeatureFlags {
	if Features == nil {
		Features = NewFeatureFlags()
	}

	return *Features
}

// Add a feature flag to the system with a callback to determine truthiness
func (ff *FeatureFlags) Add(name string, callback Rule) {
	ff.Features[name] = callback
}

// Allowed detects whether a value passes the feature flag or not
func (ff FeatureFlags) Allowed(name string, value interface{}) bool {
	rule, ok := ff.Features[name]

	log.Printf("Testing rule %s with value `%v`\n", name, value)
	if ok && rule.Run(value) {
		return true
	}
	return false
}

// Denied detects whether a value fails the feature flag or not
func (ff FeatureFlags) Denied(name string, value interface{}) bool {
	return !ff.Allowed(name, value)
}

// List all flag values
func (ff FeatureFlags) List() {
	j, _ := json.MarshalIndent(ff.Features, "", "    ")

	fmt.Println(string(j))
}
