// Licensed to Toucan Software Consulting SAS under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Toucan Software Consulting SAS licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package model

import (
	"context"
	"github.com/Fs02/rel"
	"github.com/Fs02/rel/where"
	"time"
)

// User is a model that maps to users table.
type User struct {
	ID        int64
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

// UserRepository interfaces for accessing user data
type UserRepository interface {
	Find(ctx context.Context, user *User, id int64) error
}

// actual implementation
type userRepository struct {
	repository rel.Repository
}

func (ur userRepository) Find(ctx context.Context, user *User, id int64) error {
	ur.repository.Find(ctx, user, where.Eq("id", id))
	return nil
}

// New returns a new repository
func New(repo rel.Repository) UserRepository {
	return userRepository{
		repository: repo,
	}
}
