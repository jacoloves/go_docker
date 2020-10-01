/***
	fgo ap計算お知らせツール
***/
package main

import (
	"fmt"
	"time"
)

type Hourset struct {
	min int
	sec int
}

func display(hs *Hourset, ap *int) {
	hs.sec++

	if hs.sec == 60 {
		hs.min++
		hs.sec = 0
	}

	if hs.min == 5 {
		*ap++
		hs.min = 0
	}

	// if *ap == sc {
	// 	return true
	// }

	defer fmt.Printf("\rAPが %v 回復しました。", *ap)

}

func main() {

	fmt.Println("AP測定開始")

	hs := &Hourset{}
	ap := 0
	// flg := false

	// fmt.Print("スタミナはどれだけ回復しますか？：")
	// scanner := bufio.NewScanner(os.Stdin)
	// scanner.Scan()
	// text := scanner.Text()

	// sc := strconv.Atoi(text)

	for {
		display(hs, &ap)
		time.Sleep(time.Second)

		// if flg == true {
		// 	break
		// }
	}

	fmt.Println("AP測定終了")
}
