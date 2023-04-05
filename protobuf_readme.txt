1. run codeGenerator.go
2. run proto generator with grpc
3. walk through all -non-proto- generated files, make sure no errors are there
4. run tests
5. run server
6. run UI


----------------------------------PROTOCOL BUFFERS-------------------------------------

# has to be installed to compile protoc:

go get -u github.com/golang/protobuf/protoc-gen-go@v1.3.2
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.26
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1


## examples to run protoc compiler to generate proto go files ##

# this actually works
protoc -I ./src --go_out=src/ simple/simple.proto

# with grpc
protoc -I ./src --go_out=src/ --go-grpc_out=src/ simple/simple.proto 

# https://github.com/infobloxopen/protoc-gen-gorm
# to install  gorm generator for protobuffer generators:
go get github.com/infobloxopen/protoc-gen-gorm

# generate protobuffers with gorm:
protoc -I ./src --go_out=src/ --go-grpc_out=src/ --gorm_out=src/ simple/simple.proto 


protoc -I ./proto --go_out=./ --go-grpc_out=./ proto/upm.proto 



## DART ##
# from Frontend/UPM_v0_Flutter --> where pubspec.yaml resides
dart pub get protobuf.dart/protoc_plugin/


protoc -I ./lib/proto --dart_out=grpc:lib/proto/ --proto_path=./../../backend/proto/ upm.proto