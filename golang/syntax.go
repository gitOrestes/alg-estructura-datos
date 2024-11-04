package main

import (
	"fmt"	
)

func main() {
	var w Writer 
}

type Writer interface {
	Write([]byte) (int, error)
}

type ConsoleWrite struct {}

func (cw ConsoleWrite) Write(data [] byte) (int, error){
	 n, err := fmt.Printnl(string(data))
	 return n, err
}

/*type myStruck struct {
	foo int
}

func main () {

var a int = 42
var b *int = &a
fmt.Println(a, *b)

var ms *myStruck
ms = new(myStruck)
(*ms).foo = 52
fmt.Println((*ms).foo)

}
*/

/*
//for doble

func main() {
	for i, j := 0, 0; i < 5; i, j = i+1, j+1 {
		fmt.Println(j, i)
	}

	s := []int{1, 2, 3}
	for k, v := range s {
		fmt.Println(k, v)
	}
}*/

/*func main() 
//switch type
{
	var i interface{} = 1.0
	switch i.(type) {
	case int:
		fmt.Println("int")
	case float64:
		fmt.Println("float64")
	case string:
		fmt.Println("string")
	default:
		fmt.Println("is another type")
	}
}*/

/*import "reflect"

type Animal struct {
	Name 	string `required max:100`
	Origin 	string
}

func main() {
	t := reflect.TypeOf(Animal{})
	field, _ := t.FieldByName("Name")
	fmt.Println(field.Tag)
}
*/
//maps
/*func main() {

	statePopulations := make(map[string]int)

	statePopulations = map[string]int{
		"California":	39250017,
		"Texas":		27862596,
		"Florida":		20612439,
		"New York":		19745289,
		"Pennsylvania":	12801503,
		"Illinois":		12801539,
		"Ohio":			11614373,
	}

	_, ok := statePopulations["Ohio"]
	fmt.Println(ok)

	sp := statePopulations
	delete(sp, "Texas")

	if _, ok := statePopulations["Ohio"]; ok {
		fmt.Println("oli")
	}

	delete(statePopulations, "Ohio")
	statePopulations["Georgia"] = 	10310371
	fmt.Println(statePopulations["Texas"])
}*/


/*func main(){
	a := []int{1,2,3,4,5,6,7,8,9,10}
	b := a[:]
	c := a[3:]
	d := a[:7]
	e := a[2:8]
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
}*/


/*


//enums

const (
	isAdmin = 1 << iota
	isHeadquarters
	canSeeFinancials

	canSeeAfrica
	canSeeAsia
	canSeeEurope
	canSeeNorthAmerica
	canSeeSouthAmerica
)

func main() {
	var roles byte = isHeadquarters |canSeeFinancials | canSeeEurope | canSeeNorthAmerica
	fmt.Printf("%b\n", roles)
	fmt.Printf("Is admin? %v", isAdmin & roles == isAdmin)
}*/

/*

//pila con slices

type Stack []string

func (s *Stack) isEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) push(str string) {
	*s = append(*s, str)
}

func (s *Stack) pop() (string, bool) {
	if s.isEmpty() {
		return "", false
	} else {
		index := len(*s) - 1
		element := (*s)[index]
		*s = (*s)[:index]
		return element, true
	}
}

func main() {

	var stack Stack

	stack.push("estai")
	stack.push("como")
	stack.push("oli")

	for !stack.isEmpty() {
		x, y := stack.pop()

		fmt.Println(x, y)
	}

}
*/