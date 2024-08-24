package main

import (
	"log"
	"os"
)

func main() {
	log.SetOutput(os.Stdout)

	log.Print("Log\n")
	log.Println("Log2")
	log.Printf("Log%d\n", 3)

	f, err := os.Create("test.log")
	if err != nil {
		return
	}
	log.SetOutput(f)
	log.Println("ファイルに書き込む")

	log.SetFlags(log.LstdFlags)
	log.Println("A")

	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds)
	log.Println("B")

	log.SetFlags(log.Ltime | log.Lshortfile)
	log.Println("C")

	log.SetFlags(log.Ltime | log.Llongfile)
	log.Println("D")

	log.SetFlags(log.LstdFlags)

	log.SetPrefix("[LOG]")
	log.Println("E")

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lshortfile)
	logger.Fatalln("message")

	_, err = os.Open("fdafdsafa")
	if err != nil {
		log.Fatalln("Exit", err)
	}

}
