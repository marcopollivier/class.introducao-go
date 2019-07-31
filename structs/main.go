package main

import "fmt"

type palavra struct {
	prefixo, radical, sufixo string
}

func (p palavra) derivacao() string {
	return p.prefixo + p.radical + p.sufixo
}

func main() {

	p := palavra{radical: "noite", sufixo: "cer", prefixo: "a"}

	fmt.Println(p.derivacao())

}
