package mal

import (
	"fmt"
	"log"
	"strings"
)

func read(s string) string {
	runes := []rune(s)

	tokensCh := make(chan string)
	errCh := make(chan error)
	tokenize(runes, tokensCh, errCh)

	for {
		select {
		case err := <-errCh:
			log.Fatal(err)
		case token, ok := <-tokensCh:
			if !ok {
				break
			}
			log.Printf("Token: %s", token)
		}
	}
}

func tokenize(runes []rune, tokenCh chan<- string, errCh chan<- error) {
	sb := &strings.Builder{}
	position := 0
	defer close(tokenCh)

	peek := func() rune {
		r := runes[position]
		position++
		return r
	}

	next := func() (c rune, ok bool) {
		if position >= len(runes) {
			errCh <- fmt.Errorf("unexpected EOF")
			return 0, false
		}
		c = peek()
		position++
		return c, true
	}

	finishToken := func() {
		defer sb.Reset()
		if sb.Len() > 0 {
			tokenCh <- sb.String()
		}
	}

	inString := false
	isComment := false
	for {
		c, ok := next()
		if !ok {
			break
		}

		if isComment {
			if c == '\n' {
				finishToken()
				isComment = false
				continue
			}
			sb.WriteRune(c)
			continue
		}

		switch c {
		case ';':
			sb.WriteRune(c)
			if p := peek(); p == ';' {
				next()
				isComment = true
				continue
			}
		case '\\':
			sb.WriteRune(c)
			if p := peek(); p == '"' {
				next()
				sb.WriteRune('"')
				continue
			}
		case '"':
			sb.WriteRune(c)
			if inString {
				inString = false
				finishToken()
				continue
			}
			inString = true
		case '~':
			sb.WriteString("~")
			switch peek() {
			case '@':
				next()
				sb.WriteRune('@')
				finishToken()
				continue
			}
		case '[', ']', '(', ')', '{', '}', '\'', '`', '^', '@':
			sb.WriteRune(c)
			finishToken()
		}
	}

}
