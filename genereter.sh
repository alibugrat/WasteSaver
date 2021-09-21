protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative product/product.proto

#for linux docker
docker build --platform linux/amd64 --tag alibugrat/waste-saver-grpc-server-amd64 .

#for mac m1 docker
docker build --tag alibugrat/waste-saver-grpc-server .
