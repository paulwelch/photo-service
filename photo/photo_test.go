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
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"reflect"
	"regexp"
	"testing"
)

//service mock

type testService struct{}

func (s *testService) New(context.Context, *photopb.NewRequest) (*photopb.Photo, error) {
	return nil, nil
}

func (s *testService) Get(context.Context, *photopb.GetRequest) (*photopb.Photo, error) {
	return nil, nil
}

func (s *testService) Update(context.Context, *photopb.UpdateRequest) (*photopb.Photo, error) {
	return nil, nil
}

func (s *testService) List(context.Context, *photopb.ListRequest) (*photopb.PhotoList, error) {
	return nil, nil
}

func (s *testService) Delete(context.Context, *photopb.DeleteRequest) (*photopb.Empty, error) {
	return nil, nil
}

//Tests

func TestRegisterService(t *testing.T) {
	service := &testService{}
	RegisterService(service)
	if !reflect.DeepEqual(Service(), service) {
		t.FailNow()
	}
}

func TestRegisterPhotoServer(t *testing.T) {
	serviceInterface = &testService{}
	server := grpc.NewServer()
	RegisterPhotoServer(server)
}

func TestService(t *testing.T) {
	serviceInterface = &testService{}
	if !reflect.DeepEqual(Service(), serviceInterface) {
		t.FailNow()
	}
	serviceInterface = nil
	defer func() {
		if r := recover(); r == nil {
			t.Fatal(r)
		}
	}()
	Service()
}

func TestReadMethods(t *testing.T) {
	methods := []*regexp.Regexp{
		regexp.MustCompile(baseMethod + "Get"),
		regexp.MustCompile(baseMethod + "List"),
	}
	for k, v := range ReadMethods() {
		if v.String() != methods[k].String() {
			t.FailNow()
		}
	}
}

func TestWriteMethods(t *testing.T) {
	methods := []*regexp.Regexp{
		regexp.MustCompile(baseMethod + "New"),
		regexp.MustCompile(baseMethod + "Update"),
		regexp.MustCompile(baseMethod + "Delete"),
	}
	for k, v := range WriteMethods() {
		if v.String() != methods[k].String() {
			t.FailNow()
		}
	}
}
