package main

import (
	db "com.alibugrat/database"
	pb "com.alibugrat/product"
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
)

func getAddress() string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	return fmt.Sprintf("%s:%s",os.Getenv("serverIp"), os.Getenv("serverPort"))
}

func NewProductManagementServer() *ProductManagementServer {
	return &ProductManagementServer{}
}

type ProductManagementServer struct {
	pb.UnimplementedProductManagementServer
}

func (server *ProductManagementServer ) Run() error {
	lis, err := net.Listen("tcp", getAddress())
	if err != nil{
		log.Fatalf("failed to listen: %v",err)
	}

	s :=grpc.NewServer()
	pb.RegisterProductManagementServer(s, server)
	log.Printf("Server listening at %v",lis.Addr())

	return s.Serve(lis)
}

func (server ProductManagementServer) GetProducts(_ context.Context, _ *pb.GetProductParams)  (*pb.ProductList, error) {
	var productList *pb.ProductList = &pb.ProductList{}
	rows,err := db.DbConn.Query( "select productId,manufacturer,productName from inventorydb.products")
	if err != nil{
		return nil, err
	}
	defer rows.Close()
	for rows.Next(){
		product := pb.Product{}
		err = rows.Scan(&product.ProductId, &product.Manufacturer, &product.ProductName)
		if err != nil{
			return nil, err
		}
		productList.Product = append(productList.Product, &product)
	}
	log.Print("Return Response")
	return productList, nil

}

func (server ProductManagementServer) GetProduct(ctx context.Context, in *pb.GetProductParamsWithId)  (*pb.Product, error) {
	var product *pb.Product = &pb.Product{}
	query := fmt.Sprintf("select productId,manufacturer,productName from inventorydb.products where productId=%d",in.ProductId)
	rows,err := db.DbConn.Query( query )
	log.Printf("Result: %v",rows)
	if err != nil{
		return nil, err
	}
	defer rows.Close()
	e := rows.Scan(&product.ProductId, &product.Manufacturer, &product.ProductName)
	if e != nil {
		return nil, e
	}
	for rows.Next(){
		err = rows.Scan(&product.ProductId, &product.Manufacturer, &product.ProductName)
		if err != nil{
			return nil, err
		}

	}

	return product, nil

}

func main() {
	db.SetupDatabase()
	var productManagementServer *ProductManagementServer = NewProductManagementServer()
	if err := productManagementServer.Run(); err != nil{
		log.Fatalf("failed to server: %v",err)
	}
}