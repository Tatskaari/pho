# pho
Functional stuff for go

## result.Result[T]

```golang
package main

import (
        "fmt"
        "os"
        "io/ioutil"
        "path/filepath"

        "github.com/tatskaari/pho/result"
)

func main() {
        wd := result.Wrap(os.Getwd())
        path := result.Map(wd, func(wd string) string {
                 return filepath.Join(wd, "file.txt")
        })
        file := result.ThenWrap(path, os.Open)
        defer file.Use(func(file *os.File) {
                file.Close()
        })
        bytes := result.ThenWrap(file, func(f *os.File) ([]byte, error) {
                return ioutil.ReadAll(f)
        })
        str, err := result.Map(bytes, func(b []byte) string {
                return string(b)
        }).Unwrap()

        if err != nil {
                fmt.Println("Error reading file: %v", err)
                return
        }
        fmt.Println(str)

}
```
