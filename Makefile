.PHONY: compile server clean

clean:
	rm -f main
	find . -name "*.pb.go" -delete

compile:
	protoc api/v1/*.proto \
		--gogo_out=Mgogoproto/gogo.proto=github.com/gogo/protobuf/proto:. \
		--proto_path=$$(go list -f '{{ .Dir }}' -m github.com/gogo/protobuf) \
		--proto_path=.

server:
	go build cmd/server/main.go