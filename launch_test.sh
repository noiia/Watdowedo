#!/bin/bash

export COMPOSE_BAKE=true
docker-compose -f ./docker-compose.yml --project-name test --env-file ./test/.env_test up --build 
