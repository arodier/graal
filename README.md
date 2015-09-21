# Graal
_go remote administration api for linux_

Graal is a REST api server to administer or to monitor a linux machine, programmed in Go.

- It is compiled as a native ELF binary.
- It's fast
- It's light
- It's safe (strongly typed language)
- It does not need extra librairies (php/python/perl,etc...)
- It's simple

Some services are just reading information from the server, while other services may act on it.

## Requirements
To use shared librairies, you will require at least the version 1.5 of go

## Building
- make clean: remove binaries
- make build: build the main binary and the archives
- make run: run the server, on default IP address and port (http://127.0.0.1:1188/)

## Directories
- src/formatters: output formatters: json, xml, etc...
- src/services: services, classifed by type
- bin: contains the binary, after built
- lib: contains librairies when building services as shared librairies
- docs: some documents about the API

## Running
At this time, there is no authentication implemented.
If you need it, it is suggestted to use nginx or Apache as a reverse proxy,
with an authentication & authorization backend (pam, ldap, radius, etc…)

Neither there is SSL encryption.

## Command line options
If you start the program without any option, it will listen on 127.0.0.1, on port 1188

To change the address or the port number, use the following syntax:

    graal -ip 192.168.42.42 -port 1909

To see the other options, use graal --help

## Call examples
All data is retuned in a JSON encoded object, in a ‘Data’ field

### Testing the API

Say Hello!

    GET /hello

    {"Data":"Hello, how are you?"}

### Get system time

    GET /system/time

    {"Data":"2015-09-20T08:23:58.880287421+01:00"}

### Get system statistics

    GET /system/stats

    {
      "Data": {
        "Load": [
          "0.08",
          "0.15",
          "0.21",
          "2/882",
          "18247"
        ],
        "Uptime": [
          "1768902.39",
          "2666255.90"
        ],
        "MemInfo": [
          {
            "Name": "MemTotal",
            "Value": "16392264 kB"
          },{
            "Name": "MemFree",
            "Value": "810352 kB"
          },
        ...
        ]
      }
    }
