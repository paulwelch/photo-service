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
	"github.com/google/uuid"
	"github.com/paulwelch/photo-service/photo/photopb"
	"log"
	"time"
)

//TODO: find an appropriate place for storage initialization
//CREATE KEYSPACE funphotos WITH REPLICATION = { 'class' : 'SimpleStrategy', 'replication_factor' : 1 };
//CREATE TABLE funphotos.photos (id UUID PRIMARY KEY, name TEXT, created TIMESTAMP, updated TIMESTAMP);

func NewSession() (*gocql.Session, error) {
	//TODO: parameterize connection properties
	cluster := gocql.NewCluster("127.0.0.1")
	cluster.Port = 9042
	cluster.ProtoVersion = 3
	cluster.Keyspace = "funphotos"

	var s *gocql.Session
	var err error
	if s, err = cluster.CreateSession(); err != nil {
		println("Error creating the session")
		log.Fatal(err)
		return nil, err
	}
	return s, nil
}

func Insert(s *gocql.Session, p *photopb.Photo) error {
	ts := time.Now()
	p.Id = uuid.New().String()
	p.Created = ts.Unix()
	p.Updated = ts.Unix()

	if err := s.Query("INSERT INTO funphotos.photos (id, name, created, updated) VALUES (?, ?, ?, ?)", p.Id, p.Name, ts, ts).Exec(); err != nil {
		log.Fatal(err)
		return fmt.Errorf("data access: unable to insert record: %v", err)
	}

	return nil
}

func Get(s *gocql.Session, name string) (*photopb.Photo, error) {
	var id string
	var created, updated int64

	if err := s.Query("SELECT id, name, created, updated FROM funphotos.photos WHERE name='Test Photo 123' ALLOW FILTERING").Scan(&id, &name, &created, &updated); err != nil {
		log.Fatal(err)
		return nil, err
	}

	p := &photopb.Photo{
		Id:      id,
		Name:    name,
		Created: created,
		Updated: updated,
	}

	return p, nil
}

func Update(s *gocql.Session, p *photopb.Photo) error {
	if err := s.Query("UPDATE funphotos.photos SET name=?, updated=? WHERE id=? IF EXISTS", p.Name, time.Now(), p.Id).Exec(); err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

func Delete(s *gocql.Session, p *photopb.Photo) error {
	if err := s.Query("DELETE FROM funphotos.photos WHERE id = ?", p.Id).Exec(); err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}
