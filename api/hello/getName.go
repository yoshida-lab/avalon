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

package hello

import (
	"fmt"
	"github.com/gin-gonic/gin"
	pb "github.com/yoshida-lab/avalon/protobuf/helloworld"
	"google.golang.org/grpc"
	"net/http"
)

type GetName struct{}

func (g GetName) Register(group *gin.RouterGroup) {
	group.GET("/rest/n/:hello", func(c *gin.Context) {
		name := c.Param("hello")

		// Contact the server and print out its response.
		conn := c.MustGet("rpcConn").(*grpc.ClientConn)
		client := pb.NewGreeterClient(conn)
		req := &pb.HelloRequest{Name: name}
		res, err := client.SayHello(c, req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"result": fmt.Sprint(res.Message),
		})
	})
}
