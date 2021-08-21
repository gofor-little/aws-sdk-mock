## A package for mocking interactions with AWS services

![GitHub tag (latest SemVer pre-release)](https://img.shields.io/github/v/tag/gofor-little/aws-sdk-mock?include_prereleases)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/gofor-little/aws-sdk-mock)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://raw.githubusercontent.com/gofor-little/aws-sdk-mock/main/LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/gofor-little/aws-sdk-mock)](https://goreportcard.com/report/github.com/gofor-little/aws-sdk-mock)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/gofor-little/aws-sdk-mock)](https://pkg.go.dev/github.com/gofor-little/aws-sdk-mock)

### Deprecated
At this stage there are no plans to migrate to v2 of the AWS Go SDK or continue to maintain this package.

### Introduction
* Mock interactions with AWS services
* Make unit testing quicker and easier

### Example
```go
package mock_test

import (
    "github.com/gofor-little/aws-sdk-mock"
)

func TestMockExample(t *testing.T) {
    client := mock.S3client{}

    output, err := client.GetObjectWithContext(context.Background(), &s3.GetObjectInput{})
    if err != nil {
        t.Fatal(err)
    }

    // Check output is correct...
}
```