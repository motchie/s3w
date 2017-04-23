# s3w #

A CLI for Amazon S3 website configurations.

### Description

s3w is a Command Line Interface tool that control Amazon S3 website configurations.
An auxiliary tool for aws cli.

### Features

* No JSON required to use.
* AWS CLI compatible credentials and flags.
* Reduce operations for most common use.

### Usage
```
s3w is a Command Line Interface tool that control Amazon S3 website configurations.
An auxiliary tool for aws cli.

- No JSON required to use.
- AWS CLI compatible credentials and flags.
- Reduce operations for most common use.

Usage:
  s3w [command]

Available Commands:
  delete-website                Removes the website configuration from the bucket
  get-bucket-acl                Gets the access control list (ACL) for the bucket.
  get-bucket-policy             Gets the policy of a specified bucket.
  get-location                  Returns returns the region the bucket resides in
  get-website                   Returns the website configuration for a bucket
  help                          Help about any command
  list-buckets                  List all S3 buckets
  set-bucket-acl                Sets the permissions on a bucket using access control list (ACL).
  set-bucket-policy-for-website Replaces a policy on a bucket for Web Site.
  set-redirect-all-requests     Set to redirect all website requests sent to the bucket's endpoint
  set-website                   Set the website configuration for a bucket.
  version                       Show version and Copyright.

Flags:
  -t, --toggle   Help message for toggle

Use "s3w [command] --help" for more information about a command.```
```
## Installation

## Author

[motchie](https://github.com/motchie)

## License

[MIT](https://github.com/motchie/s3w/blob/master/LICENSE)
