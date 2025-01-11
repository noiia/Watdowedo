#!/bin/bash

docker-compose -f ./docker-compose.yml --env-file ./test/.env_test up --build -d

docker-compose -f ./docker-compose.yml --env-file ./test/.env_test down
