Optional query parameter decoder for Golang


Example
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
   
    // Decode full struct
    u1 := User{}
    query := r.URL.Query()
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
    // if `is_dog` query param is not there, it'll assign default value 
    err := querydecoder.New(query).DecodeField("is_dog", true /* default value*/, &isDog)

    if err != nil {
        panic(err)
    }

}

```




Beachmark
```
cpu: Intel(R) Core(TM) i5-8257U CPU @ 1.40GHz

// Parse by struct tags
BenchmarkDecode-8   	 1629248	       742.5 ns/op	      40 B/op	       5 allocs/op

// Parse by key
BenchmarkDecodeField-8   	 7497049	       153.4 ns/op	       0 B/op	       0 allocs/op

//  manual parsing
BenchmarkManualDecode-8   	17626317	        60.22 ns/op	       0 B/op	       0 allocs/op



```