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


func ServeHTTP(w http.ResponseWriter, r *http.Request) {
   
    u1 := User{}
    query := r.URL.Query()
    err := querydecoder.New(query).Decode(&u1)
   
    if err != nil {
        panic(err)
    }
   
   log.Println(u1) 

}

```



