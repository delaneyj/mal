package mal_test

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func testMal(t *testing.T, malPath string) {
	testDir, err := filepath.Abs("/home/delaney/repos/mal/impls/dg/src")
	assert.Nil(t, err)

	fullMalPath := filepath.Join(testDir, "..", "..", "..", "tests", malPath)
	b, err := os.ReadFile(fullMalPath)
	assert.Nil(t, err)

	lines := strings.Split(string(b), "\n")
	i := 0

	clearWhitespace := func() {
		for i < len(lines) {
			line := lines[i]
			if strings.TrimSpace(line) != "" {
				break
			}
			i++
		}
	}

	for i < len(lines) {
		clearWhitespace()
		if i >= len(lines) {
			break
		}
		line := lines[i]
		if strings.HasPrefix(line, ";;") || strings.HasPrefix(line, ";>>>") {
			t.Log(line)
			i++
		} else {
			actual := ";=>" + line
			expected := lines[i+1]
			assert.Equal(t, expected, actual)
			i += 2
		}
	}
}

func TestStep0(t *testing.T) {
	testMal(t, "step0_repl.mal")
}
