go-semver - Semantic versioning library
=======================================

[![wercker status](https://app.wercker.com/status/2cdcb3415ccf131f7adfe363164357fe/s "wercker status")](https://app.wercker.com/project/bykey/2cdcb3415ccf131f7adfe363164357fe)
[![Coverage Status](https://coveralls.io/repos/zhevron/go-semver/badge.svg?branch=HEAD)](https://coveralls.io/r/zhevron/go-semver?branch=HEAD)
[![GoDoc](https://godoc.org/gopkg.in/zhevron/go-semver.v0/semver?status.svg)](https://godoc.org/gopkg.in/zhevron/go-semver.v0/semver)

**go-semver** is a [semantic versioning](http://semver.org/) library for [Go](https://golang.org/).  

For package documentation, refer to the GoDoc badge above.

## Installation

```
go get gopkg.in/zhevron/go-semver.v0/semver
```

## Usage

### Parse a version number

```go
package main

import (
  "fmt"

  "gopkg.in/zhevron/go-semver.v0/semver"
)

func main() {
  // A semver string
  versionString := "2.0.0-beta1+20150505"

  // Parse it into a semver.Version
  version, err := semver.ParseVersion(versionString)
  if err != nil {
    panic(err)
  }

  // Print the major, minor and patch versions
  fmt.Printf("Version: %d.%d.%d\n", version.Major, version.Minor, version.Patch)
}
```

### Constraints

```go
package main

import (
  "fmt"

  "gopkg.in/zhevron/go-semver.v0/semver"
)

func main() {
  // A semver string
  versionString := "2.0.0-beta1+20150505"

  // A version constraint string
  constraintString := ">=1.5.0"

  // Parse the version string into a semver.Version
  version, err := semver.ParseVersion(versionString)
  if err != nil {
    panic(err)
  }

  // Parse the constraint string into a semver.Constraint
  constraint, err := semver.ParseConstraint(constraintString)
  if err != nil {
    panic(err)
  }

  // Check if the version matches the constraint
  if constraint.Match(version) {
    fmt.Println("The version matches the constraint")
  } else {
    fmt.Println("The version does NOT match the constraint")
  }
}
```

## License

**go-semver** is licensed under the [MIT license](http://opensource.org/licenses/MIT).
