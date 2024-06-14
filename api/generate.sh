

openapi-generator generate --api-package api -i api.yaml -g go -o client --model-package models --git-user-id synology-community --git-repo-id dsm-api --additional-properties=packageName=synology,packageVersion=1.0.0,generateInterfaces=true