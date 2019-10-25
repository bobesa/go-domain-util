package domainutil

import (
	"strings"

	"golang.org/x/net/idna"
)
// HasSubdomainQuantity checks the amount of subdomains in domain.
// If quantity matches the number of subdomains in domain, this function returns true.
func HasSubdomainQuantity(domain string, quantity int) bool {
	domainSplit := SplitDomain(domain)
	if len(domainSplit) - 2 == quantity {
		return true
	}
	return false
}

// HasSubdomain reports whether domain contains any subdomain.
func HasSubdomain(domain string) bool {
	domain, top := stripURLParts(domain), Domain(domain)
	return domain != top && top != ""
}

// Subdomain returns subdomain from provided url.
// If subdomain is not found in provided url, this function returns empty string.
func Subdomain(url string) string {
	domain, top := stripURLParts(url), Domain(url)
	lt, ld := len(top), len(domain)
	if lt < ld && top != "" {
		return domain[:(ld-lt)-1]
	}
	return ""
}

// SplitDomain split domain into string array
// for example, zh.wikipedia.org will split into {"zh", "wikipedia", "org"}
func SplitDomain(url string) []string {
	domain, second, top := Subdomain(url), DomainPrefix(url), DomainSuffix(url)
	if len(top) == 0 {
		return nil
	}

	if len(second) == 0 {
		return []string{top}
	}

	if len(domain) == 0 {
		return []string{second, top}
	}

	array := strings.Split(domain, ".")
	res := append(array, second, top)
	return res
}

// DomainPrefix returns second-level domain from provided url.
// If no SLD is found in provided url, this function returns empty string.
func DomainPrefix(url string) string {
	domain := Domain(url)
	if len(domain) != 0 {
		return domain[:strings.Index(domain, ".")]
	}
	return ""
}

// DomainSuffix returns domain suffix from provided url.
// If no TLD is found in provided url, this function returns empty string.
func DomainSuffix(url string) string {
	domain := Domain(url)
	if len(domain) != 0 {
		return domain[strings.Index(domain, ".")+1:]
	}
	return ""
}

// Domain returns top level domain from url string.
// If no domain is found in provided url, this function returns empty string.
// If no TLD is found in provided url, this function returns empty string.
func Domain(url string) string {
	domain, top := stripURLParts(url), ""
	parts := strings.Split(domain, ".")
	currentTld := *tlds
	foundTld := false
	if !isDomainName(url){
		return ""
	}
	// Cycle trough parts in reverse
	if len(parts) > 1 {
		for i := len(parts) - 1; i >= 0; i-- {
			// Generate top domain output
			if top != "" {
				top = "." + top
			}
			top = parts[i] + top

			// Check for TLD
			if currentTld == nil {
				return top // Return current output because we no longer have the TLD
			} else if tldEntry, found := currentTld[parts[i]]; found {
				if tldEntry != nil {
					currentTld = *tldEntry
				} else {
					currentTld = nil
				}
				foundTld = true
				continue
			} else if foundTld {
				return top // Return current output if tld was found before
			}

			// Return empty string if no tld was found ever
			return ""
		}
	}

	return ""
}

// stripURLParts removes path, protocol & query from url and returns it.
func stripURLParts(url string) string {
	// Lower case the url
	url = strings.ToLower(url)

	// Strip protocol
	if index := strings.Index(url, "://"); index > -1 {
		url = url[index+3:]
	}

	// Now, if the url looks like this: username:password@www.example.com/path?query=?
	// we remove the content before the '@' symbol
	if index := strings.Index(url, "@"); index > -1 {
		url = url[index+1:]
	}

	// Strip path (and query with it)
	if index := strings.Index(url, "/"); index > -1 {
		url = url[:index]
	} else if index := strings.Index(url, "?"); index > -1 { // Strip query if path is not found
		url = url[:index]
	}

	// Convert domain to unicode
	if strings.Index(url, "xn--") != -1 {
		var err error
		url, err = idna.ToUnicode(url)
		if err != nil {
			return ""
		}
	}

	// Return domain
	return url
}

// Protocol returns protocol from given url
//
// If protocol is not present - return empty string
func Protocol(url string) string {
	if index := strings.Index(url, "://"); index > -1 {
		return url[:index]
	}
	return ""
}

// credentials returns credentials (user:pass) from given url
func credentials(url string) string {
	index := strings.IndexRune(url, '@')
	if index == -1 {
		return ""
	}
	if protocol := Protocol(url); protocol != "" {
		return url[len(protocol)+3 : index]
	}
	return url[:index]
}

// Username returns username from given url
//
// If username is not present - return empty string
func Username(url string) string {
	auth := strings.SplitN(credentials(url), ":", 2)
	if len(auth) == 0 {
		return ""
	}
	return auth[0]
}

// Password returns password from given url
//
// If password is not present - return empty string
func Password(url string) string {
	auth := strings.SplitN(credentials(url), ":", 2)
	if len(auth) < 2 {
		return ""
	}
	return auth[1]
}

// isDomainName checks if a string is a presentation-format domain name
// (currently restricted to hostname-compatible "preferred name" LDH labels and
// SRV-like "underscore labels"; see golang.org/issue/12421).
func isDomainName(s string) bool {
	// See RFC 1035, RFC 3696.
	// Presentation format has dots before every label except the first, and the
	// terminal empty label is optional here because we assume fully-qualified
	// (absolute) input. We must therefore reserve space for the first and last
	// labels' length octets in wire format, where they are necessary and the
	// maximum total length is 255.
	// So our _effective_ maximum is 253, but 254 is not rejected if the last
	// character is a dot.
	l := len(s)
	if l == 0 || l > 254 || l == 254 && s[l-1] != '.' {
		return false
	}

	last := byte('.')
	nonNumeric := false // true once we've seen a letter or hyphen
	partlen := 0
	for i := 0; i < len(s); i++ {
		c := s[i]
		switch {
		default:
			return false
		case 'a' <= c && c <= 'z' || 'A' <= c && c <= 'Z' || c == '_':
			nonNumeric = true
			partlen++
		case '0' <= c && c <= '9':
			// fine
			partlen++
		case c == '-':
			// Byte before dash cannot be dot.
			if last == '.' {
				return false
			}
			partlen++
			nonNumeric = true
		case c == '.':
			// Byte before dot cannot be dot, dash.
			if last == '.' || last == '-' {
				return false
			}
			if partlen > 63 || partlen == 0 {
				return false
			}
			partlen = 0
		}
		last = c
	}
	if last == '-' || partlen > 63 {
		return false
	}

	return nonNumeric
}
