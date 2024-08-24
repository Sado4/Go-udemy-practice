package main

import (
	"fmt"
	"math/rand"
	"time"
)

//4.rand
//疑似乱数を生成する昨日がまとめられた機能。
//ランタイム全体で共有される疑似乱数生成と、任意のシード値を元にした疑似乱数生成などができる。
//最も単純な使い方は、rand.Seedへ任意のint64の数値を渡してデフォルトの疑似乱数生成器のシードを設定する方法
//rand.Float64はデフォルトの疑似乱数生成器を使って、0.0<=n<1.0　の条件を満たす疑似乱数を生成する。

func main() {
	//デフォルトの疑似乱数生成器のシードを設定
	//シードに設定された数値が256に固定されているため、何度実行しても同じ内容の擬似乱数が生成される。
	rand.Seed(256)

	fmt.Println(rand.Float64())
	fmt.Println(rand.Float64())
	fmt.Println(rand.Float64())
	/*
		0.813527291469711
		0.5598026045235738
		0.6695717783859498
	*/

	//現在の時刻をシードに使った疑似乱数の生成
	//プログラムの中で毎回異なった疑似乱数生成器のシードを設定するには、現在の時刻を利用する方法が最も手軽。
	//timeパッケージを利用して、現在時刻をもとにしたシード値を設定する。
	//例
	//1970/1/1からの累積ナノ秒をシードに設定
	fmt.Println(time.Now().UnixNano())
	rand.Seed(time.Now().UnixNano())
	//0~99の間の疑似乱数
	fmt.Println(rand.Intn(100))
	//int型の疑似乱数
	fmt.Println(rand.Int())
	//int32型の疑似乱数
	fmt.Println(rand.Int31())
	//int64型の疑似乱数
	fmt.Println(rand.Int63())
	//uint32型の疑似乱数
	fmt.Println(rand.Uint32())
	//特にrand.Intn(n)を使って0以上でnより小さい乱数を生成するパターンが使われる。

	//擬似乱数生成器の生成
	//rand.Seedやrand.IntnはGoのランタイム情に用意されたデフォルトの擬似乱数生成器を共有している。
	//簡易的に使用する分には問題ないが、プログラムの意図しない場所で擬似乱数生成器を書き換えられてしまう危険性がある。
	//解決するには、疑似乱数生成期を明示的に生成して管理する必要がある。
	//rand.NewSourceにシード値を与えて、rand.Source型を生成し、
	//rand.Newから独立した疑似乱数生成器を生成することができる。
	//rand.Rand型にはrand.Intnなどの関数と同名のメソッドが定義されているため、デフォルト擬似乱数生成器を利用した場合を同様に操作する。

	//疑似乱数生成期のソースを現在時刻から生成
	src := rand.NewSource(time.Now().UnixNano())
	//ソースを元に擬似乱数生成器を生成
	rnd := rand.New(src)
	//0~99の疑似乱数
	fmt.Println(rnd.Intn(100))
}
