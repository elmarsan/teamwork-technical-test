# Teamwork technical test

This repository contains a technical test for Teamwork company

[![Go Report Card](https://goreportcard.com/badge/github.com/elmarsan/teamwork-technical-test)](https://goreportcard.com/report/github.com/elmarsan/teamwork-technical-test)
![Coverage](https://img.shields.io/badge/Coverage-96.4%25-brightgreen)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

## Requirements

- package customerimporter reads from the given customers.csv file and returns a sorted (data structure of your choice) of email domains along with the number of customers with e-mail addresses for each domain.  Any errors should be logged (or handled). Performance matters (this is only ~3k lines, but *could* be 1m lines or run on a small machine).