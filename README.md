# Learn Go with Pepper

I am currently writing a Hamcrest inspired matcher library for Go [called Pepper](https://github.com/quii/pepper). 

I figured a good way to test it, would be to [dogfood](https://en.wikipedia.org/wiki/Eating_your_own_dog_food#:~:text=Eating%20your%20own%20dog%20food%20or%20%22dogfooding%22%20is%20the%20practice,usage%20using%20product%20management%20techniques.) it, by working through my own course [Learn Go with Tests](https://github.com/quii/learn-go-with-tests). 

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

The idea of Pepper is you can use it as a basic assertion library, much like other libraries out there, to do simple assertions like the ones in Hello World, but it is also extensible, allowing you to write your own "Matchers" for your domain. 

This helps you write tests with clearer test output, more consistently, and over the longer run, cheaply. 

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

We can declaratively test properties of shapes with custom matchers, and leveraging Pepper's composition tools like `And`. 

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

## Whilst Pepper may look magical, it's just functions

I found I could be using the same combination of matchers repeatedly. All I had to do was mash `extract method` in Intellij (command+option+m), and I could reduce that boilerplate safely and easily.

```go
func ToHaveOutputted(expected string) Matcher[io.Reader] {
	return HaveString(EqualTo(expected))
}
```