go:
	protoc --proto_path=${GOPATH}/src --go_out=:${GOPATH}/src ${GOPATH}/src/github.com/CSUNetSec/netsec-protobufs/common/common.proto;
	protoc --proto_path=${GOPATH}/src --go_out=:${GOPATH}/src ${GOPATH}/src/github.com/CSUNetSec/netsec-protobufs/protocol/bgp/bgp.proto;
	protoc --proto_path=${GOPATH}/src --go_out=plugins=grpc:${GOPATH}/src ${GOPATH}/src/github.com/CSUNetSec/netsec-protobufs/bgpmon/bgpmon.proto;
	protoc --proto_path=${GOPATH}/src --go_out=plugins=grpc:${GOPATH}/src ${GOPATH}/src/github.com/CSUNetSec/netsec-protobufs/proddle/proddle.proto;
	protoc --proto_path=${GOPATH}/src --go_out=plugins=grpc:${GOPATH}/src ${GOPATH}/src/github.com/CSUNetSec/netsec-protobufs/netbrane/netbrane.proto;
	protoc --proto_path=${GOPATH}/src --go_out=plugins=grpc:${GOPATH}/src ${GOPATH}/src/github.com/CSUNetSec/netsec-protobufs/bgpmon/v2/bgpmon.proto;

python:
	protoc --proto_path=${GOPATH}/src --python_out=:${GOPATH}/src ${GOPATH}/src/github.com/CSUNetSec/netsec-protobufs/common/common.proto;
	protoc --proto_path=${GOPATH}/src --python_out=:${GOPATH}/src ${GOPATH}/src/github.com/CSUNetSec/netsec-protobufs/protocol/bgp/bgp.proto;


clean:
	find . -name "*.go" -exec rm {} \;
