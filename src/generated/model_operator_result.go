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

type OperatorResult struct {
	// Name of the used operator
	Operator string `json:"operator,omitempty"`
	// Type of the PII entity
	EntityType string `json:"entity_type"`
	// Start index of the changed text
	Start int32 `json:"start"`
	// End index in the changed text
	End int32 `json:"end"`
	// The new text returned
	Text string `json:"text,omitempty"`
}