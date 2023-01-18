package client

import (
	"errors"

	"github.com/CodeRunRepeat/presidio-go-client/generated"
)

const DEFAULT string = "DEFAULT"

type Anonymizer interface {
	getTypeName() string
	generateRequest() any
	compareWithRequest(request any) (bool, error)
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

func (as *AnonymizerSet) prepareAnonymizerSetForRequest() *generated.AnyOfAnonymizeRequestAnonymizers {
	output := new(generated.AnyOfAnonymizeRequestAnonymizers)

	// Storing just one value from the AnonymizerSet since the auto-generated type for this field
	// cannot accommodate multiple values.
	firstAnonymizer := as.Get(DEFAULT)
	if firstAnonymizer == nil {
		firstAnonymizer = as.First()
	}
	if firstAnonymizer == nil {
		return nil
	}

	switch firstAnonymizer.getTypeName() {
	case REPLACE:
		output.Replace = firstAnonymizer.generateRequest().(generated.Replace)
	case REDACT:
		output.Redact = firstAnonymizer.generateRequest().(generated.Redact)
	case MASK:
		output.Mask = firstAnonymizer.generateRequest().(generated.Mask)
	case HASH:
		output.Hash = firstAnonymizer.generateRequest().(generated.Hash)
	case ENCRYPT:
		output.Encrypt = firstAnonymizer.generateRequest().(generated.Encrypt)
	}
	return output
}

/*------------------------------ Anonymizer types ------------------------------*/
const REDACT string = "redact"
const REPLACE string = "replace"
const MASK string = "mask"
const HASH string = "hash"
const ENCRYPT string = "encrypt"

type anonymizerRequestType interface {
	generated.Replace | generated.Redact | generated.Mask | generated.Hash | generated.Encrypt
}

func checkForType[requestType anonymizerRequestType](request any) (bool, requestType, error) {
	var value requestType
	if request == nil {
		return false, value, errors.New("request is nil")
	}
	value, test := request.(requestType)
	if !test {
		return false, value, errors.New("request has incorrect type")
	}

	return true, value, nil
}

type RedactAnonymizer struct {
}

func (ra RedactAnonymizer) getTypeName() string  { return REDACT }
func (ra RedactAnonymizer) generateRequest() any { return generated.Redact{Type_: ra.getTypeName()} }
func (ra RedactAnonymizer) compareWithRequest(request any) (bool, error) {
	res, _, err := checkForType[generated.Redact](request)
	return res, err
}

type ReplaceAnonymizer struct {
	NewValue string
}

func (ra ReplaceAnonymizer) getTypeName() string { return REPLACE }
func (ra ReplaceAnonymizer) generateRequest() any {
	return generated.Replace{Type_: ra.getTypeName(), NewValue: ra.NewValue}
}
func (ra ReplaceAnonymizer) compareWithRequest(request any) (bool, error) {
	res, value, err := checkForType[generated.Replace](request)
	if res {
		return (value.NewValue == ra.NewValue), nil
	}
	return res, err
}

type MaskAnonymizer struct {
	MaskingChar string
	CharsToMask int32
	FromEnd     bool
}

func (ma MaskAnonymizer) getTypeName() string { return MASK }
func (ma MaskAnonymizer) generateRequest() any {
	return generated.Mask{
		Type_:       ma.getTypeName(),
		MaskingChar: ma.MaskingChar,
		CharsToMask: ma.CharsToMask,
		FromEnd:     ma.FromEnd,
	}
}
func (ma MaskAnonymizer) compareWithRequest(request any) (bool, error) {
	res, value, err := checkForType[generated.Mask](request)
	if res {
		return (value.MaskingChar == ma.MaskingChar && value.CharsToMask == ma.CharsToMask && value.FromEnd == ma.FromEnd), nil
	}
	return res, err
}

type HashAnonymizer struct {
	HashType string
}

func (ha HashAnonymizer) getTypeName() string { return HASH }
func (ha HashAnonymizer) generateRequest() any {
	return generated.Hash{Type_: ha.getTypeName(), HashType: ha.HashType}
}
func (ha HashAnonymizer) compareWithRequest(request any) (bool, error) {
	res, value, err := checkForType[generated.Hash](request)
	if res {
		return (value.HashType == ha.HashType), nil
	}
	return res, err
}

type EncryptAnonymizer struct {
	Key string
}

func (ea EncryptAnonymizer) getTypeName() string { return ENCRYPT }
func (ea EncryptAnonymizer) generateRequest() any {
	return generated.Encrypt{Type_: ea.getTypeName(), Key: ea.Key}
}
func (ea EncryptAnonymizer) compareWithRequest(request any) (bool, error) {
	res, value, err := checkForType[generated.Encrypt](request)
	if res {
		return (value.Key == ea.Key), nil
	}
	return res, err
}
