#!/usr/bin/env bash

bash -c "cd website && yarn run build"

go generate ./static.go