#!/usr/bin/env bash
docker run -it --link cassandra-dev:cassandra --rm cassandra:latest cqlsh cassandra
