import "fmt"

func fibonacci() func() int {
	suite := 0
	suite2 := 1
	return func () int {
		res := suite2 + suite
		suite = suite2
		suite2 = res
		return res
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
