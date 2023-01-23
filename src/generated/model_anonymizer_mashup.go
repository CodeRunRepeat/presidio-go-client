/*
This file, unlike others in this directory, is not generated. Its purpose is to provide with a
type for the AnoynimizeRequest.Anonymizers field that serializes in a compatible way with
what the Presidio API expects.
*/
package generated

type AnonymizerMashup struct {
	// redact|replace|mask|hash|encrypt
	Type_ string `json:"type"`

	/* -------------- redact -------------- */

	/* -------------- replace -------------- */
	// The string to replace with
	NewValue string `json:"new_value,omitempty"`

	/* -------------- mask -------------- */
	// The replacement character
	MaskingChar string `json:"masking_char,omitempty"`
	// The amount of characters that should be replaced
	CharsToMask int32 `json:"chars_to_mask,omitempty"`
	// Whether to mask the PII from it's end
	FromEnd bool `json:"from_end,omitempty"`

	/* -------------- hash -------------- */
	// The hashing algorithm
	HashType string `json:"hash_type,omitempty"`

	/* -------------- encrypt -------------- */
	// Cryptographic key of length 128, 192 or 256 bits, in a string format
	Key string `json:"key,omitempty"`
}

type AnonymizerMap map[string]AnonymizerMashup
