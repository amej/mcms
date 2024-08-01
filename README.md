# My Go Application mcms
 
Build a Simple Metrics Collection and Monitoring System
## Overview

This application provides various functionalities including metrics handling, alerting, and visualization.

## Download
./bin/mcms

# Building locally.
Prerequisite:  Define GOPATH & GOBIN
`go build`

# Install locally
`go install`

# Run
Prerequisite: CWD should allow creation of metrics.json file.

# Help
`mcms --help`
`mcms`

## Commands

### `mcms`

The base command for the application.

#### `mcms metrics`

Displays metrics at specified interval. ( Units: seconds )

`mcms metrics -i 20`

#### `mcms api`

Runs API service on localhost:1983

Example to override.

`mcms metrics -p 127.3.4.5:9999`

Fetch data by hitting following endpoints.
Default:  
1. http://localhost:1983/metrics
2. http://localhost:1983/metrics/aggregate


#### `mcms alert`
Enchancement
Handles alerting based on metrics.

#### `mcms visualize`
Enhancement
Visualizes metrics data.

## Package Documentation

In Github, detailed documentation for each package can be found in the sidebar generated content for  '.go' files
Locally, one can run godoc to display the description of package &/ functions.
