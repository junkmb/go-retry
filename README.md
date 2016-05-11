```
var i int
r := NewRetryer(time.Second, 5)
for r.Loop() {
	fmt.Println(i)
	i++
}
```
