# ExportIt
-----
## What is it?
Just write your function, and you'll have a CLI tool-no need to worry about flags, figuring out their names, or creating a config struct to manage them. Simply focus on the function, and let me take care of the rest.


## Usage
1. create your functions
```go
func Curl(link string) (string, error) {
	res, err := http.Get(link)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func Cat(path string) (string, error) {
	fmt.Println("mama")

	data, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	fmt.Println(string(data))
	return string(data), nil
}
```

2. use the tool
```go
import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/nasoooor29/ExportIT"
)

func main() {
	ExportIT.ExecCli(
		"appName",
		"short Desc",
		"Long Description",
		ExportIT.CliNamedParam(Curl),
		ExportIT.CliNamedParam(Cat),
	)
}

```
3. you are done!!
```sh
❯ go build
❯ ./ExportIT
Long Description

Usage:
  appName [command]

Available Commands:
  Cat         main.Cat
  Curl        main.Curl
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command

Flags:
  -h, --help   help for appName

Use "appName [command] --help" for more information about a command.
❯ 
```

## what is the caviat?
*	currently it work only with simple data types 
		- string
		- int
		- bool
*   only work with named arguments
*   can't pipe things into it 


## todos:
* better shorthand naming
