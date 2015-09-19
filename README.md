# Graal
_go remote administration api for linux_

Graal is a REST api server to administer a linux machine, programmed in Go.

Some services are just reading information from the server, while other services may act on it.

## Requirements
This project require at least the version 1.5 of go

## Directories
- src/formatters: output formatters: json, xml, etc...
- src/services: services, classifed by type
- bin: contains the binary, after built
- lib: contains librairies when building services as shared librairies
