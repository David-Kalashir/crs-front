# CRS-Front

[![Go Version](https://img.shields.io/github/go-mod/go-version/David-Kalashir/crs-front.svg)](https://golang.org/doc/go1.20)

## Table of Contents

- [Description](#description)
- [Features](#features)
- [License](#licensing)
- [Installation](#installation)
  - [Create Your Own Version](#create-your-own-version)
- [Running The Project](#running-the-project)
  - [Install Tooling](#install-tooling)
  - [Run Tests](#run-tests)
  - [Shutdown Tests](#shutdown-tests)
  - [Run Project](#run-project)
  - [Run Load](#run-load)
  - [Run Tooling](#run-tooling)
  - [Shut Project](#shut-project)
- [Credits](#credits)

## Description

This is a front-end project built with **Golang** and **HTMX**, following a **Domain-Driven Design (DDD)** architecture. The project maintains a clear separation between the different layers of the application, ensuring scalability, testability, and ease of maintenance.

Inspired by Bill Kennedy’s structure for Go projects, this repository implements best practices in Go development and architecture.

## Features

- Domain-Driven Design (DDD) architecture
- Modular and scalable code structure
- Clear separation of concerns
- Frontend developed with **HTMX**


## Licensing

```
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
```

## Installation

To clone the project, create a folder and use the git clone command. Then please read the Makefile to learn how to install all the tooling and Docker images.

```bash
$ cd $HOME
$ mkdir code
$ cd code
$ git clone https://github.com/David-Kalashir/crs-front.git
$ cd crs-front
```
## Create Your Own Version

If you want to create a version of the project for your own use, use the new gonew command.

```bash
$ go install golang.org/x/tools/cmd/gonew@latest

$ cd $HOME
$ mkdir code
$ cd code
$ gonew github.com/David-Kalashir/crs-front github.com/mydomain/myproject
$ cd myproject
$ go mod vendor
```
Now you have a copy with your own module name. All you need to do is initialize the project for git.

## Running The Project
To run the project, use the following commands:

### Install Tooling

```bash
$ make dev-gotooling
$ make dev-brew
$ make dev-docker
```
### Run Tests

```bash
$ make test
```
### Shutdown Tests

```bash
$ make test-down
```
### Run Project

```bash
$ make dev-up
$ make dev-update-apply
$ make token
$ export TOKEN=<COPY TOKEN>
$ make users
```
### Run Load

```bash
$ make load
```
### Run Tooling

```bash
$ make grafana
$ make statsviz
```
### Shut Project

```bash
$ make dev-down
```
## Credits
This project draws significant inspiration from Bill Kennedy’s work on structuring Go projects with Domain-Driven Design. His books and talks were instrumental in shaping the architecture of this repository.

___
All material is licensed under the [Apache License Version 2.0, January 2004](http://www.apache.org/licenses/LICENSE-2.0).