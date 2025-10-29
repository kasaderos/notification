package html

import (
	"bytes"
	"regexp"
)

// ProcessContent removes JS, CSS and cleans HTML content
func ProcessContent(html []byte) string {
	// Remove script tags and their content
	scriptRegex := regexp.MustCompile(`(?i)<script[^>]*>.*?</script>`)
	html = scriptRegex.ReplaceAll(html, []byte(""))

	// Remove style tags and their content
	styleRegex := regexp.MustCompile(`(?i)<style[^>]*>.*?</style>`)
	html = styleRegex.ReplaceAll(html, []byte(""))

	// Remove HTML comments
	commentRegex := regexp.MustCompile(`<!--.*?-->`)
	html = commentRegex.ReplaceAll(html, []byte(""))

	// Remove HTML tags
	tagRegex := regexp.MustCompile(`<[^>]*>`)
	html = tagRegex.ReplaceAll(html, []byte(" "))

	// Decode HTML entities
	html = bytes.ReplaceAll(html, []byte("&nbsp;"), []byte(" "))
	html = bytes.ReplaceAll(html, []byte("&amp;"), []byte("&"))
	html = bytes.ReplaceAll(html, []byte("&lt;"), []byte("<"))
	html = bytes.ReplaceAll(html, []byte("&gt;"), []byte(">"))
	html = bytes.ReplaceAll(html, []byte("&quot;"), []byte("\""))
	html = bytes.ReplaceAll(html, []byte("&#39;"), []byte("'"))

	// Clean up whitespace
	whitespaceRegex := regexp.MustCompile(`\s+`)
	html = whitespaceRegex.ReplaceAll(html, []byte(" "))

	// Trim and limit length
	html = bytes.TrimSpace(html)

	return string(html)
}

// ExtractTitle extracts title from HTML content
// todo
func ExtractTitle(html []byte) string {
	titleRegex := regexp.MustCompile(`(?is)<title[^>]*>(.*?)</title>`)
	matches := titleRegex.FindSubmatch(html)

	if len(matches) > 1 {
		// Remove nested tags from title content
		title := matches[1]
		title = regexp.MustCompile(`<[^>]*>`).ReplaceAll(title, []byte(""))
		title = bytes.TrimSpace(title)

		return string(title)
	}

	return "Untitled"
}
