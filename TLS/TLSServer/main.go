package main

import (
	"crypto/rand"
	"crypto/tls"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"sync"
)

func main() {
	fileName := "./example.mov"
	f, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	fileInfo, _ := f.Stat()
	var fileSize int64 = fileInfo.Size()
	fmt.Println("File Size :", fileSize, "bytes")

	tempBuff := make([]byte, 1024)

	chunknum := 6
	var buffnum int64

	if (fileSize/int64(chunknum))%1024 == 0 {
		buffnum = fileSize / int64(chunknum) / 1024
	} else {
		buffnum = (fileSize / int64(chunknum) / 1024) + 1
	}

	fmt.Println("number of buffers to send : ", buffnum)
	for i := 0; i < chunknum; i++ {
		chunkwrite, err := os.Create("chunk_" + strconv.FormatInt(int64(i), 10))
		if err != nil {
			log.Println(err)
		}
		for j := 0; j < int(buffnum); j++ {
			tempcnt, err := f.Read(tempBuff)
			if err != nil && err != io.EOF {
				panic(err)
			}

			if tempcnt == 0 {
				break
			}
			_, err = chunkwrite.Write(tempBuff[:tempcnt])
			if err != nil {
				panic(err)
			}

		}
	}

	// Connections
	cert, err := tls.LoadX509KeyPair("certs/server.pem", "certs/server.key")
	if err != nil {
		log.Fatalf("server: loadkeys: %s", err)
	}
	config := tls.Config{Certificates: []tls.Certificate{cert}}
	config.Rand = rand.Reader
	service := "0.0.0.0:23023"
	listener, err := tls.Listen("tcp", service, &config)
	if err != nil {
		log.Fatalf("server: listen: %s", err)
	}
	log.Print("server: listening")
	defer listener.Close()

	conn0, err := listener.Accept()
	if err != nil {
		log.Printf("server: accept: %s", err)
	}
	log.Println("conn0")
	buff := make([]byte, 1024)
	cnt, err := conn0.Read(buff) // After the TCP+TLS connection, data must be moved to enable multiple connections.
	log.Println(string(buff[:cnt]))

	conn1, err := listener.Accept()
	if err != nil {
		log.Printf("server: accept: %s", err)
	}
	log.Println("conn1")
	cnt, err = conn1.Read(buff)
	log.Println(string(buff[:cnt]))

	conn2, err := listener.Accept()
	if err != nil {
		log.Printf("server: accept: %s", err)
	}
	cnt, err = conn2.Read(buff)
	log.Println(string(buff[:cnt]))

	conn3, err := listener.Accept()
	if err != nil {
		log.Printf("server: accept: %s", err)
	}
	cnt, err = conn3.Read(buff)
	log.Println(string(buff[:cnt]))

	conn4, err := listener.Accept()
	if err != nil {
		log.Printf("server: accept: %s", err)
	}
	cnt, err = conn4.Read(buff)
	log.Println(string(buff[:cnt]))

	conn5, err := listener.Accept()
	if err != nil {
		log.Printf("server: accept: %s", err)
	}
	cnt, err = conn5.Read(buff)
	log.Println(string(buff[:cnt]))

	var wg sync.WaitGroup
	wg.Add(chunknum)

	go func() {
		buff0 := make([]byte, 1024)
		chunk, err := os.Open("./chunk_0")
		if err != nil {
			log.Println(err)
		}
		defer chunk.Close()
		for {
			ccnt0, err := chunk.Read(buff0)

			if err != nil && err != io.EOF {
				log.Panic(err)
			}
			if ccnt0 == 0 {
				log.Println("chunk_0 done")
				break
			}
			_, err = conn0.Write(buff0[:ccnt0])
		}
		wg.Done()
	}()
	go func() {
		buff0 := make([]byte, 1024)
		chunk, err := os.Open("./chunk_1")
		if err != nil {
			log.Println(err)
		}
		defer chunk.Close()
		for {
			ccnt0, err := chunk.Read(buff0)

			if err != nil && err != io.EOF {
				log.Panic(err)
			}
			if ccnt0 == 0 {
				log.Println("chunk_1 done")
				break
			}
			_, err = conn1.Write(buff0[:ccnt0])
		}
		wg.Done()
	}()
	go func() {
		buff0 := make([]byte, 1024)
		chunk, err := os.Open("./chunk_2")
		if err != nil {
			log.Println(err)
		}
		defer chunk.Close()
		for {
			ccnt0, err := chunk.Read(buff0)

			if err != nil && err != io.EOF {
				log.Panic(err)
			}
			if ccnt0 == 0 {
				log.Println("chunk_2 done")
				break
			}
			_, err = conn2.Write(buff0[:ccnt0])
		}
		wg.Done()
	}()
	go func() {
		buff0 := make([]byte, 1024)
		chunk, err := os.Open("./chunk_3")
		if err != nil {
			log.Println(err)
		}
		defer chunk.Close()
		for {
			ccnt0, err := chunk.Read(buff0)

			if err != nil && err != io.EOF {
				log.Panic(err)
			}
			if ccnt0 == 0 {
				log.Println("chunk_3 done")
				break
			}
			_, err = conn3.Write(buff0[:ccnt0])
		}
		wg.Done()
	}()
	go func() {
		buff0 := make([]byte, 1024)
		chunk, err := os.Open("./chunk_4")
		if err != nil {
			log.Println(err)
		}
		defer chunk.Close()
		for {
			ccnt0, err := chunk.Read(buff0)

			if err != nil && err != io.EOF {
				log.Panic(err)
			}
			if ccnt0 == 0 {
				log.Println("chunk_4 done")
				break
			}
			_, err = conn4.Write(buff0[:ccnt0])
		}
		wg.Done()
	}()
	go func() {
		buff0 := make([]byte, 1024)
		chunk, err := os.Open("./chunk_5")
		if err != nil {
			log.Println(err)
		}
		defer chunk.Close()
		for {
			ccnt0, err := chunk.Read(buff0)

			if err != nil && err != io.EOF {
				log.Panic(err)
			}
			if ccnt0 == 0 {
				log.Println("chunk_5 done")
				break
			}
			_, err = conn5.Write(buff0[:ccnt0])
		}
		wg.Done()
	}()
	wg.Wait()

}
