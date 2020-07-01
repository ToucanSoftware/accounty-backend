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

package v1

import (
	"context"
	"fmt"
	"github.com/Fs02/rel"
	model "github.com/ToucanSoftware/accounty-backend/pkg/users/model"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"net/http"
)

// Server represents the gRPC server for User Managemenet Services
type Server struct {
	Repo model.UserRepository
}

// CreateUser Creates a new User in the system
func (s *Server) CreateUser(ctx context.Context, in *CreateUserRequest) (*CreateUserResponse, error) {

	var user = model.User{
		Name:     in.User.Name,
		Username: in.User.Username,
		Email:    in.User.Email,
		Password: "1234",
	}

	err := s.Repo.Create(ctx, &user)

	if err != nil {
		log.Printf("Error: %v", err)
		return nil, err
	}

	return &CreateUserResponse{
		User: &User{
			Id:       user.ID,
			Name:     user.Name,
			Username: user.Username,
			Email:    user.Email,
		},
	}, nil
}

// ListUsers List users in the system
func (s *Server) ListUsers(ctx context.Context, in *ListUsersRequest) (*ListUsersResponse, error) {
	var user = model.User{}
	err := s.Repo.Find(ctx, &user, 1)

	if err != nil {
		log.Fatalf("Error %v", err)
	}

	var users = []*User{
		&User{
			Id:   user.ID,
			Name: user.Name,
		},
	}

	return &ListUsersResponse{
		Users: users}, nil
}

// StartUserManagemenetRESTServer Starts s REST Reverse proxy service for User Management
func StartUserManagemenetRESTServer(address, grpcAddress, certFile string) error {
	var err error
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()

	// Setup the client gRPC options
	opts := []grpc.DialOption{grpc.WithInsecure()}

	// Register ping
	err = RegisterUserManagementServiceHandlerFromEndpoint(ctx, mux, grpcAddress, opts)
	if err != nil {
		return fmt.Errorf("could not register service User Management Service: %s", err)
	}

	log.Printf("starting HTTP/1.1 REST server on %s", address)
	http.ListenAndServe(address, mux)

	return nil
}

// StartUserManagemenetGRPCServer Starts a gRPC Server for User Management
func StartUserManagemenetGRPCServer(address, certFile, keyFile string, repo rel.Repository) error {
	// create a listener on TCP port
	lis, err := net.Listen("tcp", address)
	if err != nil {
		return fmt.Errorf("failed to listen: %v", err)
	}

	// create a server instance
	s := Server{Repo: model.New(repo)}

	grpcServer := grpc.NewServer()

	// attach the Ping service to the server
	RegisterUserManagementServiceServer(grpcServer, &s)

	reflection.Register(grpcServer)

	// start the server
	log.Printf("starting HTTP/2 gRPC server on %s", address)
	if err := grpcServer.Serve(lis); err != nil {
		return fmt.Errorf("failed to serve: %s", err)
	}

	return nil
}
