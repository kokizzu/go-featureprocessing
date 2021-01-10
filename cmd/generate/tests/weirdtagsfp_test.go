// Code generated by go-featureprocessing DO NOT EDIT

package examplemodule

import (
	"encoding/json"
	"testing"

	"github.com/google/gofuzz"
	"github.com/stretchr/testify/assert"
)

func TestWeirdTagsFeatureTransformerTransform(t *testing.T) {
	t.Run("empty struct", func(t *testing.T) {
		var s WeirdTags
		fuzz.New().Fuzz(&s)

		tr := WeirdTagsFeatureTransformer{}
		fuzz.New().NilChance(0).NumElements(1, 1).Fuzz(&tr)

		tr.Transform(&s)

		features := tr.Transform(&s)

		assert.NotNil(t, features)
		assert.True(t, len(features) > 0)
		assert.Equal(t, tr.GetNumFeatures(), len(features))
	})

	t.Run("fuzzy struct", func(t *testing.T) {
		var s WeirdTags
		fuzz.New().Fuzz(&s)

		tr := WeirdTagsFeatureTransformer{}
		fuzz.New().NilChance(0).NumElements(1, 1).Fuzz(&tr)

		features := tr.Transform(&s)

		assert.NotNil(t, features)
		assert.True(t, len(features) > 0)
		assert.Equal(t, tr.GetNumFeatures(), len(features))
	})

	t.Run("transformer is nil and struct is not nil", func(t *testing.T) {
		var s WeirdTags
		fuzz.New().Fuzz(&s)

		var tr *WeirdTagsFeatureTransformer

		features := tr.Transform(&s)

		assert.Nil(t, features)
		assert.Equal(t, tr.GetNumFeatures(), 0)
	})

	t.Run("transformer is not nil but struct is nil", func(t *testing.T) {
		var s *WeirdTags

		tr := WeirdTagsFeatureTransformer{}
		fuzz.New().NilChance(0).NumElements(1, 1).Fuzz(&tr)

		features := tr.Transform(s)

		assert.Nil(t, features)
		assert.True(t, tr.GetNumFeatures() > 0)
	})

	t.Run("serialize and deserialize transformer", func(t *testing.T) {
		tr := WeirdTagsFeatureTransformer{}
		fuzz.New().NilChance(0).NumElements(1, 1).Fuzz(&tr)

		output, err := json.Marshal(tr)
		assert.Nil(t, err)
		assert.NotEmpty(t, output)

		var tr2 WeirdTagsFeatureTransformer
		err = json.Unmarshal(output, &tr2)
		assert.Nil(t, err)
		assert.Equal(t, tr, tr2)
	})
}

func TestWeirdTagsFeatureTransformerFit(t *testing.T) {
	t.Run("fuzzy input", func(t *testing.T) {
		s := make([]WeirdTags, 10)
		fuzz.New().NilChance(0).NumElements(1, 1).Fuzz(&s)

		trEmpty := WeirdTagsFeatureTransformer{}
		tr := WeirdTagsFeatureTransformer{}
		tr.Fit(s)

		assert.NotNil(t, tr)
		assert.NotEqual(t, tr, trEmpty)
	})

	t.Run("not nil transformer nil input", func(t *testing.T) {
		trEmpty := WeirdTagsFeatureTransformer{}
		tr := WeirdTagsFeatureTransformer{}
		tr.Fit(nil)

		assert.Equal(t, trEmpty, tr)
	})

	t.Run("nil transformer not nil input", func(t *testing.T) {
		s := make([]WeirdTags, 10)

		var tr *WeirdTagsFeatureTransformer
		tr.Fit(s)

		assert.Nil(t, tr)
	})
}

func fitTransformerWeirdTags(b *testing.B, numelem int) {
	s := make([]WeirdTags, numelem)
	fuzz.New().NilChance(0).NumElements(numelem, numelem).Fuzz(&s)

	var tr WeirdTagsFeatureTransformer

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		tr.Fit(s)
	}
}

func BenchmarkWeirdTagsFeatureTransformer_Fit_100elements(b *testing.B) {
	fitTransformerWeirdTags(b, 100)
}

func BenchmarkWeirdTagsFeatureTransformer_Fit_1000elements(b *testing.B) {
	fitTransformerWeirdTags(b, 1000)
}

func BenchmarkWeirdTagsFeatureTransformer_Fit_10000elements(b *testing.B) {
	fitTransformerWeirdTags(b, 10000)
}

func BenchmarkWeirdTagsFeatureTransformer_Fit_100000elements(b *testing.B) {
	fitTransformerWeirdTags(b, 100000)
}

func BenchmarkWeirdTagsFeatureTransformer_Transform(b *testing.B) {
	var s WeirdTags
	fuzz.New().Fuzz(&s)

	tr := WeirdTagsFeatureTransformer{}
	fuzz.New().NilChance(0).NumElements(1, 1).Fuzz(&tr)

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		tr.Transform(&s)
	}
}

func benchLargeTransformerWeirdTags(b *testing.B, numelem int) {
	var s WeirdTags
	fuzz.New().Fuzz(&s)

	tr := WeirdTagsFeatureTransformer{}
	fuzz.New().NilChance(0).NumElements(numelem, numelem).Fuzz(&tr)

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		tr.Transform(&s)
	}
}

func BenchmarkWeirdTagsFeatureTransformer_Transform_LargeComposites_100elements(b *testing.B) {
	benchLargeTransformerWeirdTags(b, 100)
}

func BenchmarkWeirdTagsFeatureTransformer_Transform_LargeComposites_1000elements(b *testing.B) {
	benchLargeTransformerWeirdTags(b, 1000)
}

func BenchmarkWeirdTagsFeatureTransformer_Transform_LargeComposites_10000elements(b *testing.B) {
	benchLargeTransformerWeirdTags(b, 10000)
}

func BenchmarkWeirdTagsFeatureTransformer_Transform_LargeComposites_100000elements(b *testing.B) {
	benchLargeTransformerWeirdTags(b, 100000)
}
