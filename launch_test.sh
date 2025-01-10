#!/bin/bash

docker-compose -f ./internal/database/test/docker-compose-test.yml --env-file ./internal/database/test/.env_test up --build -d

docker-compose -f ./internal/database/test/docker-compose-test.yml --env-file ./internal/database/test/.env_test down
