#!/bin/sh

# Create the docker environment for testing
docker build -t=steve/runit - < tests/dockerfile.runit
