# echoex

## About

echoex is a simple wrapper that you can put in front of any command you run. It
echoes out this command and then executes it, while attaching stdin, stdout and
stderr to the parent process. This makes echoex completely transparent, letting
you use the command interactively or in a script. The additional line of
printing out the actual command comes in handy for me when debugging shell
globbing or when writing into log files. Combine echoex with the `-t` option to
also get a timestamp in the output.

## Usage

```
$ echoex ls /
`ls /`:
bin  boot  dev	etc  home  lib	lib64  lost+found  media  mnt  opt  proc  root	run  sbin  srv	sys  tmp  usr  var
$ echoex -t ls /
`ls /` at "Thu Aug 12 2021 14:51:43":
bin  boot  dev	etc  home  lib	lib64  lost+found  media  mnt  opt  proc  root	run  sbin  srv	sys  tmp  usr  var
```

## Install

Install the tool directly from source using go get:

```
$ go get github.com/rohrschacht/echoex
```
