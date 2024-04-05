# Learn Go with Pepper

I am currently writing a hamcrest inspired matcher library for Go [called Pepper](https://github.com/quii/pepper). 

I figured a good way to test it, would be to dogfood it, by working through my own course [learn go with tests](https://github.com/quii/learn-go-with-tests). 

It has already helped me iron out some usability issues, so im confident this will prove the worth of the library. 

## It already feels pretty good compared to writing "raw" go tests

```go
func TestHello(t *testing.T) {
	Expect(t, Hello("", "English")).To(Equal(`Hello, world`))
	Expect(t, Hello("Chris", "English")).To(Equal(`Hello, Chris`))
	Expect(t, Hello("Elodie", "Spanish")).To(Equal(`Hola, Elodie`))
	Expect(t, Hello("Milo", "French")).To(Equal(`Bonjour, Milo`))
}
```
