# Book inventory

This example is heavily based on <https://www.alexedwards.net/blog/serverless-api-with-go-and-aws-lambda> and should provide a good starting point for working with DynamoDB.

Usage:

```shell
curl -i -H "Content-Type: application/json" -X POST \
-d '{"isbn":"978-0141439587", "title":"Emma", "author": "Jane Austen"}' \
"https://nn4c0kk2t5.execute-api.eu-central-1.amazonaws.com/Prod/books"
```

```shell
curl "https://nn4c0kk2t5.execute-api.eu-central-1.amazonaws.com/Prod/books?isbn=978-0141439587"
```

## Requirements

* AWS CLI already configured with Administrator permission (including `aws sam`)
* [Docker installed](https://www.docker.com/community-edition)
* [Golang](https://golang.org)

## Setup process

### Installing dependencies

In this example we use the built-in `go get` and the only dependencie we need are AWS (Lambda) Go SDK:

```shell
go get -u github.com/aws/aws-lambda-go/...
go get github.com/aws/aws-sdk-go
```

**NOTE:** As you change your application code as well as dependencies during development, you might want to research how to handle dependencies in Golang at scale.

### Building

Golang is a statically compiled language, meaning that in order to run it you have to build the executable target.

You can issue the following command in a shell to build it:

```shell
GOOS=linux GOARCH=amd64 go build -o hello-world/hello-world ./hello-world
```

Alternatively you can use `make build` or `sam build`.

**NOTE**: If you're not building the function on a Linux machine, you will need to specify the `GOOS` and `GOARCH` environment variables, this allows Golang to build your function for another system architecture and ensure compatibility.

### Local development

**Invoking function locally through local API Gateway**

```bash
sam local start-api
```

If the previous command ran successfully you should now be able to hit the following local endpoint to invoke your function `http://localhost:3000/books`

**SAM CLI** is used to emulate both Lambda and API Gateway locally and uses our `template.yaml` to understand how to bootstrap this environment (runtime, where the source code is, etc.).

**NOTE**: Using DynamoDB resources required extra fiddling with SAM Local. There are open issues on GitHub.

## Packaging and deployment

AWS Lambda Golang runtime requires a flat folder with the executable generated on build step. SAM will use `CodeUri` property to know where to look up for the application:

Use `sam build` to build.

To deploy your application for the first time, run the following in your shell:

```bash
sam deploy --guided
```

The command will package and deploy your application to AWS, with a series of prompts:

* **Stack Name**: The name of the stack to deploy to CloudFormation. This should be unique to your account and region, and a good starting point would be something matching your project name.
* **AWS Region**: The AWS region you want to deploy your app to.
* **Confirm changes before deploy**: If set to yes, any change sets will be shown to you before execution for manual review. If set to no, the AWS SAM CLI will automatically deploy application changes.
* **Allow SAM CLI IAM role creation**: Many AWS SAM templates, including this example, create AWS IAM roles required for the AWS Lambda function(s) included to access AWS services. By default, these are scoped down to minimum required permissions. To deploy an AWS CloudFormation stack which creates or modified IAM roles, the `CAPABILITY_IAM` value for `capabilities` must be provided. If permission isn't provided through this prompt, to deploy this example you must explicitly pass `--capabilities CAPABILITY_IAM` to the `sam deploy` command.
* **Save arguments to samconfig.toml**: If set to yes, your choices will be saved to a configuration file inside the project, so that in the future you can just re-run `sam deploy` without parameters to deploy changes to your application.

You can find your API Gateway Endpoint URL in the output values displayed after deployment.

### Testing

No testing anymore...

# Appendix

### Golang installation

Please ensure Go 1.x (where 'x' is the latest version) is installed as per the instructions on the official golang website: https://golang.org/doc/install

A quickstart way would be to use Homebrew, chocolatey or your linux package manager.

#### Homebrew (Mac)

Issue the following command from the terminal:

```shell
brew install golang
```

If it's already installed, run the following command to ensure it's the latest version:

```shell
brew update
brew upgrade golang
```

#### Chocolatey (Windows)

Issue the following command from the powershell:

```shell
choco install golang
```

If it's already installed, run the following command to ensure it's the latest version:

```shell
choco upgrade golang
```

## Bringing to the next level

Here are a few ideas that you can use to get more acquainted as to how this overall process works:

* Create an additional API resource (e.g. /hello/{proxy+}) and return the name requested through this new path
* Update unit test to capture that
* Package & Deploy

Next, you can use the following resources to know more about beyond hello world samples and how others structure their Serverless applications:

* [AWS Serverless Application Repository](https://aws.amazon.com/serverless/serverlessrepo/)
