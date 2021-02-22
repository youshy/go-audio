package sine

import "testing"

/*
func TestGenerateSine(t *testing.T) {
	t.Run("Generate sine well, runs", func(t *testing.T) {
		GenerateSine()
	})
}
*/

func TestGenerateTunedSine(t *testing.T) {
	t.Run("Generate tuned sine well, runs and generates 440hz sine", func(t *testing.T) {
		GenerateTunedSine(30)
	})
	/*
		t.Run("Generate tuned sine well, runs and generates 31hz sine", func(t *testing.T) {
			GenerateTunedSine(31)
		})
		t.Run("Generate tuned sine well, runs and generates 12000hz sine", func(t *testing.T) {
			GenerateTunedSine(12000)
		})
	*/
}
