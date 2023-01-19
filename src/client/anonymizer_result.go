package client

import (
	"github.com/CodeRunRepeat/presidio-go-client/generated"
)

type AnonymizerMatch struct {
	// Start is the index of the match's first character in the text.
	Start int32

	// End is the index of the match's end (first non-character) in the text.
	End int32

	// EntityType is the name of the PII entity changed by the AnonymizerAction.
	EntityType string

	// AnonymizerAction is the action taken from an anonymizer against a part of the text.
	AnonymizerAction string

	// This is the Text that was changed by the AnonymizerAction.
	Text string
}

type AnonymizerResult struct {
	Matches []AnonymizerMatch
}

func NewAnonymizerResult(numMatches int) *AnonymizerResult {
	var result = new(AnonymizerResult)
	result.Matches = make([]AnonymizerMatch, numMatches)
	return result
}

//lint:ignore U1000 Will be fixed when we resolve the anonymizer serialization issue with the Presidio team
func transformToAnonymizerResult(input []generated.OperatorResult) AnonymizerResult {
	var result = NewAnonymizerResult(len(input))

	for index, r := range input {
		result.Matches[index] = AnonymizerMatch{
			Text:             r.Text,
			Start:            r.Start,
			End:              r.End,
			EntityType:       r.EntityType,
			AnonymizerAction: r.Operator,
		}
	}

	return *result
}

//lint:ignore U1000 Will be fixed when we resolve the anonymizer serialization issue with the Presidio team
func transformFromAnonymizerResult(input AnonymizerResult) []generated.OperatorResult {
	var result = make([]generated.OperatorResult, len(input.Matches))
	for index, m := range input.Matches {
		result[index] = generated.OperatorResult{
			Text:       m.Text,
			Start:      m.Start,
			End:        m.End,
			EntityType: m.EntityType,
			Operator:   m.AnonymizerAction,
		}
	}

	return result
}
