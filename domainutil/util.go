package domainutil

import (
	"strings"
	"golang.org/x/net/idna"
)

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

	//Cycle trough parts in reverse
	if len(parts) > 1 {
		for i := len(parts) - 1; i >= 0; i-- {
			//Generate top domain output
			if top != "" {
				top = "." + top
			}
			top = parts[i] + top

			//Check for TLD
			if currentTld == nil {
				return top //Return current output because we no longer have the TLD
			} else if tldEntry, found := currentTld[parts[i]]; found {
				if tldEntry != nil {
					currentTld = *tldEntry
				} else {
					currentTld = nil
				}
				foundTld = true
				continue
			} else if foundTld {
				return top //Return current output if tld was found before
			}

			//Return empty string if no tld was found ever
			return ""
		}
	}

	return ""
}

// stripURLParts removes path, protocol & query from url and returns it.
func stripURLParts(url string) string {
	//Lower case the url
	url = strings.ToLower(url)

	//Strip protocol
	if index := strings.Index(url, "://"); index > -1 {
		url = url[index+3:]
	}

	//Strip path (and query with it)
	if index := strings.Index(url, "/"); index > -1 {
		url = url[:index]
	} else if index := strings.Index(url, "?"); index > -1 { //Strip query if path is not found
		url = url[:index]
	}
	
	//Convert domain to unicode
	if strings.Index(url, "xn--") != -1 {
		var err error
		url, err = idna.ToUnicode(url)
		if err != nil {
			return ""
		}
	}

	//Return domain
	return url
}
