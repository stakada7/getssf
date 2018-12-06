# getssf

## Status
Development only.

## Installation

To compile from source, you need Go1.10 or later (including $GOPATH setup) and glide for dependency management. After setup, then clone the source code by running the following command,

```
$ mkdir -p $GOPATH/src/github.com/stakada7
$ cd $GOPATH/src/github.com/stakada7
$ git clone https://github.com/stakada7/getssf
```

To fetch dependencies and build, run the following make tasks,

```
make bundle
make
```

## Usage

To run getssf, you must representation configuration file (See CONFIGURATION.md about details),

## Configuration

see CONFIGURATION.md about details.

NOTE: The default configuration is not specify Twitter account access token and account list. you should create twitter app and you get twitter access token.

## Specification

## License
MIT License

## TODO

* [x] Whether the target data was received correctly with multiple connections.
    * done. If it does not reach rate limiting errors etc, it seems to be able to receive without any problem..
* [x] Implementation of functions that can dynamically change by changing the file in the account list
* [x] Golang-like folder structure
* [x] Understand 420 error and reconnection logic in the implementation of Anaconda library
    * In the Anaconda library, we are handling rate limiting errors, so there is no need for the user side to handle it.
    * In Throttling you can specify it with SetDelay method, but it is now off. In the future, it may be turned on by default.(by anaconda)

* [ ] Think about the unit of the output file. 1 day, 1 hour, 5 minutes, etc. (What is a handy unit?)
* [ ] dependency manager change to dep.
