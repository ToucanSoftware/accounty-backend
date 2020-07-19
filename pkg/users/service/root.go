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

package service

import (
	"context"
	"fmt"

	"github.com/Fs02/rel"
	model "github.com/ToucanSoftware/accounty-backend/pkg/users/model"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/go-playground/validator.v9"
)

var (
	logger, _ = zap.NewProduction(zap.Fields(zap.String("type", "service")))
)

// UserService Define services for User Management
type UserService interface {
	// Create a User
	CreateUser(ctx context.Context, request *model.User) (*model.User, error)
}

type userService struct {
	Repo      model.UserRepository
	Validator *validator.Validate
}

func (s *userService) CreateUser(ctx context.Context, request *model.User) (*model.User, error) {
	// Validate the user
	err := s.Validator.Struct(request)

	if err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			fmt.Println(e)
		}
		logger.Error(err.Error(), zap.Error(err))
		return nil, err
	}

	password := []byte(request.Password)

	// Hashing the password with the default cost of 10
	hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		logger.Fatal(err.Error(), zap.Error(err))
		return nil, err
	}
	fmt.Println(string(hashedPassword))

	request.Password = string(hashedPassword)

	err = s.Repo.Create(ctx, request)
	if err != nil {
		logger.Fatal(err.Error(), zap.Error(err))
		return nil, err
	}

	return request, nil
}

// NewUserService Create a new User Service
func NewUserService(repo rel.Repository) UserService {
	return &userService{
		Repo:      model.New(repo),
		Validator: validator.New(),
	}
}
