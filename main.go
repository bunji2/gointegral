// シンプソン法を使って定積分の近似値を計算するプログラム
// 積分対象となる関数や積分区間などは JavaScript で与える。

package main

import (
	"fmt"
	"os"
)

const (
	usageFmt = "%s f.js"
)

func main() {
	os.Exit(run())
}

func run() int {

	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, usageFmt, os.Args[0])
		return 1
	}

	err := runJS(os.Args[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	return 0
}
