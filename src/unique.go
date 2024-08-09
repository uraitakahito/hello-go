//go:build ignore

// マップを利用して重複排除処理を実装する
package main

import "fmt"

func main() {
	followers := []string{"John", "Richard", "John", "Jane", "Jane", "Alan"}

	// 全てユニークであることを考慮して、対象のスライスの長さ分の容量を確保する
	// 繰り返し処理が実行される中で、スライスの容量拡張を抑えることができる
	unique := make([]string, 0, len(followers))

	// 存在有無チェック機構としてマップを利用
	// 値を空の構造体にすることで、メモリのアロケーションをゼロにできる
	m := make(map[string]struct{})
	for _, v := range followers {
		if _, ok := m[v]; ok {
			continue
		}
		unique = append(unique, v)
		m[v] = struct{}{}
	}

	fmt.Println(unique)
	// Output:[John Richard Jane Alan]
}
