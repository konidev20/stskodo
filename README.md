# stskodo
A simple CLI tool to fetch STS tokens for debugging activities.

# Usage
1. Set environment variables
```
export AWS_ACCESS_KEY_ID = <aws-access-key-id>
export AWS_SECRET_ACCESS_KEY = <aws-secret-access-key>
```
2. Call the command
```
stskodo <sts-endpoint>
```
For example:
```
stskodo sts.amazonaws.com
```
