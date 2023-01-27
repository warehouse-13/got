package template

import "fmt"

func Generate(pkg string) []byte {
	contents := `package %s_test

import (
	"testing"

	. "github.com/onsi/gomega"
)

// Rename
func Test_simpleTest(t *testing.T) {
	g := NewWithT(t)

	t.Cleanup(func() {
		// Add cleanup here
	})

	// Test the code
	g.Expect(true).To(BeTrue())
}

// Rename
func Test_tableTest(t *testing.T) {
	g := NewWithT(t)

	t.Cleanup(func() {
		// Add cleanup here
	})

	// Edit the test variables
	tt := []struct {
		name     string
		input    string
		expected func(*WithT)
	}{
		// Create some test cases
		{
			name:  "explain what the test is showing",
			input: "some-value",
			expected: func(g *WithT) {
				g.Expect(true).To(BeTrue())
			},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			// Test the code
			tc.expected(g)
		})
	}
}
`

	return []byte(fmt.Sprintf(contents, pkg))
}
