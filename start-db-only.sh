#!/bin/bash

cd db
docker build -t wiki-db .
docker run -p 5432:5432 --env-file .env wiki-db 