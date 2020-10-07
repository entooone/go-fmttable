# go-fmttable

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
