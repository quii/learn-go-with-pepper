package reading_files

import (
	"fmt"
	. "github.com/quii/pepper"
	. "github.com/quii/pepper/matchers/comparable"
	"github.com/quii/pepper/matchers/slice"
	"testing"
	"testing/fstest"
)

func TestNewBlogPosts(t *testing.T) {
	fs := fstest.MapFS{
		"hello-world.md": {Data: []byte(`Title: Post 1
Description: Description 1
Tags: tdd, go
---
Hello
World`)},

		"hello-world2.md": {Data: []byte(`Title: Post 2
Description: Description 2
Tags: rust, borrow-checker
---
B
L
M`)},
	}

	posts, err := NewPostsFromFS(fs)

	ExpectNoError(t, err)
	Expect(t, posts).To(slice.HaveSize[Post](EqualTo(2)))
	Expect(t, posts).To(slice.ContainItem(PostEqualTo(
		Post{
			Title:       "Post 1",
			Description: "Description 1",
			Tags:        []string{"tdd", "go"},
			Body: `Hello
World`,
		}),
	))
}

func PostEqualTo(want Post) Matcher[Post] {
	return func(post Post) MatchResult {
		return HaveTitle(want.Title).
			And(HaveDescription(want.Description).
				And(HaveTags(want.Tags)).
				And(HaveBody(want.Body)))(post)
	}
}

func HaveBody(want string) Matcher[Post] {
	return func(actual Post) MatchResult {
		return MatchResult{
			Description: fmt.Sprintf("have body of %q", want),
			Matches:     actual.Body == want,
			But:         fmt.Sprintf("has body %q", actual.Body),
		}
	}
}

func HaveTitle(title string) Matcher[Post] {
	return func(actual Post) MatchResult {
		return MatchResult{
			Description: fmt.Sprintf("have title of %q", title),
			Matches:     actual.Title == title,
			But:         fmt.Sprintf("has %q", actual.Title),
		}
	}
}

func HaveDescription(description string) Matcher[Post] {
	return func(actual Post) MatchResult {
		return MatchResult{
			Description: fmt.Sprintf("have description of %q", description),
			Matches:     actual.Description == description,
			But:         fmt.Sprintf("has %q", actual.Description),
		}
	}
}

func HaveTags(tags []string) Matcher[Post] {
	return func(actual Post) MatchResult {
		result := slice.ShallowEquals(tags)(actual.Tags)
		result.Description = fmt.Sprintf("have tags of %q", tags)
		result.But = fmt.Sprintf("has %q", actual.Tags)
		return result
	}
}
