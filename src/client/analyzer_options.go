package client

import "presidio.org/generated"

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

	if o.request.AdHocRecognizers == nil {
		o.request.AdHocRecognizers = make([]generated.PatternRecognizer, 1)
		o.request.AdHocRecognizers[0] = recognizer
	} else {
		o.request.AdHocRecognizers = append(o.request.AdHocRecognizers, recognizer)
	}

	return o
}
