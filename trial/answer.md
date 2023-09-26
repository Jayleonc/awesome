# answer

[toc]

## 1 defer, panic and recover
### 1.1
```
2
1
0
```

### 1.2
```
1
```

### 1.3
```
1
```

### 1.4
```
2
```

### 1.5
```
3
```

### 1.6
```
2
```

### 1.7
```
1
```

### 1.8
```
in f3 111
110
```

### 1.9
```
b 1 2 3
d 0 2 2
c 0 2 2
a 1 3 4
```

### 1.10
```
No, because recover is only useful inside deferred functions. 

The following will be printed:

panic: 110

goroutine 1 [running]:
[stack trace omitted]
```

### 1.11
```
No, The following will be printed:

panic: 119

goroutine 1 [running]:
[stack trace omitted]
```

### 1.12
```
catch panic: 120
hello golang
```

### 1.13
```
Calling g.
Printing in g 0
Printing in g 1
Panicking!
Defer in g 1
Defer in g 0
Recovered in f 2
Returned normally from f.
```

### 1.14
```
Calling g.
Printing in g 0
Printing in g 1
Panicking!
Defer in g 1
Defer in g 0
panic: 2

goroutine 1 [running]:
[stack trace omitted]
```

## Reference
* [The Go Blog Defer Panic, and Recover](https://blog.golang.org/defer-panic-and-recover)
