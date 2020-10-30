localhost-proxy
=======================
[![GitHub version](https://badge.fury.io/gh/dersnek%2Flocalhost-proxy.svg)](https://badge.fury.io/gh/dersnek%2Flocalhost-proxy)

#### Listens to local network and re-routes requests to localhost. Useful for WSL.

Why is it needed? Servers running in WSL are not accessible from local network without black magic, so this might be the simplest way to visit your Ruby on Rails app running in WSL from another PC. Maybe there are other uses, this what I need it for ¯\\\_(ツ)\_/¯

It's super barebones, no error handling, just a quick and dirty fix to a small problem.

### Installation

Just get the latest binary and run it

Alternatively, clone the repo and `go run .` (you'll need golang installed in this case, of course).

### Usage

When you start the program, it will ask you which port to re-route requests to. This is the port which your server is listening to WSL.

Then **localhost-proxy** will start listening at `:8080`. Send requests from any PC to your local IP address, something like `192.168.1.x:8080`.
Don't forget to allow connections in your firewall.
