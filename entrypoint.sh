#!/bin/bash

envsubst < ./config/yaml/config.yaml.template > ./config/yaml/config.yaml
 
# Execute
exec "$@"