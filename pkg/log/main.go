package main

import "log"

func main() {

	log.SetFlags(log.LstdFlags)
	log.Println("A")

	//　時刻
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds)
	log.Println("B")

	// 時刻とファイルの行番号
	log.SetFlags(log.Ltime | log.Lshortfile)
	log.Println("C")

	//時刻とファイルの行番号 フルパス
	log.SetFlags(log.Ltime | log.Llongfile)
	log.Println("D")

	log.SetFlags(log.LstdFlags)

	//ログのプリフィックスを設定
	log.SetPrefix("[LOG] ")
	log.Println("E")
}
