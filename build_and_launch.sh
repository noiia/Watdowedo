#!/bin/bash

export COMPOSE_BAKE=true
docker-compose -f ./docker-compose.yml --project-name prod --env-file ./.env up --build 
