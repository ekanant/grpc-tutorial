protoc *.proto --go_out="../server"
protoc --grpc-gateway_out="logtostderr=true:../server" *.proto
