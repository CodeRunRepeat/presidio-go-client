package main

import (
	"encoding/json"
	"fmt"
)

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

func main() {
	var config map[string]AnonymizerMashup = make(map[string]AnonymizerMashup)
	config["PERSON"] = AnonymizerMashup{Type_: "redact"}
	config["PHONE_NUMBER"] = AnonymizerMashup{Type_: "replace", NewValue: "ANONYMIZED"}

	result, _ := json.Marshal(config)
	fmt.Printf("%v\n", string(result))

	unmarshal()
}

func unmarshal() {
	var source = `{
		"PERSON": {
			"type": "redact"
		},
		"PHONE_NUMBER": {
			"type": "replace",
			"new_value": "ANONYMIZED"
		}
	}`

	var target any
	json.Unmarshal([]byte(source), &target)

	fmt.Printf("%t", target)
}
