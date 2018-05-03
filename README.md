# Lambda Hello World in Go lang
Hello World app for AWS Lambda in Go!

Assumptions:
You have already created an S3 bucket, an IAM role that can invoke the lambda function (with policy `AWSLambdaBasicExecutionRole`) and the function itself -- by selecting `Go 1.x` as the runtime environment -- using the AWS console or CLI.

Start by cloning the repo.

Go ahead and build the the code:
`GOOS=linux go build lambdaHello.go`

Next, ZIP the file up:
`zip lambdaHello.zip lambdaHello`

Upload the ZIP file to an S3 bucket:
`aws s3 cp lambdaHello.zip s3://<my-bucket-name>`

Invoke the function:
`aws lambda invoke --function-name <lambda-function-name> o.txt`