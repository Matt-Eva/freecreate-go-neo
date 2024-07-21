#!/bin/bash

go build cmd/seed/seed_data.go

./seed_data "$@"