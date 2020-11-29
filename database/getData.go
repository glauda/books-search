package main

import (
    //"bytes"
    //"encoding/json"
    "fmt"
    "io/ioutil"
	"net/http"
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
)

func GenRandomBytes(size int) (blk []byte, err error) {
    blk = make([]byte, size)
	_, err = rand.Read(blk)
    return
}

/*	// generate random timestamp
ts, err := GenRandomBytes(10)

if err != nil {
	panic(err)
} else {
	sblk := string(blk)
	fmt.Printf("sblk: %s\n", sblk)
	fmt.Printf("hex : %x\n", blk)
}
*/

func CreateAPIRequest(path string, option string) (request string){
	
	fmt.Printf("Creating request for %s \n", path)

	// API variables 
	publicKey := []byte("acc1c4efc8246d5b7ff81d2b9b2e00ce")
	privateKey := []byte("67909a94a20b1231835d360b55d26c16634ee256")
	ts := []byte("10")

	hash := md5.Sum( append( append(ts, privateKey...), publicKey...) )

	domainName := "https://gateway.marvel.com:443" 

	// log messages
	fmt.Printf("ts=%s \n", ts)
	fmt.Printf("apikey=%s \n", publicKey)
	fmt.Printf("hash=%x \n", hash)

	request = domainName + path + "?" +
		option + 
		"&ts=" + string(ts) +
		"&apikey=" + string(publicKey) +
		"&hash=" + string(hex.EncodeToString(hash[:]))

	return 
}


func main() {
	
	// http://gateway.marvel.com/v1/public/comics?ts=10&apikey=acc1c4efc8246d5b7ff81d2b9b2e00ce&hash=900246b409580e3ac228b27e13977a2b

	fmt.Println("Starting the application...")

	
	// GET request
	
	path, option := "/v1/public/series", "limit=50"
	//fmt.Printf(CreateAPIRequest(path, option))
	
	response, err := http.Get(CreateAPIRequest(path, option))
	
	if err != nil {
        fmt.Printf("The HTTP request failed with error %s\n", err)
    } else {
		data, _ := ioutil.ReadAll(response.Body)
		_ = ioutil.WriteFile("./data/resGetSeries.json", data, 0644)
		//fmt.Println(string(data))
		
    }
	
	fmt.Println("Terminating the application...")
}