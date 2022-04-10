module github.com/ekkinox/go-grpc/greet/client

go 1.18

replace github.com/ekkinox/go-grpc/greet/proto => ../proto

require (
	github.com/ekkinox/go-grpc v0.0.0-20220409191643-d95599159138
	github.com/ekkinox/go-grpc/greet/proto v0.0.0-20220410120422-074034508dcf
	google.golang.org/grpc v1.45.0
)

require (
	github.com/golang/protobuf v1.5.2 // indirect
	golang.org/x/net v0.0.0-20200822124328-c89045814202 // indirect
	golang.org/x/sys v0.0.0-20200323222414-85ca7c5b95cd // indirect
	golang.org/x/text v0.3.0 // indirect
	google.golang.org/genproto v0.0.0-20200526211855-cb27e3aa2013 // indirect
	google.golang.org/protobuf v1.28.0 // indirect
)
