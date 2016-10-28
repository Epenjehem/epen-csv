# epen-csv
CSV file manipulation package for Go (Golang)

## Constructor
```go
func NewCSV(filePath string, settings map[string]interface{}) (*EpenCSV, error)
```
See usage section

## Settings
| Types            | Default | Data Type | Description                                         |
| ---------------- |:------- | ---------:| --------------------------------------------------- |
| comma            | ','     | *rune*    | field delimiter                                     |
| trimLeadingSpace | true    | *bool*    | If true, leading white space in a field is ignored. |

## Features
Currently supported features:

### GetMean
```go
func (r *EpenCSV) GetMean(column int) (float64, error)
```
Return a mean of a single column from CSV file

### GetMedian
```go
func (r *EpenCSV) GetMedian(column int) (float64, error)
```
Return a median of a single column from CSV file

### Print
```go
func (r *EpenCSV) Print()
```
Print a csv in console

## Usage
``` go
import (
	"fmt"
	"github.com/epenjehem/epen_csv"
	"log"
)

func main() {
	settings := map[string]interface{}{
		"comma": '\t',
		"trimLeadingSpace": false,
	}

	r, err := epen_csv.NewReport("test-data.csv", settings)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(r.GetMean(2))
}
```
