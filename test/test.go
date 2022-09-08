package main

func main() {

		aString :="Hello"

		aStringChangeable := []byte(aString)

		aStringChangeable[0] = 'A'

		aStringFinalized := string(aStringChangeable)

		fmt.Println(aString)

		fmt.Println(aStringChangeable)

		fmt.Println(aStringFinalized)
}
