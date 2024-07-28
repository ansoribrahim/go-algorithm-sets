build:
	@echo "Generating files..."
	mkdir generated || true
	cd openapi && oapi-codegen  -config config.yaml openapi.yaml
