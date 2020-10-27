# gh-actions-test
![Integration Tests](https://github.com/3ll34ndr0/gh-actions-test/workflows/Integration%20Tests/badge.svg)

## Exercise ##
* Create Terraform code to create a AWS S3 bucket with two files: test1.txt and test2.txt. The content of these files must be the timestamp when the code was executed.
* Using Terratest, create the test automation for the Terraform code, validating that both files and the bucket are created successfully.
* Setup Github Actions to run a pipeline to validate this code.

## Running automated tests against this task definition 

1. Sign up for [AWS](https://aws.amazon.com/).
1. Configure your AWS credentials using one of the [supported methods for AWS CLI
   tools](https://docs.aws.amazon.com/cli/latest/userguide/cli-chap-getting-started.html), such as setting the
   `AWS_ACCESS_KEY_ID` and `AWS_SECRET_ACCESS_KEY` environment variables. 
1. Install [Terraform](https://www.terraform.io/) and make sure it's on your `PATH`.
1. Install [Golang](https://golang.org/) and make sure this code is checked out into your `GOPATH`.
1. Clone this repository and `cd` to it
1. `cd test`
1. `go mod init "<MODULE_NAME>"`
1. `go test -v `
