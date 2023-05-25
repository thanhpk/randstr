package randstr

import (
	"fmt"
)

func ExampleBase64() {
	for i := 0; i < 5; i++ {
		token := Base64(16)
		fmt.Println(token)
	}
	// Output:
	// jEivTuka++OZsN4w
	// /sbBr4mBF/lgnh3e
	// 0jhwQUFEUYw0Y8FG
	// CIy0D4/diZmD8WW8
	// 40zek8qIvfn26akN
}

func ExampleBase62() {
	for i := 0; i < 5; i++ {
		token := Base62(16)
		fmt.Println(token)
	}
	// Output:
	// DVosDCoMShcJ3G1X
	// uxF6JGne4gI5M74K
	// fed6PPJPApmrt5p4
	// JwQ1QUMuNxPkKqgN
	// DkiEUP32DqgBUwej
}

func ExampleHex() {
	for i := 0; i < 5; i++ {
		token := Hex(16)
		fmt.Println(token)
	}
	// Output:
	// e83f2b2af67d616a
	// deb3430f4e827df8
	// c1da2016675b8efe
	// f41b0d7600cb15fe
	// 685458e8c898545d
}

func ExampleString() {
	for i := 0; i < 5; i++ {
		token := String(16)
		fmt.Println(token)
	}
	// Output:
	// 7EbxkrHc1l3Ahmyr
	// I5XH2gc1EEHgbmGI
	// GlCycMpsxGkn9cDQ
	// U2OfBDQoak0z8FwV
	// kDX1m81u14YwEiCY
}

func ExampleDec() {
	for i := 0; i < 5; i++ {
		token := Dec(16)
		fmt.Println(token)
	}
	// Output:
	// 1232392418047380
	// 9160917876815937
	// 6629264107419930
	// 0271037110897873
	// 0337735480322223
}
