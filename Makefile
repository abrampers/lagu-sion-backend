PROTO_DIR=lagusion
PROTO_GEN_DIR=lagusion

protobuf:
	mkdir -p ${PROTO_GEN_DIR}
	protoc -I ${PROTO_DIR} lagu_sion.proto \
		--go_out=plugins=grpc:${PROTO_GEN_DIR}
