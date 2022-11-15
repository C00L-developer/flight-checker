# Flight Path Tracker
This repo is a simple flight tracker implementation

## Context

There are over 100,000 flights a day, with millions of people and cargo being transferred around the world. With so many people and different carrier/agency groups, it can be hard to track where a person might be. In order to determine the flight path of a person, we must sort through all of their flight records. To create a simple microservice API that can help us understand and track how a particular person's flight path may be queried.

## Architecture

- Provide `gRPC` and `REST` apis:

```cmd
gRPC port: 9090
REST port: 8080
```

- REST api endpoint: `calculate` uses `post` method to trasfer JSON body

- [`protobuf` schema](./proto/flight/v1/flight.proto)

## Running the service

To run the service:

```bash
make run
```

To run the unit-test:

```bash
make test
```