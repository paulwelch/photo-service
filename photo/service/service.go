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
	"github.com/gocql/gocql"
	"github.com/paulwelch/photo-service/photo/data"
	"github.com/paulwelch/photo-service/photo/photopb"
	"golang.org/x/net/context"
)

const ns = "photo"

type photos []*photopb.Photo

func (p *photos) GetNamespace() string { return ns }

type photo struct {
	photopb.Photo
}

func (p *photo) GetNamespace() string { return ns }

func (p *photo) SetId(id string) { p.Id = id }

func (p *photo) SetCreated(t int64) { p.Created = t }

func (p *photo) setUpdated(t int64) { p.Updated = t }

type PhotoService struct{}

func (s *PhotoService) New(ctx context.Context, req *photopb.NewRequest) (*photopb.Photo, error) {
	// TODO: validate data before inserting
	p := photo{
		Photo: photopb.Photo{
			Name: req.GetName(),
		},
	}

	var sess *gocql.Session
	var err error
	if sess, err = data.NewSession(); err != nil {
		return nil, err
	}
	if err = data.Insert(sess, &p.Photo); err != nil {
		return nil, err
	}
	sess.Close()

	return &p.Photo, nil
}

func (s *PhotoService) Get(ctx context.Context, req *photopb.GetRequest) (*photopb.Photo, error) {
	p := &photo{
		Photo: photopb.Photo{
			Id: req.Id,
		},
	}

	var sess *gocql.Session
	var err error
	var r *photopb.Photo
	if sess, err = data.NewSession(); err != nil {
		return nil, err
	}
	if r, err = data.Get(sess, p.Photo.Name); err != nil {
		return nil, err
	}
	sess.Close()

	return r, nil
}

func (s *PhotoService) List(ctx context.Context, req *photopb.ListRequest) (*photopb.PhotoList, error) {
	p := &photos{}

	n := 0
	// TODO: implement data.List
	// n, err := data.List(p, object.ListOpt{
	// 	Limit: req.GetLimit(),
	// 	Page:  req.GetPage(),
	// 	Sort:  object.SortNatural,
	// })

	return &photopb.PhotoList{Photos: *p, Count: int32(n)}, nil
}

func (s *PhotoService) Update(ctx context.Context, req *photopb.UpdateRequest) (*photopb.Photo, error) {
	p := &photo{
		Photo: photopb.Photo{
			Id: req.Id,
		},
	}

	if x := req.GetName(); x != "" {
		p.Photo.Name = x
	}

	var sess *gocql.Session
	var err error
	if sess, err = data.NewSession(); err != nil {
		return nil, err
	}
	if err = data.Update(sess, &p.Photo); err != nil {
		return nil, err
	}
	sess.Close()

	return &p.Photo, nil
}

func (s *PhotoService) Delete(ctx context.Context, req *photopb.DeleteRequest) (*photopb.Empty, error) {
	p := &photo{
		Photo: photopb.Photo{
			Id: req.Id,
		},
	}

	var sess *gocql.Session
	var err error
	if sess, err = data.NewSession(); err != nil {
		return nil, err
	}
	if err = data.Delete(sess, &p.Photo); err != nil {
		return nil, err
	}
	sess.Close()

	return &photopb.Empty{}, nil
}
