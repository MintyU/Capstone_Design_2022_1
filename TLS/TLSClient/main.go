package main

import (
	"crypto/tls"
	"io"
	"log"
	"os"
	"strconv"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	start := time.Now()
	cert, err := tls.LoadX509KeyPair("certs/client.pem", "certs/client.key")
	if err != nil {
		log.Fatalf("server: loadkeys: %s", err)
	}
	config := tls.Config{Certificates: []tls.Certificate{cert}, InsecureSkipVerify: true}

	// Connections
	conn0, err := tls.Dial("tcp", "127.0.0.1:23023", &config)
	if err != nil {
		log.Fatalf("client0: dial: %s", err)
	}
	defer conn0.Close()
	// After the TCP+TLS connection, data must be moved to enable multiple connections.
	conn0.Write([]byte("connect"))

	conn1, err := tls.Dial("tcp", "127.0.0.1:23023", &config)
	if err != nil {
		log.Fatalf("client0: dial: %s", err)
	}
	defer conn1.Close()
	conn1.Write([]byte("connect"))

	conn2, err := tls.Dial("tcp", "127.0.0.1:23023", &config)
	if err != nil {
		log.Fatalf("client0: dial: %s", err)
	}
	defer conn2.Close()
	conn2.Write([]byte("connect"))

	conn3, err := tls.Dial("tcp", "127.0.0.1:23023", &config)
	if err != nil {
		log.Fatalf("client0: dial: %s", err)
	}
	defer conn3.Close()
	conn3.Write([]byte("connect"))

	conn4, err := tls.Dial("tcp", "127.0.0.1:23023", &config)
	if err != nil {
		log.Fatalf("client0: dial: %s", err)
	}
	defer conn4.Close()
	conn4.Write([]byte("connect"))

	conn5, err := tls.Dial("tcp", "127.0.0.1:23023", &config)
	if err != nil {
		log.Fatalf("client0: dial: %s", err)
	}
	defer conn5.Close()
	conn5.Write([]byte("connect"))

	log.Println("client: connected to: ", conn0.RemoteAddr())
	elapsed := time.Since(start)
	log.Println(elapsed)
	chunknum := 6
	wg.Add(chunknum)
	go func() {
		buff0 := make([]byte, 1024)
		cw0, err := os.Create("./chunk_0")
		defer cw0.Close()

		if err != nil {
			log.Println(err)
		}
		for {
			n0, err := conn0.Read(buff0)
			if err == io.EOF {
				break
			}
			_, err = cw0.Write(buff0[:n0])
		}
		wg.Done()
	}()
	go func() {
		buff1 := make([]byte, 1024)
		cw1, err := os.Create("./chunk_1")
		defer cw1.Close()

		if err != nil {
			log.Println(err)
		}
		for {
			n1, err := conn1.Read(buff1)
			if err == io.EOF {
				break
			}
			_, err = cw1.Write(buff1[:n1])
		}
		wg.Done()
	}()
	go func() {
		buff1 := make([]byte, 1024)
		cw1, err := os.Create("./chunk_2")
		defer cw1.Close()

		if err != nil {
			log.Println(err)
		}
		for {
			n1, err := conn2.Read(buff1)
			if err == io.EOF {
				break
			}
			_, err = cw1.Write(buff1[:n1])
		}
		wg.Done()
	}()
	go func() {
		buff1 := make([]byte, 1024)
		cw1, err := os.Create("./chunk_3")
		defer cw1.Close()

		if err != nil {
			log.Println(err)
		}
		for {
			n1, err := conn3.Read(buff1)
			if err == io.EOF {
				break
			}
			_, err = cw1.Write(buff1[:n1])
		}
		wg.Done()
	}()
	go func() {
		buff1 := make([]byte, 1024)
		cw1, err := os.Create("./chunk_4")
		defer cw1.Close()

		if err != nil {
			log.Println(err)
		}
		for {
			n1, err := conn4.Read(buff1)
			if err == io.EOF {
				break
			}
			_, err = cw1.Write(buff1[:n1])
		}
		wg.Done()
	}()
	go func() {
		buff1 := make([]byte, 1024)
		cw1, err := os.Create("./chunk_5")
		defer cw1.Close()

		if err != nil {
			log.Println(err)
		}
		for {
			n1, err := conn5.Read(buff1)
			if err == io.EOF {
				break
			}
			_, err = cw1.Write(buff1[:n1])
		}
		wg.Done()
	}()
	wg.Wait()
	elapsed = time.Since(start)
	log.Println(elapsed)

	merge, err := os.Create("./example.mov")
	tempBuff := make([]byte, 1024)
	for i := 0; i < chunknum; i++ {
		f, err := os.Open("./chunk_" + strconv.FormatInt(int64(i), 10))
		if err != nil {
			log.Println(err)
		}
		for {
			cnt, err := f.Read(tempBuff)

			if err == io.EOF {
				break
			}
			_, err = merge.Write(tempBuff[:cnt])

		}

	}

	elapsed = time.Since(start)
	log.Println(elapsed)
}
