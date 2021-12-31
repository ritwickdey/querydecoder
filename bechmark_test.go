package querydecoder_test

import (
	"encoding/json"
	"net/url"
	"strconv"
	"testing"

	"github.com/ritwickdey/querydecoder"
)

func BenchmarkDecode(t *testing.B) {
	queryValues := url.Values{}
	queryValues.Add("is_super_user", "true")
	queryValues.Add("user_id", "12")
	queryValues.Add("user_name", "ritwick")
	queryValues.Add("height", "1.0322")

	u1 := user{
		IsSuperUser: false,
		RandomField: "random_value",
	}

	for i := 0; i < t.N; i++ {
		if err := querydecoder.New(queryValues).Decode(&u1); err != nil {
			panic(err)
		}
	}

}

func BenchmarkDecodeField(t *testing.B) {
	queryValues := url.Values{}
	queryValues.Add("is_super_user", "true")
	queryValues.Add("user_id", "12")
	queryValues.Add("user_name", "ritwick")
	queryValues.Add("height", "1.0322")

	var isSuperUser bool
	var userId int64
	var userName string
	var height float32

	for i := 0; i < t.N; i++ {
		decoder := querydecoder.New(queryValues)
		decoder.DecodeField("is_super_user", false, &isSuperUser)
		decoder.DecodeField("user_id", 0, &userId)
		decoder.DecodeField("user_name", false, &userName)
		decoder.DecodeField("height", 0.1, &height)
	}

}

func BenchmarkManualDecode(t *testing.B) {
	queryValues := url.Values{}
	queryValues.Add("is_super_user", "true")
	queryValues.Add("user_id", "12")
	queryValues.Add("user_name", "ritwick")
	queryValues.Add("height", "1.0322")

	u1 := user{
		IsSuperUser: false,
		RandomField: "random_value",
	}

	for i := 0; i < t.N; i++ {
		var err error
		u1.IsSuperUser = queryValues.Get("is_super_user") == "true"

		uIdStr := queryValues.Get("user_id")
		u1.UserID, err = strconv.ParseInt(uIdStr, 16, 64)
		if err != nil {
			panic(err)
		}

		u1.UserName = queryValues.Get("user_name")

		heightStr := queryValues.Get("height")
		f64, err := strconv.ParseFloat(heightStr, 32)
		if err != nil {
			panic(err)
		}

		u1.Height = float32(f64)
	}

}

func BenchmarkJsonUnmarshalField(t *testing.B) {
	u1 := user{
		IsSuperUser: false,
		RandomField: "random_value",
		UserName:    "Ritwick Dey",
		UserID:      231,
		Height:      1.202,
	}

	jsonBytes, _ := json.Marshal(u1)

	for i := 0; i < t.N; i++ {
		json.Unmarshal(jsonBytes, &user{})
	}

}
