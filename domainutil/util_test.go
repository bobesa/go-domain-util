package domainutil

import (
	"fmt"
	"reflect"
	"testing"
)

func ExampleHasSubdomain() {
	fmt.Println(HasSubdomain("google.com"))
	fmt.Println(HasSubdomain("keep.google.com"))
	// Output: false
	// true
}

func TestSplitDomain(t *testing.T) {
	cases := map[string][]string{
		"http://zh.wikipedia.org":                          {"zh", "wikipedia", "org"},
		"zh.wikipedia.org":                                 {"zh", "wikipedia", "org"},
		"https://zh.wikipedia.org/wiki/%E5%9F%9F%E5%90%8D": {"zh", "wikipedia", "org"},
		"wikipedia.org":                                    {"wikipedia", "org"},
		".org":                                             {"org"},
		"org":                                              nil,
		"a.b.c.d.wikipedia.org": {"a", "b", "c", "d", "wikipedia", "org"},
	}

	for url, array := range cases {
		arrVal := SplitDomain(url)
		if !reflect.DeepEqual(array, arrVal) {
			t.Errorf("Url (%q) return %v for SplitDomain, bug %v was expected", url, arrVal, array)
		}
	}
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

// TestDomainPrefix tests DomainPrefix function
func TestDomainPrefix(t *testing.T) {
	//Test cases
	cases := map[string]string{
		"http://google.com":           "google",
		"http://google.com/ding?true": "google",
		"google.com/?ding=false":      "google",
		"google.com?ding=false":       "google",
		"google.com":                  "google",
		"google.co.uk":                "google",
		"gama.google.com":             "google",
		"gama.google.co.uk":           "google",
		"beta.gama.google.co.uk":      "google",
	}

	for url, expectedPrefix := range cases {
		domainPrefix := DomainPrefix(url)
		if domainPrefix != expectedPrefix {
			t.Errorf("Url (%q) returned %q for DomainPrefix(), but %q was expected", url, domainPrefix, expectedPrefix)
		}
	}
}

func BenchmarkDomainPrefix(b *testing.B) {
	for i := 0; i < b.N; i++ {
		DomainPrefix("https://beta.gama.google.co.uk?test=true")
	}
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
		"http://google.com":                                    "google.com",
		"http://google.com/ding?true":                          "google.com",
		"google.com/?ding=false":                               "google.com",
		"google.com?ding=false":                                "google.com",
		"nonexist.***":                                         "nonexist.***",
		"google.com":                                           "google.com",
		"google.co.uk":                                         "google.co.uk",
		"gama.google.com":                                      "gama.google.com",
		"gama.google.co.uk":                                    "gama.google.co.uk",
		"beta.gama.google.co.uk":                               "beta.gama.google.co.uk",
		"https://beta.gama.google.co.uk":                       "beta.gama.google.co.uk",
		"xn--n3h.example":                                      "☃.example",
		"xn--äää":                                              "",
		"http://admin:adminpw@google.com":                      "google.com",
		"admin:adminpw@gama.google.com":                        "gama.google.com",
		"https://admin:adminpw@gama.google.com/path?key=value": "gama.google.com",
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

func ExampleProtocol() {
	fmt.Printf("%q\n", Protocol("google.com"))
	fmt.Printf("%q\n", Protocol("ftp://google.com"))
	fmt.Printf("%q\n", Protocol("http://google.com"))
	fmt.Printf("%q\n", Protocol("https://google.com"))
	fmt.Printf("%q\n", Protocol("https://user@google.com"))
	fmt.Printf("%q\n", Protocol("https://user:pass@google.com"))
	// Output: ""
	// "ftp"
	// "http"
	// "https"
	// "https"
	// "https"
}

// TestProtocol tests Protocol() function
func TestProtocol(t *testing.T) {
	for _, testCase := range []struct{ URL, Expected string }{
		{"google.com", ""},
		{"ftp://google.com", "ftp"},
		{"http://google.com", "http"},
		{"https://google.com", "https"},
		{"https://user@google.com", "https"},
		{"https://user:pass@google.com", "https"},
	} {
		if result := Protocol(testCase.URL); result != testCase.Expected {
			t.Errorf(`Url (%q) returned %q for Protocol(), but %q was expected`, testCase.URL, result, testCase.Expected)
		}
	}
}

// BenchmarkProtocol benchmarks Protocol() function
func BenchmarkProtocol(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Protocol("https://user:pass@beta.gama.google.co.uk?test=true")
	}
}

func ExampleUsername() {
	fmt.Printf("%q\n", Username("user:pass@google.com"))
	fmt.Printf("%q\n", Username("https://user:pass@google.com"))
	fmt.Printf("%q\n", Username("https://user@google.com"))
	fmt.Printf("%q\n", Username("https://google.com"))
	fmt.Printf("%q\n", Username("google.com"))
	// Output: "user"
	// "user"
	// "user"
	// ""
	// ""
}

// TestUsername tests Username() function
func TestUsername(t *testing.T) {
	for _, testCase := range []struct{ URL, Expected string }{
		{"user:pass@google.com", "user"},
		{"https://user:pass@google.com", "user"},
		{"https://user@google.com", "user"},
		{"https://google.com", ""},
		{"google.com", ""},
	} {
		if result := Username(testCase.URL); result != testCase.Expected {
			t.Errorf(`Url (%q) returned %q for Username(), but %q was expected`, testCase.URL, result, testCase.Expected)
		}
	}
}

// BenchmarkUsername benchmarks Username() function
func BenchmarkUsername(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Username("https://user:pass@beta.gama.google.co.uk?test=true")
	}
}

func ExamplePassword() {
	fmt.Printf("%q\n", Password("user:pass@google.com"))
	fmt.Printf("%q\n", Password("https://user:pass@google.com"))
	fmt.Printf("%q\n", Password("https://user@google.com"))
	fmt.Printf("%q\n", Password("https://google.com"))
	fmt.Printf("%q\n", Password("google.com"))
	// Output: "pass"
	// "pass"
	// ""
	// ""
	// ""
}

// TestPassword tests Password() function
func TestPassword(t *testing.T) {
	for _, testCase := range []struct{ URL, Expected string }{
		{"user:pass@google.com", "pass"},
		{"https://user:pass@google.com", "pass"},
		{"https://user@google.com", ""},
		{"https://google.com", ""},
		{"google.com", ""},
	} {
		if result := Password(testCase.URL); result != testCase.Expected {
			t.Errorf(`Url (%q) returned %q for Password(), but %q was expected`, testCase.URL, result, testCase.Expected)
		}
	}
}

// BenchmarkPassword benchmarks Password() function
func BenchmarkPassword(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Password("https://user:pass@beta.gama.google.co.uk?test=true")
	}
}
