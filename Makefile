PACKAGE            = trax
DATE              ?= $(shell date +%FT%T%z)
VERSION           ?= $(shell echo $(shell cat $(PWD)/.version)-$(shell git describe --tags --always))
DIR                = $(strip $(shell dirname $(realpath $(lastword $(MAKEFILE_LIST)))))

GO                 = go
GODOC              = godoc
GOFMT              = gofmt

# For CI
ifneq ($(wildcard ./bin/golangci-lint),)
	GOLINT = ./bin/golangci-lint
else
	GOLINT = golangci-lint
endif

V                 = 0
Q                 = $(if $(filter 1,$V),,@)
M                 = $(shell printf "\033[0;35m▶\033[0m")

ADMIN             = admin
API               = api
AUTH              = auth
WEB_CLIENT        = web_client
CLIENT            = client

#  Migration directory for goose postgres
MIGRATION_DIR     = db/postgres
GOOSE_DRIVER      = postgres
GOOSE_DBSTRING    = postgres://username:password@localhost:5432/postgres

# Static directory name for client
STATIC            = static

ROOT               = .
GEN_PB_GO          = protoc -I=$(ROOT) --gogoslick_out=plugins=grpc,paths=source_relative:$(ROOT)
GEN_PB_TS          = cd cmd/client && npx protoc -I=../../$(ROOT) --ts_out ../../$(ROOT) --ts_opt generate_dependencies,optimize_speed,long_type_string

.PHONY: all
all: admin api auth client web_client

.PHONY: admin
admin:  ## Build admin binary
	$(info $(M) building executable admin) @
	$Q cd cmd/$(ADMIN) && $(GO) build \
		-mod=readonly \
		-tags release \
		-ldflags '-X main.version=$(VERSION)' \
		-o ../../bin/$(PACKAGE)_$(ADMIN)_$(VERSION)
	$Q cp bin/$(PACKAGE)_$(ADMIN)_$(VERSION) bin/$(PACKAGE)_$(ADMIN)

.PHONY: api
api:  ## Build api binary
	$(info $(M) building executable api) @
	$Q cd cmd/$(API) && $(GO) build \
		-mod=readonly \
		-tags release \
		-ldflags '-X main.version=$(VERSION)' \
		-o ../../bin/$(PACKAGE)_$(API)_$(VERSION)
	$Q cp bin/$(PACKAGE)_$(API)_$(VERSION) bin/$(PACKAGE)_$(API)

.PHONY: auth
auth:  ## Build auth binary
	$(info $(M) building executable auth) @
	$Q cd cmd/$(AUTH) && $(GO) build \
		-mod=readonly \
		-tags release \
		-ldflags '-X main.version=$(VERSION)' \
		-o ../../bin/$(PACKAGE)_$(AUTH)_$(VERSION)
	$Q cp bin/$(PACKAGE)_$(AUTH)_$(VERSION) bin/$(PACKAGE)_$(AUTH)

.PHONY: client
client:  ## Build client content
	$(info $(M) building bundle client) @
	$Q cd cmd/$(CLIENT) && CONFIG_FILE=$${CONFIG_FILE:-../../config/client/local.json} npx vite build
	$Q mkdir -p bin && rm -rf bin/$(STATIC) && mkdir -p bin/$(CLIENT)/$(STATIC)/
	$Q yes | cp -rf cmd/$(CLIENT)/dist/. bin/$(CLIENT)/$(STATIC)/

.PHONY: web_client
web_client: ## Build web_client binary
	$(info $(M) building executable web_client) @
	$Q cd cmd/$(WEB_CLIENT) && $(GO) build \
		-mod=readonly \
		-tags release \
		-ldflags '-X main.version=$(VERSION)' \
		-o ../../bin/$(PACKAGE)_$(WEB_CLIENT)_$(VERSION)
	$Q yes | cp -rf bin/$(PACKAGE)_$(WEB_CLIENT)_$(VERSION) bin/$(PACKAGE)_$(WEB_CLIENT)

# NPM
.PHONY: npm
npm: ## Install npm dependencies
	$(info $(M) running npm install) @
	$Q cd cmd/$(CLIENT) && npm install

# Proto lang
.PHONY: proto-go proto-ts
proto-go:    PB_LANG = GO
proto-ts:    PB_LANG = TS
proto-go proto-ts: ## Regenerate protobuf files
	$(info $(M) running protobuf $(PB_LANG)) @
	$(info $(M) generate pkg) @
	$Q $(GEN_PB_$(PB_LANG)) pkg/gogoproto/gogo.proto
	$Q $(GEN_PB_$(PB_LANG)) pkg/paginate/paginate.proto
	$Q $(GEN_PB_$(PB_LANG)) pkg/pbtypes/empty.proto
	$Q $(GEN_PB_$(PB_LANG)) pkg/pbtypes/string.proto
	$Q $(GEN_PB_$(PB_LANG)) pkg/cookie/keys.proto
	$(info $(M) generate internal) @
	$Q $(GEN_PB_$(PB_LANG)) internal/user/claims.proto
	$Q $(GEN_PB_$(PB_LANG)) internal/user/entity.proto
	$Q $(GEN_PB_$(PB_LANG)) internal/user/role.proto
	$Q $(GEN_PB_$(PB_LANG)) internal/user/user.proto
	$(info $(M) generate clients) @
	$(info $(M) generate dto) @
	$Q $(GEN_PB_$(PB_LANG)) internal/user/dto/entity.proto
	$Q $(GEN_PB_$(PB_LANG)) internal/user/dto/role.proto
	$Q $(GEN_PB_$(PB_LANG)) internal/user/dto/user.proto
	$(info $(M) generate services) @
	$Q $(GEN_PB_$(PB_LANG)) cmd/$(ADMIN)/grpc/$(ADMIN).proto
	$Q $(GEN_PB_$(PB_LANG)) cmd/$(API)/grpc/$(API).proto
	$Q $(GEN_PB_$(PB_LANG)) cmd/$(AUTH)/grpc/$(AUTH).proto

# Proto
.PHONY: proto
proto: proto-go proto-ts

# Vendor
.PHONY: vendor
vendor:
	$(info $(M) running go mod vendor) @
	$Q $(GO) mod vendor

# Tidy
.PHONY: tidy
tidy:
	$(info $(M) running go mod tidy) @
	$Q $(GO) mod tidy

# Check
.PHONY: check
check: vendor lint test

# Lint
.PHONY: lint
lint:
	$(info $(M) running $(GOLINT))
	$Q $(GOLINT) run

# Test
.PHONY: test
test:
	$(info $(M) running go test) @
	$Q $(GO) test -cover -race -v ./...

# Clean
.PHONY: clean
clean:
	$(info $(M) cleaning binaries) @
	$Q rm -rf bin/$(PACKAGE)_$(ADMIN)*
	$Q rm -rf bin/$(PACKAGE)_$(API)*
	$Q rm -rf bin/$(PACKAGE)_$(AUTH)*
	$Q rm -rf bin/$(PACKAGE)_$(WEB_CLIENT)*
	$Q rm -rf bin/$(PACKAGE)_$(CLIENT)*
	$Q rm -rf bin/$(CLIENT)/$(STATIC)

# Clean protobuf generated files
.PHONY: clean-proto
clean-proto:
	$(info $(M) cleaning protobuf generated files) @
	$Q find . -name "*_pb.d.ts" -type f -delete
	$Q find . -name "*_pb.js" -type f -delete
	$Q find . -name "*.pb.go" -type f -delete
	$Q find . -name "*pb_test.go" -type f -delete
	$Q find . -name "*ServiceClientPb.ts" -type f -delete

## Helpers

# Goose helpers
.PHONY: goose-create
goose-create: ## Create a new migration
	$(info $(M) create new migration file using NAME=$(NAME)) @
	$Q cd $(MIGRATION_DIR) && GOOSE_DRIVER=$(GOOSE_DRIVER) GOOSE_DBSTRING=$(GOOSE_DBSTRING) goose create $(NAME) sql

.PHONY: goose-up
goose-up: ## Create a new migration
	$(info $(M) migration up) @
	$Q cd $(MIGRATION_DIR) && GOOSE_DRIVER=$(GOOSE_DRIVER) GOOSE_DBSTRING=$(GOOSE_DBSTRING) goose up

# Goose helpers
.PHONY: goose-down
goose-down: ## Create a new migration
	$(info $(M) migration down) @
	$Q cd $(MIGRATION_DIR) && GOOSE_DRIVER=$(GOOSE_DRIVER) GOOSE_DBSTRING=$(GOOSE_DBSTRING) goose down

.PHONY: go-version
go-version: ## Print go version used in this makefile
	$Q echo $(GO)

.PHONY: fmt
fmt: ## Format code
	$(info $(M) running $(GOFMT)) @
	$Q $(GOFMT) ./...

.PHONY: doc
doc: ## Generate project documentation
	$(info $(M) running $(GODOC)) @
	$Q $(GODOC) ./...

.PHONY: version
version: ## Print current project version
	@echo $(VERSION)

.PHONY: help
help: ## Print this
	@grep -E '^[ a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-15s\033[0m %s\n", $$1, $$2}'
