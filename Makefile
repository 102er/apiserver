GOPATH:=$(shell go env GOPATH)

.PHONY: proto1
# generate go protoc code
# 官方带的proto-gen
gp:
	protoc  --go_out=paths=source_relative:. \
			 api/v1/hello_world.proto

# gogo-proto-gen
.PHONY: gpp
ggp:
	protoc  --gogo_out=paths=source_relative:. \
			 api/v1/hello_world.proto

.PHONY: errors
# generate errcode code
errors:
	protoc --proto_path=. \
      	   --proto_path=./third_party \
       	   --gogo_out=paths=source_relative:. \
           --go-errors_out=paths=source_relative:. \
    ./errors/errors_v1.proto

generate:
	wire ./cmd