gen-cal:
	protoc ./calculator/proto/calculator.proto --go_out=./calculator/proto --go-grpc_out=./calculator/proto --go_opt=module=github.com/Chungbien/udemy-grpc-course/calculator/proto --go-grpc_opt=module=github.com/Chungbien/udemy-grpc-course/calculator/proto
run-server:
	go run calculator/server/main.go
run-client:
	go run calculator/client/main.go
gen-stream:
	protoc ./stream/proto/stream.proto --go_out=./stream/proto --go-grpc_out=./stream/proto --go_opt=module=github.com/Chungbien/udemy-grpc-course/stream/proto --go-grpc_opt=module=github.com/Chungbien/udemy-grpc-course/stream/proto
run-server-stream:
	go run stream/server/main.go
run-client-stream:
	go run stream/client/main.go