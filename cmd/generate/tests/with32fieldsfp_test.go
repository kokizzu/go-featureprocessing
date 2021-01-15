// Code generated by go-featureprocessing DO NOT EDIT

package examplemodule

import (
	"encoding/json"
	"testing"

	"github.com/google/gofuzz"
	"github.com/stretchr/testify/assert"
)

// makeMock creates some valid With32FieldsFeatureTransformer by fitting on fuzzy data.
// This function is handy for tests.
func makeMockWith32FieldsFeatureTransformer() *With32FieldsFeatureTransformer {
	s := make([]With32Fields, 1000000)
	fuzz.New().NilChance(0).NumElements(10, 10).Fuzz(&s)

	tr := With32FieldsFeatureTransformer{}
	tr.Fit(s)
	return &tr
}

func TestWith32FieldsFeatureTransformerFeatureNames(t *testing.T) {
	validTransformer := makeMockWith32FieldsFeatureTransformer()

	fuzzyTransformer := With32FieldsFeatureTransformer{}
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
		tr := With32FieldsFeatureTransformer{}
		names := tr.FeatureNames()
		assert.True(t, len(names) > 0)
		assert.Equal(t, len(names), tr.NumFeatures())
	})

	t.Run("feature name transformer is nil", func(t *testing.T) {
		var tr *With32FieldsFeatureTransformer
		names := tr.FeatureNames()
		assert.Nil(t, names)
	})
}

func TestWith32FieldsFeatureTransformerTransform(t *testing.T) {
	t.Run("empty struct fuzzy transformer", func(t *testing.T) {
		s := With32Fields{}

		tr := With32FieldsFeatureTransformer{}
		fuzz.New().NilChance(0).NumElements(1, 1).Fuzz(&tr)

		tr.Transform(&s)

		features := tr.Transform(&s)

		assert.NotNil(t, features)
		assert.True(t, len(features) > 0)
		assert.Equal(t, tr.NumFeatures(), len(features))
	})

	t.Run("fuzzy struct", func(t *testing.T) {
		var s With32Fields
		fuzz.New().Fuzz(&s)

		tr := With32FieldsFeatureTransformer{}
		fuzz.New().NilChance(0).NumElements(1, 1).Fuzz(&tr)

		features := tr.Transform(&s)

		assert.NotNil(t, features)
		assert.True(t, len(features) > 0)
		assert.Equal(t, tr.NumFeatures(), len(features))
	})

	t.Run("transformer is nil and struct is not nil", func(t *testing.T) {
		var s With32Fields
		fuzz.New().Fuzz(&s)

		var tr *With32FieldsFeatureTransformer

		features := tr.Transform(&s)

		assert.Nil(t, features)
		assert.Equal(t, tr.NumFeatures(), 0)
	})

	t.Run("transformer is not nil but struct is nil", func(t *testing.T) {
		var s *With32Fields

		tr := With32FieldsFeatureTransformer{}
		fuzz.New().NilChance(0).NumElements(1, 1).Fuzz(&tr)

		features := tr.Transform(s)

		assert.Nil(t, features)
		assert.True(t, tr.NumFeatures() > 0)
	})

	t.Run("serialize and deserialize transformer", func(t *testing.T) {
		tr := With32FieldsFeatureTransformer{}
		fuzz.New().NilChance(0).NumElements(1, 1).Fuzz(&tr)

		output, err := json.Marshal(tr)
		assert.Nil(t, err)
		assert.NotEmpty(t, output)

		var tr2 With32FieldsFeatureTransformer
		err = json.Unmarshal(output, &tr2)
		assert.Nil(t, err)
		assert.Equal(t, tr, tr2)
	})

	t.Run("inplace transform does not run when destination does not match num features", func(t *testing.T) {
		var s With32Fields
		fuzz.New().Fuzz(&s)

		tr := With32FieldsFeatureTransformer{}
		fuzz.New().NilChance(0).NumElements(1, 1).Fuzz(&tr)

		features := make([]float64, 1000)
		features[0] = 123456789.0
		tr.TransformInplace(features, &s)

		assert.Equal(t, 123456789.0, features[0])
	})
}

func TestWith32FieldsFeatureTransformerTransformAll(t *testing.T) {
	t.Run("transform all", func(t *testing.T) {
		s := make([]With32Fields, 100)
		fuzz.New().NilChance(0).NumElements(100, 100).Fuzz(&s)

		tr := With32FieldsFeatureTransformer{}
		fuzz.New().NilChance(0).NumElements(100, 100).Fuzz(&tr)

		features := tr.TransformAll(s)
		assert.Equal(t, len(s)*tr.NumFeatures(), len(features))
	})

	t.Run("transform all parallel 1 worker", func(t *testing.T) {
		s := make([]With32Fields, 100)
		fuzz.New().NilChance(0).NumElements(100, 100).Fuzz(&s)

		tr := With32FieldsFeatureTransformer{}
		fuzz.New().NilChance(0).NumElements(100, 100).Fuzz(&tr)

		features := tr.TransformAllParallel(s, 1)
		assert.Equal(t, len(s)*tr.NumFeatures(), len(features))
	})

	t.Run("transform all parallel 4 workers", func(t *testing.T) {
		s := make([]With32Fields, 100)
		fuzz.New().NilChance(0).NumElements(100, 100).Fuzz(&s)

		tr := With32FieldsFeatureTransformer{}
		fuzz.New().NilChance(0).NumElements(100, 100).Fuzz(&tr)

		features := tr.TransformAllParallel(s, 4)
		assert.Equal(t, len(s)*tr.NumFeatures(), len(features))
	})
}

func TestWith32FieldsFeatureTransformerFit(t *testing.T) {
	t.Run("fuzzy input", func(t *testing.T) {
		s := make([]With32Fields, 10)
		fuzz.New().NilChance(0).NumElements(1, 1).Fuzz(&s)

		trEmpty := With32FieldsFeatureTransformer{}
		tr := With32FieldsFeatureTransformer{}
		tr.Fit(s)

		assert.NotNil(t, tr)
		assert.NotEqual(t, tr, trEmpty)
	})

	t.Run("not nil transformer nil input", func(t *testing.T) {
		trEmpty := With32FieldsFeatureTransformer{}
		tr := With32FieldsFeatureTransformer{}
		tr.Fit(nil)

		assert.Equal(t, trEmpty, tr)
	})

	t.Run("nil transformer not nil input", func(t *testing.T) {
		s := make([]With32Fields, 10)

		var tr *With32FieldsFeatureTransformer
		tr.Fit(s)

		assert.Nil(t, tr)
	})
}

func fitTransformerWith32Fields(b *testing.B, numelem int) {
	s := make([]With32Fields, numelem)
	fuzz.New().NilChance(0).NumElements(numelem, numelem).Fuzz(&s)

	var tr With32FieldsFeatureTransformer

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		tr.Fit(s)
	}
}

func BenchmarkWith32FieldsFeatureTransformer_Fit_100elements(b *testing.B) {
	fitTransformerWith32Fields(b, 100)
}

func BenchmarkWith32FieldsFeatureTransformer_Fit_1000elements(b *testing.B) {
	fitTransformerWith32Fields(b, 1000)
}

func BenchmarkWith32FieldsFeatureTransformer_Fit_10000elements(b *testing.B) {
	fitTransformerWith32Fields(b, 10000)
}

func BenchmarkWith32FieldsFeatureTransformer_Transform(b *testing.B) {
	var s With32Fields
	fuzz.New().Fuzz(&s)

	tr := With32FieldsFeatureTransformer{}
	fuzz.New().NilChance(0).NumElements(1, 1).Fuzz(&tr)

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		tr.Transform(&s)
	}
}

func BenchmarkWith32FieldsFeatureTransformer_Transform_Inplace(b *testing.B) {
	var s With32Fields
	fuzz.New().Fuzz(&s)

	tr := With32FieldsFeatureTransformer{}
	fuzz.New().NilChance(0).NumElements(1, 1).Fuzz(&tr)

	features := make([]float64, tr.NumFeatures())

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		tr.TransformInplace(features, &s)
	}
}

func benchTransformAllWith32Fields(b *testing.B, numelem int) {
	s := make([]With32Fields, numelem)
	fuzz.New().NilChance(0).NumElements(numelem, numelem).Fuzz(&s)

	tr := With32FieldsFeatureTransformer{}
	fuzz.New().NilChance(0).NumElements(100, 100).Fuzz(&tr)

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		tr.TransformAll(s)
	}
}

func BenchmarkWith32FieldsFeatureTransformer_TransformAll_100elems(b *testing.B) {
	benchTransformAllWith32Fields(b, 100)
}

func BenchmarkWith32FieldsFeatureTransformer_TransformAll_1000elems(b *testing.B) {
	benchTransformAllWith32Fields(b, 1000)
}

func BenchmarkWith32FieldsFeatureTransformer_TransformAll_10000elems(b *testing.B) {
	benchTransformAllWith32Fields(b, 10000)
}

func BenchmarkWith32FieldsFeatureTransformer_TransformAll_100000elems(b *testing.B) {
	benchTransformAllWith32Fields(b, 100000)
}

func BenchmarkWith32FieldsFeatureTransformer_TransformAll_1000000elems(b *testing.B) {
	benchTransformAllWith32Fields(b, 1000000)
}

func benchTransformAllParallelWith32Fields(b *testing.B, numelem int, nworkers uint) {
	s := make([]With32Fields, numelem)
	fuzz.New().NilChance(0).NumElements(numelem, numelem).Fuzz(&s)

	tr := With32FieldsFeatureTransformer{}
	fuzz.New().NilChance(0).NumElements(100, 100).Fuzz(&tr)

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		tr.TransformAllParallel(s, nworkers)
	}
}

func BenchmarkWith32FieldsFeatureTransformer_TransformAll_100elems_8workers(b *testing.B) {
	benchTransformAllParallelWith32Fields(b, 100, 8)
}

func BenchmarkWith32FieldsFeatureTransformer_TransformAll_1000elems_8workers(b *testing.B) {
	benchTransformAllParallelWith32Fields(b, 1000, 8)
}

func BenchmarkWith32FieldsFeatureTransformer_TransformAll_10000elems_8workers(b *testing.B) {
	benchTransformAllParallelWith32Fields(b, 10000, 8)
}

func BenchmarkWith32FieldsFeatureTransformer_TransformAll_100000elems_8workers(b *testing.B) {
	benchTransformAllParallelWith32Fields(b, 100000, 8)
}

func BenchmarkWith32FieldsFeatureTransformer_TransformAll_1000000elems_8workers(b *testing.B) {
	benchTransformAllParallelWith32Fields(b, 1000000, 8)
}
