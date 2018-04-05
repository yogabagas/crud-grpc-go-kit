package main

import (
	"context"
	"fmt"
	"time"

	cli "MiniProject/kampus.id/denda/katdenda/endpoint"
	opt "MiniProject/kampus.id/denda/util/grpc"
	util "MiniProject/kampus.id/denda/util/microservice"
	tr "github.com/opentracing/opentracing-go"
)

func main() {
	logger := util.Logger()
	tracer := tr.GlobalTracer()
	option := opt.ClientOption{Retry: 3, RetryTimeout: 500 * time.Millisecond, Timeout: 30 * time.Second}

	client, err := cli.NewGRPCDendaClient([]string{"127.0.0.1:2181"}, nil, option, tracer, logger)
	if err != nil {
		logger.Log("error", err)
	}

	//Add Mahasiswa
	//client.AddDendaService(context.Background(), svc.Denda{ID: "KD004", Jenis: "3", Jumlah: "500", CreatedBy: "yogabagas"})

	//Get Mahasiswa By Nim No
	//cusID, _ := client.ReadDendaByIDService(context.Background(), "KD004")
	//fmt.Println("denda based on id:", cusID)

	//List Customer
	//cuss, _ := client.ReadDendaService(context.Background())
	//fmt.Println("all denda:", cuss)

	//GetByKet
	cusKet, _ := client.ReadDendaByKeteranganService(context.Background(), "Ri%")
	fmt.Println("all denda", cusKet)

	//Update Customer
	//client.UpdateDendaService(context.Background(), svc.Denda{ID: "KD002", Jenis: "2", Jumlah: "500", Status: 1, UpdatedBy: "yoga"})

	/*
		//Get Customer By Email
		cusEmail, _ := client.ReadCustomerByEmailService(context.Background(), "joko@gmail.com")
		fmt.Println("customer based on email:", cusEmail)
	*/
}
