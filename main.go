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
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	usersv1 "github.com/ToucanSoftware/accounty-backend/api/users/v1"
	_ "github.com/dimiro1/banner/autoload"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	"github.com/Fs02/rel"
	"github.com/Fs02/rel/adapter/postgres"
	_ "github.com/lib/pq"
	"go.uber.org/zap"
)

var (
	logger, _ = zap.NewProduction(zap.Fields(zap.String("type", "main")))
	shutdowns []func() error
)

func initRepository() rel.Repository {
	var (
		logger, _ = zap.NewProduction(zap.Fields(zap.String("type", "repository")))
		dsn       = fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
			os.Getenv("POSTGRESQL_USERNAME"),
			os.Getenv("POSTGRESQL_PASSWORD"),
			os.Getenv("POSTGRESQL_HOST"),
			os.Getenv("POSTGRESQL_PORT"),
			os.Getenv("POSTGRESQL_DATABASE"))
	)

	adapter, err := postgres.Open(dsn)
	if err != nil {
		logger.Fatal(err.Error(), zap.Error(err))
	}
	// add to graceful shutdown list.
	shutdowns = append(shutdowns, adapter.Close)

	repository := rel.New(adapter)
	repository.Instrumentation(func(ctx context.Context, op string, message string) func(err error) {
		// no op for rel functions.
		if strings.HasPrefix(op, "rel-") {
			return func(error) {}
		}

		t := time.Now()

		return func(err error) {
			duration := time.Since(t)
			if err != nil {
				logger.Error(message, zap.Error(err), zap.Duration("duration", duration), zap.String("operation", op))
			} else {
				logger.Info(message, zap.Duration("duration", duration), zap.String("operation", op))
			}
		}
	})

	return repository
}

func gracefulShutdown(ctx context.Context, shutdown chan struct{}) {
	var (
		sigint = make(chan os.Signal, 1)
	)

	signal.Notify(sigint, os.Interrupt, syscall.SIGTERM)
	<-sigint

	logger.Info("shutting down server gracefully")

	// // stop receiving any request.
	// if err := server.Shutdown(ctx); err != nil {
	// 	logger.Fatal("shutdown error", zap.Error(err))
	// }

	// close any other modules.
	for i := range shutdowns {
		shutdowns[i]()
	}

	close(shutdown)
}

func main() {
	var (
		ctx = context.Background()
		//port       = os.Getenv("PORT")
		repository = initRepository()
		//mux        = api.NewMux(repository)
		// server     = http.Server{
		// 	Addr:    ":" + port,
		// 	Handler: mux,
		// }
		shutdown = make(chan struct{})
	)

	m, err := migrate.New(
		"file://db/migrations",
		"postgres://toucan:password@localhost:5432/accounty?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("Unable to migrate up to the latest database schema - %v", err)
	}

	grpcAddress := fmt.Sprintf("%s:%d", "0.0.0.0", 9090)
	restAddress := fmt.Sprintf("%s:%d", "0.0.0.0", 8081)
	certFile := "cert/server.crt"
	keyFile := "cert/server.key"

	// select {}
	go gracefulShutdown(ctx, shutdown)

	// fire the gRPC server in a goroutine
	go func() {
		err := usersv1.StartUserManagemenetGRPCServer(grpcAddress, certFile, keyFile, repository)
		if err != nil {
			log.Fatalf("failed to start gRPC server: %s", err)
		}
	}()

	// fire the REST server in a goroutine
	go func() {
		err := usersv1.StartUserManagemenetRESTServer(restAddress, grpcAddress, certFile)
		if err != nil {
			log.Fatalf("failed to start gRPC server: %s", err)
		}
	}()

	// infinite loop
	log.Printf("Entering infinite loop")

	<-shutdown
	select {}
}
