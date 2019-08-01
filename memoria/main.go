package main

func main() {
	f()
}

//go:noinline
func f() *int {
	variavelComNomeGrandeParaChamarAtencao := 10
	return &variavelComNomeGrandeParaChamarAtencao
}
