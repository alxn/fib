package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
)

func fib(x int64) int64 {
	if x < 2 {
		return x
	}
	return fib(x-2) + fib(x-1)
}

func fibMemoImpl(x int64, memo []int64) int64 {
	if x < 2 {
		return x
	}
	if memo[x] != 0 {
		return memo[x]
	}
	memo[x] = fibMemoImpl(x-2, memo) + fibMemoImpl(x-1, memo)
	return memo[x]
}

func fibMemo(x int64) int64 {
	return fibMemoImpl(x, make([]int64, x+1))
}

func fibDynamic(x int64) int64 {
	tbl := []int64{0, 1}
	for ; x > 0; x-- {
		tbl[0] += tbl[1]
		tbl[0], tbl[1] = tbl[1], tbl[0]
	}
	return tbl[0]
}

func usage() {
	fmt.Println("calculate fib to <x>")
	flag.PrintDefaults()
	os.Exit(1)
}

func main() {
	var algo int
	flag.IntVar(&algo, "algo", 0,
		"Algorithm:\n\t0: recursive\n\t1: memo\n\t2: dynamic")
	flag.Parse()
	if flag.NArg() != 1 {
		usage()
	}

	x, err := strconv.ParseInt(flag.Arg(0), 10, 0)
	if err != nil {
		usage()
	}

	var f int64
	switch algo {
	case 0:
		f = fib(x)
		break
	case 1:
		f = fibMemo(x)
		break
	case 2:
		f = fibDynamic(x)
		break
	default:
		usage()
	}

	fmt.Printf("fib(%v) => %v\n", x, f)
}
