package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

type arrayFlags []string

func (i *arrayFlags) String() string {
	return "my string representation"
}

func (i *arrayFlags) Set(value string) error {
	*i = append(*i, strings.TrimSpace(value))
	return nil
}

func findPrimes(ctx context.Context, s string, wg *sync.WaitGroup, id int, ch chan string) {
	defer wg.Done()
	var a, b int
	var err error
	arr := strings.Split(s, ":")
	if a, err = strconv.Atoi(arr[0]); err != nil {
		log.Fatal(err)
	}
	if b, err = strconv.Atoi(arr[1]); err != nil {
		log.Fatal(err)
	}
	out := make([]int, 0)
	for i := a; i <= b; i++ {
		select {
		case <-ctx.Done():
			fmt.Println("Timeout for job ", id)
			ch <- fmt.Sprint("range ", s, " ", out)
			return
		default:
			isPrime := true
			for j := 2; j < i; j++ {
				if i%j == 0 {
					isPrime = false
				}
			}
			if isPrime {
				out = append(out, i)
			}
		}
	}
	fmt.Printf("Job %d done\n", id)
	ch <- fmt.Sprint("range ", s, " ", out)
}

func writePrimesToFile(file *os.File, ch chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	s := <-ch
	if _, err := file.WriteString(s + "\n"); err != nil {
		log.Fatal(err)
	}
}

func main() {
	var ranges arrayFlags
	ch := make(chan string)
	flag.Var(&ranges, "range", "ranges of primes")
	timeout := flag.Int("timeout", 2, "integer value in seconds")
	filename := flag.String("file", "output.txt", "file of primes")
	file, _ := os.Create(*filename)
	defer file.Close()
	flag.Parse()
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(*timeout)*time.Second)
	defer cancel()
	wg := sync.WaitGroup{}
	for i, s := range ranges {
		s := s
		wg.Add(2)
		go findPrimes(ctx, s, &wg, i, ch)
		go writePrimesToFile(file, ch, &wg)
	}
	wg.Wait()
	fmt.Println("Job done")
}
