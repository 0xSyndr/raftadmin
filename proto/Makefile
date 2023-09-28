raftadmin.pb.go: raftadmin.proto
	protoc --go_out=../ --go-grpc_out=../ raftadmin.proto

force:
	make raftadmin.pb.go
