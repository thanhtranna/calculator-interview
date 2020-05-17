package calculator

import "testing"

func TestCalculator(t *testing.T) {
	c := Calculator{}
	t.Run("addition", func(t *testing.T) {
		output, err := c.Evaluate("1 + 2")
		assert(t, "3", output, err)
	})
	t.Run("multiplication", func(t *testing.T) {
		output, err := c.Evaluate("2 * 3 + 4")
		assert(t, "10", output, err)
	})
	t.Run("parentheses", func(t *testing.T) {
		output, err := c.Evaluate("2 * (3 + 4)")
		assert(t, "14", output, err)
	})
	t.Run("negation", func(t *testing.T) {
		output, err := c.Evaluate("-(2 + 1)")
		assert(t, "-3", output, err)
	})

	// bonus point
	t.Run("assignment", func(t *testing.T) {
		output, err := c.Evaluate("A = 8 + 2")
		assert(t, "10", output, err)
	})
	t.Run("variable", func(t *testing.T) {
		output, err := c.Evaluate("A * 2")
		assert(t, "20", output, err)
	})
}

func assert(t *testing.T, expected string, actual string, err error) {
	if err != nil {
		t.Errorf("unexpected error: %v", err)
		return
	}
	if expected != actual {
		t.Errorf("expect `%v`, got `%v`", expected, actual)
	}
}
