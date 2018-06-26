/*
 * Copyright 2018 Paul Welch
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with the License. You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the specific language governing permissions and limitations under the License.
 */

package main

import (
	"github.com/paulwelch/photo-service/photo/photopb"
	"github.com/paulwelch/photo-service/photo/service"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"testing"
)

var photoService = &service.PhotoService{}

func TestClientNew(t *testing.T) {

	go main()

	testName := "Test Main Service"

	c, err := grpc.Dial("127.0.0.1:8080", grpc.WithInsecure())
	if err != nil {
		t.Fatal(err)
	}
	defer c.Close()

	client := photopb.NewPhotoServiceClient(c)
	resp, err := client.New(context.Background(), &photopb.NewRequest{Name: testName})
	if err != nil {
		t.Fatal(err)
	}
	println(resp)

}
