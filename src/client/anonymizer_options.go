package client

import (
	"github.com/CodeRunRepeat/presidio-go-client/generated"
)

const DEFAULT string = "DEFAULT"

type Anonymizer interface {
	getTypeName() string
	generateRequest() generated.AnonymizerMashup
	compareWithRequest(request generated.AnonymizerMashup) bool
	generateReverseRequest() any
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

func getFirstAnonymizer(as *AnonymizerSet) Anonymizer {
	// Storing just one value from the AnonymizerSet since the auto-generated type for this field
	// cannot accommodate multiple values.
	firstAnonymizer := as.Get(DEFAULT)
	if firstAnonymizer == nil {
		firstAnonymizer = as.First()
	}
	return firstAnonymizer
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
func (as *AnonymizerSet) prepareAnonymizerSetForReverseRequest() *generated.AnyOfDeanonymizeRequestDeanonymizers {
	output := new(generated.AnyOfDeanonymizeRequestDeanonymizers)

	firstAnonymizer := getFirstAnonymizer(as)
	if firstAnonymizer == nil {
		return nil
	}

	reverse := firstAnonymizer.generateReverseRequest()
	if reverse == nil {
		return nil
	}
	output.Decrypt = reverse.(generated.Decrypt) // The only reversible anonymization method available is encryption
	return output
}

/*------------------------------ Anonymizer types ------------------------------*/
const (
	REDACT  string = "redact"
	REPLACE string = "replace"
	MASK    string = "mask"
	HASH    string = "hash"
	ENCRYPT string = "encrypt"
	DECRYPT string = "decrypt"
)

func checkForType(request generated.AnonymizerMashup, typeName string) bool {
	return request.Type_ == typeName
}

type RedactAnonymizer struct {
}

func (ra RedactAnonymizer) getTypeName() string { return REDACT }
func (ra RedactAnonymizer) generateRequest() generated.AnonymizerMashup {
	return generated.AnonymizerMashup{Type_: ra.getTypeName()}
}
func (ra RedactAnonymizer) compareWithRequest(request generated.AnonymizerMashup) bool {
	return checkForType(request, ra.getTypeName())
}
func (ra RedactAnonymizer) generateReverseRequest() any { return nil }

type ReplaceAnonymizer struct {
	NewValue string
}

func (ra ReplaceAnonymizer) getTypeName() string { return REPLACE }
func (ra ReplaceAnonymizer) generateRequest() generated.AnonymizerMashup {
	return generated.AnonymizerMashup{Type_: ra.getTypeName(), NewValue: ra.NewValue}
}
func (ra ReplaceAnonymizer) compareWithRequest(request generated.AnonymizerMashup) bool {
	if checkForType(request, ra.getTypeName()) {
		return (request.NewValue == ra.NewValue)
	}
	return false
}
func (ra ReplaceAnonymizer) generateReverseRequest() any { return nil }

type MaskAnonymizer struct {
	MaskingChar string
	CharsToMask int32
	FromEnd     bool
}

func (ma MaskAnonymizer) getTypeName() string { return MASK }
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
func (ma MaskAnonymizer) generateReverseRequest() any { return nil }

type HashAnonymizer struct {
	HashType string
}

func (ha HashAnonymizer) getTypeName() string { return HASH }
func (ha HashAnonymizer) generateRequest() generated.AnonymizerMashup {
	return generated.AnonymizerMashup{Type_: ha.getTypeName(), HashType: ha.HashType}
}
func (ha HashAnonymizer) compareWithRequest(request generated.AnonymizerMashup) bool {
	if checkForType(request, ha.getTypeName()) {
		return (request.HashType == ha.HashType)
	}
	return false
}
func (ha HashAnonymizer) generateReverseRequest() any { return nil }

type EncryptAnonymizer struct {
	Key string
}

func (ea EncryptAnonymizer) getTypeName() string { return ENCRYPT }
func (ea EncryptAnonymizer) generateRequest() generated.AnonymizerMashup {
	return generated.AnonymizerMashup{Type_: ea.getTypeName(), Key: ea.Key}
}
func (ea EncryptAnonymizer) compareWithRequest(request generated.AnonymizerMashup) bool {
	if checkForType(request, ea.getTypeName()) {
		return (request.Key == ea.Key)
	}
	return false
}
func (ea EncryptAnonymizer) generateReverseRequest() any {
	return generated.Decrypt{Type_: DECRYPT, Key: ea.Key}
}
