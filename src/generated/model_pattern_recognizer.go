/*
 * Presidio
 *
 * Context aware, pluggable and customizable PII anonymization service for text and images.
 *
 * API version: 2.0
 * Contact: presidio@microsoft.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package generated

// A regular expressions or deny-list based recognizer
type PatternRecognizer struct {
	// Name of recognizer
	Name string `json:"name,omitempty"`
	// Language code supported by this recognizer
	SupportedLanguage string `json:"supported_language,omitempty"`
	// List of type Pattern containing regex expressions with additional metadata.
	Patterns []Pattern `json:"patterns,omitempty"`
	// List of words to be returned as PII if found.
	DenyList []string `json:"deny_list,omitempty"`
	// List of words to be used to increase confidence if found in the vicinity of detected entities.
	Context []string `json:"context,omitempty"`
	// The name of entity this ad hoc recognizer detects
	SupportedEntity string `json:"supported_entity,omitempty"`
}
