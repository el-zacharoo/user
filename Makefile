# load env vars
# -include .env
# export mongoUri := $(value mongoUri)

# proto generates code from the most recent proto file(s)
.PHONY: proto
proto:
	cd proto && buf mod update
	buf lint
	# buf breaking --against './.git#branch=main,ref=HEAD~1'
	buf build
	buf generate
	cd proto && buf push

.PHONY: run
rungo:
	go run main.go

.PHONY: run
run:
	dapr run \
		--app-id messaging \
		--app-port 8081 \
		--app-protocol grpc \
		--config ./.dapr/config.yaml \
		--components-path ./.dapr/components \
		go run .

.PHONY: kill
kill:
	lsof -P -i TCP -s TCP:LISTEN | grep 8081 | awk '{print $2}' | { read pid; kill -9 ${pid}; }
	lsof -P -i TCP -s TCP:LISTEN | grep 9090 | awk '{print $2}' | { read pid; kill -9 ${pid}; }

.PHONY: test
test:
	go test -v ./handler/...
