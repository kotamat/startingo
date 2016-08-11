package main

import (
	"fmt"
	"log"
	"os"
)

func fileopen() {
	f, err := os.Open("foo.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// os.File
	bs := make([]byte, 128)
	n, err := f.Read(bs)
	fmt.Println(n, bs)

	// ステータス
	fi, err := f.Stat()
	fmt.Println(
		fi.Name(),
		fi.IsDir(),
		fi.Size(),
		fi.Mode(),
		fi.ModTime(),
	)
}

func filecreate() {
	f, _ := os.Create("bar.txt")

	//書き込み
	f.Write([]byte("Hello, World;\n")) // []byte型で書き込み
	f.WriteAt([]byte("Golang"), 7)     //オフセットを指定する
	f.Seek(0, os.SEEK_END)             //ファイルの末尾にオフセットを移動
	f.WriteString("Yeah!")             // 文字列をファイルに書き込む

}

func fileopenstrict() {
	f, err := os.OpenFile("foo.txt", os.O_RDONLY, 0666)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()
}

func main() {
	fileopen()

	filecreate()

	fileopenstrict()
}
