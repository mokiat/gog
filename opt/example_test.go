package opt_test

import (
	"fmt"

	"github.com/mokiat/gog/opt"
)

func ExampleT() {
	type Search struct {
		Country opt.T[string]
		Age     opt.T[int]
	}

	usersCache := make(map[Search][]string)
	usersCache[Search{
		Country: opt.V("Bulgaria"),
		Age:     opt.V(35),
	}] = []string{"John Doe", "Jane Doe"}

	users := usersCache[Search{
		Country: opt.V("Bulgaria"),
		Age:     opt.V(35),
	}]
	fmt.Printf("%#v\n", users)

	users = usersCache[Search{
		Country: opt.V("Bulgaria"),
		Age:     opt.Unspecified[int](),
	}]
	fmt.Printf("%#v\n", users)

	// Output:
	// []string{"John Doe", "Jane Doe"}
	// []string(nil)
}
