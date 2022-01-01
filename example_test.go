package querydecoder_test

import (
	"log"
	"net/http"

	"github.com/ritwickdey/querydecoder"
)

type User struct {
	IsSuperUser bool   `query:"is_super_user"`
	UserName    string `query:"user_name"`
	UserID      int64  `query:"user_id"`
}

func handler1(w http.ResponseWriter, r *http.Request) {

	u1 := User{}
	// Decode into struct
	err := querydecoder.New(r.URL.Query()).Decode(&u1)

	if err != nil {
		panic(err)
	}

	log.Println(u1)

}

func handler2(w http.ResponseWriter, r *http.Request) {

	var isDog bool
	err := querydecoder.New(r.URL.Query()).DecodeField("is_dog", &isDog)
	// if `is_dog` query param is not there, it'll not modify the value of the variable.

	if err != nil {
		panic(err)
	}

}

func Example() {
	http.HandleFunc("/foo1", handler1)
	http.HandleFunc("/foo2", handler2)
	http.ListenAndServe(":8090", nil)
}
