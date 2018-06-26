/*
 * Copyright 2018 Paul Welch
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with the License. You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the specific language governing permissions and limitations under the License.
 */

package service

import (
	"github.com/paulwelch/photo-service/photo/photopb"
	"golang.org/x/net/context"
	"testing"
)

var service = &PhotoService{}

func TestPhotoService_New(t *testing.T) {
	testName := "test service"

	p, err := service.New(context.Background(), &photopb.NewRequest{
		Name: testName,
	})
	if err != nil {
		t.Fatal(err)
	}

	if p.Name != testName {
		t.Error("New photo name does not match")
	}
}
