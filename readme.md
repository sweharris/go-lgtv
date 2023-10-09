# GO-LGTV

## sweharris fork

This fork is my version which adds slightly new functionality for my OLED
TV.  I attempted to PR to upstream but it didn't get any response.  So...

I've added the ability to send Toast messages and additional checking on
the power state of the TV ('cos OLED TVs go into "Active Standby" and that
caused the original code to turn _on_ when we wanted to turn off).

The library now uses wss on port 3001 since it seems LG is removing the
ws on port 3000 (at least it no longer works on my TV).

Everything else is as per the original author, except the "install" instructions
below.

## Introduction

Go-LgTv is a Golang package for discovering and controlling LG Smart TVs. This has been built against version 3.5 of WebOS, which uses a websocket connection for communication. In the past, LG's Smart TVs used their UDAP protocol for communication, but this changed at some undisclosed point. I therefore can't say exactly which versions of WebOS this will work with - it has been tested on version 3.5 running on a C7V 2017 Model TV.

## Installation

Get it:

```
go get github.com/sweharris/go-lgtv
```

## Examples

```
package main

import (
	"github.com/sweharris/go-lgtv/control"
	"github.com/sweharris/go-lgtv/discovery"
)

func main() {
	// Discover the TV on your local network using the IP address of your local gateway
	tv, err := discovery.Discover("192.168.1.1")

	// Or if you already know the IP address of your TV, create an instance of it directly
	tv, err = control.NewTV("192.168.1.129", "", "")

	// In order to use the TurnOn functionality, the TVs MAC address and the subnet mask of the local network must also be supplied
	tv, err = control.NewTV("192.168.1.129", "38:8C:50:6B:CD:B1", "255.255.255.0")

	// If you don't already have a client key, connect to it with an empty key and it will create a new one.
	// This call will block until the request to connect has been accepted on the TV.
	// The timeout value is in milliseconds.
	clientKey, err := tv.Connect("", 1000)

	// Or if you already have a client key from before, you can specify it to connect immediately
	_, err = tv.Connect("7668cb15d16a1a319f3731a9264b700b", 1000)

	// TurnOn uses WOL, and so relies on the TV being connected using ethernet
	err = tv.TurnOn()

	// Once connected, you can perform operations like play, pause, launch an app etc.
	err = tv.Play()
	err = tv.SetChannel(1)
	err = tv.SetVolume(30)
	_, err = tv.LaunchApp("netflix")
	err = tv.SwitchInput("HDMI_1")

	// Various things can be queried from the TV like getting a list of channels, installed apps, external inputs etc.
	channels, err := tv.ListChannels()
	inputs, err := tv.ListExternalInputs()
	apps, err := tv.ListInstalledApps()

	// You can switch to a certain channel/input/app directly from that object
	err = channels[0].Watch()
	_, err = apps[0].Launch()
	err = inputs[0].Switch()

	// Disconnect from the TV once you're done with it
	err = tv.Disconnect()
}
```

## A note on `TurnOn()`

This package uses Wake-On-LAN functionality to turn the TV on, as the normal networking stack is shut down when WebOS TVs are put in to standby. For this to work, the TV must be connected to the local network over ethernet (*not* Wifi) and WOL must be enabled in the TV's settings.

## A note on discovery

During development, I've been unable to get the TV to respond to SSDP requests (which is how other people seem to be doing discovery). As a result, discovery instead polls every IP address on the local network on port 1426, which WebOS 3.5 keeps open to HTTP GET requests. In response is sends some XML, which is used to identify the TV on the network. This method of doing discovery is slow, inefficient and at time unreliable (I've seen response times to the GET take up to 250ms, which makes the overall search slow). For that reason, if at all possible I recommend you use the IP address of the TV directly instead of using discovery, which will need to be re-written at some point.

Also note that this method makes a couple of assumptions:

* There is only one TV on the network that would respond to this request (it will return the first one found)
* Your local network is a simple 256 address network with a single gateway like a router

UPDATE:
As of V1.1, it appears discovery seems to be completely broken. I need to redo it before it will be in a usable state.

## Dependencies

This package uses Gorilla's websocket implementation (https://github.com/gorilla/websocket) for the underlying websocket management. It also uses ghthor's Wake-On-LAN package (https://github.com/ghthor/gowol) for WOL functionality. Everything else used is standard.

## Release notes

**V1.1.1**

- Made `IsConnected` publically available.
- Made `ClientKey` publically available.
- Made `Connect()` threadsafe.

**V1.1**

- Added the ability to use Wake-On-LAN to turn on the TV using the `TurnOn()` function.

**V1.0.1**

- Added a missing permission in order to be able to use media control.

**V1.0**

- The first version adds support for the following operations after connecting to the TV:
    - Volume controls (up, down, set (0-100), get)
    - Mute controls (get, set)
    - Media controls (play, pause, stop, rewind, fastforward)
    - Channel controls (up, down, set, get current, list all)
    - Get program list for current channel
    - External input controls (list inputs, switch input)
    - App controls (List installed apps, Launch app)
    - Turn off
