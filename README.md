# Nanobox Config

Parses config files and command line arguments into a singe map for ease of use. The config file is parsed first and loaded into the map, then the command line arguments are layered over the top of the config map.

The resultant map contains defaults specified in the code, replacements from the config file, and replacements from the command line args.

**Example**:

```go
package main

import (
  "github.com/pagodabox/na-config" config
)

func main() {
  defaults := map[string]string{
    "default": "true",
    "key": "old"
  }
  config.Load(defaults)
  fmt.Printf("%v\n", config.all())
}
```

**Usage**:
```bash
$ cat > ./test.config <<EOF
# this is an example config file
path=testing
EOF
$ ./program ./test.config -key=value
{
  "key": "value"
  "path": "testing"
  "default": "true"
} 
```