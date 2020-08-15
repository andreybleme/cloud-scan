# cloud-scan

A tool for scanning your cloud resources security.

It displays the current security state of your AWS S3 buckets, by showing buckets that may be completely public, or have public files.

## Install

### Download from releases page

1. Download the latest binary for your OS on the [releases page](https://github.com/andreybleme/cloud-scan/releases).
2. Move the binary to a folder on your `PATH`. E.g.: `mv cloud-scan_darwin_amd64 /usr/local/bin/cloud-scan`.
3. Add execute permissions to the binary. E.g.: `chmod u+x /usr/local/bin/cloud-scan`.
4. Test to check if it installed correctly: `cloud-scan --help`.

### Download from source code

1. Clone the repo `git clone git@github.com:andreybleme/cloud-scan.git`
2. Install the Go package by running `go install` at the root of the cloned repo.

## Usage

Simply running `cloud-scan s3` will start the process of scanning your S3 buckets. A list of buckets will be shown with it's respective condition: secure or unsecure.

![Cloud Scan](https://github.com/andreybleme/cloud-scan/blob/master/cloud-scan.gif)

### Credentials

This CLI relies on the default AWS credentials file. This file should be stored at `~/.aws/credentials` in Unix-based systems and in `$HOME/.aws/credentials`.

Check here how to get your AWS credentials [here](https://docs.aws.amazon.com/general/latest/gr/aws-sec-cred-types.html), and see the details on where is you credentials file [here](https://docs.aws.amazon.com/cli/latest/userguide/cli-configure-files.html).

