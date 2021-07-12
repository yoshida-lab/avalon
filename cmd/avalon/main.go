/*
 * Copyright (c) 2021 TsumiNa
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/yoshida-lab/avalon/api"
	"github.com/yoshida-lab/avalon/rpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	// any error should be shutdown gracefully
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Error on: `%v`, close all resources", r)
		}
	}()

	// -----------------------------------------------------
	// gRPC part
	// -----------------------------------------------------
	sockFile, err := os.CreateTemp("", "rpc-*-.sock")
	if err != nil {
		log.Fatalf("UDS file creation error: %v", err)
	}

	// we need to remove this file
	// because listener will create new one itself
	if err := os.Remove(sockFile.Name()); err != nil {
		log.Fatalf("Remove socket file error: %v", err)
	}

	// listen on the UDS
	lis, err := net.Listen("unix", sockFile.Name())
	if err != nil {
		log.Fatalf("gRPC server failed to listen on %v; Error: %v", sockFile.Name(), err)
	}

	// new rpc server instance
	rpcSrv := grpc.NewServer()

	// register services
	rpc.RegisterService(rpcSrv)

	// Register reflection service on gRPC server.
	reflection.Register(rpcSrv)

	// Initializing the server in a goroutine
	go func() {
		if err := rpcSrv.Serve(lis); err != nil && errors.Is(err, http.ErrServerClosed) {
			log.Printf("gRPC server listen: %s\n", err)
		}
	}()

	// --------------------------------------------------
	// API server part
	// --------------------------------------------------

	// gRPC client
	conn, err := grpc.Dial("unix:///"+sockFile.Name(), grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer func(conn *grpc.ClientConn) {
		_ = conn.Close()
	}(conn)

	router := gin.Default()
	router.Use(RPCConnection(conn))

	// Simple group: v1
	v1 := router.Group("/v1")
	api.RegisterRoute(v1)

	// Run http server
	apiServer := &http.Server{
		Addr:         ":8052",
		Handler:      router,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	// Initializing the server in a goroutine
	go func() {
		if err := apiServer.ListenAndServe(); err != nil && errors.Is(err, http.ErrServerClosed) {
			log.Printf("%s\n", err)
		}
	}()

	// graceful shutdown
	<-quit
	log.Println("Shutting down servers...")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	if err := apiServer.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	rpcSrvDone := make(chan bool)
	go func() {
		rpcSrv.GracefulStop()
		rpcSrvDone <- true

	}()
	select {
	case <-ctx.Done():
		log.Println("gRPC server shutting down out of time, close it immediately")
	case <-rpcSrvDone:
		log.Println("gRPC: Server closed")
	}

	log.Println("Bye!!")
}
