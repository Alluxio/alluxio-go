// +build ignore

package main

import (
	"bytes"
	"crypto/rand"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"sync"
	"time"

	alluxio "github.com/Alluxio/alluxio-go"
	"github.com/Alluxio/alluxio-go/option"
)

func metadataWorker(fs *alluxio.Client, path string) {
	for i := 0; i < numOperationsFlag; i++ {
		if _, err := fs.Exists(path, &option.Exists{}); err != nil {
			log.Fatal(err)
		}
	}
}

func metadataBenchmark() {
	clients, path := []*alluxio.Client{}, "/test_file"
	for i := 0; i < numThreadsFlag; i++ {
		clients = append(clients, alluxio.NewClient("localhost", 39999-i, time.Minute))
	}
	totalRPCs := numOperationsFlag * numThreadsFlag
	var totalTime int64
	data := make([]byte, bufferSizeFlag)
	if _, err := rand.Read(data); err != nil {
		log.Fatal(err)
	}
	if ok, err := clients[0].Exists(path, &option.Exists{}); err != nil {
		log.Fatal(err)
	} else if ok {
		if err := clients[0].Delete(path, &option.Delete{}); err != nil {
			log.Fatal(err)
		}
	}
	upload(clients[0], data, path)
	for j := 0; j < numIterationsFlag; j++ {
		start := time.Now()
		wg := sync.WaitGroup{}
		wg.Add(numThreadsFlag)
		for i := 0; i < numThreadsFlag; i++ {
			i := i
			go func() {
				defer wg.Done()
				metadataWorker(clients[i], path)
			}()
		}
		wg.Wait()
		end := time.Now()
		fmt.Printf("Iteration #%d: %d RPCs in %v.\n", j+1, totalRPCs, end.Sub(start))
		totalTime += end.Sub(start).Nanoseconds()
	}
	average := time.Duration(totalTime / int64(numIterationsFlag))
	fmt.Printf("Average time %v, average throughput %.2f RPC/s\n", average, float64(totalRPCs)/average.Seconds())
}

func readBenchmark() {
	clients, path := []*alluxio.Client{}, "/test_file"
	for i := 0; i < numThreadsFlag; i++ {
		clients = append(clients, alluxio.NewClient("localhost", 39999-i, time.Minute))
	}
	totalBytes := float64(numOperationsFlag*numThreadsFlag*fileSizeFlag) / float64(GB)
	var totalTime int64
	data := make([]byte, bufferSizeFlag)
	if _, err := rand.Read(data); err != nil {
		log.Fatal(err)
	}
	if ok, err := clients[0].Exists(path, &option.Exists{}); err != nil {
		log.Fatal(err)
	} else if ok {
		if err := clients[0].Delete(path, &option.Delete{}); err != nil {
			log.Fatal(err)
		}
	}
	upload(clients[0], data, path)
	for j := 0; j < numIterationsFlag; j++ {
		start := time.Now()
		wg := sync.WaitGroup{}
		wg.Add(numThreadsFlag)
		for i := 0; i < numThreadsFlag; i++ {
			i := i
			go func() {
				defer wg.Done()
				readWorker(clients[i], path)
			}()
		}
		wg.Wait()
		end := time.Now()
		fmt.Printf("Iteration #%d: read %.2fGBs in %v.\n", j+1, totalBytes, end.Sub(start))
		totalTime += end.Sub(start).Nanoseconds()
	}
	average := time.Duration(totalTime / int64(numIterationsFlag))
	fmt.Printf("Average time %v, average throughput %.2f GB/s\n", average, float64(totalBytes)/average.Seconds())
}

func readWorker(fs *alluxio.Client, path string) {
	for i := 0; i < numOperationsFlag; i++ {
		id, err := fs.OpenFile(path, &option.OpenFile{})
		if err != nil {
			log.Fatal(err)
		}
		read(fs, id)
	}
}

func read(fs *alluxio.Client, id int) {
	r, err := fs.Read(id)
	if err != nil {
		log.Fatal(err)
	}
	defer r.Close()
	defer fs.Close(id)
	if _, err := io.Copy(ioutil.Discard, r); err != nil {
		log.Fatal(err)
	}
}

func writeBenchmark() {
	fs, path := alluxio.NewClient("localhost", 39999, time.Minute), "/test-file"
	totalBytes := float64(numOperationsFlag*numThreadsFlag*fileSizeFlag) / float64(GB)
	var totalTime int64
	for j := 0; j < numIterationsFlag; j++ {
		for i := 0; i < numOperationsFlag; i++ {
			for k := 0; k < numThreadsFlag; k++ {
				if ok, err := fs.Exists(fmt.Sprintf("%s-%d-%d", path, i, k), &option.Exists{}); err != nil {
					log.Fatal(err)
				} else if ok {
					if err := fs.Delete(fmt.Sprintf("%s-%d-%d", path, i, k), &option.Delete{}); err != nil {
						log.Fatal(err)
					}
				}
			}
		}
		data := make([]byte, bufferSizeFlag)
		if _, err := rand.Read(data); err != nil {
			log.Fatal(err)
		}
		start := time.Now()
		wg := sync.WaitGroup{}
		wg.Add(numThreadsFlag)
		for i := 0; i < numThreadsFlag; i++ {
			i := i
			go func() {
				defer wg.Done()
				writeWorker(fs, data, path, i)
			}()
		}
		wg.Wait()
		end := time.Now()
		fmt.Printf("Iteration #%d: wrote %.2fGBs in %v.\n", j+1, totalBytes, end.Sub(start))
		totalTime += end.Sub(start).Nanoseconds()
	}
	average := time.Duration(totalTime / int64(numIterationsFlag))
	fmt.Printf("Average time %v, average throughput %.2f GB/s\n", average, float64(totalBytes)/average.Seconds())
}

func writeWorker(fs *alluxio.Client, data []byte, path string, id int) {
	for i := 0; i < numOperationsFlag; i++ {
		upload(fs, data, fmt.Sprintf("%s-%d-%d", path, i, id))
	}
}

func upload(fs *alluxio.Client, data []byte, path string) {
	bytesToWrite := fileSizeFlag
	id, err := fs.CreateFile(path, &option.CreateFile{})
	if err != nil {
		log.Fatal(err)
	}
	for bytesToWrite > 0 {
		n, err := fs.Write(id, bytes.NewBuffer(data))
		if err != nil {
			log.Fatal(err)
		}
		bytesToWrite -= n
	}
	fs.Close(id)
}

const (
	KB = 1024
	MB = KB * KB
	GB = KB * KB * KB
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	flag.IntVar(&benchmarkFlag, "benchmark", 0, "the benchmark to run")
	flag.IntVar(&fileSizeFlag, "file-size", MB<<8, "the size of the file to use")
	flag.IntVar(&bufferSizeFlag, "buffer-size", MB, "the size of the buffer to use")
	flag.IntVar(&numOperationsFlag, "num-operations", 1, "the numbers of operations to run")
	flag.IntVar(&numIterationsFlag, "num-iterations", 10, "the number of iterations to run")
	flag.IntVar(&numThreadsFlag, "num-threads", 1, "the number of threads to run")
	flag.Parse()
}

var (
	benchmarkFlag     int
	bufferSizeFlag    int
	fileSizeFlag      int
	numOperationsFlag int
	numIterationsFlag int
	numThreadsFlag    int
)

func main() {
	switch benchmarkFlag {
	case 0:
		readBenchmark()
	case 1:
		writeBenchmark()
	case 2:
		metadataBenchmark()
	}
}
