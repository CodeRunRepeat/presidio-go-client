package client

import (
	"github.com/CodeRunRepeat/presidio-go-client/generated"
)

const DEFAULT string = "DEFAULT"

type Anonymizer interface {
	getTypeName() string
	generateRequest() generated.AnonymizerMashup
	compareWithRequest(request generated.AnonymizerMashup) bool
	generateReverseRequest() generated.AnonymizerMashup
}

// An AnonymizerSet holds configuration on how to anonymize different entities.
type AnonymizerSet map[string]Anonymizer

func CreateAnonymizerSet() *AnonymizerSet {
	set := make(AnonymizerSet)
	return &set
}

func (as *AnonymizerSet) Count() int {
	return len(*as)
}

func (as *AnonymizerSet) Get(entityName string) Anonymizer {
	return (*as)[entityName]
}

func (as *AnonymizerSet) First() Anonymizer {
	if len(*as) == 0 {
		return nil
	}
	for _, value := range *as {
		return value
	}

	return nil
}

func (as *AnonymizerSet) AddAnonymizer(entityName string, anonymizer Anonymizer) *AnonymizerSet {
	if anonymizer == nil {
		panic("AddAnonymizer: anonymizer should not be nil")
	}
	(*as)[entityName] = anonymizer
	return as
}

func (as *AnonymizerSet) AddDefaultAnonymizer(anonymizer Anonymizer) *AnonymizerSet {
	return as.AddAnonymizer(DEFAULT, anonymizer)
}

func (as *AnonymizerSet) RemoveAnonymizer(entityName string) *AnonymizerSet {
	delete(*as, entityName)
	return as
}

func (as *AnonymizerSet) prepareAnonymizerSetForRequest() *generated.AnonymizerMap {
	output := make(generated.AnonymizerMap, as.Count())

	for entityType, an := range *as {
		var mashup = an.generateRequest()
		output[entityType] = mashup
	}

	return &output
}

//lint:ignore U1000 Will be fixed when we resolve the anonymizer serialization issue with the Presidio team
func (as *AnonymizerSet) prepareAnonymizerSetForReverseRequest() *generated.AnonymizerMap {
	output := make(generated.AnonymizerMap, as.Count())

	for entityType, an := range *as {
		var mashup = an.generateReverseRequest()
		output[entityType] = mashup
	}

	return &output
}

/*------------------------------ Anonymizer types ------------------------------*/
func checkForType(request generated.AnonymizerMashup, typeName string) bool {
	return request.Type_ == typeName
}

type RedactAnonymizer struct {
}

func (ra RedactAnonymizer) getTypeName() string { return generated.REDACT }
func (ra RedactAnonymizer) generateRequest() generated.AnonymizerMashup {
	return generated.AnonymizerMashup{Type_: ra.getTypeName()}
}
func (ra RedactAnonymizer) compareWithRequest(request generated.AnonymizerMashup) bool {
	return checkForType(request, ra.getTypeName())
}
func (ra RedactAnonymizer) generateReverseRequest() generated.AnonymizerMashup {
	return generated.AnonymizerMashup{Type_: generated.NIL}
}

type ReplaceAnonymizer struct {
	NewValue string
}

func (ra ReplaceAnonymizer) getTypeName() string { return generated.REPLACE }
func (ra ReplaceAnonymizer) generateRequest() generated.AnonymizerMashup {
	return generated.AnonymizerMashup{Type_: ra.getTypeName(), NewValue: ra.NewValue}
}
func (ra ReplaceAnonymizer) compareWithRequest(request generated.AnonymizerMashup) bool {
	if checkForType(request, ra.getTypeName()) {
		return (request.NewValue == ra.NewValue)
	}
	return false
}
func (ra ReplaceAnonymizer) generateReverseRequest() generated.AnonymizerMashup {
	return generated.AnonymizerMashup{Type_: generated.NIL}
}

type MaskAnonymizer struct {
	MaskingChar string
	CharsToMask int32
	FromEnd     bool
}

func (ma MaskAnonymizer) getTypeName() string { return generated.MASK }
func (ma MaskAnonymizer) generateRequest() generated.AnonymizerMashup {
	return generated.AnonymizerMashup{
		Type_:       ma.getTypeName(),
		MaskingChar: ma.MaskingChar,
		CharsToMask: ma.CharsToMask,
		FromEnd:     ma.FromEnd,
	}
}
func (ma MaskAnonymizer) compareWithRequest(request generated.AnonymizerMashup) bool {
	if checkForType(request, ma.getTypeName()) {
		return (request.MaskingChar == ma.MaskingChar && request.CharsToMask == ma.CharsToMask && request.FromEnd == ma.FromEnd)
	}
	return false
}
func (ma MaskAnonymizer) generateReverseRequest() generated.AnonymizerMashup {
	return generated.AnonymizerMashup{Type_: generated.NIL}
}

type HashAnonymizer struct {
	HashType string
}

func (ha HashAnonymizer) getTypeName() string { return generated.HASH }
func (ha HashAnonymizer) generateRequest() generated.AnonymizerMashup {
	return generated.AnonymizerMashup{Type_: ha.getTypeName(), HashType: ha.HashType}
}
func (ha HashAnonymizer) compareWithRequest(request generated.AnonymizerMashup) bool {
	if checkForType(request, ha.getTypeName()) {
		return (request.HashType == ha.HashType)
	}
	return false
}
func (ha HashAnonymizer) generateReverseRequest() generated.AnonymizerMashup {
	return generated.AnonymizerMashup{Type_: generated.NIL}
}

type EncryptAnonymizer struct {
	Key string
}

func (ea EncryptAnonymizer) getTypeName() string { return generated.ENCRYPT }
func (ea EncryptAnonymizer) generateRequest() generated.AnonymizerMashup {
	return generated.AnonymizerMashup{Type_: ea.getTypeName(), Key: ea.Key}
}
func (ea EncryptAnonymizer) compareWithRequest(request generated.AnonymizerMashup) bool {
	if checkForType(request, ea.getTypeName()) {
		return (request.Key == ea.Key)
	}
	return false
}
func (ea EncryptAnonymizer) generateReverseRequest() generated.AnonymizerMashup {
	return generated.AnonymizerMashup{Type_: generated.DECRYPT, Key: ea.Key}
}
