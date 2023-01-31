/*
This file, unlike others in this directory, is not generated. Its purpose is to provide with a
type for the AnoynimizeRequest.Anonymizers field that serializes in a compatible way with
what the Presidio API expects.
*/
package generated

import (
	"encoding/json"
	"fmt"
	"reflect"
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

const (
	REDACT  string = "redact"
	REPLACE string = "replace"
	MASK    string = "mask"
	HASH    string = "hash"
	ENCRYPT string = "encrypt"
	DECRYPT string = "decrypt"
	NIL     string = ""
)

func (a AnonymizerMashup) MarshalJSON() ([]byte, error) {
	switch a.Type_ {
	case REDACT:
		return json.Marshal(Redact{Type_: a.Type_})
	case REPLACE:
		return json.Marshal(Replace{Type_: a.Type_, NewValue: a.NewValue})
	case MASK:
		return json.Marshal(Mask{Type_: a.Type_, MaskingChar: a.MaskingChar, CharsToMask: a.CharsToMask, FromEnd: a.FromEnd})
	case ENCRYPT:
		return json.Marshal(Encrypt{Type_: a.Type_, Key: a.Key})
	case HASH:
		return json.Marshal(Hash{Type_: a.Type_, HashType: a.HashType})
	case DECRYPT:
		return json.Marshal(Decrypt{Type_: a.Type_, Key: a.Key})
	default:
		return nil, &json.MarshalerError{Type: reflect.TypeOf(a), Err: fmt.Errorf("unknown type %v", a.Type_)}
	}
}

type AnonymizerMap map[string]AnonymizerMashup
