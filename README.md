# Lambda Hello World in Go lang
Hello World app for AWS Lambda in Go!

Assumptions:
You have already created an S3 bucket, an IAM role that can invoke the lambda function (with policy `AWSLambdaBasicExecutionRole`) and the function itself -- by selecting `Go 1.x` as the runtime environment -- using the AWS console or CLI.

Start by cloning the repo.

Go ahead and build the the code (IMPORTANT: Note the param `GOOS=linux`, this is needed for the code to execute on AWS. Otherwise you'll see `exec format error` in your CloudWatch logs:
`GOOS=linux go build lambdaHello.go`

Next, ZIP the file up:
`zip lambdaHello.zip lambdaHello`

Upload the ZIP file to an S3 bucket:
`aws s3 cp lambdaHello.zip s3://<my-bucket-name>`

Invoke the function:
`aws lambda invoke --function-name <lambda-function-name> o.txt`

Assuming everything went well, you should see the following output:
`{
    "ExecutedVersion": "$LATEST", 
    "StatusCode": 200
}`

You can always `cat` the file `o.txt` and see what the function returned. In our case, it will be `null` as the code doesn't retun anything but `nil`.