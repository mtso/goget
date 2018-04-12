# goget

[summary]::
`go get` **all** the repos.

## Install

```
$ go get github.com/mtso/goget
```

## Usage

`goget` uses `GITPATH` to determine where to save the repository to.

```
$ GITPATH=~/dev goget github.com/mtso/goget
```

To change directories into the repository right after cloning:

```
$ cd `goget github.com/mtso/cha-0`
```
