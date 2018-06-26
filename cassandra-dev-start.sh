#!/usr/bin/env bash
docker run --rm --name cassandra-dev -e CASSANDRA_BROADCAST_ADDRESS=$(ifconfig en0 | grep inet | grep -v inet6 | cut -d ' ' -f2) -v $(pwd)/cassandra-data:/var/lib/cassandra -p 7000:7000 -p 9042:9042 -p 9160:9160 -d cassandra:latest
