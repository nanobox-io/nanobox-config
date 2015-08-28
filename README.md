# Nanobox Config

Parses config files and command line arguments into a singe map for ease of use. The config file is parsed first and layered over the default map, then the command line arguments are layered over that.

The resultant map contains defaults specified in the code, replacements from the config file, and replacements from the command line args.

**Example**:

```go
package main

import (
  config "github.com/pagodabox/nanobox-config"
  "os"
)

func main() {
  defaults := map[string]string{
    "default": "true",
    "key": "old"
  }
  config.Load(defaults, os.Args[1:])
  fmt.Printf("%v\n", config.Config)
}
```

**Usage**:
```bash
$ cat > ./test.config <<EOF
# this is an example config file
path testing
EOF
$ ./program ./test.config -key=value
{
  "key": "value"
  "path": "testing"
  "default": "true"
} 
```