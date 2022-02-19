#!/usr/bin/env bash
APP_NAME=webbe
DEPLOY_CONNECT=ryu@167.71.214.240
# Digital

echo "Downloading packages ...."
go mod tidy

echo "Compiling"
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app

echo "Docker build ..."
docker buildx build -t ${APP_NAME} -f ./Dockerfile . --platform linux/amd64

echo "Docker saving..."
docker save -o ${APP_NAME}.tar ${APP_NAME}

echo "Deploying..."
scp -o StrictHostKeyChecking=no ./${APP_NAME}.tar ${DEPLOY_CONNECT}:~
ssh -o StrictHostKeyChecking=no ${DEPLOY_CONNECT} 'bash -s' < ./deploy/stg.sh

echo "Cleaning..."
rm -f ./${APP_NAME}.tar
#docker rmi $(docker images -qa -f 'dangling=true')
echo "Done"