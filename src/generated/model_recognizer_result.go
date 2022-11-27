/*
 * Presidio
 *
 * Context aware, pluggable and customizable PII anonymization service for text and images.
 *
 * API version: 2.0
 * Contact: presidio@microsoft.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type RecognizerResult struct {
	// Where the PII starts
	Start int32 `json:"start"`
	// Where the PII ends
	End int32 `json:"end"`
	// The PII detection score
	Score float64 `json:"score"`
	EntityType string `json:"entity_type"`
	RecognitionMetadata *RecognizedMetadata `json:"recognition_metadata,omitempty"`
}
