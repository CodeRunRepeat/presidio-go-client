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

// Replace with an encrypted value
type Encrypt struct {
	// encrypt
	Type_ string `json:"type"`
	// Cryptographic key of length 128, 192 or 256 bits, in a string format
	Key string `json:"key"`
}
