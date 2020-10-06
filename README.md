scotch
===

Compare scope string

Usage
---

### Basic Usage

```go
import "github.com/sawadashota/scotch"

func main() {
    required := scotch.New("project:foo=read")
    
    if required.Satisfy("project:*") {
        // Authorized
    }
}
```

Installation
---

```
go get -u github.com/sawadashota/scotch
```
