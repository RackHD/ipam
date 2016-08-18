[![Build Status](https://travis-ci.org/RackHD/ipam.svg?branch=master)](https://travis-ci.org/RackHD/ipam)

# Overview

IPAM intends to fill a gap around dynamic configuration of resources in datacenter management workflows.  A flexible solution providing extensible configuration, reservation, and auditing of IP resources will provide a capability to more easily integrate with different operational designs.

To that end we are thinking of IPAM as more of a resource allocation service with a strong API & flexible data model while providing less out of the box integration with network services like DHCP & DNS or concerns such as routing topologies.

Currently IPAM development is driving towards providing the initial technology for it's REST API and underlying data model.  As such the plan of record is to implement all of the API end points and their associated data models in a first pass while following up with specific business logic to implement the initial lease allocation and reservation process.  The next pass will incorporate extensibility features and flush out Consul integration.

# Design Overview

## Configuration

Users will be able to create Pools & Subnets through the API. A Pool is a container for one or more Subnets. A Subnet will define a range of IP Addresses that it manages for reservations. Metadata for both objects is a free form field represented as a JSON object to provide storage for Pool/Subnet specific application data.

## Reservation

Users will be able to reserve IP's from either Pools or Subnets. A request against a Subnet will only attempt to reserve addresses in the requested Subnet. A request against a Pool will attempt to reserve addresses in any of the Subnets a Pool manages.

The Reservation object is created upon request and is used to atomically reserve free addresses in the Subnet(s) requested by updating free addresses with a reference to the reservation object.
A list of obtained Leases is returned along with the reservation object to the requester upon completion.

## Inspection
Users can view Pool, Subnet, and Reservation data via the API. Reservations will provide metadata from the Pool & Subnet to which it belongs so that data can be leveraged for additional application specific scenarios.

## Extensibility

The IPAM data model is designed to be extensible allowing consumers to set application specific data on any of the entities provided. IPAM considers this data to be opaque and will provide it upon request to the consumer.
IPAM will also provide an event based notification model over a stateless message bus when configured to do so. Events will be generated for CRUD based events regarding Pools & Subnets as well as reservation events.
Consumers can utilize the events and the available API's to customize their workflows for their application scenarios.

For example Leases has an update API which is provided to allow consumers to store additional data related to individual allocations based on their application scenario. For instance if an allocated Lease is assigned to a particular entity by MAC address the consumer could place that data into the Lease metadata object for tracking purposes.

# Getting Started

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

1. git clone git@github.com:RackHD/ipam.git
2. cd ipam
3. make
4. make run
5. http://localhost:8000/pools

# Details

## Planned API End Points

### Pool Routes

* GET /pools
* GET /pools/{id}
* POST /pools
* DELETE /pools/{id}
* PATCH /pools/{id}

### Subnet Routes

* GET /pools/{id}/subnets
* GET /subnets/{id}
* POST /pools/{id}/subnets
* DELETE /subnets/{id}
* PATCH /subnets/{id}

### Reservation Routes

* GET /subnets/{id}/reservations
* GET /pools/{id}/reservations
* GET /reservations/{id}
* POST /subnets/{id}/reservations
* POST /pools/{id}/reservations
* DELETE /reservations/{id}
* PATCH /reservations/{id}

### Lease Routes

* GET /reservations/{id}/leases
* GET /leases/{id}
* PATCH /leases/{id}

## Planned Object model

### Pool

Properties
* ID
* Name
* Tags
* Metadata

Relationships
* Subnets

### Subnet

Properties
* ID
* Name
* Tags
* Metadata

Relationships
* Pool

### Reservation

Properties
* ID
* Name
* Tags
* Metadata

Relationships
* Pool/Subnet
* Leases

### Lease

Properties

* ID
* Address
* Tags
* Metadata

Relationships
* Reservation

## Extensibility Interface

Initial extensibility will be provided through the tagging and metadata fields on IPAM data models through the API. An event based notification system will be supported via a message bus with initial support for RabbitMQ and/or Nats (http://nats.io).

As users interact with the API events based on CRUD actions will be sent to the message bus to provide asynchronous feedback to consumers. Further extensibility may be achieved by providing asynchronous request/response patterns to bus consumers for pre-commit hooks such as data validation.

## Consul Integration

IPAM will provide the ability to integrate with Consul for service discovery. Integration will include both identifying a MongoDB service for persistence of the IPAM data model as well as registration of the IPAM service end points. The Consul integration should be considered optional and as such IPAM will offer a command line configuration mechanism suitable for physical and containerized deployments.

* --mongo = List of comma separated MongoDB servers which will be ignored if the consul flag is present.
* --consul = Connection info for a suitable Consul agent.
* --consul-mongo = Service name for the MongoDB to be utilized by IPAM (defaults to mongodb).
* --consul-service = Service name for IPAM to register itself as with Consul (defaults to ipam).

## MVP Deployment Model

The initial IPAM deployment model will consist of a development/demo environment based on Docker Compose. The composed IPAM application will consist of a single IPAM service in a container paired with a single MongoDB service in a container. Configuration of the container based environment will leverage the MongoDB command line configuration parameter.

A Consul based deployment (either physical, container, etc) will leverage the Consul command line configuration parameters to instruct IPAM which MongoDB service to locate and what the IPAM service parameters are for service registration with a similar mechanism will be put in place for the message bus extensibility.
