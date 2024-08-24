package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	_, err := os.Open("foo")
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(os.Args[0])
	fmt.Println(os.Args[1])
	fmt.Println(os.Args[2])
	fmt.Println(os.Args[3])

	fmt.Printf("length=%d\n", len(os.Args))
	for i, v := range os.Args {
		fmt.Println(i, v)
	}
	f, err := os.Open("test.log")
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	bs := make([]byte, 128)
	n, err := f.Read(bs)
	fmt.Println(n)
	fmt.Println(string(bs))

	nn, err := f.ReadAt(bs, 10)
	fmt.Println(nn)
	fmt.Println(string(bs))

	//offset, err := f.Seek(10, os.SEEK_SET)
	//offset, err = f.Seek(-2, os.SEEK_CUR)
	//offset, err = f.Seek(0, os.SEEK_END)

	fi, err := f.Stat()
	fmt.Println(fi.Name())
	fi.Size()
	fi.Mode()
	fi.ModTime()
	fi.IsDir()
	f, _ = os.Create("foo.txt")
	f.Write([]byte("Hello\n"))
	f.WriteAt([]byte("Golang"), 7)
	f.Seek(0, os.SEEK_END)
	f.WriteString("Yaah")
	f, err = os.OpenFile("foo.txt", os.O_RDONLY, 0666)

	//err := os.Remove("foo.txt")
	//err := os.Remove("foo")

	//err := os.RemoveAll("foo")

	err = os.Rename("foo.txt", "bar.txt")
	err = os.Rename("bar.txt", "foo.txt")

	//err = os.Rename("foo", "bar")

	//err = os.Rename("foo.txt", "bar/foo.txt")

	dir, err := os.Getwd()
	fmt.Println(dir)

	//err = os.Chdir("path/to/dir")

	err = os.Mkdir("test", 0775)
	//err = os.MkdirAll("foo/fooo/foooo", 0775)

	f, err = os.Open(".")
	if err != nil {
		log.Println(err)
	}
	defer f.Close()

	fis, err := f.Readdir(0)
	for _, fi := range fis {
		if fi.IsDir() {
			fmt.Println(fi)
		}
	}

	host, err := os.Hostname()
	fmt.Println(host)

	for _, v := range os.Environ() {
		fmt.Println(v)
	}

	fmt.Println(os.Getpid())
	os.Getppid()
	os.Getuid()
	os.Geteuid()
	os.Getgid()
	os.Getegid()
}
