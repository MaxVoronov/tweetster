## help: Show current help
help: Makefile
	@echo "Choose a command run:"
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'

## lint: Check source code by linters
lint:
	@echo "Checking go vet..." && go vet ./... && echo "Done!\n"
	@echo "Checking golint..." && golint ./... && echo "Done!\n"
	@echo "Checking golangci-lint..." && golangci-lint run ./... && echo "Done!"

## proto-gen: Generate protobuf files for all services
proto-gen: proto-clean
	@printf "Generating protobuf files... "
	@find $(CURDIR)/api/proto/ -name '*.proto' -exec \
		protoc \
			--proto_path=$(CURDIR)/api/proto/ \
			--go_out=plugins=grpc:$(CURDIR)/internal/tweets/pb \
			{} \;
	@echo "Done"

## proto-clean: Remove generated protobuf files
proto-clean:
	@printf "Cleaning protobuf files... "
	@rm -rf $(CURDIR)/internal/tweets/pb/*
	@echo "Done"
