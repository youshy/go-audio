package sine

import "testing"

func TestGenerateSine(t *testing.T) {
	t.Run("Generate sine runs 50 times", func(t *testing.T) {
		GenerateSine(50)
	})
}
