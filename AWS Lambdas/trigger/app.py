import json
import boto3
import os
import random

# boto3 S3 initialization
s3_client = boto3.client("s3")


def lambda_handler(event, context):
   destination_bucket_name = os.environ['LamqBucket']
   source_bucket_name = destination_bucket_name

   # event contains all information about uploaded object
   # Filename of object (with path)
   file_key_name_in = "image.png"
   file_key_name_out = f"image_{random.random()}.png"

   # Copy Source Object
   copy_source_object = {'Bucket': source_bucket_name, 'Key': file_key_name_in}

   # S3 copy object operation
   s3_client.copy_object(CopySource=copy_source_object, Bucket=destination_bucket_name, Key=file_key_name_out)

   return {
       'statusCode': 200,
       'body': json.dumps('Hello from S3 events Lambda!')
   }