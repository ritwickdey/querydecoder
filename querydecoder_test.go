package querydecoder_test

import (
	"log"
	"net/url"
	"testing"

	"github.com/ritwickdey/querydecoder"
)

func TestDecodeField(t *testing.T) {

	queryValues := url.Values{}
	queryValues.Add("is_super_user", "true")
	queryValues.Add("user_id", "12")
	queryValues.Add("user_name", "ritwick")

	var err error

	{ //test for pointer reciver

		field := ""
		err = querydecoder.New(queryValues).DecodeField("field", "", field)
		if err == nil {
			log.Panicln("should have error: Pointer error")
		}

	}

	{ // Test for boolean
		var isSuperUser bool

		err = querydecoder.New(queryValues).DecodeField("is_super_user", false, &isSuperUser)
		if err != nil {
			log.Panicln("should not have errors", err)
		}

		if isSuperUser == false {
			log.Panicln("super user flag should be true")
		}

	}

	{ // Test for int
		var userId int64
		err = querydecoder.New(queryValues).DecodeField("user_id", 0, &userId)
		if err != nil {
			log.Panicln("should not have errors", err)
		}
		if userId != 12 {
			log.Panicln("userId should be 12")
		}
	}

	{ // Test for string
		var userName string
		err = querydecoder.New(queryValues).DecodeField("user_name", "defaultStr", &userName)
		if err != nil {
			log.Panicln("should not have errors", err)
		}
		if userName != "ritwick" {
			log.Panicln("userName should be ritwick")
		}
	}

	{
		// Test for random field (default value test)
		var randomField string
		err = querydecoder.New(queryValues).DecodeField("random_str", "defaultStr", &randomField)
		if err != nil {
			log.Panicln("should not have errors", err)
		}
		if randomField != "defaultStr" {
			log.Panicln("random_str should be defaultStr")
		}
	}

}

func TestDecode(t *testing.T) {
	queryValues := url.Values{}
	queryValues.Add("is_super_user", "true")
	queryValues.Add("user_id", "12")
	queryValues.Add("user_name", "ritwick")

	u1 := user{
		IsSuperUser: false,
		RandomField: "random_value",
	}

	err := querydecoder.New(queryValues).Decode(&u1)

	if err != nil {
		log.Panicln("err should be nil", err)
	}

	if u1.IsSuperUser != true {
		log.Panicln("IsSuperUser should be true")
	}

	if u1.UserID != 12 {
		log.Panicln("UserID should be 12")
	}

	if u1.UserName != "ritwick" {
		log.Panicln("UserName should be ritwick")
	}

	if u1.RandomField != "random_value" {
		log.Panicln("RandomField should be random_value")
	}

}

type user struct {
	IsSuperUser bool   `query:"is_super_user"`
	UserName    string `query:"user_name"`
	UserID      int64  `query:"user_id"`
	RandomField string `query:"random_field"`
}
