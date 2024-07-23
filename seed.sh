#!/bin/bash

go build cmd/seed/db_seeds.go

./db_seeds "$@"