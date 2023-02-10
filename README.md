# AWS subnet exporter
Fetch AWS subnet available IP count and expose it as Prometheus metrics. Why? Because AWS does not expose these in CloudWatch and you don't want to run out of available IP's in your subnets.

## Metrics exported
```
# Curl metrics example
curl http://localhost:8080/metrics
# HELP aws_subnet_exporter_available_ips Available IPs in subnets
# TYPE aws_subnet_exporter_available_ips gauge
aws_subnet_exporter_available_ips{az="eu-west-1a",cidrblock="10.103.0.0/28",name="eks_clu_eu-west-1a",subnetid="subnet-XXX",vpcid="vpc-YYY"} 10
...

# HELP aws_subnet_exporter_max_ips Max host IPs in subnet
# TYPE aws_subnet_exporter_max_ips gauge
aws_subnet_exporter_max_ips{az="eu-west-1a",cidrblock="10.103.0.0/28",name="eks_clu_eu-west-1a",subnetid="subnet-XXX",vpcid="vpc-YYY"} 14
...
```

## Assumptions
This service assumes that you subnets have a tag "Name" and that you have exported your AWS access key and secret.

## AWS policy required
You require this policy to your user/role (use roles for best practice) in order to fetch AWS subnet data.
```json
{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Sid": "VisualEditor0",
            "Effect": "Allow",
            "Action": "ec2:DescribeSubnets",
            "Resource": "*"
        }
    ]
}
```

## Running
You have to provide an AWS context and I will not cover how to do this here.

```bash
docker run -p 8080:8080 -e AWS_ACCESS_KEY_ID=xyz -e AWS_SECRET_ACCESS_KEY=aaa ghcr.io/wcarlsen/aws-subnet-exporter:latest ./aws-subnet-exporter --port="8080" --region="eu-west-1" --filter="*" --period="60" --debug
```