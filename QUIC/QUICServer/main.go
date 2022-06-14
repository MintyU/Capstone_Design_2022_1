package main

import (
	"context"
	"crypto/rand"
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io"
	"log"
	"math/big"
	"os"
	"strconv"
	"sync"

	"github.com/lucas-clemente/quic-go"
)

const addr = "localhost:9000"

func main() {
	var wg sync.WaitGroup

	// Open file and split into chunks
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
	listener, err := quic.ListenAddr(addr, generateTLSConfig(), nil)
	if err != nil {
		log.Println(err, "Listener")
	}
	conn, err := listener.Accept(context.Background())
	if err != nil {
		log.Println(err, "Connection")
	}
	checkstream, err := conn.AcceptStream(context.Background())
	if err != nil {
		log.Println(err, "checkstream")
	}
	stream0, err := conn.AcceptStream(context.Background())
	if err != nil {
		log.Println(err, "stream0")
	}
	stream1, err := conn.AcceptStream(context.Background())
	if err != nil {
		log.Println(err, "stream1")
	}
	stream2, err := conn.AcceptStream(context.Background())
	if err != nil {
		log.Println(err, "stream2")
	}
	stream3, err := conn.AcceptStream(context.Background())
	if err != nil {
		log.Println(err, "stream3")
	}
	stream4, err := conn.AcceptStream(context.Background())
	if err != nil {
		log.Println(err, "stream4")
	}
	stream5, err := conn.AcceptStream(context.Background())
	if err != nil {
		log.Println(err, "stream5")
	}

	wg.Add(chunknum + 1)

	go func() {
		// checkstream prevents the server from closing before the file transfer is finished
		buff := make([]byte, 1024)
		checkstream.Read(buff)
		checkstream.Close()
		wg.Done()
	}()

	// File split transfer using goroutine
	go func() {
		buff0 := make([]byte, 1024)
		chunk0, err := os.Open("./chunk_0")
		if err != nil {
			log.Println(err)
		}
		defer chunk0.Close()

		for {
			ccnt0, err := chunk0.Read(buff0)

			if err != nil && err != io.EOF {
				log.Panic(err)
			}
			if ccnt0 == 0 {
				log.Println("chunk 0 done")
				break
			}
			_, err = stream0.Write(buff0[:ccnt0])
		}
		stream0.Close()
		wg.Done()
	}()
	go func() {

		buff1 := make([]byte, 1024)
		chunk1, err := os.Open("./chunk_1")
		if err != nil {
			log.Println(err)
		}
		defer chunk1.Close()

		for {
			ccnt1, err := chunk1.Read(buff1)

			if err != nil && err != io.EOF {
				log.Panic(err)
			}
			if ccnt1 == 0 {
				log.Println("chunk 1 done")

				break
			}
			_, err = stream1.Write(buff1[:ccnt1])
		}
		stream1.Close()
		wg.Done()
	}()
	go func() {

		buff1 := make([]byte, 1024)
		chunk1, err := os.Open("./chunk_2")
		if err != nil {
			log.Println(err)
		}
		defer chunk1.Close()

		for {
			ccnt1, err := chunk1.Read(buff1)

			if err != nil && err != io.EOF {
				log.Panic(err)
			}
			if ccnt1 == 0 {
				log.Println("chunk 2 done")

				break
			}
			_, err = stream2.Write(buff1[:ccnt1])
		}
		stream2.Close()
		wg.Done()
	}()
	go func() {

		buff3 := make([]byte, 1024)
		chunk3, err := os.Open("./chunk_3")
		if err != nil {
			log.Println(err)
		}
		defer chunk3.Close()

		for {
			ccnt3, err := chunk3.Read(buff3)

			if err != nil && err != io.EOF {
				log.Panic(err)
			}
			if ccnt3 == 0 {
				log.Println("chunk 3 done")

				break
			}
			_, err = stream3.Write(buff3[:ccnt3])
		}
		stream3.Close()
		wg.Done()
	}()
	go func() {

		buff3 := make([]byte, 1024)
		chunk3, err := os.Open("./chunk_4")
		if err != nil {
			log.Println(err)
		}
		defer chunk3.Close()

		for {
			ccnt3, err := chunk3.Read(buff3)

			if err != nil && err != io.EOF {
				log.Panic(err)
			}
			if ccnt3 == 0 {
				log.Println("chunk 4 done")

				break
			}
			_, err = stream4.Write(buff3[:ccnt3])
		}
		stream4.Close()
		wg.Done()
	}()
	go func() {

		buff3 := make([]byte, 1024)
		chunk3, err := os.Open("./chunk_5")
		if err != nil {
			log.Println(err)
		}
		defer chunk3.Close()

		for {
			ccnt3, err := chunk3.Read(buff3)

			if err != nil && err != io.EOF {
				log.Panic(err)
			}
			if ccnt3 == 0 {
				log.Println("chunk 5 done")

				break
			}
			_, err = stream5.Write(buff3[:ccnt3])
		}
		stream5.Close()
		wg.Done()
	}()
	wg.Wait()

}

func generateTLSConfig() *tls.Config {
	key, err := rsa.GenerateKey(rand.Reader, 1024)
	if err != nil {
		panic(err)
	}
	template := x509.Certificate{SerialNumber: big.NewInt(1)}
	certDER, err := x509.CreateCertificate(rand.Reader, &template, &template, &key.PublicKey, key)
	if err != nil {
		panic(err)
	}
	keyPEM := pem.EncodeToMemory(&pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(key)})
	certPEM := pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: certDER})

	tlsCert, err := tls.X509KeyPair(certPEM, keyPEM)
	if err != nil {
		panic(err)
	}
	return &tls.Config{
		Certificates: []tls.Certificate{tlsCert},
		NextProtos:   []string{"quic-echo-example"},
	}
}
