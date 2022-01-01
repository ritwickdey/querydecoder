# querydecoder

`querydecoder` is a Go Package to populating struct from optional query parameters using 'query' tag. This package can be used to populate value into primitive variable.


#### Example
```go
import (
	"github.com/ritwickdey/querydecoder"
)

type User struct {
	IsSuperUser bool   `query:"is_super_user"`
	UserName    string `query:"user_name"`
	UserID      int64  `query:"user_id"`
}

// Parse by struct tags
func ServeHTTP(w http.ResponseWriter, r *http.Request) {
    
    u1 := User{}
    query := r.URL.Query()
    // Decode into struct
    err := querydecoder.New(query).Decode(&u1)
   
    if err != nil {
        panic(err)
    }
   
   log.Println(u1) 

}

// OR,
// Parse by key
func ServeHTTP2(w http.ResponseWriter, r *http.Request) {
  
    var isDog bool
    err := querydecoder.New(query).DecodeField("is_dog", &isDog)
    // if `is_dog` query param is not there, it'll not modity variable.

    if err != nil {
        panic(err)
    }

}

```


#### Benchmark 

```
goos: darwin
goarch: arm64 (Mac M1 Chip)
pkg: github.com/ritwickdey/querydecoder

BenchmarkDecode-8                2019590                603.1 ns/op //Parse by struct tags
BenchmarkDecodeField-8          10521158                113.1 ns/op //Parse by key
BenchmarkManualDecode-8         25947067                46.74 ns/op //Manual parsing

BenchmarkJsonUnmarshal-8          883681                 1413 ns/op //json.Unmarshal - Unrelated, but added to compare.


```




