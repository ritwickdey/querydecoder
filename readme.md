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

cpu: Intel(R) Core(TM) i5-8257U CPU @ 1.40GHz
BenchmarkDecode-8                1677639               711.2 ns/op // Parse by struct tags
BenchmarkDecodeField-8           7735048               152.3 ns/op // Parse by key
BenchmarkManualDecode-8         19119616                60.06 ns/op //  manual parsing

BenchmarkJsonUnmarshal-8              737768              1746 ns/op //json.Unmarshal - Unrelated, but added to compare.


```