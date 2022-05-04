gen-cal:
	protoc ./calculator/proto/calculator.proto --go_out=./calculator/proto --go-grpc_out=./calculator/proto --go_opt=module=github.com/Chungbien/udemy-grpc-course/calculator/proto --go-grpc_opt=module=github.com/Chungbien/udemy-grpc-course/calculator/proto
run-server:
	go run calculator/server/main.go
run-client:
	go run calculator/client/main.go