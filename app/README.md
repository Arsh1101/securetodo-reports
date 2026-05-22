# SecureTodo Reports App

A small Go web application for a Terraform and DevSecOps demo.

## Purpose

The app allows a single user to:

- Create todo items
- Mark todos as completed
- Generate a JSON report
- Save reports locally in phase 1
- Save reports to S3 in a later phase

## TDD Strategy

The project is structured so business logic can be tested before infrastructure or database work.

Recommended TDD order:

1. Todo validation tests
2. Todo service tests
3. Report service tests
4. Repository tests
5. Handler tests
6. SQLite integration tests
7. S3 repository tests in a later phase

## Run tests

go test ./...

