example
```
var i int
r := NewCountRetryer(time.Second, 5)
for r.Loop() {
	fmt.Println(i)
	i++
}
```
