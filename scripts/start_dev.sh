#!/bin/sh

docker compose -f scripts/docker-compose.yml --env-file .env up -d
