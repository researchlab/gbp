#!/bin/sh

if [ -z $PROJECT_PATH ]; then
    echo "Environment PROJECT_PATH required"
    exit 1
else
  cd "$PROJECT_PATH" || exit
fi

if [ -z $CONFIG_FILE ]; then
    echo "Config file not specified"
else
    CONFIG_FILE="-- $CONFIG_FILE"
    export CONFIG_FILE
fi

exec "$@"