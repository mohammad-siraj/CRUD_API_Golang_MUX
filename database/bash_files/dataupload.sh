#!bin/sh
docker cp SQL_files/car.sql crud_api:/car.sql
docker exec -u root crud_api psql root root -f car.sql