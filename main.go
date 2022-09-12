package main

import "fmt"

// 型の成約をインターフェースとして宣言する。
// これを使用することでジェネリクス関数をもっと簡素にできるみたい。
// Number型として作成する
type Number interface {
	int64 | float64
}

func main() {
	// 整数値のmapのサンプルデータ
	ints := map[string]int64{
		"first":  34,
		"second": 12,
	}

	// floatのmapのサンプルデータ
	floats := map[string]float64{
		"first":  35.98,
		"second": 26.99,
	}

	// 出力する。ジェネリクス関数不使用の方
	fmt.Printf(
		"ジェネリクスを使用していない関数での計算: %v and %v\n",
		SumInts(ints),
		SumFloats(floats),
	)

	// ジェネリクス関数（あえて型指定してる）を使用して計算結果を出力
	fmt.Printf(
		"ジェネリクス関数を使用しての計算： %v and %v\n",
		SumIntsOrFloats[string, int64](ints),
		SumIntsOrFloats[string, float64](floats),
	)

	// ジェネリクス関数（こっちは型指定なし。）で計算結果を出力
	// これが出来るのは、ジェネリクス関数が引数をもってるから。
	// 引数を持ってないジェネリクス関数は型推論ができないみたい
	fmt.Printf(
		"型を指定せず、型推論でジェネリクス関数を使用しての計算： %v and %v\n",
		SumIntsOrFloats(ints),
		SumIntsOrFloats(floats),
	)

	// インターフェースのNumber型を使ったジェネリクス関数を使用しての計算を出力
	fmt.Printf(
		"Number型を使ったジェネリクス関数での計算： %v and %v\n",
		SumNumbers(ints),
		SumNumbers(floats),
	)
}

// ジェネリクスを使用しないでの関数------------------------------------------------------
// map(ハッシュ＝連想配列)のm の値（キー：string　値：int64）を足していく。ただし型がint64
func SumInts(m map[string]int64) int64 {
	var s int64
	for _, v := range m {
		s += v
	}
	return s
}

// mapm の値を追加する。ただし型がfloat64
func SumFloats(m map[string]float64) float64 {
	var s float64
	for _, v := range m {
		s += v
	}
	return s
}

// ジェネリクス不使用　ここまで----------------------------------------------------------

// これはジェネリクス関数
func SumIntsOrFloats[K comparable, V int64 | float64](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}

// インターフェースのNumber型を使用したジェネリクス関数
func SumNumbers[K comparable, V Number](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}
