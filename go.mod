module github.com/silverliningco/ahri_grpc

go 1.16

replace github.com/silverliningco/ahri_grpc => ../ahri_grpc

require (
	github.com/lib/pq v1.10.1
	google.golang.org/grpc v1.42.0
	google.golang.org/protobuf v1.27.1
)
