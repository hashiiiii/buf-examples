module github.com/bufbuild/buf-examples/bsr/quickstart/client

go 1.24.0

require (
	buf.build/gen/go/aquarium/common/protocolbuffers/go v1.36.11-20260129015210-58039600ab3e.1
	buf.build/gen/go/aquarium/invoice/connectrpc/gosimple v1.19.1-20260129020012-08e19ed5c664.2
	buf.build/gen/go/aquarium/invoice/protocolbuffers/go v1.36.11-20260129020012-08e19ed5c664.1
)

require (
	buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go v1.36.11-20251209175733-2a1774d88802.1 // indirect
	connectrpc.com/connect v1.19.1 // indirect
	google.golang.org/protobuf v1.36.11 // indirect
)
