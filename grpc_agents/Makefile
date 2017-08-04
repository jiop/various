build:
	protoc -I status/ status/status.proto --go_out=plugins=grpc:status
	protoc -I advertise/ advertise/advertise.proto --go_out=plugins=grpc:advertise