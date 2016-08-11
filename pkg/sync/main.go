package main

import (
	"fmt"
	"sync"
	"time"
)

// 各ごルーチンが共有するパッケージ変数
var st struct{ A, B, C int }

// ミューテックスを保持するパッケージ変数
var mutex *sync.Mutex

func UpdateAndPrint(n int) {
	// ロック
	mutex.Lock()
	defer mutex.Unlock()
	// stの各フィールドをスリープをはさみながら更新
	st.A = n
	time.Sleep(time.Microsecond)
	st.B = n
	time.Sleep(time.Microsecond)
	st.C = n
	time.Sleep(time.Microsecond)
	fmt.Println(st.A, st.B, st.C)

}

func main() {
	mutex = new(sync.Mutex)
	for i := 0; i < 5; i++ {
		go func() {
			for i := 0; i < 1000; i++ {
				UpdateAndPrint(i)

			}
		}()

	}
	for {
	}

}
