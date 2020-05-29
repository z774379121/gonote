package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	// many more fields…
}

func jprint(format string, obj interface{}) {
	bytes, _ := json.Marshal(obj)
	fmt.Printf("%s:\n %s\n", format, string(bytes))
}

func main() {
	// 忽略字段
	t := &User{
		Email:    "email!",
		Password: "omit me!",
	}
	jprint("original", t)

	jprint("emit filed", struct {
		*User
		JustOmitMe bool `json:"password,omitempty"`
	}{User: t})

	jprint("additional field for now", struct {
		*User
		Name string `json:"name"`
	}{
		t,
		"name!",
	})
	type Value struct {
		string
	}
	type CacheItem struct {
		Key    string `json:"key"`
		MaxAge int    `json:"cacheAge"`
		Value  Value  `json:"cacheValue"`
	}

	c := &CacheItem{
		Key:    "22",
		MaxAge: 30,
		Value:  Value{"d"},
	}
	jprint("original", c)
	jprint("change name for now", struct {
		*CacheItem

		// Omit bad keys
		OmitMaxAge int `json:"cacheAge,omitempty"`
		OmitValue  int `json:"cacheValue,omitempty"`

		// Add nice keys
		MaxAge int    `json:"max_age"`
		Value  *Value `json:"value"`
	}{
		CacheItem: c,

		// Set the int by value:
		MaxAge: 10,

		// Set the nested struct by reference, avoid making a copy:
		Value: &Value{"f"},
	})
}
