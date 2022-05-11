# Create grpc
protoc *.proto --go_out="../server" --go-grpc_out="../server"

#Create grpc gateway
protoc --grpc-gateway_out="logtostderr=true:../server" *.proto