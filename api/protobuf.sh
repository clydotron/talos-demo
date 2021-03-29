protoc api/cluster/cluster.proto --go_out=plugins=grpc:.

// better
 protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    cluster/cluster.proto