.PHONY: generate

BUILD_DIR   := build

clean:
	go clean
	rm -rf ./${BUILD_DIR}

generate:
	go run ./internal/cmd/generate

generate-antlr:
	mkdir -p fhirpath/internal/parser
	find fhirpath/internal/parser -type f ! -name 'fhirpath.g4' -delete
	cd fhirpath/internal/parser && antlr fhirpath.g4 -Dlanguage=Go -no-listener -no-visitor -o .

	# make generated type names more go like
	sed -i '' \
		-e 's/\([ 	&]\)Fhirpath/\1FHIRPath/g' \
		-e 's/\([a-z]\)fhirpath/\1FHIRPath/g' \
		-e 's/fhirpathParser\([ )]\)/FHIRPathParser\1/g' \
		fhirpath/internal/parser/*.go
