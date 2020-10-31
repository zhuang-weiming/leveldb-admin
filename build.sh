#!/usr/bin/env bash

bash -c "cd website && yarn install && yarn run build"

go generate ./static.go