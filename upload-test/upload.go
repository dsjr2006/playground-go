package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	pb "gopkg.in/cheggaaa/pb.v1"
)

func main() {
	path := "/Users/dsjr2006/Downloads/debitOverdraftServiceDisclosure.htm.pdf"
	//path := "/Users/dsjr2006/Downloads/ed-109pg_16088_7.pdf"
	//path := "/Users/dsjr2006/Downloads/Postman-osx-4.9.3.zip"
	file, e := os.Open(path)
	if e != nil {
		log.Fatal("File Error")
	}
	defer file.Close()
	fi, e := file.Stat()
	if e != nil {
		log.Fatal("File Stat Error")
	}
	body := bytes.NewBuffer(nil)
	if _, err := io.Copy(body, file); err != nil {
		log.Fatal("Could not create part upload buffer")
	}

	bar := pb.New(int(fi.Size()))
	prBody := bar.NewProxyReader(body)

	bar.Start()

	client := &http.Client{}
	req, e := http.NewRequest("POST", "https://posttestserver.com/post.php", prBody)
	//req, e := http.NewRequest("POST", "http://httptest.dlsmi.com:8088/upload/file", prBody)
	req.Header.Add("Content-Length", string(fi.Size()))
	if e != nil {
		log.Fatal("Request Error")
	}
	resp, e := client.Do(req)
	if e != nil {
		log.Fatalf("Response Error: %v", e)
	}
	bar.Finish()
	respBody, e := ioutil.ReadAll(resp.Body)
	if e != nil {
		log.Fatal("Response Body Read Error")
	}
	fmt.Println(string(respBody))

	return
}
