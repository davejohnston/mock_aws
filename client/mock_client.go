package client

import (
	"encoding/json"

	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/aws/aws-sdk-go/service/ec2/ec2iface"
)

type MockEC2Client struct {
	ec2iface.EC2API
}

func (m *MockEC2Client) DescribeRegions(*ec2.DescribeRegionsInput) (*ec2.DescribeRegionsOutput, error) {
	data := &ec2.DescribeRegionsOutput{}
	json.Unmarshal([]byte(regionData), data)
	return data, nil
}

func (m *MockEC2Client) DescribeInstances(*ec2.DescribeInstancesInput) (*ec2.DescribeInstancesOutput, error) {
	data := &ec2.DescribeInstancesOutput{}
	json.Unmarshal([]byte(computeData), data)
	return data, nil
}

const regionData = `{
  "Regions": [
    {
      "Endpoint": "ec2.ap-south-1.amazonaws.com",
      "RegionName": "ap-south-1"
    },
    {
      "Endpoint": "ec2.eu-west-2.amazonaws.com",
      "RegionName": "eu-west-2"
    },
    {
      "Endpoint": "ec2.eu-west-1.amazonaws.com",
      "RegionName": "eu-west-1"
    },
    {
      "Endpoint": "ec2.ap-northeast-2.amazonaws.com",
      "RegionName": "ap-northeast-2"
    },
    {
      "Endpoint": "ec2.ap-northeast-1.amazonaws.com",
      "RegionName": "ap-northeast-1"
    },
    {
      "Endpoint": "ec2.sa-east-1.amazonaws.com",
      "RegionName": "sa-east-1"
    },
    {
      "Endpoint": "ec2.ca-central-1.amazonaws.com",
      "RegionName": "ca-central-1"
    },
    {
      "Endpoint": "ec2.ap-southeast-1.amazonaws.com",
      "RegionName": "ap-southeast-1"
    },
    {
      "Endpoint": "ec2.ap-southeast-2.amazonaws.com",
      "RegionName": "ap-southeast-2"
    },
    {
      "Endpoint": "ec2.eu-central-1.amazonaws.com",
      "RegionName": "eu-central-1"
    },
    {
      "Endpoint": "ec2.us-east-1.amazonaws.com",
      "RegionName": "us-east-1"
    },
    {
      "Endpoint": "ec2.us-east-2.amazonaws.com",
      "RegionName": "us-east-2"
    },
    {
      "Endpoint": "ec2.us-west-1.amazonaws.com",
      "RegionName": "us-west-1"
    },
    {
      "Endpoint": "ec2.us-west-2.amazonaws.com",
      "RegionName": "us-west-2"
    }
  ]
}`

const computeData = `{
  "NextToken": null,
  "Reservations": [
    {
      "Groups": null,
      "Instances": [
        {
          "AmiLaunchIndex": 0,
          "Architecture": "x86_64",
          "BlockDeviceMappings": [
            {
              "DeviceName": "/dev/sda1",
              "Ebs": {
                "AttachTime": "2017-10-06T08:37:29Z",
                "DeleteOnTermination": true,
                "Status": "attached",
                "VolumeId": "vol-046c9c82044fd7cd9"
              }
            }
          ],
          "ClientToken": "rSNQA1507279047650",
          "EbsOptimized": false,
          "ElasticGpuAssociations": null,
          "EnaSupport": true,
          "Hypervisor": "xen",
          "IamInstanceProfile": null,
          "ImageId": "ami-bb9a6bc2",
          "InstanceId": "i-09abe73024efd0341",
          "InstanceLifecycle": null,
          "InstanceType": "t2.micro",
          "KernelId": null,
          "KeyName": "dajohnst",
          "LaunchTime": "2017-10-06T08:37:28Z",
          "Monitoring": {
            "State": "disabled"
          },
          "NetworkInterfaces": [
            {
              "Association": {
                "IpOwnerId": "amazon",
                "PublicDnsName": "ec2-34-240-22-140.eu-west-1.compute.amazonaws.com",
                "PublicIp": "34.240.22.140"
              },
              "Attachment": {
                "AttachTime": "2017-10-06T08:37:28Z",
                "AttachmentId": "eni-attach-797d1f05",
                "DeleteOnTermination": true,
                "DeviceIndex": 0,
                "Status": "attached"
              },
              "Description": "",
              "Groups": [
                {
                  "GroupId": "sg-292d2751",
                  "GroupName": "SSH Only"
                }
              ],
              "Ipv6Addresses": null,
              "MacAddress": "0a:7e:25:32:7a:1e",
              "NetworkInterfaceId": "eni-257f5824",
              "OwnerId": "689951665833",
              "PrivateDnsName": "ip-172-31-4-2.eu-west-1.compute.internal",
              "PrivateIpAddress": "172.31.4.2",
              "PrivateIpAddresses": [
                {
                  "Association": {
                    "IpOwnerId": "amazon",
                    "PublicDnsName": "ec2-34-240-22-140.eu-west-1.compute.amazonaws.com",
                    "PublicIp": "34.240.22.140"
                  },
                  "Primary": true,
                  "PrivateDnsName": "ip-172-31-4-2.eu-west-1.compute.internal",
                  "PrivateIpAddress": "172.31.4.2"
                }
              ],
              "SourceDestCheck": true,
              "Status": "in-use",
              "SubnetId": "subnet-5b06031d",
              "VpcId": "vpc-193dd07c"
            }
          ],
          "Placement": {
            "Affinity": null,
            "AvailabilityZone": "eu-west-1b",
            "GroupName": "",
            "HostId": null,
            "SpreadDomain": null,
            "Tenancy": "default"
          },
          "Platform": null,
          "PrivateDnsName": "ip-172-31-4-2.eu-west-1.compute.internal",
          "PrivateIpAddress": "172.31.4.2",
          "ProductCodes": null,
          "PublicDnsName": "ec2-34-240-22-140.eu-west-1.compute.amazonaws.com",
          "PublicIpAddress": "34.240.22.140",
          "RamdiskId": null,
          "RootDeviceName": "/dev/sda1",
          "RootDeviceType": "ebs",
          "SecurityGroups": [
            {
              "GroupId": "sg-292d2751",
              "GroupName": "SSH Only"
            }
          ],
          "SourceDestCheck": true,
          "SpotInstanceRequestId": null,
          "SriovNetSupport": null,
          "State": {
            "Code": 16,
            "Name": "running"
          },
          "StateReason": null,
          "StateTransitionReason": "",
          "SubnetId": "subnet-5b06031d",
          "Tags": [
            {
              "Key": "Purpose",
              "Value": "Puppet Tester"
            }
          ],
          "VirtualizationType": "hvm",
          "VpcId": "vpc-193dd07c"
        }
      ],
      "OwnerId": "689951665833",
      "RequesterId": null,
      "ReservationId": "r-04a4bfeaea11c0cac"
    },
    {
      "Groups": null,
      "Instances": [
        {
          "AmiLaunchIndex": 0,
          "Architecture": "x86_64",
          "BlockDeviceMappings": [
            {
              "DeviceName": "/dev/xvda",
              "Ebs": {
                "AttachTime": "2016-04-05T09:04:05Z",
                "DeleteOnTermination": true,
                "Status": "attached",
                "VolumeId": "vol-42eb52b0"
              }
            }
          ],
          "ClientToken": "SApRC1459847044363",
          "EbsOptimized": false,
          "ElasticGpuAssociations": null,
          "EnaSupport": null,
          "Hypervisor": "xen",
          "IamInstanceProfile": null,
          "ImageId": "ami-31328842",
          "InstanceId": "i-34a424be",
          "InstanceLifecycle": null,
          "InstanceType": "t2.micro",
          "KernelId": null,
          "KeyName": "dajohnst",
          "LaunchTime": "2016-08-05T16:21:42Z",
          "Monitoring": {
            "State": "disabled"
          },
          "NetworkInterfaces": [
            {
              "Association": {
                "IpOwnerId": "689951665833",
                "PublicDnsName": "ec2-52-50-147-44.eu-west-1.compute.amazonaws.com",
                "PublicIp": "52.50.147.44"
              },
              "Attachment": {
                "AttachTime": "2016-04-05T09:04:04Z",
                "AttachmentId": "eni-attach-0e7d34cc",
                "DeleteOnTermination": true,
                "DeviceIndex": 0,
                "Status": "attached"
              },
              "Description": "",
              "Groups": [
                {
                  "GroupId": "sg-a16a87c6",
                  "GroupName": "wordpress"
                }
              ],
              "Ipv6Addresses": null,
              "MacAddress": "06:80:56:1a:41:b7",
              "NetworkInterfaceId": "eni-dfcaf294",
              "OwnerId": "689951665833",
              "PrivateDnsName": "ip-172-31-41-200.eu-west-1.compute.internal",
              "PrivateIpAddress": "172.31.41.200",
              "PrivateIpAddresses": [
                {
                  "Association": {
                    "IpOwnerId": "689951665833",
                    "PublicDnsName": "ec2-52-50-147-44.eu-west-1.compute.amazonaws.com",
                    "PublicIp": "52.50.147.44"
                  },
                  "Primary": true,
                  "PrivateDnsName": "ip-172-31-41-200.eu-west-1.compute.internal",
                  "PrivateIpAddress": "172.31.41.200"
                }
              ],
              "SourceDestCheck": true,
              "Status": "in-use",
              "SubnetId": "subnet-6313f714",
              "VpcId": "vpc-193dd07c"
            }
          ],
          "Placement": {
            "Affinity": null,
            "AvailabilityZone": "eu-west-1a",
            "GroupName": "",
            "HostId": null,
            "SpreadDomain": null,
            "Tenancy": "default"
          },
          "Platform": null,
          "PrivateDnsName": "ip-172-31-41-200.eu-west-1.compute.internal",
          "PrivateIpAddress": "172.31.41.200",
          "ProductCodes": null,
          "PublicDnsName": "ec2-52-50-147-44.eu-west-1.compute.amazonaws.com",
          "PublicIpAddress": "52.50.147.44",
          "RamdiskId": null,
          "RootDeviceName": "/dev/xvda",
          "RootDeviceType": "ebs",
          "SecurityGroups": [
            {
              "GroupId": "sg-a16a87c6",
              "GroupName": "wordpress"
            }
          ],
          "SourceDestCheck": true,
          "SpotInstanceRequestId": null,
          "SriovNetSupport": null,
          "State": {
            "Code": 16,
            "Name": "running"
          },
          "StateReason": null,
          "StateTransitionReason": "",
          "SubnetId": "subnet-6313f714",
          "Tags": [
            {
              "Key": "Owner",
              "Value": "davej"
            },
            {
              "Key": "Name",
              "Value": "wordpress_production"
            }
          ],
          "VirtualizationType": "hvm",
          "VpcId": "vpc-193dd07c"
        }
      ],
      "OwnerId": "689951665833",
      "RequesterId": null,
      "ReservationId": "r-b00a7509"
    }
  ]
}
`
