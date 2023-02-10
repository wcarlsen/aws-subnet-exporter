#!/bin/ash

./aws-subnet-exporter \
    --port=${PORT:":8080"} \
    --region=${REGION} \
    --filter=${FILTER:"*eks*"} \
    --period=${PERIOD:"60"}