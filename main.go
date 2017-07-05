package main

import (
	"bufio"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"fmt"
	"math/big"
	"net"
	"os"
	"strconv"
)

const (
	tcpProtocol     = "tcp4"      //Using protocol
	keySize         = 1024        //Length of rsa-key
	readWriteerSize = keySize / 8 //Max length of key (byte)
)

type remoteConn struct {
	c    *net.TCPConn
	pubK *rsa.PublicKey
} // "c" - connection, "pubK" - key of connection

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func getRemoteConn(c *net.TCPConn) *remoteConn {
	return &remoteConn{c: c, pubK: waitPubKey(bufio.NewReader(c))}
}

func waitPubKey(buf *bufio.Reader) *rsa.PublicKey {
	line, _, err := buf.ReadLine()
	checkErr(err)
	if string(line) == "CONNECT" {
		pubKey := rsa.PublicKey{N: big.NewInt(0)}
		pubKey.N.SetString(string(line), 10)
		line, _, err = buf.ReadLine()
		checkErr(err)
		pubKey.E, err = strconv.Atoi(string(line))
		checkErr(err)
		return &pubKey
	} else {
		fmt.Println("Error: inknow command ", string(line))
		os.Exit(1)
	}
	return nil
}

func (rConn *remoteConn) sendCommand(comm string) {
	eComm, err := rsa.EncryptOAEP(sha1.New(), rand.Reader, rConn.pubK, []byte(comm), nil)
	checkErr(err)
	rConn.c.Write(eComm)
}

func listen() {
	l, err := net.ListenTCP(tcpProtocol, listenAddr)
	checkErr(err)
	fmt.Println("Listen port: ", l.Addr().(*net.TCPAddr).Port)
	c, err := l.AcceptTCP()
	checkErr(err)
	fmt.Println("Connect from:", c.RemoteAddr())
	rConn := getRemoteConn(c)
	rConn.sendCommand("Hello!")
}

var listenAddr = &net.TCPAddr{IP: net.IPv4(127, 0, 0, 1), Port: 0}

func main() {
	listen()
}
