module grpc

go 1.20


replace (
	hello => ./hello
)

require google.golang.org/grpc v1.64.0

require google.golang.org/protobuf v1.33.0 // indirect

