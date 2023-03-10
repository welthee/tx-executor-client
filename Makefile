ASYNCAPI_CODEGEN := asyncapi
API_SCHEMA_DIR := api
API_GENERATED_DIR := api/generated
API_API_SCHEMA_FILE := $(API_SCHEMA_DIR)/tx-api.yaml



oapi: ## build oapi
	mkdir -pv $(API_GENERATED_DIR)/types
	$(ASYNCAPI_CODEGEN) generate models golang $(API_API_SCHEMA_FILE) -o $(API_GENERATED_DIR)/types --packageName=types
