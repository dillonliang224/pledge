package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime/pprof"
	"sync"
)

func counter() {
	slice := [100000]int{0}
	c := 1
	for i := 0; i < 100000; i++ {
		c = i + 15
		fmt.Println(c)
		slice[i] = c
	}
}

func workOnce(wg *sync.WaitGroup) {
	counter()
	wg.Done()
}

func main() {
	var cpuProfile = flag.String("cpuprofile", "", "write cpu profile to file")
	var memProfile = flag.String("memprofile", "", "write mem profile to file")
	flag.Parse()

	if *cpuProfile != "" {
		f, err := os.Create(*cpuProfile)
		if err != nil {
			log.Fatal(err)
		}

		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	var wg sync.WaitGroup
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go workOnce(&wg)
	}

	wg.Wait()

	if *memProfile != "" {
		f, err := os.Create(*memProfile)
		if err != nil {
			log.Fatal(err)
		}

		pprof.WriteHeapProfile(f)
		f.Close()
	}
}
