/*
 * Copyright 2018 Paul Welch
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with the License. You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the specific language governing permissions and limitations under the License.
 */

package data

import (
	"fmt"
	"github.com/gocql/gocql"
	"github.com/paulwelch/photo-service/photo/photopb"
	"testing"
	"time"
)

type testData struct{}

func TestInsert(t *testing.T) {
	p := &photopb.Photo{
		Name: "Test Photo 123",
	}

	var s *gocql.Session
	var err error
	if s, err = NewSession(); err != nil {
		t.Errorf("Error creating session: %v", err)
	}
	Insert(s, p)

	fmt.Println("Inserted: " + p.Id)

	s.Close()
}

func TestGet(t *testing.T) {
	var s *gocql.Session
	var err error
	var p *photopb.Photo
	if s, err = NewSession(); err != nil {
		t.Errorf("Error creating session: %v", err)
	}
	name := "Test Photo 123"
	if p, err = Get(s, name); err != nil {
		t.Errorf("Error getting record: %v", err)
	}

	fmt.Println("Found: " + p.Id)

	s.Close()
}

func TestUpdate(t *testing.T) {
	var s *gocql.Session
	var err error
	var p *photopb.Photo
	if s, err = NewSession(); err != nil {
		t.Errorf("Error creating session: %v", err)
	}
	name := "New Name 123"
	if p, err = Get(s, name); err != nil {
		t.Errorf("Error getting record: %v", err)
	}
	newName := "Test Photo 123"
	p.Name = newName
	Update(s, p)

	//verify
	p, err = Get(s, newName)
	fmt.Println("Updated Name: " + p.Name)
	fmt.Println("Updated Time: " + time.Unix(p.Updated, 0).String())

	if p.Name != newName {
		t.Error("New name not set")
	}

	s.Close()
}

func TestDelete(t *testing.T) {
	var s *gocql.Session
	var err error
	var p *photopb.Photo
	if s, err = NewSession(); err != nil {
		t.Errorf("Error creating session: %v", err)
	}
	name := "Test Photo 123"
	if p, err = Get(s, name); err != nil {
		t.Errorf("Error deleting record: %v", err)
	}
	Delete(s, p)

	s.Close()
}
