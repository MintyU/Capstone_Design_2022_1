package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strconv"
	"sync"
)

func main() {
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
	port := ":9000"

	s, err := net.Listen("tcp", port)
	if err != nil {
		log.Println(err)
	}
	log.Println("Server is listing on port", port)
	defer s.Close()

	conn1, err := s.Accept()
	if err != nil {
		log.Println(err)
	}
	log.Println("TCP Connection #1 connected")

	conn2, err := s.Accept()
	if err != nil {
		log.Println(err)
	}
	log.Println("TCP Connection #2 connected")

	conn3, err := s.Accept()
	if err != nil {
		log.Println(err)
	}
	log.Println("TCP Connection #3 connected")

	conn4, err := s.Accept()
	if err != nil {
		log.Println(err)
	}
	log.Println("TCP Connection #4 connected")

	conn5, err := s.Accept()
	if err != nil {
		log.Println(err)
	}
	log.Println("TCP Connection #5 connected")

	conn6, err := s.Accept()
	if err != nil {
		log.Println(err)
	}
	log.Println("TCP Connection #6 connected")

	var wg sync.WaitGroup
	wg.Add(chunknum)

	go func() {
		buff1 := make([]byte, 1024)
		chunk1, err := os.Open("./chunk_0")
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
				log.Println("chunk_0 done")
				break
			}
			_, err = conn1.Write(buff1[:ccnt1])
		}
		wg.Done()
	}()

	go func() {
		buff2 := make([]byte, 1024)
		chunk2, err := os.Open("./chunk_1")
		if err != nil {
			log.Println(err)
		}
		defer chunk2.Close()
		for {
			ccnt2, err := chunk2.Read(buff2)

			if err != nil && err != io.EOF {
				log.Panic(err)
			}
			if ccnt2 == 0 {
				log.Println("chunk_1 done")
				break
			}
			_, err = conn2.Write(buff2[:ccnt2])
		}
		wg.Done()
	}()

	go func() {
		buff3 := make([]byte, 1024)
		chunk3, err := os.Open("./chunk_2")
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
				log.Println("chunk_2 done")
				break
			}
			_, err = conn3.Write(buff3[:ccnt3])
		}
		wg.Done()
	}()
	go func() {
		buff4 := make([]byte, 1024)
		chunk4, err := os.Open("./chunk_3")
		if err != nil {
			log.Println(err)
		}
		defer chunk4.Close()

		for {
			ccnt4, err := chunk4.Read(buff4)

			if err != nil && err != io.EOF {
				log.Panic(err)
			}
			if ccnt4 == 0 {
				log.Println("chunk_3 done")
				break
			}
			_, err = conn4.Write(buff4[:ccnt4])
		}
		wg.Done()
	}()
	go func() {
		buff5 := make([]byte, 1024)
		chunk5, err := os.Open("./chunk_4")
		if err != nil {
			log.Println(err)
		}
		defer chunk5.Close()

		for {
			ccnt5, err := chunk5.Read(buff5)

			if err != nil && err != io.EOF {
				log.Panic(err)
			}
			if ccnt5 == 0 {
				log.Println("chunk_4 done")
				break
			}
			_, err = conn5.Write(buff5[:ccnt5])
		}
		wg.Done()
	}()

	go func() {
		buff6 := make([]byte, 1024)
		chunk6, err := os.Open("./chunk_5")
		if err != nil {
			log.Println(err)
		}
		defer chunk6.Close()

		for {
			ccnt6, err := chunk6.Read(buff6)

			if err != nil && err != io.EOF {
				log.Panic(err)
			}
			if ccnt6 == 0 {
				log.Println("chunk_5 done")
				break
			}
			_, err = conn6.Write(buff6[:ccnt6])
		}
		wg.Done()
	}()
	wg.Wait()
	conn1.Close()
	conn2.Close()
	conn3.Close()
	conn4.Close()
	conn5.Close()
	conn6.Close()
}
