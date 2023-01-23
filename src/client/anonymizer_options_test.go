package client

import "testing"

func TestAnonymizerSetCreate(t *testing.T) {
	set := CreateAnonymizerSet()
	if set.Count() != 0 {
		t.Error("AnonymizerSet should hold 0 items")
	}
}

func TestAnonymizerSetAddOne(t *testing.T) {
	set := CreateAnonymizerSet()
	set.AddDefaultAnonymizer(HashAnonymizer{HashType: "SHA-256"})

	if set.Count() != 1 {
		t.Error("AnonymizerSet.Count() should return 1")
	}

	ha := set.Get(DEFAULT)
	if ha == nil {
		t.Error("AnonymizerSet should contain the DEFAULT anonymizer")
	}
}

func TestAnonymizerSetNotFound(t *testing.T) {
	set := CreateAnonymizerSet()

	a := set.Get(DEFAULT)
	if a != nil {
		t.Error("Empty AnonymizerSet should not return anything with Get()")
	}

	set.AddDefaultAnonymizer(HashAnonymizer{HashType: "SHA-256"})
	a = set.Get("PERSON_NAME")
	if a != nil {
		t.Error("AnonymizerSet should not return unknown anonymizer with Get()")
	}
}

func TestAnonymizerSetPrepare(t *testing.T) {
	set := CreateAnonymizerSet()

	var anonymizer Anonymizer = RedactAnonymizer{}
	set.AddDefaultAnonymizer(anonymizer)
	p := (*set.prepareAnonymizerSetForRequest())[DEFAULT]
	if p.Type_ != anonymizer.getTypeName() {
		t.Errorf("AnonymizerSet should produce a %v request", anonymizer.getTypeName())
	}

	anonymizer = ReplaceAnonymizer{NewValue: "test"}
	set.AddDefaultAnonymizer(anonymizer)
	p = (*set.prepareAnonymizerSetForRequest())[DEFAULT]
	if p.Type_ != anonymizer.getTypeName() {
		t.Errorf("AnonymizerSet should produce a %v request", anonymizer.getTypeName())
	}
	if !anonymizer.compareWithRequest(p) {
		t.Errorf("AnonymizerSet produced a %v request with incorrect values", anonymizer.getTypeName())
	}

	anonymizer = MaskAnonymizer{MaskingChar: "*", CharsToMask: 4, FromEnd: false}
	set.AddDefaultAnonymizer(anonymizer)
	p = (*set.prepareAnonymizerSetForRequest())[DEFAULT]
	if p.Type_ != anonymizer.getTypeName() {
		t.Errorf("AnonymizerSet should produce a %v request", anonymizer.getTypeName())
	}
	if !anonymizer.compareWithRequest(p) {
		t.Errorf("AnonymizerSet produced a %v request with incorrect values", anonymizer.getTypeName())
	}

	anonymizer = HashAnonymizer{HashType: "SHA256"}
	set.AddDefaultAnonymizer(anonymizer)
	p = (*set.prepareAnonymizerSetForRequest())[DEFAULT]
	if p.Type_ != anonymizer.getTypeName() {
		t.Errorf("AnonymizerSet should produce a %v request", anonymizer.getTypeName())
	}
	if !anonymizer.compareWithRequest(p) {
		t.Errorf("AnonymizerSet produced a %v request with incorrect values", anonymizer.getTypeName())
	}

	anonymizer = EncryptAnonymizer{Key: "key"}
	set.AddDefaultAnonymizer(anonymizer)
	p = (*set.prepareAnonymizerSetForRequest())[DEFAULT]
	if p.Type_ != anonymizer.getTypeName() {
		t.Errorf("AnonymizerSet should produce a %v request", anonymizer.getTypeName())
	}
	if !anonymizer.compareWithRequest(p) {
		t.Errorf("AnonymizerSet produced a %v request with incorrect values", anonymizer.getTypeName())
	}
}
