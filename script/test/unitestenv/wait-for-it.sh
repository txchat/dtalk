#!/bin/bash

echo "waiting for message queue..."
is_healthy() {
    service="$1"
    echo "$service"
    container_id="$(docker compose -f components.compose.yaml ps -q "$service")"
    health_status="$(docker inspect -f "{{.State.Health.Status}}" "$container_id")"

    if [ "$health_status" = "healthy" ]; then
        return 0
    else
        return 1
    fi
}

while ! is_healthy "mysql"; do sleep 1; done

echo "all start"
