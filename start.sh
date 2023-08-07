#!/bin/bash

echo -e "\033[0;34m*** Run Docker***\033[0;0m"; \
docker-compose up -d;
echo -e "\033[0;32mFINISHED\033[0;0m"

export XE_POSTGRES_HOST="localhost:3306";
export XE_POSTGRES_DB="nam_0508";
export XE_POSTGRES_USER="root";
export XE_POSTGRES_PASSWORD="mysql_db";

echo -e "\033[0;34m*** Run Migrate***\033[0;0m"; \
migrate -path sql -database "mysql://root:mysql_db@tcp(localhost:3306)/nam_0508?net_write_timeout=6000" -verbose up ;
echo -e "\033[0;32mFINISHED\033[0;0m"

