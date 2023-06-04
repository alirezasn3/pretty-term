# Pretty Terminal

Use this package to change text color, background color, text decoration and cursor postion in windows terminal app.

Note that this will not work normal cmd or powershell terminals and they must be opened in windows terminal app.

# Usage

Import the package and use it like this:

```go
import (
	"fmt"
	prettyterm "github.com/alirezasn3/pretty-term"
)

func main() {
    prettyterm.ClearTerminal()
    prettyterm.SetCursor(0, 0)
    fmt.Print("Hello World in ")
    prettyterm.SetColor(prettyterm.Yellow)
    fmt.Print("Yellow")
    fmt.Prinln("!")
}
```