# GoArguments

![GitHub repo size](https://img.shields.io/github/repo-size/danny270793/GoArguments)
![GitHub code size in bytes](https://img.shields.io/github/languages/code-size/danny270793/GoArguments)

![GitHub commit activity](https://img.shields.io/github/commit-activity/m/danny270793/GoArguments)
![GitHub Downloads (all assets, all releases)](https://img.shields.io/github/downloads/danny270793/GoArguments/total)

Library to parse CMD arguments

## Ussage

Import library

```go
import (
	"fmt"
	"os"

	"github.com/danny270793/goarguments"
)
```

Create an array of avaialable `Command`s in the CMD
Every `Command` have:

- `Name` and `Description` to be show on the `help` command
- `Arguments` which is a map (key-value) of the arguments (with they default value) that the `Command` can receive
- `Callback` which is the function that will be called when the `Command` name's matches, it receives the `Arguments` as argument of the function
- Finally, can receive a list of sub`Commands`


```go
commands := []goarguments.Command{
    {
        Name: "import",
        Description: "Import container images from filesystem",
        Arguments: map[string]string{"--folder": "./images"},
        Callback: func(arguments map[string]string) {
            fmt.Println(arguments)
        },
        Commands: []goarguments.Command{},
    },
    ...
}
```

Then can parse the `os.Args` to the array of available `command`, if there is an error show the help, which is generated by default with all the `command`s available

```go
chain, err := goarguments.ParseCommand(commands, os.Args[1:])
if err != nil {
    fmt.Printf("%s\n\n", err)
    goarguments.ShowHelp(commands)
    os.Exit(1)
}
```

`goarguments.ParseCommand` will return a array of cllbacks that will be called passing its respective arguments

```go
for _, item := range chain {
    item.Callback(item.Arguments)
}
```

Can check more examples [here](https://github.com/danny270793/GoArguments/tree/master/examples)

## Follow me

[![YouTube](https://img.shields.io/badge/YouTube-%23FF0000.svg?style=for-the-badge&logo=YouTube&logoColor=white)](https://www.youtube.com/channel/UC5MAQWU2s2VESTXaUo-ysgg)
[![GitHub](https://img.shields.io/badge/github-%23121011.svg?style=for-the-badge&logo=github&logoColor=white)](https://www.github.com/danny270793/)
[![LinkedIn](https://img.shields.io/badge/linkedin-%230077B5.svg?logo=data:image/svg%2bxml;base64,PHN2ZyB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciIHhtbG5zOnhsaW5rPSJodHRwOi8vd3d3LnczLm9yZy8xOTk5L3hsaW5rIiB2aWV3Qm94PSIwLDAsMjU2LDI1NiIgd2lkdGg9IjUwcHgiIGhlaWdodD0iNTBweCIgZmlsbC1ydWxlPSJub256ZXJvIj48ZyBmaWxsPSIjZmZmZmZmIiBmaWxsLXJ1bGU9Im5vbnplcm8iIHN0cm9rZT0ibm9uZSIgc3Ryb2tlLXdpZHRoPSIxIiBzdHJva2UtbGluZWNhcD0iYnV0dCIgc3Ryb2tlLWxpbmVqb2luPSJtaXRlciIgc3Ryb2tlLW1pdGVybGltaXQ9IjEwIiBzdHJva2UtZGFzaGFycmF5PSIiIHN0cm9rZS1kYXNob2Zmc2V0PSIwIiBmb250LWZhbWlseT0ibm9uZSIgZm9udC13ZWlnaHQ9Im5vbmUiIGZvbnQtc2l6ZT0ibm9uZSIgdGV4dC1hbmNob3I9Im5vbmUiIHN0eWxlPSJtaXgtYmxlbmQtbW9kZTogbm9ybWFsIj48ZyB0cmFuc2Zvcm09InNjYWxlKDUuMTIsNS4xMikiPjxwYXRoIGQ9Ik00MSw0aC0zMmMtMi43NiwwIC01LDIuMjQgLTUsNXYzMmMwLDIuNzYgMi4yNCw1IDUsNWgzMmMyLjc2LDAgNSwtMi4yNCA1LC01di0zMmMwLC0yLjc2IC0yLjI0LC01IC01LC01ek0xNywyMHYxOWgtNnYtMTl6TTExLDE0LjQ3YzAsLTEuNCAxLjIsLTIuNDcgMywtMi40N2MxLjgsMCAyLjkzLDEuMDcgMywyLjQ3YzAsMS40IC0xLjEyLDIuNTMgLTMsMi41M2MtMS44LDAgLTMsLTEuMTMgLTMsLTIuNTN6TTM5LDM5aC02YzAsMCAwLC05LjI2IDAsLTEwYzAsLTIgLTEsLTQgLTMuNSwtNC4wNGgtMC4wOGMtMi40MiwwIC0zLjQyLDIuMDYgLTMuNDIsNC4wNGMwLDAuOTEgMCwxMCAwLDEwaC02di0xOWg2djIuNTZjMCwwIDEuOTMsLTIuNTYgNS44MSwtMi41NmMzLjk3LDAgNy4xOSwyLjczIDcuMTksOC4yNnoiPjwvcGF0aD48L2c+PC9nPjwvc3ZnPg==&logoColor=white&style=for-the-badge)](https://www.linkedin.com/in/danny270793)

## LICENSE

[![GitHub License](https://img.shields.io/github/license/danny270793/TSFramework)](license.md)

## Version

![GitHub Tag](https://img.shields.io/github/v/tag/danny270793/GoArguments)
![GitHub Release](https://img.shields.io/github/v/release/danny270793/GoArguments)

Last update 07/08/2024

Last update 07/11/2023
