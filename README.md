usort
---
a simpler golang slice sort library based on Go 1.18+ Generics and Go sort library.

使用 Go 1.18+ 范型和 sort 包实现的自定义切片排序库。

### Quick Start

```go
s1 := []int{1,4,3,2}

usort.USort(s1, func(i, j int) bool {
    return s1[i] > s1[j]
})
```

### Complex type
```go
type People struct {
    Name string
    Age  int
}

tom := People{"tom", 31}
libi := People{"libi", 30}
jerry := People{"jerry", 25}

s2 := []People{jerry, tom, libi}

usort.USort(s2, func(i, j int) bool {
    return s2[i].Age > s2[j].Age
})

fmt.Println(s2)
// output : [{tom 31} {libi 30} {jerry 25}]

```
