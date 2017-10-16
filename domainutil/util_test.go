package domainutil

import (
	"fmt"
	"testing"
)

func ExampleHasSubdomain() {
	fmt.Println(HasSubdomain("google.com"))
	fmt.Println(HasSubdomain("keep.google.com"))
	// Output: false
	// true
}

// TestHasSubdomain tests HasSubdomain() function
func TestHasSubdomain(t *testing.T) {
	//Test cases
	cases := map[string]bool{
		"http://google.com":           false,
		"http://google.com/ding?true": false,
		"google.com/?ding=false":      false,
		"google.com?ding=false":       false,
		"nonexist.***":                false,
		"google.com":                  false,
		"google.co.uk":                false,
		"gama.google.com":             true,
		"gama.google.co.uk":           true,
		"beta.gama.google.co.uk":      true,
	}

	//Test each domain, some should fail (expected)
	for url, shouldHaveSubdomain := range cases {
		hasSubdomain := HasSubdomain(url)
		if hasSubdomain != shouldHaveSubdomain {
			t.Errorf("Url (%q) returned %v for HasSubdomain(), but %v was expected", url, hasSubdomain, shouldHaveSubdomain)
		}
	}
}

// BenchmarkHasSubdomain benchmarks HasSubdomain() function
func BenchmarkHasSubdomain(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HasSubdomain("https://beta.gama.google.co.uk?test=true")
	}
}

func ExampleSubdomain() {
	fmt.Printf("%q %q", Subdomain("google.com"), Subdomain("keep.google.com"))
	// Output: "" "keep"
}

// TestSubdomain tests Subdomain() function
func TestSubdomain(t *testing.T) {
	//Test cases
	cases := map[string]string{
		"http://google.com":           "",
		"http://google.com/ding?true": "",
		"google.com/?ding=false":      "",
		"google.com?ding=false":       "",
		"nonexist.***":                "",
		"google.com":                  "",
		"google.co.uk":                "",
		"gama.google.com":             "gama",
		"gama.google.co.uk":           "gama",
		"beta.gama.google.co.uk":      "beta.gama",
		"": "",
	}

	//Test each domain, some should fail (expected)
	for url, expectedSubdomain := range cases {
		subdomain := Subdomain(url)
		if subdomain != expectedSubdomain {
			t.Errorf("Url (%q) returned %q for Subdomain(), but %q was expected", url, subdomain, expectedSubdomain)
		}
	}
}

// BenchmarkSubdomain benchmarks Subdomain() function
func BenchmarkSubdomain(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Subdomain("https://beta.gama.google.co.uk?test=true")
	}
}

func ExampleDomainSuffix() {
	fmt.Println(DomainSuffix("google.co.uk"))
	fmt.Println(DomainSuffix("keep.google.com"))
	// Output: co.uk
	// com
}

// TestDomainSuffix tests DomainSuffix() function
func TestDomainSuffix(t *testing.T) {
	//Test cases
	cases := map[string]string{
		"http://google.com":           "com",
		"http://google.com/ding?true": "com",
		"google.com/?ding=false":      "com",
		"google.com?ding=false":       "com",
		"nonexist.***":                "",
		"google.com":                  "com",
		"google.co.uk":                "co.uk",
		"gama.google.com":             "com",
		"gama.google.co.uk":           "co.uk",
		"beta.gama.google.co.uk":      "co.uk",
	}

	//Test each domain, some should fail (expected)
	for url, expectedSuffix := range cases {
		domainSuffix := DomainSuffix(url)
		if domainSuffix != expectedSuffix {
			t.Errorf("Url (%q) returned %q for DomainSuffix(), but %q was expected", url, domainSuffix, expectedSuffix)
		}
	}
}

// BenchmarkDomainSuffix benchmarks DomainSuffix() function
func BenchmarkDomainSuffix(b *testing.B) {
	for i := 0; i < b.N; i++ {
		DomainSuffix("https://beta.gama.google.co.uk?test=true")
	}
}

func ExampleDomain() {
	fmt.Println(Domain("google.co.uk"))
	fmt.Println(Domain("keep.google.com"))
	// Output: google.co.uk
	// google.com
}

// TestDomain tests Domain() function
func TestDomain(t *testing.T) {
	//Test cases
	cases := map[string]bool{
		"http://google.com":           true,
		"http://google.com/ding?true": true,
		"google.com/?ding=false":      true,
		"google.com?ding=false":       true,
		"nonexist.***":                false,
		"google.com":                  true,
		"google.co.uk":                true,
		"gama.google.com":             true,
		"gama.google.co.uk":           true,
		"beta.gama.google.co.uk":      true,
		"something.blogspot.com":      true,
		"something.blogspot.co.uk":    true,
	}

	//Test each domain, some should fail (expected)
	for url, shouldNotBeEmpty := range cases {
		domain := Domain(url)
		if domain == "" && shouldNotBeEmpty {
			t.Errorf("Url (%q) returned empty string as domain, but expected non-empty string", url)
		} else if domain != "" && !shouldNotBeEmpty {
			t.Errorf("Url (%q) returned (%s) as domain, but expected empty string", url, domain)
		}
	}
}

// BenchmarkDomain benchmarks Domain() function
func BenchmarkDomain(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Domain("https://beta.gama.google.co.uk?test=true")
	}
}

// TestStripURLParts tests stripURLParts() function
func TestStripURLParts(t *testing.T) {
	//Test cases
	cases := map[string]string{
		"http://google.com":              "google.com",
		"http://google.com/ding?true":    "google.com",
		"google.com/?ding=false":         "google.com",
		"google.com?ding=false":          "google.com",
		"nonexist.***":                   "nonexist.***",
		"google.com":                     "google.com",
		"google.co.uk":                   "google.co.uk",
		"gama.google.com":                "gama.google.com",
		"gama.google.co.uk":              "gama.google.co.uk",
		"beta.gama.google.co.uk":         "beta.gama.google.co.uk",
		"https://beta.gama.google.co.uk": "beta.gama.google.co.uk",
		"xn--n3h.example":                "☃.example",
		"xn--äää":                        "",
	}

	//Test each domain, some should fail (expected)
	for url, expectedStripped := range cases {
		stripped := stripURLParts(url)
		if stripped != expectedStripped {
			t.Errorf("Url (%q) returned %q for StripURLParts(), but %q was expected", url, stripped, expectedStripped)
		}
	}
}

// BenchmarkStripURLParts benchmarks StripURLParts() function
func BenchmarkStripURLParts(b *testing.B) {
	for i := 0; i < b.N; i++ {
		stripURLParts("https://beta.gama.google.co.uk?test=true")
	}
}
