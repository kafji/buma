#!/bin/bash

docker run -e POSTGRES_PASSWORD=password -p 5432:5432 -it postgres:14
