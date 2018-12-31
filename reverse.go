package main

// 1. https://godoc.org/github.com/pa-m/sklearn/cluster
// 2. https://www.cheatography.com/chewxy/cheat-sheets/data-science-in-go-a/

import (
	"bytes"
	"flag"
	"io"
	"io/ioutil"
	"log"
	"net"
	"sync"
	"time"

	"github.com/komuw/dbscan/heart"
	"github.com/pkg/errors"
)

/*
usage:
  1.
  go run -race reverse.go -p localhost:7777 -r localhost:3333
  curl -vkL localhost:7777
  echo -n "test out the server" | nc localhost 7777

  2.
  go run -race reverse.go -p localhost:7777 -r httpbin.org:443
  curl -vkL \
	-H "accept: application/json" \
	-H "Content-Type: application/json" \
	-H "Host: httpbin.org" \
	-d '{"name":"komu"}' \
	localhost:7777/post
*/

// TODO: do the same for responses

type myStruct struct {
	l                   sync.RWMutex
	noOfAllRequests     int
	allRequests         []float64
	lengthOfEachRequest int
}

var thisStruct myStruct

func handleRequest(requestBuf []byte) {
	thisStruct.l.Lock()
	defer thisStruct.l.Unlock()
	thisStruct.lengthOfEachRequest = len(requestBuf)

	for _, v := range requestBuf {
		thisStruct.allRequests = append(thisStruct.allRequests, float64(v))
	}
}

const nulByte = "\x00"

func forward(reverseProxyConn net.Conn, remoteAddr string) {
	defer reverseProxyConn.Close()
	err := reverseProxyConn.SetDeadline(time.Now().Add(5 * time.Second))
	if err != nil {
		err = errors.Wrap(err, "Reverse Unable to set reverseProxyConn deadline")
		log.Fatalf("\n%+v", err)
	}

	// TODO: make the buffer growable
	// TODO: use ioutil.ReadAll() for this
	requestBuf := make([]byte, 512)
	reqLen, err := reverseProxyConn.Read(requestBuf)
	if err != nil {
		err = errors.Wrap(err, "Reverse Error reading")
		log.Fatalf("\n%+v", err)
	}
	_ = reqLen
	requestBuf = bytes.Trim(requestBuf, nulByte)
	log.Println("Reverse read::", requestBuf)
	log.Println("Reverse read2::", string(requestBuf))

	handleRequest(requestBuf)

	// TODO: since we also want to dbscan the responses, we should make a copy here also.
	backendConn, err := net.Dial("tcp", remoteAddr)
	if err != nil {
		err = errors.Wrap(err, "Reverse Dial failed for address"+remoteAddr)
		log.Fatalf("\n%+v", err)
	}
	defer backendConn.Close()
	err = backendConn.SetDeadline(time.Now().Add(5 * time.Second))
	if err != nil {
		err = errors.Wrap(err, "Reverse Unable to set backendConn deadline")
		log.Fatalf("\n%+v", err)
	}
	log.Printf("reverseProxyConnected to localhost %v\n", reverseProxyConn)

	var backendBuf bytes.Buffer
	backendTee := io.TeeReader(backendConn, &backendBuf)
	io.Copy(backendConn, bytes.NewReader(requestBuf))
	io.Copy(reverseProxyConn, backendTee)

	backendBytes, err := ioutil.ReadAll(&backendBuf)
	if err != nil {
		err = errors.Wrap(err, "Reverse Unable to read backendBuf")
		log.Fatalf("\n%+v", err)
	}
	log.Println("backendBytes:::", backendBytes)
	log.Println("backendBytes2:::", string(backendBytes))

	thisStruct.l.Lock()
	thisStruct.noOfAllRequests++
	log.Println("allRequests:", thisStruct.allRequests)
	log.Println("lengthOfEachRequest:", thisStruct.lengthOfEachRequest)
	thisStruct.l.Unlock()

	// mat.NewDense(noOfAllRequests, lengthOfEachRequest, allRequests)

}

func cooler() {
	thisStruct.l.Lock()
	defer thisStruct.l.Unlock()
	heart.Run(thisStruct.noOfAllRequests, thisStruct.lengthOfEachRequest, thisStruct.allRequests, 3.0, 1.0, false)
}

func main() {
	var p string
	var r string
	flag.StringVar(
		&p,
		"p",
		"localhost:7777",
		"the IP:PORT pair to bind the proxy to.")
	flag.StringVar(
		&r,
		"r",
		"localhost:3333",
		"the IP:PORT pair to proxy connections to ie the remote server.")
	flag.Parse()

	// {
	// 	go proxyd.Run(r)
	// }

	listener, err := net.Listen("tcp", p)
	if err != nil {
		err = errors.Wrapf(err, "Reverse failed to setup listener %v", p)
		log.Fatalf("\n%+v", err)
	}
	log.Println("Reverse Listening on " + p)

	{
		time.AfterFunc(23*time.Second, cooler)
	}

	for {
		reverseProxyConn, err := listener.Accept()
		if err != nil {
			err = errors.Wrapf(err, "Reverse failed to accept listener %v", p)
			log.Fatalf("\n%+v", err)
		}
		log.Printf("Accepted reverseProxyConnection %v\n", reverseProxyConn)
		go forward(reverseProxyConn, r)
	}

}