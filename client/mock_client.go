package client

import (
	"encoding/json"

	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/aws/aws-sdk-go/service/ec2/ec2iface"
	"github.com/davejohnston/mock_aws/client/fixtures"
)

type MockEC2Client struct {
	ec2iface.EC2API
}

func (m *MockEC2Client) DescribeRegions(*ec2.DescribeRegionsInput) (*ec2.DescribeRegionsOutput, error) {
	data := &ec2.DescribeRegionsOutput{}

	LoadFixtureFromAsset("client/fixtures/regions.json", data)

	//json.Unmarshal([]byte(regionData), data)
	return data, nil
}

func (m *MockEC2Client) DescribeInstances(*ec2.DescribeInstancesInput) (*ec2.DescribeInstancesOutput, error) {
	data := &ec2.DescribeInstancesOutput{}
	LoadFixtureFromAsset("client/fixtures/instances.json", data)
	//json.Unmarshal([]byte(computeData), data)
	return data, nil
}

// LoadFixtureFromAsset loads the specified file and decodes the JSON into the provided struct.  It uses the assets generated from a go-bindata command
func LoadFixtureFromAsset(filename string, fixture interface{}) error {
	raw, err := fixtures.Asset(filename)
	if err != nil {
		// Asset was not found.
		return err
	}
	json.Unmarshal(raw, &fixture)
	return nil
}
