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

// Package repository Implements repository access for accounty
package repository

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/Fs02/rel"
	"github.com/Fs02/rel/adapter/postgres"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres" // PostgreSQL migration
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq" // for migration
	"go.uber.org/zap"
)

var (
	logger, _ = zap.NewProduction(zap.Fields(zap.String("type", "repository")))
	shutdowns []func() error
)

// InitRepository Initialize Repositories.
// Connect to accounty database, run migrations and initialize the ropository
func InitRepository() rel.Repository {
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

	m, err := migrate.New(
		"file://db/migrations", dsn)
	if err != nil {
		log.Fatal(err)
	}
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("Unable to migrate up to the latest database schema - %v", err)
	}

	return repository
}
