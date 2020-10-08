# go-fmttable

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
