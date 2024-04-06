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

## Custom matchers for the win!

The idea of Pepper is you can use it as a basic assertion library, much like other libraries out there, to do simple assertions like the ones in Hello World, but it is also extensible, allowing you to write your own "Matchers" for your domain. This helps you write tests with clearer test output, more consistently, and over the longer run, cheaply. 

See how nicely this works for the shapes testing

```go
func TestShapes(t *testing.T) {
	Expect[Shape](t, Rectangle{10.0, 10.0}).To(
		HavePerimeter(40.0).And(HaveArea(100.0)),
	)
	Expect[Shape](t, EquilateralTriangle{Width: 10.0}).To(
		HaveArea(43.3012701892).And(HavePerimeter(30.0)),
	)
}
```

We can declaratively test properties of shapes with custom matchers, and leveraging Pepper's composition tools with `And`. 

Here are the matchers

```go
func HaveArea(want float64) Matcher[Shape] {
	return func(got Shape) MatchResult {
		return MatchResult{
			Description: fmt.Sprintf("to have area of %.5f", want),
			Matches:     almostEqual(got.Area(), want),
			But:         fmt.Sprintf("it had area of %.5f", got.Area()),
			SubjectName: "Shape",
		}
	}
}

func HavePerimeter(want float64) Matcher[Shape] {
	return func(got Shape) MatchResult {
		return MatchResult{
			Description: fmt.Sprintf("to have perimeter of %.5f", want),
			Matches:     got.Perimeter() == want,
			But:         fmt.Sprintf("it had perimeter of %.5f", got.Perimeter()),
			SubjectName: "Shape",
		}
	}
}
```

If we were to add more shapes, we can re-use the matchers to test them. 