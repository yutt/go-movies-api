#!/bin/bash

go mod download 
swag init --parseInternal --parseDependency
