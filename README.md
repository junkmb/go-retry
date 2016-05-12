example
```
var i int
r := NewCountRetryer(5, time.Second)
for r.Loop() {
	fmt.Println(i)
	i++
}
```
