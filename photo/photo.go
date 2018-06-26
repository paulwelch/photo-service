/*
 * Copyright 2018 Paul Welch
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with the License. You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the specific language governing permissions and limitations under the License.
 */

package photo

import (
	"github.com/paulwelch/photo-service/photo/photopb"
	servicePackage "github.com/paulwelch/photo-service/photo/service"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"os/signal"
	"regexp"
	"syscall"
)

const baseMethod = "^(.photopb.PhotoService/)"

var serviceInterface Interface

func init() {
	RegisterService(&servicePackage.PhotoService{})
}

type Interface interface {
	photopb.PhotoServiceServer
}

func RegisterService(p Interface) {
	if serviceInterface != nil {
		panic("PhotoService already registered")
	}
	serviceInterface = p
}

func Service() Interface {
	if serviceInterface == nil {
		panic("PhotoService not registered")
	}
	return serviceInterface
}

func RegisterPhotoServer(server *grpc.Server) {
	photopb.RegisterPhotoServiceServer(server, Service())
}

func ReadMethods() []*regexp.Regexp {
	return []*regexp.Regexp{
		regexp.MustCompile(baseMethod + "Get"),
		regexp.MustCompile(baseMethod + "List"),
	}
}

func WriteMethods() []*regexp.Regexp {
	return []*regexp.Regexp{
		regexp.MustCompile(baseMethod + "New"),
		regexp.MustCompile(baseMethod + "Update"),
		regexp.MustCompile(baseMethod + "Delete"),
	}
}

func Run() {

	lis, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	g := grpc.NewServer()
	RegisterPhotoServer(g)
	//TODO: ssl, auth, logging, etc

	c := make(chan os.Signal, 2)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		for sig := range c {
			// handle ^C
			log.Println("Shutdown Server : Signal %s", sig.String())
			g.GracefulStop()
		}
	}()
	if err := g.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
