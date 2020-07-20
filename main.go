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

package main

import (
	"fmt"

	utils "github.com/ToucanSoftware/accounty-backend/internal/utils"
	usersRepository "github.com/ToucanSoftware/accounty-backend/pkg/users/repository"
	usersv1 "github.com/ToucanSoftware/accounty-backend/pkg/users/v1"

	_ "github.com/dimiro1/banner/autoload"
	_ "github.com/lib/pq"
	"go.uber.org/zap"
)

var (
	logger, _ = zap.NewProduction(zap.Fields(zap.String("type", "main")))
)

const (
	// Default port for REST Web API
	defaultRestPort = "8081"

	// Default port for gRPC API
	defaultGpcpPort = "9090"
)

func main() {
	var (
		repository = usersRepository.InitRepository()
		grpcPort   = utils.GetEnv("GRPC_PORT", defaultGpcpPort)
		restPort   = utils.GetEnv("REST_PORT", defaultRestPort)
	)

	grpcAddress := fmt.Sprintf("%s:%s", "0.0.0.0", grpcPort)
	restAddress := fmt.Sprintf("%s:%s", "0.0.0.0", restPort)

	// fire the gRPC server in a goroutine
	go func() {
		err := usersv1.StartUserManagemenetGRPCServer(grpcAddress, repository)
		if err != nil {
			logger.Error(err.Error(), zap.Error(err))
		}
	}()

	// fire the REST server in a goroutine
	go func() {
		err := usersv1.StartUserManagemenetRESTServer(restAddress, grpcAddress)
		if err != nil {
			logger.Error(err.Error(), zap.Error(err))
		}
	}()

	// infinite loop
	logger.Info("Entering infinite loop")

	select {}
}
