package main
import ("fmt";"time";"math/rand")

type Token struct {
	data      string
	recipient int
}

func do() {
	tok := <- a
	fmt.Println(tok.recipient)
  	if tok.recipient-1 < 0 {
    		fmt.Println(tok.data)
        } else {
  		a <- Token{tok.data, tok.recipient-1}
  	}
}

var N int = 100
var a chan Token = make(chan Token)

func main() {
	for i := 1; i <= N; i++ {
    		go do()
	}
	a <- Token{"some text", 10}
	random:= time.Duration(rand.Intn(250))
	time.Sleep(time.Millisecond * random)
}