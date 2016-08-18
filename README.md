# IPAM Hello World Demo

## Prerequisites

IPAM leverages Docker for it's development & demonstration environments so you'll need to install
the latest Docker (1.12+) to try it out.

**Mac**

https://docs.docker.com/docker-for-mac/

**Windows 10 Only**

https://docs.docker.com/docker-for-windows/

**Ubuntu**

https://docs.docker.com/engine/installation/linux/ubuntulinux/

In addition IPAM is using make to provide an abstraction for complex Docker commands.  On Mac/Linux any version of GNU make is likely suitable.  On Windows something like http://gnuwin32.sourceforge.net/packages/make.htm may be suitable.  Otherwise the Docker commands can be
run directly using the Makefile as a guide for their format.

Once you have Docker and make available you can move on to **Trying it out...**


## Trying it out...

1. git clone git@github.com:RackHD/ipam.git
2. cd ipam
3. make
4. make run
5. http://localhost:8000/pools
