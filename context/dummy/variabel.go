package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"sync"
	"time"
)

const totalFile = 3000
const contentLength = 5000

var tempPath = filepath.Join(os.Getenv("TEMP"), "chapter-A.61-pipeline-cancellation-context")

type FileInfo struct {
	Index       int
	FileName    string
	WorkerIndex int
	Err         error
}

func main() {
	log.Println("start")
	start := time.Now()

	generateFiles()

	duration := time.Since(start)
	log.Println("done in", duration.Seconds(), "seconds")
}

func randomString(lenght int) string {
	randomizer := rand.New(rand.NewSource(time.Now().Unix()))
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	b := make([]rune, lenght)
	for i := range b {
		s := randomizer.Intn(len(letters))
		b[i] = letters[s]
	}
	return string(b)
}

func generateFiles() {
	os.RemoveAll(tempPath)
	os.MkdirAll(tempPath, os.ModePerm)

	// pipeline 1: job distribution
	chanFileIndex := generateFileIndexes()

	// pipeline 2: the main logic (creating files)
	createFileWorker := 100
	chanFileResult := createFiles(chanFileIndex, createFileWorker)

	// tarck and print output
	counterTotal := 0
	counterSucces := 0
	for fileResult := range chanFileResult {
		if fileResult.Err != nil {
			log.Printf("error creating file %s. stack trace: %s", fileResult.FileName, fileResult.Err)
		} else {
			counterSucces++
		}
		counterTotal++
	}
	log.Printf("%d/%d of total files created", counterSucces, counterTotal)
}

func generateFileIndexes() <-chan FileInfo {
	chanOut := make(chan FileInfo)

	go func() {
		for i := 0; i < totalFile; i++ {
			chanOut <- FileInfo{
				Index:    i,
				FileName: fmt.Sprintf("file-%d.txt", i),
			}
		}
		close(chanOut)
	}()
	return chanOut
}

func createFiles(chanIn <-chan FileInfo, numberOfWorkers int) <-chan FileInfo {
	chanOut := make(chan FileInfo)

	wg := new(sync.WaitGroup)
	wg.Add(numberOfWorkers)

	go func() {
		for WorkerIndex := 0; WorkerIndex < numberOfWorkers; WorkerIndex++ {
			go func(WorkerIndex int) {
				for job := range chanIn {
					filePath := filepath.Join(tempPath, job.FileName)
					content := randomString(contentLength)
					err := ioutil.WriteFile(filePath, []byte(content), os.ModePerm)

					log.Println("worker", WorkerIndex, "working on", job.FileName, "file generation")

					chanOut <- FileInfo{
						FileName:    job.FileName,
						WorkerIndex: WorkerIndex,
						Err:         err,
					}
				}

				wg.Done()
			}(WorkerIndex)
		}
	}()

	go func() {
		wg.Wait()
		close(chanOut)
	}()
	return chanOut
}
