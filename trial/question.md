# question

[toc]

## 1 defer, panic and recover
### 1.1
```go
package main

import (
	"fmt"
)

func main() {
	f0()
}

func f0() {
	for i := 0; i < 3; i++ {
		defer fmt.Println(i)
	}
}
//2
//1
//0
```

### 1.2
```go
package main

import (
	"fmt"
)

func main() {
	f0()
}

func f0() {
	i := 1
	defer fmt.Println(i)
	i++
}

// 直觉告诉我是 2 XX
// +++++
// 1
```

### 1.3
```go
package main

import (
	"fmt"
)

func main() {
	f1()
}

func f1() {
	i := 1
	defer func(n int) {
		fmt.Println(n)
	}(i)
	i++
}

// 2
// +++
// 1

```

### 1.4
```go
package main

import (
	"fmt"
)

func main() {
	f2()
}

func f2() {
	i := 1
	defer func() {
		fmt.Println(i)
	}()
	i++
}

// 2
```

### 1.5
```go
package main

import (
	"fmt"
)

func main() {
	fmt.Println(f0())
}

func f0() (i int) {
	defer func() {
		i++
	}()
	return 2
}

// 感觉会报错吧
// +++
// 3
```

### 1.6
```go
package main

import (
	"fmt"
)

func main() {
	fmt.Println(f1())
}

func f1() (i int) {
	defer func(i int) {
		i++
	}(i)
	return 2
}

// 2

```

### 1.7
```go
package main

import (
	"fmt"
)

func main() {
	fmt.Println(f2())
}

func f2() (i int) {
	defer func() {
		i++
	}()
	return i
}

// 这个肯定报错
// +++
// 1
```

### 1.8
```go
package main

import (
	"fmt"
)

func main() {
	fmt.Println(f3())
}

func f3() int {
	i := 110
	defer func() {
		i++
		fmt.Println("in f3", i)
	}()
	return i
}

// 110
// in f3 111
// +++
// 反过来
```

### 1.9
```go
package main

import (
	"fmt"
)

func main() {
	a := 1
	b := 2
	defer calc("a", a, calc("b", a, b))
	a = 0
	defer calc("c", a, calc("d", a, b))
	b = 1
}

func calc(prefix string, a, b int) int {
	res := a + b
	fmt.Println(prefix, a, b, res)
	return res
}
// 不会，放弃这一题
```

### 1.10
* question: Will "hello golang" be printed when the code below is executed?
```go
package main

import (
	"fmt"
)

func main() {
	f()
	fmt.Println("hello golang")
}

func f() {
	if p := recover(); p != nil {
		fmt.Println("catch panic:", p)
	}
	panic("110")
}
// yes
// hello golang
// catch panic: 110
// +++
// NO
```

### 1.11
* question: Will "hello golang" be printed when the code below is executed?
```go
package main

import (
	"fmt"
)

func main() {
	f()
	fmt.Println("hello golang")
}

func f() {
	defer recover()
	panic("119")
}
// yes
// hello golang
// +++
// NO
```

### 1.12
```go
package main

import (
	"fmt"
)

func main() {
	f()
	fmt.Println("hello golang")
}

func f() {
	defer func() {
		if p := recover(); p != nil {
			fmt.Println("catch panic:", p)
		}
	}()
	panic("120")
}
// hello golang
// catch panic: 120
// +++
// 反过来
```

### 1.13
```go
package main

import (
	"fmt"
)

func main() {
	f()
	fmt.Println("Returned normally from f.")
}

func f() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
		}
	}()
	fmt.Println("Calling g.")
	g(0)
	fmt.Println("Returned normally from g.")
}

func g(i int) {
	if i > 1 {
		fmt.Println("Panicking!")
		panic(fmt.Sprintf("%v", i))
	}
	defer fmt.Println("Defer in g", i)
	fmt.Println("Printing in g", i)
	g(i + 1)
}


```

### 1.14
```go
package main

import (
	"fmt"
)

func main() {
	f()
	fmt.Println("Returned normally from f.")
}

func f() {
	fmt.Println("Calling g.")
	g(0)
	fmt.Println("Returned normally from g.")
}

func g(i int) {
	if i > 1 {
		fmt.Println("Panicking!")
		panic(fmt.Sprintf("%v", i))
	}
	defer fmt.Println("Defer in g", i)
	fmt.Println("Printing in g", i)
	g(i + 1)
}

// Calling g.
// Printing in g 0
// Printing in g 1
// Panicking!
// 2
// Defer in g 2
// Returned normally from g.
// Returned normally from f.
```

