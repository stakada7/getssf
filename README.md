# getssf

## Status
Development only.

## Installation

To compile from source, you need Go1.01 or later (including $GOPATH setup) and glide for dependency management. After setup, then clone the source code by running the following command,

```
$ mkdir -p $GOPATH/src/github.com/tribalmedia
$ cd $GOPATH/src/github.com/tribalmedia
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

## TODO

* [x] Whether the target data was acquired correctly with multiple connections.
    * done. If it does not reach rate limiting errors etc, it seems to be able to acquire without any problem..
* [x] Identify, design, and implement accounts and data that should be targeted in addition to your own account
    * The account you should acquire with stream filter is your company's account. It is necessary to confirm in part because it stated that it is necessary to acquire in the contribution list / impression of the effect measurement.

* [x] Implementation of functions that can dynamically change by changing the file in the account list
* [x] Golang-like folder structure
* [ ] Think about the unit of the output file 1 day, 1 hour, 5 minutes, etc. (What is a handy unit?)

* [ ] Think about at what stage we will grant the first release tag

* [x] Understand 420 error and reconnection logic in the implementation of Anaconda library
    * In the Anaconda library, we are handling rate limiting errors, so there is no need for the user side to handle it.
    * In Throttling you can specify it with SetDelay method, but it is now off. In the future, it may be turned on by default.
