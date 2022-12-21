package client

import (
	"github.com/google/uuid"

	"github.com/CodeRunRepeat/presidio-go-client/generated"
)

// The AnalyzerOptions type describes all the optional settings we can use
// with the Presidio analyzer service.
type AnalyzerOptions struct {
	request generated.AnalyzeRequest
}

func (o *AnalyzerOptions) SetCorrelationId(correlationId string) *AnalyzerOptions {
	o.request.CorrelationId = correlationId
	return o
}

func (o *AnalyzerOptions) SetScoreThreshold(threshold float64) *AnalyzerOptions {
	o.request.ScoreThreshold = threshold
	return o
}

func (o *AnalyzerOptions) SetEntities(entities []string) *AnalyzerOptions {
	o.request.Entities = entities
	return o
}

func (o *AnalyzerOptions) SetContext(context []string) *AnalyzerOptions {
	o.request.Context = context
	return o
}

func (o *AnalyzerOptions) AddDenyList(name string, entity string, language string, denyList []string, context []string) *AnalyzerOptions {
	var recognizer = generated.PatternRecognizer{
		Name:              name,
		SupportedEntity:   entity,
		SupportedLanguage: language,
		DenyList:          denyList,
		Context:           context,
	}

	appendRecognizer(&o.request, &recognizer)
	return o
}

func (o *AnalyzerOptions) AddPattern(name string, entity string, language string, regexes []string, scores []float64, context []string) *AnalyzerOptions {
	if regexes == nil || scores == nil {
		return o
	}

	var recognizer = generated.PatternRecognizer{
		Name:              name,
		SupportedEntity:   entity,
		SupportedLanguage: language,
		Context:           context,
	}

	length := len(regexes)
	if length > len(scores) {
		length = len(scores)
	}
	recognizer.Patterns = make([]generated.Pattern, length)
	for i := 0; i < length; i++ {
		recognizer.Patterns[i].Name = uuid.New().String()
		recognizer.Patterns[i].Regex = regexes[i]
		recognizer.Patterns[i].Score = scores[i]
	}

	appendRecognizer(&o.request, &recognizer)
	return o
}

func appendRecognizer(request *generated.AnalyzeRequest, recognizer *generated.PatternRecognizer) {
	if request.AdHocRecognizers == nil {
		request.AdHocRecognizers = make([]generated.PatternRecognizer, 1)
		request.AdHocRecognizers[0] = *recognizer
	} else {
		request.AdHocRecognizers = append(request.AdHocRecognizers, *recognizer)
	}
}
