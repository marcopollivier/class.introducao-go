package word

import (
	"fmt"
	"testing"
)

func TestPalindrome(t *testing.T) {
	if !isPalindrome("kayak") {
		t.Error(`IsPalindrome("kayak") = false`)
	}
}

func BenchmarkPalindrome(b *testing.B) {
	isPalindrome("kayak")
}

func BenchmarkHello(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fmt.Sprintf("hello")
	}
}

func ExampleHello() {
	fmt.Println("hello")
	// Output: hello
}

func ExampleSalutations() {
	fmt.Println("hello, and")
	fmt.Println("goodbye")
	// Output:
	// hello, and
	// goodbye
}

type Pessoa struct {
	nome string
}

func f() Pessoa {
	return Pessoa{nome: "Marco"}
}

func ExampleName() {
	f()
	// Output:
	// { Marco }
}
