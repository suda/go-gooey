build-example:
	echo "Building '${EXAMPLE}' example..."
	@go build -o build/examples/${EXAMPLE} ./examples/${EXAMPLE}

examples:
	@mkdir -p build/examples
	@EXAMPLE=binding $(MAKE) -s build-example
	@EXAMPLE=simple $(MAKE) -s build-example

.PHONY: examples build-example