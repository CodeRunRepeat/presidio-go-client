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

type DeanonymizeRequest struct {
	// The anonymized text
	Text string `json:"text"`
	// Object where the key is DEFAULT or the ENTITY_TYPE and the value is decrypt since it is the only one supported
	Deanonymizers AnonymizerMap `json:"deanonymizers"`
	// Array of anonymized PIIs
	AnonymizerResults []OperatorResult `json:"anonymizer_results"`
}
