package main

import (
    //"bytes"
    //"encoding/json"
    "fmt"
//    "io/ioutil"
//	"net/http"
	"crypto/md5"
	//"math/rand"
	//"time"
	//"strconv"
)
/*
const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func RandStringBytes(n int) string {
    b := make([]byte, n)
    for i := range b {
        b[i] = letterBytes[rand.Intn(len(letterBytes))]
    }
    return string(b)
}
*/
func main() {
	
	// TODO: mettre des strings?

	// http://gateway.marvel.com/v1/public/comics?ts=10&apikey=acc1c4efc8246d5b7ff81d2b9b2e00ce&hash=900246b409580e3ac228b27e13977a2b

	// API variables 
	publicKey := []byte("acc1c4efc8246d5b7ff81d2b9b2e00ce")
	privateKey := []byte("67909a94a20b1231835d360b55d26c16634ee256")
	ts := []byte("10")
	//ts := []byte(strconv.Itoa(time.Now().Unix()))

	fmt.Println("Starting the application...")

	fmt.Printf("ts: %s \n", ts)
	

	//fmt.Printf("%s \n", append( append(ts, privateKey...), publicKey...) )

	//data := []byte("These pretzels are making me thirsty.")
	//fmt.Printf("%x \n", md5.Sum( append( append(ts, privateKey...), publicKey...) ))

	hash := md5.Sum( append( append(ts, privateKey...), publicKey...) )
	fmt.Printf("%x \n", hash)


	// GET request
	/*
	response, err := http.Get("https://httpbin.org/ip")
	
	if err != nil {
        fmt.Printf("The HTTP request failed with error %s\n", err)
    } else {
        data, _ := ioutil.ReadAll(response.Body)
        fmt.Println(string(data))
    }
	*/

	/*
	jsonData := map[string]string{"firstname": "Nic", "lastname": "Raboy"}
    jsonValue, _ := json.Marshal(jsonData)
	
	response, err = http.Post("https://httpbin.org/post", "application/json", bytes.NewBuffer(jsonValue))
	
	if err != nil {
        fmt.Printf("The HTTP request failed with error %s\n", err)
    } else {
        data, _ := ioutil.ReadAll(response.Body)
        fmt.Println(string(data))
    }
	*/	
	fmt.Println("Terminating the application...")
}