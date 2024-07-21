#!/bin/bash

go build cmd/seed/seed.go

./seed "$@"