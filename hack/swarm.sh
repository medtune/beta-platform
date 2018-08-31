#!/bin/bash

# Deploy on docker swarm
function start-swarm() {
    docker swarm init 
    docker service create --name registry --publish published=5000,target=5000 registry:2 

    docker-compose up -d
    docker-compose down --volumes 
    docker-compose push 

    docker stack deploy --compose-file docker-compose.yml beta-medtune
    docker stack services beta-medtune 
}

# Kill swarm
function kill-swarm() {
    docker stack rm beta-medtune
    docker service rm registry
    docker swarm leave --force
}


$@