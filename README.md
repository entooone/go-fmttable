# go-fmttable

[![test](https://github.com/entooone/go-fmttable/workflows/test/badge.svg)](https://github.com/entooone/go-fmttable/actions?query=workflow%3Atest)
[![codecov](https://codecov.io/gh/entooone/go-fmttable/branch/master/graph/badge.svg)](https://codecov.io/gh/entooone/go-fmttable)
[![Go Report Card](https://goreportcard.com/badge/github.com/entooone/go-fmttable)](https://goreportcard.com/report/github.com/entooone/go-fmttable)

## Installation

```
$ go get -u github.com/entooone/go-ftable/cmd/goft
```

## Usage

```
$ cat << EOF | goft
| abc | def | hjk |
| hello | world | gopher|
EOF
```

Output

```
| abc   | def   | hjk    |
| hello | world | gopher |
```
