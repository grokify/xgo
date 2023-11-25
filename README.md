# SoGo

[![Build Status][build-status-svg]][build-status-link]
[![Go Report Card][goreport-svg]][goreport-link]
[![Docs][docs-godoc-svg]][docs-godoc-link]
[![License][license-svg]][license-link]

## Description

Sogo `database/document` provides a storage abstraction layer for various storage engines with a common configuration struct.

Initially, it's goal is to provide services for simple key value string storage. This use case is very simple but often times you may want to use a different solution depending on your provider, e.g. DynamoDB for AWS, Redis for Heroku, and the local file system for testing.

Supported engines include:

* DynamoDB
* Files
* Redis

 [build-status-svg]: https://github.com/grokify/sogo/workflows/build/badge.svg
 [build-status-link]: https://github.com/grokify/sogo/actions
 [goreport-svg]: https://goreportcard.com/badge/github.com/grokify/sogo
 [goreport-link]: https://goreportcard.com/report/github.com/grokify/sogo
 [docs-godoc-svg]: https://pkg.go.dev/badge/github.com/grokify/sogo
 [docs-godoc-link]: https://pkg.go.dev/github.com/grokify/sogo
 [license-svg]: https://img.shields.io/badge/license-MIT-blue.svg
 [license-link]: https://github.com/grokify/sogo/blob/master/LICENSE
 
`sogo/database/document`` is designed to provide a generic interface to storage engines using a common config and setter/getter interface.
