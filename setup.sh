#!/bin/sh

# Setup AWS credentials
mkdir ~/.aws
echo '[default]'            > ~/.aws/config
echo 'output = json'        >> ~/.aws/config
echo 'region = us-east-1'   >> ~/.aws/config

# IAM User: lambda-deploy-for-stats-lambda-v2
echo '[default]'            > ~/.aws/credentials
echo "aws_access_key_id = $AWS_LAMBDA_DEPLOY_ACCESS_TOKEN" >> ~/.aws/credentials
echo "aws_secret_access_key = $AWS_LAMBDA_DEPLOY_ACCESS_TOKEN_SECRET" >> ~/.aws/credentials

# Set upstream
git branch --set-upstream-to=origin/master $CIRCLE_BRANCH
