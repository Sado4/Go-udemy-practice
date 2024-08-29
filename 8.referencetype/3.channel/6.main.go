package main

import (
	"fmt"
	"time"
)

// チャネル&ゴルーチンを使ったジョブキューの例

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Printf("Worker %d started job %d\n", id, j)
		time.Sleep(time.Second) // ジョブの処理をシミュレート
		results <- j * 2
		fmt.Printf("Worker %d finished job %d\n", id, j)
	}
}

func main() {
	jobs := make(chan int, 5)
	results := make(chan int, 5)

	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	close(jobs) // これ以上ジョブが送られてこないことを知らせる

	for a := 1; a <= 5; a++ {
		fmt.Println(<-results) // チャネルから結果を受信
	}
}
