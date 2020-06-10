These commands are required for setting up the resources (s3 bucket and lambdas). It is required that the AWS SAM CLI is installed. Instructions to do so can be read [here](https://docs.aws.amazon.com/serverless-application-model/latest/developerguide/serverless-sam-cli-install.html).

After installation, the function for deployment can be downloaded with:

```$sam init 2 <github folder> ```


Creating the bucket is done with:

```$aws s3 mb s3://<bucketname>```

`<bucketname>` is user determined, and must be original.

Packaging and deploying the resources is done with:
```
$sam package \
    --template-file template.yaml \
    --output-template-file packaged.yaml \
    --s3-bucket <bucketname>

$sam deploy \
    --template-file packaged.yaml \
    --stack-name aws-sam-ocr \
    --capabilities CAPABILITY_IAM \
    --region <region>
```
Here, `<bucketname>` has to be the same name chosen as before. For `<region>`, please fill in the region in which you want your resources to be deployed, e.g. us-east-2.
