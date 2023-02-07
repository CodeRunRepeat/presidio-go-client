ALL_SAMPLES = ./analyzer-advanced ./anonymizer ./basic ./serialization

all: build check

build:
	for dirname in $(ALL_SAMPLES); do cd $$dirname; go mod tidy; go build .; cd .. ; done

check:
	for dirname in $(ALL_SAMPLES); do cd $$dirname; go mod tidy; staticcheck .; cd .. ; done