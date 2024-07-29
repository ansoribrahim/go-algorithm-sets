MAIN_FILE = openapi/openapi.yaml
COMPONENTS_FILE = openapi/model.yaml
PATHS_FILE = openapi/algorithmBinarySearch.yaml
OUTPUT_FILE = openapi/merged-openapi.yaml

CHECK_YQ = $(shell command -v yq 2>/dev/null)

install-yq:
	@if [ -z "$(CHECK_YQ)" ]; then \
		echo "Installing yq..."; \
		brew install yq || { echo "Failed to install yq. Please install it manually."; exit 1; }; \
	else \
		echo "yq is already installed."; \
	fi

merge:
	@make install-yq
	@echo "Merging files..."
	yq eval-all 'select(fileIndex == 0) * select(fileIndex == 1)' $(MAIN_FILE) $(COMPONENTS_FILE) > $(OUTPUT_FILE)

build:
	@echo "Generating merged openapi.yaml..."
	@make merge
	@echo "Generating files..."
	mkdir generated || true
	cd openapi && oapi-codegen  -config config.yaml merged-openapi.yaml
