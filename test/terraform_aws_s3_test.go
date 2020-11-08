package test

import (
    "os"
	"fmt"
	"testing"
    "time"

	"github.com/gruntwork-io/terratest/modules/aws"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
    "github.com/gruntwork-io/terratest/modules/retry"
)

func TestTerraformAwsS3(t *testing.T) {
	t.Parallel()

    ts := time.Now().Format(time.UnixDate)
    fmt.Printf("time.Time %s\n", ts)
    s3_bucket_name := "ghactionstest"

	// Pick a random AWS region to test in. This helps ensure your code works in all regions.
	//awsRegion := aws.GetRandomStableRegion(t, nil, nil)
    awsRegion := "us-east-1"
    /*
    Error: Run variables are currently not supported in the cloud terraform cli
    Deleted all code related to that remote backend in commit 21045cbe16e59d9eb3cd85e34ef59c95df8c6c29
    */
	terraformOptions := &terraform.Options{
		// The path to where our Terraform code is located
		TerraformDir: "../",

		// Environment variables to set when running Terraform
		EnvVars: map[string]string{
			"AWS_DEFAULT_REGION": awsRegion,
            "AWS_ACCESS_KEY_ID": os.Getenv("AWS_ACCESS_KEY_ID"),
            "AWS_SECRET_ACCESS_KEY": os.Getenv("AWS_SECRET_ACCESS_KEY"),
		},
        Vars: map[string]interface{}{
			"execution_timestamp": ts,
            "s3_bucket": s3_bucket_name,
	    },
    }

	// Run `terraform destroy` to clean up any resources that were created after this test
	defer terraform.Destroy(t, terraformOptions)

	// `terraform init` and `terraform apply`
	terraform.InitAndApply(t, terraformOptions)

    maxRetries := 30
    timeBetweenRetries := 5 * time.Second
    description := fmt.Sprintf("Checking if S3 Bucket %s exists", s3_bucket_name)
    retry.DoWithRetry(t, description, maxRetries, timeBetweenRetries, func() (string, error) {
        err := aws.AssertS3BucketExistsE(t, awsRegion, s3_bucket_name)
        return "", err
    })
    // Verify Timestamp on files test1.txt and test2.txt are correct
    key_name := "test1.txt"
    description = fmt.Sprintf("Checking if timestamp on %s/%s is %s", s3_bucket_name, key_name, ts)
    retry.DoWithRetry(t, description, maxRetries, timeBetweenRetries, func() (string, error) {
        actual_value, err := aws.GetS3ObjectContentsE(t, awsRegion, s3_bucket_name, key_name)
	    assert.Equal(t, ts, actual_value)
        return "", err
    })
    key_name = "test2.txt"
    description = fmt.Sprintf("Checking if timestamp on %s/%s is %s", s3_bucket_name, key_name, ts)
    retry.DoWithRetry(t, description, maxRetries, timeBetweenRetries, func() (string, error) {
        actual_value, err := aws.GetS3ObjectContentsE(t, awsRegion, s3_bucket_name, key_name)
	    assert.Equal(t, ts, actual_value)
        return "", err
    })

}
