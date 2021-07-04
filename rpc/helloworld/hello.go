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

package helloworld

import (
	"context"
	pb "github.com/yoshida-lab/avalon/protobuf/helloworld"
	"google.golang.org/grpc"
	"log"
)

// Service is used to implement helloworld.GreeterServer.
type Service struct {
	pb.UnimplementedGreeterServer
}

// SayHello implements helloworld.GreeterServer
func (srv *Service) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &pb.HelloReply{Message: "Hello " + in.GetName()}, nil
}

func (srv *Service) Register(s grpc.ServiceRegistrar) {
	pb.RegisterGreeterServer(s, srv)
}
