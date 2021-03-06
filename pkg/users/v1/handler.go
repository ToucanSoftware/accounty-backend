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
	"net"
	"net/http"
	"strconv"

	model "github.com/ToucanSoftware/accounty-backend/pkg/users/model"
	security "github.com/ToucanSoftware/accounty-backend/pkg/users/security"
	service "github.com/ToucanSoftware/accounty-backend/pkg/users/service"

	"github.com/Fs02/rel"
	"github.com/golang/protobuf/proto"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
)

// Server represents the gRPC server for User Managemenet Services
type Server struct {
	Repo    model.UserRepository
	Service service.UserService
}

const (
	authenticationHeader = "X-ACCOUNTY-ACCESS-TOKEN"
)

var (
	logger, _ = zap.NewProduction(zap.Fields(zap.String("type", "v1")))
)

// CreateUser Creates a new User in the system
func (s *Server) CreateUser(ctx context.Context, in *CreateUserRequest) (*CreateUserResponse, error) {

	var user = model.User{
		Name:     in.User.Name,
		Username: in.User.Username,
		Email:    in.User.Email,
		Password: "1234",
	}

	_, err := s.Service.CreateUser(ctx, &user)

	if err != nil {
		logger.Error(err.Error(), zap.Error(err))
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
	var result []model.User

	err := s.Repo.FindAll(ctx, &result)

	if err != nil {
		logger.Error(err.Error(), zap.Error(err))
		return nil, err
	}

	var users = []*User{}

	for _, user := range result {
		users = append(users, UnmarsallUser(&user))
	}

	return &ListUsersResponse{
		Users: users}, nil
}

// GetUser Get an existing users
func (s *Server) GetUser(ctx context.Context, in *GetUserRequest) (*GetUserResponse, error) {
	var result model.User

	var id = in.Id

	err := s.Repo.Find(ctx, &result, id)

	if err != nil {
		if err == rel.ErrNotFound {
			return nil, status.Error(codes.NotFound,
				fmt.Sprintf("User with id %d not found.", id))
		}
		logger.Error(err.Error(), zap.Error(err))
		return nil, err
	}

	var user = UnmarsallUser(&result)

	return &GetUserResponse{
		User: user,
	}, nil
}

// UpdateUser Updates an existing users
func (s *Server) UpdateUser(ctx context.Context, in *UpdateUserRequest) (*UpdateUserResponse, error) {
	var result model.User
	var id = in.Id
	var err error

	err = s.Repo.Find(ctx, &result, id)

	if err != nil {
		if err == rel.ErrNotFound {
			return nil, status.Error(codes.NotFound,
				fmt.Sprintf("User with id %d not found.", id))
		}
		logger.Error(err.Error(), zap.Error(err))
		return nil, err
	}
	var user = MarsallUser(in.User)
	user.ID = id

	err = s.Repo.Update(ctx, user)

	if err != nil {
		logger.Error(err.Error(), zap.Error(err))
		return nil, err
	}

	return &UpdateUserResponse{
		User: &User{
			Id:       user.ID,
			Name:     user.Name,
			Username: user.Username,
			Email:    user.Email,
		},
	}, nil
}

// DeleteUser Deletes an existing user
func (s *Server) DeleteUser(ctx context.Context, in *DeleteUserRequest) (*DeleteUserResponse, error) {
	var result model.User
	var id = in.Id
	var err error

	err = s.Repo.Find(ctx, &result, id)

	if err != nil {
		if err == rel.ErrNotFound {
			return nil, status.Error(codes.NotFound,
				fmt.Sprintf("User with id %d not found.", id))
		}
		logger.Error(err.Error(), zap.Error(err))
		return nil, err
	}

	err = s.Repo.Delete(ctx, &result)

	if err != nil {
		logger.Error(err.Error(), zap.Error(err))
		return nil, err
	}

	return &DeleteUserResponse{}, nil
}

// AuthenticateUser Authenticate a User in the system
func (s *Server) AuthenticateUser(ctx context.Context, request *AuthenticateUserRequest) (*AuthenticateUserResponse, error) {
	var result model.User
	var err error

	err = s.Repo.FindByUsername(ctx, &result, request.Username)

	if err != nil {
		if err == rel.ErrNotFound {
			return nil, status.Error(codes.Unauthenticated, "Invalid username or password.")
		}
		logger.Error(err.Error(), zap.Error(err))
		return nil, err
	}

	hashedPassword := []byte(result.Password)
	password := []byte(request.Password)

	// Comparing the password with the hash
	err = bcrypt.CompareHashAndPassword(hashedPassword, password)

	if err != nil {
		if err == bcrypt.ErrMismatchedHashAndPassword {
			return nil, status.Error(codes.Unauthenticated, "Invalid username or password.")
		}
		logger.Error(err.Error(), zap.Error(err))
		return nil, err
	}

	accessToken, err := security.GenerateAccessToken(&result)

	if err != nil {
		logger.Error(err.Error(), zap.Error(err))
		return nil, err
	}

	// Create and send header.
	header := metadata.New(map[string]string{authenticationHeader: accessToken})
	grpc.SendHeader(ctx, header)

	return &AuthenticateUserResponse{User: UnmarsallUser(&result)}, nil
}

// StartUserManagemenetRESTServer Starts s REST Reverse proxy service for User Management
func StartUserManagemenetRESTServer(address, grpcAddress string) error {
	var err error
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux(runtime.WithForwardResponseOption(httpResponseModifier))

	// Setup the client gRPC options
	opts := []grpc.DialOption{grpc.WithInsecure()}

	// Register ping
	err = RegisterUserManagementServiceHandlerFromEndpoint(ctx, mux, grpcAddress, opts)
	if err != nil {
		return fmt.Errorf("could not register service User Management Service: %s", err)
	}

	logger.Info(fmt.Sprintf("Starting HTTP/1.1 REST server on %s", address))
	http.ListenAndServe(address, mux)

	return nil
}

// StartUserManagemenetGRPCServer Starts a gRPC Server for User Management
func StartUserManagemenetGRPCServer(address string, repo rel.Repository) error {
	// create a listener on TCP port
	lis, err := net.Listen("tcp", address)
	if err != nil {
		return fmt.Errorf("failed to listen: %v", err)
	}

	// create a server instance
	s := Server{
		Repo:    model.New(repo),
		Service: service.NewUserService(repo),
	}

	grpcServer := grpc.NewServer()

	// attach the Ping service to the server
	RegisterUserManagementServiceServer(grpcServer, &s)

	reflection.Register(grpcServer)

	// start the server
	logger.Info(fmt.Sprintf("Starting HTTP/2 gRPC server on %s", address))

	if err := grpcServer.Serve(lis); err != nil {
		return fmt.Errorf("failed to serve: %s", err)
	}

	return nil
}

// MarsallUser marsall model Object into gRPC object
func MarsallUser(value *User) *model.User {
	return &model.User{
		ID:       value.Id,
		Name:     value.Name,
		Username: value.Username,
		Email:    value.Email,
	}
}

// UnmarsallUser unmarsall gRPC User model object into User
func UnmarsallUser(value *model.User) *User {
	return &User{
		Id:       value.ID,
		Name:     value.Name,
		Username: value.Username,
		Email:    value.Email,
	}
}

func httpResponseModifier(ctx context.Context, w http.ResponseWriter, p proto.Message) error {
	md, ok := runtime.ServerMetadataFromContext(ctx)
	if !ok {
		return nil
	}

	// set http status code
	if vals := md.HeaderMD.Get("x-http-code"); len(vals) > 0 {
		code, err := strconv.Atoi(vals[0])
		if err != nil {
			return err
		}
		w.WriteHeader(code)
		// delete the headers to not expose any grpc-metadata in http response
		delete(md.HeaderMD, "x-http-code")
		delete(w.Header(), "Grpc-Metadata-X-Http-Code")
	}
	if vals := md.HeaderMD.Get(authenticationHeader); len(vals) > 0 {
		delete(md.HeaderMD, authenticationHeader)
		delete(w.Header(), "Grpc-Metadata-X-Accounty-Access-Token")
		w.Header().Add(authenticationHeader, vals[0])
	}

	return nil
}
