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
		err = querydecoder.New(queryValues).DecodeField("field", field)
		if err == nil {
			log.Panicln("should have error: Pointer error")
		}

	}

	{ // Test for boolean
		var isSuperUser bool = false

		err = querydecoder.New(queryValues).DecodeField("is_super_user", &isSuperUser)
		if err != nil {
			log.Panicln("should not have errors", err)
		}

		if isSuperUser == false {
			log.Panicln("super user flag should be true")
		}

	}

	{ // Test for int
		var userId int64
		err = querydecoder.New(queryValues).DecodeField("user_id", &userId)
		if err != nil {
			log.Panicln("should not have errors", err)
		}
		if userId != 12 {
			log.Panicln("userId should be 12")
		}
	}

	{ // Test for string
		var userName string
		err = querydecoder.New(queryValues).DecodeField("user_name", &userName)
		if err != nil {
			log.Panicln("should not have errors", err)
		}
		if userName != "ritwick" {
			log.Panicln("userName should be ritwick")
		}
	}

	{
		// Test for random field (default value test)
		var randomField string = "defaultStr"
		err = querydecoder.New(queryValues).DecodeField("random_str", &randomField)
		if err != nil {
			log.Panicln("should not have errors", err)
		}
		if randomField != "defaultStr" {
			log.Panicln("random_str should be defaultStr")
		}
	}

}

func TestDecodePointer(t *testing.T) {

	queryValues := url.Values{}
	queryValues.Add("random_pointer", "123")

	{
		u1 := user{}
		err := querydecoder.New(queryValues).Decode(&u1)
		if err != nil {
			log.Panicln("err should be nil", err)
		}
		if u1.RandomPointer == nil {
			log.Panicln("RandomPointer is nil")
		}
		if *u1.RandomPointer != 123 {
			log.Panicln("RandomPointer incorrect value")
		}
	}

	{
		var num *int32
		err := querydecoder.New(queryValues).DecodeField("random_pointer", &num)
		if err != nil {
			log.Panicln("err should be nil", err)
		}

		if num == nil {
			log.Panicln("num is nil")
		}
		if *num != 123 {
			log.Panicln("num incorrect value")
		}

	}

	{
		var num *int32
		err := querydecoder.New(queryValues).DecodeField("random_pointer_4343", &num)
		if err != nil {
			log.Panicln("err should be nil", err)
		}
		if num != nil {
			log.Panicln("num should be nil")
		}
	}

}

func TestDecode(t *testing.T) {
	queryValues := url.Values{}
	queryValues.Add("is_super_user", "true")
	queryValues.Add("user_id", "12")
	queryValues.Add("user_name", "ritwick")
	queryValues.Add("height", "1.0322") // random value :p

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

	if u1.Height != 1.0322 { // I know, this is not perface way to check float equal
		log.Panicln("Height should be 1.0322")
	}

}

type user struct {
	IsSuperUser   bool    `query:"is_super_user"`
	UserName      string  `query:"user_name"`
	UserID        int32   `query:"user_id"`
	Height        float32 `query:"height"`
	RandomField   string  `query:"random_field"`
	RandomPointer *int64  `query:"random_pointer"`
}
