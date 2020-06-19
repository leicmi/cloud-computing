
const s3Util = require('./s3-util'),
	childProcessPromise = require('./child-process-promise'),
	path = require('path'),
	os = require('os'),
	EXTENSION = process.env.EXTENSION,
	THUMB_WIDTH = process.env.THUMB_WIDTH,
	OUTPUT_BUCKET = process.env.OUTPUT_BUCKET,
	MIME_TYPE =  process.env.MIME_TYPE;

// Create a DocumentClient that represents the query to add an item
const dynamodb = require('aws-sdk/clients/dynamodb');
const docClient = new dynamodb.DocumentClient();

// Get the DynamoDB table name from environment variables
const tableName = process.env.DynamoDBTable;

exports.handler = function (eventObject, context) {
	const eventRecord = eventObject.Records && eventObject.Records[0],
		inputBucket = eventRecord.s3.bucket.name,
		key = eventRecord.s3.object.key,
		id = context.awsRequestId,
		resultKey = key.replace(/\.[^.]+$/, EXTENSION),
		workdir = os.tmpdir(),
		inputFile = path.join(workdir,  id + path.extname(key)),
		outputFile = path.join(workdir, 'converted-' + id + EXTENSION);

    // Creates a new item, or replaces an old item with a new item
    // https://docs.aws.amazon.com/AWSJavaScriptSDK/latest/AWS/DynamoDB/DocumentClient.html#put-property
	var params = {TableName : tableName, Item: { id : key, jobStatus: "CONVERTING" }};
	docClient.put(params, function(err, data) {
		if (err) console.log(err);
		//else console.log(data);
	  });

	console.log('converting', inputBucket, key, 'using', inputFile);

	result = s3Util.downloadFileFromS3(inputBucket, key, inputFile)
	.then(() => childProcessPromise.spawn(
		'/opt/bin/convert',
		[inputFile, '-resize', `${THUMB_WIDTH}x`, outputFile],
		{env: process.env, cwd: workdir}
	))
	.then(() => s3Util.uploadFileToS3(OUTPUT_BUCKET, resultKey, outputFile, MIME_TYPE));

    // Creates a new item, or replaces an old item with a new item
    // https://docs.aws.amazon.com/AWSJavaScriptSDK/latest/AWS/DynamoDB/DocumentClient.html#put-property
	var params = {TableName : tableName, Item: { id : key, jobStatus: "FINISHED" }};
	docClient.put(params, function(err, data) {
		if (err) console.log(err);
		//else console.log(data);
	  });

	return result
};
