package io_test

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"

	io_ "github.com/searKing/golang/go/io"
)

func ExampleSniffReader() {
	r := strings.NewReader("MSG: some io.Reader stream to be read\n")
	sniff := io_.SniffReader(r)

	printSniff := func(r io.Reader, n int) {
		b := make([]byte, n)
		_, err := r.Read(b)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%s", b)
	}

	printall := func(r io.Reader) {
		b, err := ioutil.ReadAll(r)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%s", b)
	}

	// start sniffing
	sniff.Sniff(true)
	// sniff "MSG:\n"
	printSniff(sniff, len("MSG:"))
	fmt.Printf("\n")

	// stop sniffing
	sniff.Sniff(false)
	printall(sniff)

	// Output:
	// MSG:
	// MSG: some io.Reader stream to be read
}

func ExampleEOFReader() {
	r := io_.EOFReader()

	printall := func(r io.Reader) {
		b, err := ioutil.ReadAll(r)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%s", b)
	}

	printall(r)

	// Output:
	//
}

func ExampleWatchReader() {
	r := strings.NewReader("some io.Reader stream to be read\n")
	watch := io_.WatchReader(r, io_.WatcherFunc(func(p []byte, n int, err error) (int, error) {
		if err != nil && err != io.EOF {
			log.Fatal(err)
		}
		fmt.Printf("%s", p[:n])
		return n, err
	}))

	printall := func(r io.Reader) {
		b, err := ioutil.ReadAll(r)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%s", b)
	}

	printall(watch)

	// Output:
	// some io.Reader stream to be read
	// some io.Reader stream to be read
}

func ExampleLimitReadSeeker() {
	r := strings.NewReader("some io.Reader stream to be read\n")
	if _, err := io.Copy(os.Stdout, r); err != nil {
		log.Fatal(err)
	}

	limit := io_.LimitReadSeeker(r, int64(len("some io.Reader stream")))

	_, _ = limit.Seek(int64(len("some io.Reader ")), io.SeekStart)
	if _, err := io.Copy(os.Stdout, limit); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("\n")

	_, _ = limit.Seek(int64(-len("stream to be read\n")), io.SeekEnd)
	if _, err := io.Copy(os.Stdout, limit); err != nil {
		log.Fatal(err)
	}

	// Output:
	// some io.Reader stream to be read
	// stream
	// stream
}

func ExampleDynamicReadSeeker() {
	r := strings.NewReader("some io.Reader stream to be read\n")
	if _, err := io.Copy(os.Stdout, r); err != nil {
		log.Fatal(err)
	}

	ignoreOff := len("some ")

	// dynamic behaves like a reader for "io.Reader stream to be read\n"
	dynamic := io_.DynamicReadSeeker(func(off int64) (reader io.Reader, e error) {
		if off >= 0 {
			off += int64(ignoreOff)
		}
		_, err := r.Seek(off, io.SeekStart)
		// to omit r's io.Seeker
		return io.MultiReader(r), err
	}, r.Size()-int64(ignoreOff))

	_, _ = dynamic.Seek(int64(len("io.Reader ")), io.SeekStart)
	if _, err := io.Copy(os.Stdout, dynamic); err != nil {
		log.Fatal(err)
	}

	_, _ = dynamic.Seek(int64(-len("stream to be read\n")), io.SeekEnd)
	if _, err := io.Copy(os.Stdout, dynamic); err != nil {
		log.Fatal(err)
	}

	// Output:
	// some io.Reader stream to be read
	// stream to be read
	// stream to be read

}
