run/broker:
	sh ./tools/runbroker.sh

run/info:
	sh ./tools/runinfo.sh

req/rpc:
	curl localhost:3000/rpc

req/grpc:
	curl localhost:3000/grpc