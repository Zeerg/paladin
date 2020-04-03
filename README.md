<p align="center"> 
  <img src="images/logo.png" width="250" title="paladin" align="center">
  <p align="center">A blueteam tool to simulate attacks using Atomic Red Team and other techniques.</p>
</p>

---

## Getting started

### Ping Exfil

The ping exfil subcommand is working and implements a basic ping exfil attack scenario.

### Atomic Red Team


### DNS Exfil
WIP

### Infection Gopher
WIP

## Local Development

This project requires Go to be installed. On OS X with Homebrew you can just run `brew install go`.

Building locally requires statik 

```console
$ go get github.com/rakyll/statik
```

Running it then should be as simple as:

```console
$ make
$ ./bin/paladin
```