gen_dal:
	cd dal && go run .

gen_api:
	cd hr-api && goctl api go --api hr.api --dir .
	cd hr-admin-api && goctl api go --api hr-admin.api --dir .

gen_swagger:
	cd hr-api && goctl api plugin -plugin goctl-swagger="swagger --filename hr-swagger.json" --api hr.api -dir .
	cd hr-admin-api && goctl api plugin -plugin goctl-swagger="swagger --filename hr-admin-swagger.json" --api hr-admin.api -dir .

gen_service:
	cd hr-service && goctl rpc protoc hr.proto --go_out=./pb --go-grpc_out=./pb --zrpc_out=.
	cd hr-admin-service && goctl rpc protoc admin.proto --go_out=./pb --go-grpc_out=./pb --zrpc_out=.

build_go:
	cd hr-api && go build -o ../build/hr-api -ldflags="-w -s" .
	cd hr-admin-api && go build -o ../build/hr-admin-api -ldflags="-w -s" .
	cd hr-service && go build -o ../build/hr-service -ldflags="-w -s" .
	cd hr-admin-service && go build -o ../build/hr-admin-service -ldflags="-w -s" .

docker:
	cd hr-api && goctl docker --version 1.20 && cd ..
	cd hr-admin-api && goctl docker --version 1.20 && cd ..
	cd hr-service && goctl docker --version 1.20 && cd ..
	cd hr-admin-service && goctl docker --version 1.20 && cd ..

init_config:
	cp common/config.yaml.example common/config.yaml
