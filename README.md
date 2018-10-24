# GoStor

[![Build Status][build-status-svg]][build-status-link]
[![Go Report Card][goreport-svg]][goreport-link]
[![Docs][docs-godoc-svg]][docs-godoc-link]
[![License][license-svg]][license-link]

## Description

GoStor provides a storage abstraction layer for various storage engines with a common configuration struct.

Initially, it's goal is to provide services for simple key value string storage.

Supported engines include:

* DynamoDB
* Files
* Redis

 [build-status-svg]: https://api.travis-ci.org/grokify/gostor.svg?branch=master
 [build-status-link]: https://travis-ci.org/grokify/gostor
 [goreport-svg]: https://goreportcard.com/badge/github.com/grokify/gostor
 [goreport-link]: https://goreportcard.com/report/github.com/grokify/gostor
 [docs-godoc-svg]: https://img.shields.io/badge/docs-godoc-blue.svg
 [docs-godoc-link]: https://godoc.org/github.com/grokify/gostor
 [license-svg]: https://img.shields.io/badge/license-MIT-blue.svg
 [license-link]: https://github.com/grokify/gostor/blob/master/LICENSE
 
GoStor is designed to provide a generic interface to storage engines using a common config and setter/getter interface.
