#!/bin/sh

`ssh -i ./sugilanon_key.pem ec2-user@$PRIVATE_DNS`

# sudo yum update -y
# sudo yum install -y docker

# sudo usermod -aG docker ec2-user

# cd /go/src/github/XanderDwyl/sugilanon

# docker build -t web .

# docker run -d -e DB_HOST=$DB_HOST -e DB_NAME=$DB_NAME -e DB_USER=$DB_USER -e DB_PASS=$DB_PASS -e DB_PORT=3306 -e MODE=production -e PORT=80  -p 80:80 --rm web