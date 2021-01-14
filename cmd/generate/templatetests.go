package main

const templateTests = `
// Code generated by go-featureprocessing DO NOT EDIT

package {{$.PackageName}}

import (
	"encoding/json"
	"testing"

	"github.com/google/gofuzz"
	"github.com/stretchr/testify/assert"
)

// makeMock creates some valid {{$.StructName}}FeatureTransformer by fitting on fuzzy data.
// This function is handy for tests.
func makeMock{{$.StructName}}FeatureTransformer() *{{$.StructName}}FeatureTransformer {
	s := make([]{{$.StructName}}, 1000000)
	fuzz.New().NilChance(0).NumElements(10, 10).Fuzz(&s)
	
	tr := {{$.StructName}}FeatureTransformer{}
	tr.Fit(s)
	return &tr
}

func Test{{$.StructName}}FeatureTransformerFeatureNames(t *testing.T) {
	validTransformer := makeMock{{$.StructName}}FeatureTransformer()

	fuzzyTransformer := {{$.StructName}}FeatureTransformer{} 
	fuzz.New().NilChance(0).NumElements(10, 10).Fuzz(&fuzzyTransformer)

	t.Run("feature names", func(t *testing.T) {
		names := validTransformer.FeatureNames()
		assert.True(t, len(names) > 0)
		assert.Equal(t, len(names), validTransformer.GetNumFeatures())
	})

	t.Run("feature names fuzzy transformer has some feature names", func(t *testing.T) {
		names := fuzzyTransformer.FeatureNames()
		assert.True(t, len(names) > 0)
	})

	t.Run("feature name transformer is empty", func(t *testing.T) {
		tr := {{$.StructName}}FeatureTransformer{}
		names := tr.FeatureNames()
		assert.True(t, len(names) > 0)
		assert.Equal(t, len(names), tr.GetNumFeatures())
	})

	t.Run("feature name transformer is nil", func(t *testing.T) {
		var tr *{{$.StructName}}FeatureTransformer
		names := tr.FeatureNames()
		assert.Nil(t, names)
	})
}

func Test{{$.StructName}}FeatureTransformerTransform(t *testing.T) {
	t.Run("empty struct fuzzy transformer", func(t *testing.T) {
		s := {{$.StructName}}{}
		
		tr := {{$.StructName}}FeatureTransformer{}
		fuzz.New().NilChance(0).NumElements(1, 1).Fuzz(&tr)

		tr.Transform(&s)
		
		features := tr.Transform(&s)

		assert.NotNil(t, features)
		assert.True(t, len(features) > 0)
		assert.Equal(t, tr.GetNumFeatures(), len(features))
	})

	t.Run("fuzzy struct", func(t *testing.T) {
		var s {{$.StructName}}
		fuzz.New().Fuzz(&s)
		
		tr := {{$.StructName}}FeatureTransformer{}
		fuzz.New().NilChance(0).NumElements(1, 1).Fuzz(&tr)

		features := tr.Transform(&s)

		assert.NotNil(t, features)
		assert.True(t, len(features) > 0)
		assert.Equal(t, tr.GetNumFeatures(), len(features))
	})

	t.Run("transformer is nil and struct is not nil", func(t *testing.T) {
		var s {{$.StructName}}
		fuzz.New().Fuzz(&s)
		
		var tr *{{$.StructName}}FeatureTransformer

		features := tr.Transform(&s)

		assert.Nil(t, features)
		assert.Equal(t, tr.GetNumFeatures(), 0)
	})

	t.Run("transformer is not nil but struct is nil", func(t *testing.T) {
		var s *{{$.StructName}}
		
		tr := {{$.StructName}}FeatureTransformer{}
		fuzz.New().NilChance(0).NumElements(1, 1).Fuzz(&tr)
		
		features := tr.Transform(s)
		
		assert.Nil(t, features)
		assert.True(t, tr.GetNumFeatures() > 0)
	})

	t.Run("serialize and deserialize transformer", func(t *testing.T) {
		tr := {{$.StructName}}FeatureTransformer{}
		fuzz.New().NilChance(0).NumElements(1, 1).Fuzz(&tr)

		output, err := json.Marshal(tr)
		assert.Nil(t, err)
		assert.NotEmpty(t, output)

		var tr2 {{$.StructName}}FeatureTransformer
		err = json.Unmarshal(output, &tr2)
		assert.Nil(t, err)
		assert.Equal(t, tr, tr2)
	})

	t.Run("inplace transform does not run when destination does not match num features", func(t *testing.T) {
		var s {{$.StructName}}
		fuzz.New().Fuzz(&s)
		
		tr := {{$.StructName}}FeatureTransformer{}
		fuzz.New().NilChance(0).NumElements(1, 1).Fuzz(&tr)

		features := make([]float64, 1000)
		features[0] = 123456789.0
		tr.TransformInplace(features, &s)

		assert.Equal(t, 123456789.0, features[0])
	})
}

func Test{{$.StructName}}FeatureTransformerFit(t *testing.T) {
	t.Run("fuzzy input", func(t *testing.T) {
		s := make([]{{$.StructName}}, 10)
		fuzz.New().NilChance(0).NumElements(1, 1).Fuzz(&s)
		
		trEmpty := {{$.StructName}}FeatureTransformer{}
		tr := {{$.StructName}}FeatureTransformer{}
		tr.Fit(s)

		assert.NotNil(t, tr)
		assert.NotEqual(t, tr, trEmpty)
	})

	t.Run("not nil transformer nil input", func(t *testing.T) {
		trEmpty := {{$.StructName}}FeatureTransformer{}
		tr := {{$.StructName}}FeatureTransformer{}
		tr.Fit(nil)

		assert.Equal(t, trEmpty, tr)
	})

	t.Run("nil transformer not nil input", func(t *testing.T) {
		s := make([]{{$.StructName}}, 10)
		
		var tr *{{$.StructName}}FeatureTransformer
		tr.Fit(s)

		assert.Nil(t, tr)
	})
}

func fitTransformer{{$.StructName}}(b *testing.B, numelem int) {
	s := make([]{{$.StructName}}, numelem)
	fuzz.New().NilChance(0).NumElements(numelem, numelem).Fuzz(&s)
	
	var tr {{$.StructName}}FeatureTransformer

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		tr.Fit(s)
	}
}

func Benchmark{{$.StructName}}FeatureTransformer_Fit_100elements(b *testing.B) {
	fitTransformer{{$.StructName}}(b, 100)
}

func Benchmark{{$.StructName}}FeatureTransformer_Fit_1000elements(b *testing.B) {
	fitTransformer{{$.StructName}}(b, 1000)
}

func Benchmark{{$.StructName}}FeatureTransformer_Fit_10000elements(b *testing.B) {
	fitTransformer{{$.StructName}}(b, 10000)
}

func Benchmark{{$.StructName}}FeatureTransformer_Transform(b *testing.B) {
	var s {{$.StructName}}
	fuzz.New().Fuzz(&s)
	
	tr := {{$.StructName}}FeatureTransformer{}
	fuzz.New().NilChance(0).NumElements(1, 1).Fuzz(&tr)

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		tr.Transform(&s)
	}
}

func Benchmark{{$.StructName}}FeatureTransformer_Transform_Inplace(b *testing.B) {
	var s {{$.StructName}}
	fuzz.New().Fuzz(&s)
	
	tr := {{$.StructName}}FeatureTransformer{}
	fuzz.New().NilChance(0).NumElements(1, 1).Fuzz(&tr)

	features := make([]float64, tr.GetNumFeatures())

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		tr.TransformInplace(features, &s)
	}
}

{{if $.HasLargeTransformers}}

func benchLargeTransformer{{$.StructName}}(b *testing.B, numelem int) {
	var s {{$.StructName}}
	fuzz.New().Fuzz(&s)
	
	tr := {{$.StructName}}FeatureTransformer{}
	fuzz.New().NilChance(0).NumElements(numelem, numelem).Fuzz(&tr)

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		tr.Transform(&s)
	}
}

func Benchmark{{$.StructName}}FeatureTransformer_Transform_LargeComposites_100elements(b *testing.B) {
	benchLargeTransformer{{$.StructName}}(b, 100)
}

func Benchmark{{$.StructName}}FeatureTransformer_Transform_LargeComposites_1000elements(b *testing.B) {
	benchLargeTransformer{{$.StructName}}(b, 1000)
}

func Benchmark{{$.StructName}}FeatureTransformer_Transform_LargeComposites_10000elements(b *testing.B) {
	benchLargeTransformer{{$.StructName}}(b, 10000)
}

func Benchmark{{$.StructName}}FeatureTransformer_Transform_LargeComposites_100000elements(b *testing.B) {
	benchLargeTransformer{{$.StructName}}(b, 100000)
}

{{end}}
`
