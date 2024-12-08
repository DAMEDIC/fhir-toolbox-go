.PHONY: generate

BUILD_DIR   := build

clean:
	go clean
	rm -rf ./${BUILD_DIR}

generate:
	go run ./internal/cmd/generate

generate-antlr:
	rm -rf fhirpath/parser/gen
	cd fhirpath/parser && antlr fhirpath.g4 -Dlanguage=Go -no-listener -no-visitor -o gen

	# make generated type names more go like
	sed -i '' \
		-e 's/\([ 	&]\)Fhirpath/\1FHIRPath/g' \
		-e 's/\([a-z]\)fhirpath/\1FHIRPath/g' \
		fhirpath/parser/gen/*