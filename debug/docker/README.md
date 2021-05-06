# Dockerized Golang DELVE debugger

This repository about golang remote debug based on [go-delve/delve](https://github.com/go-delve/delve)

## Getting started:

Firstly configure your IDE to remote debugging Go application on port :5813
then start your application
```shell script
cd Go_Project_Folder
docker run -v $(pwd):$(pwd) -e PROJECT_PATH=$(pwd) -p 5813:5813 researchboy/debuger:latest
```

## Options:

Environments variables:

```.env
    PROJECT_PATH - path to project with main entrypoint
    CONFIG_FILE - path to config file of golang application
```

Full example:

```shell script
docker run -v $(pwd):$(pwd) -e PROJECT_PATH=$(pwd) -e CONFIG_FILE=./conf/app.yaml -p 5813:5813 researchboy/debuger:latest 
```
