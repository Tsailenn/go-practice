package main

// GET
// import (
// 	"fmt"
// 	"io/ioutil"
// 	"log"
// 	"net/http"
// )

// func main() {

// 	resp, err := http.Get("http://webcode.me")

// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	defer resp.Body.Close()

// 	body, err := ioutil.ReadAll(resp.Body)

// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	fmt.Println(string(body))
// }

// POST with form
// import (
// 	"encoding/json"
// 	"fmt"
// 	"log"
// 	"net/http"
// 	"net/url"
// )

// func main() {

// 	data := url.Values{
// 		"name":       {"John Doe"},
// 		"occupation": {"gardener"},
// 	}

// 	resp, err := http.PostForm("https://httpbin.org/post", data)

// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	var res map[string]interface{}

// 	json.NewDecoder(resp.Body).Decode(&res)

// 	fmt.Println(res["form"])
// }

// POST with json

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {

	values := map[string]string{"name": "John Doe", "occupation": "gardener"}
	json_data, err := json.Marshal(values)

	if err != nil {
		log.Fatal(err)
	}

	resp, err := http.Post("https://httpbin.org/pos", "application/json",
		bytes.NewBuffer(json_data))

	if err != nil {
		log.Fatal(err)
	}

	var res map[string]interface{}

	json.NewDecoder(resp.Body).Decode(&res)

	fmt.Println(res["json"])
}
