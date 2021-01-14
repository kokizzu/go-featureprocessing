// Code generated by go-featureprocessing DO NOT EDIT

package examplemodule

import (
	"encoding/json"
	"testing"

	"github.com/google/gofuzz"
	"github.com/stretchr/testify/assert"
)

// makeMock creates some valid LargeMemoryTransformerFeatureTransformer by fitting on fuzzy data.
// This function is handy for tests.
func makeMockLargeMemoryTransformerFeatureTransformer() *LargeMemoryTransformerFeatureTransformer {
	s := make([]LargeMemoryTransformer, 1000000)
	fuzz.New().NilChance(0).NumElements(10, 10).Fuzz(&s)

	tr := LargeMemoryTransformerFeatureTransformer{}
	tr.Fit(s)
	return &tr
}

func TestLargeMemoryTransformerFeatureTransformerFeatureNames(t *testing.T) {
	validTransformer := makeMockLargeMemoryTransformerFeatureTransformer()

	fuzzyTransformer := LargeMemoryTransformerFeatureTransformer{}
	fuzz.New().NilChance(0).NumElements(10, 10).Fuzz(&fuzzyTransformer)

	t.Run("feature names", func(t *testing.T) {
		names := validTransformer.FeatureNames()
		assert.True(t, len(names) > 0)
		assert.Equal(t, len(names), validTransformer.NumFeatures())
	})

	t.Run("feature names fuzzy transformer has some feature names", func(t *testing.T) {
		names := fuzzyTransformer.FeatureNames()
		assert.True(t, len(names) > 0)
	})

	t.Run("feature name transformer is empty", func(t *testing.T) {
		tr := LargeMemoryTransformerFeatureTransformer{}
		names := tr.FeatureNames()
		assert.True(t, len(names) > 0)
		assert.Equal(t, len(names), tr.NumFeatures())
	})

	t.Run("feature name transformer is nil", func(t *testing.T) {
		var tr *LargeMemoryTransformerFeatureTransformer
		names := tr.FeatureNames()
		assert.Nil(t, names)
	})
}

func TestLargeMemoryTransformerFeatureTransformerTransform(t *testing.T) {
	t.Run("empty struct fuzzy transformer", func(t *testing.T) {
		s := LargeMemoryTransformer{}

		tr := LargeMemoryTransformerFeatureTransformer{}
		fuzz.New().NilChance(0).NumElements(1, 1).Fuzz(&tr)

		tr.Transform(&s)

		features := tr.Transform(&s)

		assert.NotNil(t, features)
		assert.True(t, len(features) > 0)
		assert.Equal(t, tr.NumFeatures(), len(features))
	})

	t.Run("fuzzy struct", func(t *testing.T) {
		var s LargeMemoryTransformer
		fuzz.New().Fuzz(&s)

		tr := LargeMemoryTransformerFeatureTransformer{}
		fuzz.New().NilChance(0).NumElements(1, 1).Fuzz(&tr)

		features := tr.Transform(&s)

		assert.NotNil(t, features)
		assert.True(t, len(features) > 0)
		assert.Equal(t, tr.NumFeatures(), len(features))
	})

	t.Run("transformer is nil and struct is not nil", func(t *testing.T) {
		var s LargeMemoryTransformer
		fuzz.New().Fuzz(&s)

		var tr *LargeMemoryTransformerFeatureTransformer

		features := tr.Transform(&s)

		assert.Nil(t, features)
		assert.Equal(t, tr.NumFeatures(), 0)
	})

	t.Run("transformer is not nil but struct is nil", func(t *testing.T) {
		var s *LargeMemoryTransformer

		tr := LargeMemoryTransformerFeatureTransformer{}
		fuzz.New().NilChance(0).NumElements(1, 1).Fuzz(&tr)

		features := tr.Transform(s)

		assert.Nil(t, features)
		assert.True(t, tr.NumFeatures() > 0)
	})

	t.Run("serialize and deserialize transformer", func(t *testing.T) {
		tr := LargeMemoryTransformerFeatureTransformer{}
		fuzz.New().NilChance(0).NumElements(1, 1).Fuzz(&tr)

		output, err := json.Marshal(tr)
		assert.Nil(t, err)
		assert.NotEmpty(t, output)

		var tr2 LargeMemoryTransformerFeatureTransformer
		err = json.Unmarshal(output, &tr2)
		assert.Nil(t, err)
		assert.Equal(t, tr, tr2)
	})

	t.Run("inplace transform does not run when destination does not match num features", func(t *testing.T) {
		var s LargeMemoryTransformer
		fuzz.New().Fuzz(&s)

		tr := LargeMemoryTransformerFeatureTransformer{}
		fuzz.New().NilChance(0).NumElements(1, 1).Fuzz(&tr)

		features := make([]float64, 1000)
		features[0] = 123456789.0
		tr.TransformInplace(features, &s)

		assert.Equal(t, 123456789.0, features[0])
	})
}

func TestLargeMemoryTransformerFeatureTransformerFit(t *testing.T) {
	t.Run("fuzzy input", func(t *testing.T) {
		s := make([]LargeMemoryTransformer, 10)
		fuzz.New().NilChance(0).NumElements(1, 1).Fuzz(&s)

		trEmpty := LargeMemoryTransformerFeatureTransformer{}
		tr := LargeMemoryTransformerFeatureTransformer{}
		tr.Fit(s)

		assert.NotNil(t, tr)
		assert.NotEqual(t, tr, trEmpty)
	})

	t.Run("not nil transformer nil input", func(t *testing.T) {
		trEmpty := LargeMemoryTransformerFeatureTransformer{}
		tr := LargeMemoryTransformerFeatureTransformer{}
		tr.Fit(nil)

		assert.Equal(t, trEmpty, tr)
	})

	t.Run("nil transformer not nil input", func(t *testing.T) {
		s := make([]LargeMemoryTransformer, 10)

		var tr *LargeMemoryTransformerFeatureTransformer
		tr.Fit(s)

		assert.Nil(t, tr)
	})
}

func fitTransformerLargeMemoryTransformer(b *testing.B, numelem int) {
	s := make([]LargeMemoryTransformer, numelem)
	fuzz.New().NilChance(0).NumElements(numelem, numelem).Fuzz(&s)

	var tr LargeMemoryTransformerFeatureTransformer

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		tr.Fit(s)
	}
}

func BenchmarkLargeMemoryTransformerFeatureTransformer_Fit_100elements(b *testing.B) {
	fitTransformerLargeMemoryTransformer(b, 100)
}

func BenchmarkLargeMemoryTransformerFeatureTransformer_Fit_1000elements(b *testing.B) {
	fitTransformerLargeMemoryTransformer(b, 1000)
}

func BenchmarkLargeMemoryTransformerFeatureTransformer_Fit_10000elements(b *testing.B) {
	fitTransformerLargeMemoryTransformer(b, 10000)
}

func BenchmarkLargeMemoryTransformerFeatureTransformer_Transform(b *testing.B) {
	var s LargeMemoryTransformer
	fuzz.New().Fuzz(&s)

	tr := LargeMemoryTransformerFeatureTransformer{}
	fuzz.New().NilChance(0).NumElements(1, 1).Fuzz(&tr)

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		tr.Transform(&s)
	}
}

func BenchmarkLargeMemoryTransformerFeatureTransformer_Transform_Inplace(b *testing.B) {
	var s LargeMemoryTransformer
	fuzz.New().Fuzz(&s)

	tr := LargeMemoryTransformerFeatureTransformer{}
	fuzz.New().NilChance(0).NumElements(1, 1).Fuzz(&tr)

	features := make([]float64, tr.NumFeatures())

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		tr.TransformInplace(features, &s)
	}
}

func benchLargeTransformerLargeMemoryTransformer(b *testing.B, numelem int) {
	var s LargeMemoryTransformer
	fuzz.New().Fuzz(&s)

	tr := LargeMemoryTransformerFeatureTransformer{}
	fuzz.New().NilChance(0).NumElements(numelem, numelem).Fuzz(&tr)

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		tr.Transform(&s)
	}
}

func BenchmarkLargeMemoryTransformerFeatureTransformer_Transform_LargeComposites_100elements(b *testing.B) {
	benchLargeTransformerLargeMemoryTransformer(b, 100)
}

func BenchmarkLargeMemoryTransformerFeatureTransformer_Transform_LargeComposites_1000elements(b *testing.B) {
	benchLargeTransformerLargeMemoryTransformer(b, 1000)
}

func BenchmarkLargeMemoryTransformerFeatureTransformer_Transform_LargeComposites_10000elements(b *testing.B) {
	benchLargeTransformerLargeMemoryTransformer(b, 10000)
}

func BenchmarkLargeMemoryTransformerFeatureTransformer_Transform_LargeComposites_100000elements(b *testing.B) {
	benchLargeTransformerLargeMemoryTransformer(b, 100000)
}
