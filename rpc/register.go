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

package rpc

import (
	"github.com/yoshida-lab/avalon/rpc/helloworld"
	"google.golang.org/grpc"
)

type Service interface {
	Register(s grpc.ServiceRegistrar)
}

var services []Service

// RegisterService register gRPC services into gRPC server
func RegisterService(s grpc.ServiceRegistrar) {
	for _, srv := range services {
		srv.Register(s)
	}
}

// Gather all services
func init() {
	services = append(services, &helloworld.Service{})
}
