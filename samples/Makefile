ALL_SAMPLES = ./analyzer-advanced ./anonymizer ./basic ./serialization

all: build check

build:
	for dirname in $(ALL_SAMPLES); do cd $$dirname; go build .; cd .. ; done

check:
	for dirname in $(ALL_SAMPLES); do cd $$dirname; staticcheck .; cd .. ; done