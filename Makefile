go:
	protoc --proto_path=${GOPATH}/src --go_out=:${GOPATH}/src ${GOPATH}/src/github.com/CSUNetSec/netsec-protobufs/common/common.proto;
	protoc --proto_path=${GOPATH}/src --go_out=:${GOPATH}/src ${GOPATH}/src/github.com/CSUNetSec/netsec-protobufs/protocol/bgp/bgp.proto;

clean:
	find . -name "*.go" -exec rm {} \;
