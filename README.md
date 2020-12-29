# GoStor

[![Build Status][build-status-svg]][build-status-link]
[![Go Report Card][goreport-svg]][goreport-link]
[![Docs][docs-godoc-svg]][docs-godoc-link]
[![License][license-svg]][license-link]

## Description

GoStor provides a storage abstraction layer for various storage engines with a common configuration struct.

Initially, it's goal is to provide services for simple key value string storage. This use case is very simple but often times you may want to use a different solution depending on your provider, e.g. DynamoDB for AWS, Redis for Heroku, and the local file system for testing.

Supported engines include:

* DynamoDB
* Files
* Redis

 [build-status-svg]: https://github.com/grokify/gostor/workflows/build/badge.svg
 [build-status-link]: https://github.com/grokify/gostor/actions
 [goreport-svg]: https://goreportcard.com/badge/github.com/grokify/gostor
 [goreport-link]: https://goreportcard.com/report/github.com/grokify/gostor
 [docs-godoc-svg]: https://pkg.go.dev/badge/github.com/grokify/gostor
 [docs-godoc-link]: https://pkg.go.dev/github.com/grokify/gostor
 [license-svg]: https://img.shields.io/badge/license-MIT-blue.svg
 [license-link]: https://github.com/grokify/gostor/blob/master/LICENSE
 
GoStor is designed to provide a generic interface to storage engines using a common config and setter/getter interface.
