package main

import (
	"io"
	"log"
	"net"
	"os"
	"strconv"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	// Connections
	var addr = "127.0.0.1:9000"

	start := time.Now()
	c1, err := net.Dial("tcp", addr)
	if err != nil {
		log.Println(err)
	}
	c2, err := net.Dial("tcp", addr)
	if err != nil {
		log.Println(err)
	}
	c3, err := net.Dial("tcp", addr)
	if err != nil {
		log.Println(err)
	}
	c4, err := net.Dial("tcp", addr)
	if err != nil {
		log.Println(err)
	}
	c5, err := net.Dial("tcp", addr)
	if err != nil {
		log.Println(err)
	}
	c6, err := net.Dial("tcp", addr)
	if err != nil {
		log.Println(err)
	}
	elapsed := time.Since(start)
	log.Println(elapsed)
	chunknum := 6
	wg.Add(chunknum)
	go func() {
		buff1 := make([]byte, 1024)
		cw1, err := os.Create("./chunk_0")
		defer cw1.Close()
		if err != nil {
			log.Println(err)
		}
		for {
			cnt1, err := c1.Read(buff1)
			if err == io.EOF {
				break
			}
			_, err = cw1.Write(buff1[:cnt1])
		}
		wg.Done()
	}()

	go func() {
		buff2 := make([]byte, 1024)
		cw2, err := os.Create("./chunk_1")
		defer cw2.Close()
		if err != nil {
			log.Println(err)
		}
		for {
			cnt2, err := c2.Read(buff2)
			if err == io.EOF {
				break
			}
			_, err = cw2.Write(buff2[:cnt2])
		}
		wg.Done()
	}()

	go func() {
		buff3 := make([]byte, 1024)
		cw3, err := os.Create("./chunk_2")
		defer cw3.Close()
		if err != nil {
			log.Println(err)
		}
		for {
			cnt3, err := c3.Read(buff3)
			if err == io.EOF {
				break
			}
			_, err = cw3.Write(buff3[:cnt3])
		}
		wg.Done()
	}()

	go func() {
		buff4 := make([]byte, 1024)
		cw4, err := os.Create("./chunk_3")
		defer cw4.Close()
		if err != nil {
			log.Println(err)
		}
		for {
			cnt4, err := c4.Read(buff4)
			if err == io.EOF {
				break
			}
			_, err = cw4.Write(buff4[:cnt4])
		}
		wg.Done()
	}()

	go func() {
		buff5 := make([]byte, 1024)
		cw5, err := os.Create("./chunk_4")
		defer cw5.Close()
		if err != nil {
			log.Println(err)
		}
		for {
			cnt5, err := c5.Read(buff5)
			if err == io.EOF {
				break
			}
			_, err = cw5.Write(buff5[:cnt5])
		}
		wg.Done()
	}()

	go func() {
		buff6 := make([]byte, 1024)
		cw6, err := os.Create("./chunk_5")
		defer cw6.Close()
		if err != nil {
			log.Println(err)
		}
		for {
			cnt6, err := c6.Read(buff6)
			if err == io.EOF {
				break
			}
			_, err = cw6.Write(buff6[:cnt6])
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
