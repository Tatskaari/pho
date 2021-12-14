# pho
Functional stuff for go

## result.Result[T]

```golang
import (
	"fmt"
	"os"
	"ioutil"

	"github.com/tatskaari/pho/result"
)

func main() {
	wd := result.Wrap(os.Wd)
	path := result.Map(wd, func(wd string) string {
		 return filepath.Join(wd, "file.txt")
	))
	file := result.ThenWrap(path, os.Open)
	defer file.Use(func(file os.File) {
		file.Close()
	})
	bytes := result.ThenWrap(file, ioutil.ReadAll)
	str := Map(bytes, func(b []byte) string {
		return string(b)
	})
	str, err := str.Unwrap()
	if err != nil {
		fmt.Println("Error reading file: %v", err)	
		return 
	}
	fmt.Println(str)

}

```
