package main

import (
	"context"
	"crypto/tls"
	"io"
	"log"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/lucas-clemente/quic-go"
)

const addr = "localhost:9000"
const message = "foobar"

func main() {
	var wg sync.WaitGroup
	start := time.Now()
	tlsConf := &tls.Config{
		InsecureSkipVerify: true,
		NextProtos:         []string{"quic-echo-example"},
	}
	chunknum := 6
	conn, err := quic.DialAddr(addr, tlsConf, nil)
	if err != nil {
		log.Println(err)
	}

	// Stream Connections
	checkstream, err := conn.OpenStreamSync(context.Background())
	if err != nil {
		log.Println(err)
	}

	// A message must travel once through the stream to confirm the connection.
	_, err = checkstream.Write([]byte(message))
	if err != nil {
		log.Println(err)
	}

	stream0, err := conn.OpenStreamSync(context.Background())
	if err != nil {
		log.Println(err)
	}
	_, err = stream0.Write([]byte(message))
	if err != nil {
		log.Println(err)
	}

	stream1, err := conn.OpenStreamSync(context.Background())
	if err != nil {
		log.Println(err)
	}
	_, err = stream1.Write([]byte(message))
	if err != nil {
		log.Println(err)
	}

	stream2, err := conn.OpenStreamSync(context.Background())
	if err != nil {
		log.Println(err)
	}
	_, err = stream2.Write([]byte(message))
	if err != nil {
		log.Println(err)
	}
	stream3, err := conn.OpenStreamSync(context.Background())
	if err != nil {
		log.Println(err)
	}
	_, err = stream3.Write([]byte(message))
	if err != nil {
		log.Println(err)
	}
	stream4, err := conn.OpenStreamSync(context.Background())
	if err != nil {
		log.Println(err)
	}
	_, err = stream4.Write([]byte(message))
	if err != nil {
		log.Println(err)
	}
	stream5, err := conn.OpenStreamSync(context.Background())
	if err != nil {
		log.Println(err)
	}
	_, err = stream5.Write([]byte(message))
	if err != nil {
		log.Println(err)
	}
	wg.Add(chunknum)
	elapsed := time.Since(start)
	log.Println(elapsed)
	go func() {

		writeBuff0 := make([]byte, 1024)
		fw0, err := os.Create("./chunk_0")
		if err != nil {
			log.Println(err)
		}
		defer fw0.Close()
		defer wg.Done()

		for {
			wcnt0, err := stream0.Read(writeBuff0)
			if err != nil && err != io.EOF {
				log.Println("0: ", err)
				break
			}

			if wcnt0 == 0 {

				break
			}
			_, err = fw0.Write(writeBuff0[:wcnt0])
			if err != nil {
				log.Println(err)
				break
			}
		}
	}()
	go func() {

		writeBuff1 := make([]byte, 1024)
		fw1, err := os.Create("./chunk_1")
		if err != nil {
			log.Println(err)
		}
		defer fw1.Close()
		defer wg.Done()

		for {
			wcnt1, err := stream1.Read(writeBuff1)
			if err != nil && err != io.EOF {
				log.Println("1: ", err)
				break
			}

			if wcnt1 == 0 {

				break
			}

			_, err = fw1.Write(writeBuff1[:wcnt1])
			if err != nil {
				log.Println(err)
				break
			}

		}

	}()
	go func() {

		writeBuff2 := make([]byte, 1024)
		fw2, err := os.Create("./chunk_2")
		if err != nil {
			log.Println(err)
		}
		defer fw2.Close()
		defer wg.Done()

		for {
			wcnt2, err := stream2.Read(writeBuff2)
			if err != nil && err != io.EOF {
				log.Println("2: ", err)
				break
			}

			if wcnt2 == 0 {

				break
			}

			_, err = fw2.Write(writeBuff2[:wcnt2])
			if err != nil {
				log.Println(err)
				break
			}

		}

	}()
	go func() {

		writeBuff2 := make([]byte, 1024)
		fw2, err := os.Create("./chunk_3")
		if err != nil {
			log.Println(err)
		}
		defer fw2.Close()
		defer wg.Done()

		for {
			wcnt2, err := stream3.Read(writeBuff2)
			if err != nil && err != io.EOF {
				log.Println("3: ", err)
				break
			}

			if wcnt2 == 0 {

				break
			}

			_, err = fw2.Write(writeBuff2[:wcnt2])
			if err != nil {
				log.Println(err)
				break
			}

		}

	}()
	go func() {

		writeBuff2 := make([]byte, 1024)
		fw2, err := os.Create("./chunk_4")
		if err != nil {
			log.Println(err)
		}
		defer fw2.Close()
		defer wg.Done()

		for {
			wcnt2, err := stream4.Read(writeBuff2)
			if err != nil && err != io.EOF {
				log.Println("4: ", err)
				break
			}

			if wcnt2 == 0 {

				break
			}

			_, err = fw2.Write(writeBuff2[:wcnt2])
			if err != nil {
				log.Println(err)
				break
			}

		}

	}()
	go func() {

		writeBuff2 := make([]byte, 1024)
		fw2, err := os.Create("./chunk_5")
		if err != nil {
			log.Println(err)
		}
		defer fw2.Close()
		defer wg.Done()

		for {
			wcnt2, err := stream5.Read(writeBuff2)
			if err != nil && err != io.EOF {
				log.Println("5: ", err)
				break
			}

			if wcnt2 == 0 {

				break
			}

			_, err = fw2.Write(writeBuff2[:wcnt2])
			if err != nil {
				log.Println(err)
				break
			}

		}

	}()
	wg.Wait()
	// When the transfer is complete, it closes the connection by sending a message to ckeckstream.
	checkstream.Write([]byte("ACK"))
	merge, err := os.Create("./example.mov")
	if err != nil {
		log.Println(err)
	}
	tempBuff := make([]byte, 1024)
	for i := 0; i < chunknum; i++ {
		f, err := os.Open("./chunk_" + strconv.FormatInt(int64(i), 10))
		if err != nil {
			log.Println(err)
		}
		for {
			cnt, err := f.Read(tempBuff)
			if err != nil {
				// log.Println(err)
				break
			}
			_, err = merge.Write(tempBuff[:cnt])

		}
	}
	elapsed = time.Since(start)
	log.Println(elapsed)
}
