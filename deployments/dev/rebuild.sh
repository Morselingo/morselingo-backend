#!/bin/bash

# echo "Stopping and removing containers, images, and orphans..."
# docker compose down --rmi all --remove-orphans

echo "Rebuilding images without cache..."
docker compose build --no-cache

echo "Starting services in detached mode..."
docker compose up