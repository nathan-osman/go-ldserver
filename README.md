## go-ldserver

[![GoDoc](https://godoc.org/github.com/nathan-osman/go-ldserver?status.svg)](https://godoc.org/github.com/nathan-osman/go-ldserver)
[![MIT License](http://img.shields.io/badge/license-MIT-9370d8.svg?style=flat)](http://opensource.org/licenses/MIT)

Given text files with timing information for lights, ldserver will run the light display in realtime.

### Building

To build the `ldserver` binary, ensure that Docker is installed and run:

    make

### Usage

ldserver expects a single command-line argument &mdash; a JSON configuration file. The following example specifies three different light sources (a debug light, a GPIO pin, and a WebSocket light):

    {
      "manager": {
        "lights": {
          "debug": {
            "names": ["debug1"]
          },
          "gpio": {
            "pins": [
              {
                "number": 17,
                "name": "gpio17"
              }
            ]
          },
          "ws": {
            "host": "ws://localhost:5000",
            "names": ["ws1"]
          }
        }
      },
      "server": {
        "addr": ":8000"
      }
    }
