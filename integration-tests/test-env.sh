#!/bin/bash

# export POSTGRES_DB=testdb
# export POSTGRES_USER=testuser
export POSTGRES_PASSWORD=$(hostname | md5sum | cut -d' ' -f 1)
