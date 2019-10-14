#!/usr/bin/env bash
sudo iptables -I INPUT -p tcp --dport 80 -j ACCEPT
sudo docker network create -d bridge turitea-net
cd turitea/
sudo docker build -t turitea:v1 .
cd ..
sudo docker run -d --rm --name pg -v `pwd .`/sqldata:/var/lib/postgresql/data --network turitea-net postgres:11.0
sudo docker run -d --rm --name turitea -v `pwd .`/files:/var/local -p 80:80 --network turitea-net turitea:v1
