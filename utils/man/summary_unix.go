//go:build !windows
// +build !windows

package man

import (
	"bufio"
	"compress/gzip"
	"io"
	"os"
	"strings"

	"github.com/lmorg/murex/utils/cache"
)

// ParseSummary runs the parser to locate a summary
func ParseSummary(paths []string) string {
	var summary string
	for i := range paths {
		if !rxMatchManSection.MatchString(paths[i]) {
			continue
		}

		if cache.Read(cache.MAN_SUMMARY, paths[i], &summary) {
			return summary
		}

		summary = parseSummary(paths[i])
		if summary != "" {
			cache.Write(cache.MAN_SUMMARY, paths[i], summary, cache.Days(30))
			return summary
		}
	}

	return ""
}

func parseSummary(filename string) string {
	file, err := os.Open(filename)
	if err != nil {
		return ""
	}
	defer file.Close()

	if len(filename) > 3 && filename[len(filename)-3:] == ".gz" {
		gz, err := gzip.NewReader(file)
		if err != nil {
			return ""
		}
		defer gz.Close()

		return parseSummaryFile(gz)
	}

	return parseSummaryFile(file)
}

func parseSummaryFile(file io.Reader) string {
	scanner := bufio.NewScanner(file)

	var (
		read bool
		desc string
	)

	for scanner.Scan() {
		s := scanner.Text()

		if strings.Contains(s, "SYNOPSIS") {
			if len(desc) > 0 && desc[len(desc)-1] == '-' {
				desc = desc[:len(desc)-1]
			}
			return strings.TrimSpace(desc)
		}

		if read {
			// Tidy up man pages generated from reStructuredText
			if strings.HasPrefix(s, `\\n[rst2man-indent`) ||
				strings.HasPrefix(s, `\\$1 \\n`) ||
				strings.HasPrefix(s, `level \\n`) ||
				strings.HasPrefix(s, `level margin: \\n`) {
				continue
			}

			s = strings.Replace(s, ".Nd ", " - ", -1)
			s = strings.Replace(s, "\\(em ", " - ", -1)
			s = strings.Replace(s, " , ", ", ", -1)
			s = strings.Replace(s, "\\fB", "", -1)
			s = strings.Replace(s, "\\fR", "", -1)
			if strings.HasSuffix(s, " ,") {
				s = s[:len(s)-2] + ", "
			}
			s = rxReplaceMarkup.ReplaceAllString(s, "")
			s = strings.Replace(s, "\\", "", -1)

			if strings.HasPrefix(s, `.`) {
				continue
			}

			desc += s
		}

		if strings.Contains(s, "NAME") {
			read = true
		}
	}

	return ""
}
