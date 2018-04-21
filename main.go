package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	//標準入力 > 標準出力
	if _, err := io.Copy(os.Stdout, os.Stdin); err != nil {
		fmt.Printf("err=%v", err.Error())
		os.Exit(1)
	}

	// ファイル > 標準出力
	f, err := os.Open("main.go")
	if err != nil {
		fmt.Printf("err=%v", err.Error())
		os.Exit(1)
	}
	defer f.Close()

	if _, err := io.Copy(os.Stdout, f); err != nil {
		fmt.Printf("err=%v", err.Error())
		os.Exit(1)
	}

	//標準入力 > ファイル
	piyo, err := os.Create("piyo.txt")
	if err != nil {
		fmt.Printf("err=%v", err.Error())
		os.Exit(1)
	}
	defer piyo.Close()

	if _, err := io.Copy(piyo, os.Stdin); err != nil {
		fmt.Printf("err=%v", err.Error())
		os.Exit(1)
	}

	//メモリ > 標準出力
	if _, err := io.Copy(os.Stdout, bytes.NewReader([]byte("ほげええええええええ"))); err != nil {
		fmt.Printf("err=%v", err.Error())
		os.Exit(1)
	}

	//ネットワーク > 標準出力
	resp, err := http.Get("https://yusukemisa.github.io/my-app5/")
	if err != nil {
		fmt.Printf("err=%v", err.Error())
		os.Exit(1)
	}

	if _, err := io.Copy(os.Stdout, resp.Body); err != nil {
		fmt.Printf("err=%v", err.Error())
		os.Exit(1)
	}

}
