#!/bin/bash
docker build -t app .
aws ecr get-login-password --region ap-northeast-1 --profile xxxxxxxx | docker login --username AWS --password-stdin xxxxxxxx.dkr.ecr.ap-northeast-1.amazonaws.com/tryapprunner-ecr
docker tag app:latest xxxxxxxx.dkr.ecr.ap-northeast-1.amazonaws.com/tryapprunner-ecr:latest
docker push xxxxxxxx.dkr.ecr.ap-northeast-1.amazonaws.com/tryapprunner-ecr:latest