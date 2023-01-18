#!/bin/ash

./aws-subnet-exporter \
--port=${PORT} \
--region=${REGION} \
--filter=${FILTER} \
--period=${PERIOD}
