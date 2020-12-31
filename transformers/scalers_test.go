package transformers_test

import (
	"testing"

	. "github.com/nikolaydubina/go-featureprocessing/transformers"
	"github.com/stretchr/testify/assert"
)

func TestIdentity(t *testing.T) {
	samples := []struct {
		name   string
		input  float64
		output float64
	}{
		{"basic", 42, 42},
		{"negative", -42, -42},
		{"zero", 0, 0},
		{"fraction", 0.5, 0.5},
	}
	for _, s := range samples {
		t.Run(s.name, func(t *testing.T) {
			encoder := Identity{}
			features := encoder.Transform((s.input))
			assert.Equal(t, s.output, features)
		})
	}

	t.Run("fit", func(t *testing.T) {
		encoder := Identity{}
		encoder.Fit(nil)
		assert.Equal(t, Identity{}, encoder)
	})
}

func TestMinMaxScalerTransform(t *testing.T) {
	samples := []struct {
		name   string
		min    float64
		max    float64
		input  float64
		output float64
	}{
		{"basic", 1, 101, 51, 0.5},
		{"basic", 1, 101, 71, 0.7},
		{"bellow", 1, 101, 0.5, 0},
		{"above", 1, 101, 102, 1},
		{"negative", 1, 101, -1, 0},
		{"zero", 1, 101, 0, 0},
		{"same1", 1, 1, 1, 0},
		{"same2", 1, 1, 0.5, 0},
		{"same2", 1, 1, 2, 0},
	}
	for _, s := range samples {
		t.Run(s.name, func(t *testing.T) {
			encoder := MinMaxScaler{Min: s.min, Max: s.max}
			features := encoder.Transform((s.input))
			assert.Equal(t, s.output, features)
		})
	}
}

func TestMinMaxScalerFit(t *testing.T) {
	samples := []struct {
		name string
		min  float64
		max  float64
		vals []float64
	}{
		{"noinput", 0, 0, nil},
		{"basic", 1, 101, []float64{1, 101}},
		{"negative_1", -1, 101, []float64{-1, 101}},
		{"negative_2", -10, -1, []float64{-10, -1}},
		{"zero", 0, 0, []float64{0, 0}},
		{"same", 1, 1, []float64{1, 1}},
		{"reverse_order", 1, 10, []float64{10, 1}},
		{"reverse_order_negative", -10, -1, []float64{-1, -10}},
	}
	for _, s := range samples {
		t.Run(s.name, func(t *testing.T) {
			encoder := MinMaxScaler{}
			encoder.Fit(s.vals)
			assert.Equal(t, MinMaxScaler{Min: s.min, Max: s.max}, encoder)
		})
	}
}

func TestMaxAbsScalerTransform(t *testing.T) {
	samples := []struct {
		name   string
		max    float64
		input  float64
		output float64
	}{
		{"basic", 100, 50, 0.5},
		{"basic", 100, 70, 0.7},
		{"above", 100, 102, 1},
		{"above_negative", 100, -102, -1},
		{"negative", 100, -50, -0.5},
		{"zero1", 100, 0, 0},
		{"zero2", 0, 0, 0},
	}
	for _, s := range samples {
		t.Run(s.name, func(t *testing.T) {
			encoder := MaxAbsScaler{Max: s.max}
			features := encoder.Transform((s.input))
			assert.Equal(t, s.output, features)
		})
	}
}

func TestMaxAbsScalerFit(t *testing.T) {
	samples := []struct {
		name string
		max  float64
		vals []float64
	}{
		{"noinput", 0, nil},
		{"basic", 100, []float64{1, 100}},
		{"negative", 100, []float64{-1, -100}},
		{"zero", 0, []float64{0, 0}},
		{"same", 1, []float64{1, 1}},
		{"reverse_order", 10, []float64{10, 1}},
		{"reverse_order_negative", 10, []float64{-1, -10}},
	}
	for _, s := range samples {
		t.Run(s.name, func(t *testing.T) {
			encoder := MaxAbsScaler{}
			encoder.Fit(s.vals)
			assert.Equal(t, MaxAbsScaler{Max: s.max}, encoder)
		})
	}
}

func TestStandardScalerTransform(t *testing.T) {
	samples := []struct {
		name   string
		mean   float64
		std    float64
		input  float64
		output float64
	}{
		{"basic_0", 100, 50, 100, 0},
		{"basic_-0.5", 100, 50, 75, -0.5},
		{"basic_0.5", 100, 50, 125, 0.5},
		{"basic_-1", 100, 50, 50, -1},
		{"basic_+1", 100, 50, 150, 1},
		{"basic_-2", 100, 50, 0, -2},
		{"basic_+2", 100, 50, 200, 2},
		{"basic_-3", 100, 50, -50, -3},
		{"basic_+3", 100, 50, 250, 3},
	}
	for _, s := range samples {
		t.Run(s.name, func(t *testing.T) {
			encoder := StandardScaler{Mean: s.mean, STD: s.std}
			assert.Equal(t, s.output, encoder.Transform(s.input))
		})
	}
}

func TestStandardScalerFit(t *testing.T) {
	samples := []struct {
		name string
		mean float64
		std  float64
		vals []float64
	}{
		{"noinput", 0, 0, nil},
		{"basic", 50.5, 70.0035713374682, []float64{1, 100}},
		{"negative", -50.5, 70.0035713374682, []float64{-1, -100}},
		{"zero", 0, 0, []float64{0, 0}},
		{"same", 1, 0, []float64{1, 1, 1, 1}},
	}
	for _, s := range samples {
		t.Run(s.name, func(t *testing.T) {
			encoder := StandardScaler{}
			encoder.Fit(s.vals)
			assert.Equal(t, StandardScaler{Mean: s.mean, STD: s.std}, encoder)
		})
	}
}

func TestQuantileScalerTransform(t *testing.T) {
	samples := []struct {
		name      string
		n         int
		quantiles []float64
		input     float64
		output    float64
	}{
		{"basic1", 4, []float64{25, 50, 75, 100}, 0, 0.25},
		{"basic2", 4, []float64{25, 50, 75, 100}, 11, 0.25},
		{"basic3", 4, []float64{25, 50, 75, 100}, 25, 0.25},
		{"basic4", 4, []float64{25, 50, 75, 100}, 40, 0.5},
		{"basic5", 4, []float64{25, 50, 75, 100}, 50, 0.5},
		{"basic6", 4, []float64{25, 50, 75, 100}, 80, 1},
		{"basic7", 4, []float64{25, 50, 75, 100}, 101, 1},
		{"empty", 0, nil, 10, 0},
	}
	for _, s := range samples {
		t.Run(s.name, func(t *testing.T) {
			encoder := QuantileScaler{NQuantiles: s.n, Quantiles: s.quantiles}
			features := encoder.Transform((s.input))
			assert.Equal(t, s.output, features)
		})
	}
}

func TestQuantileScalerFit(t *testing.T) {
	samples := []struct {
		name      string
		n         int
		quantiles []float64
		vals      []float64
	}{
		{"noinput", 1000, nil, nil},
		{"basic", 4, []float64{25, 50, 75, 100}, []float64{25, 50, 75, 100}},
		{"reverse_order", 4, []float64{25, 50, 75, 100}, []float64{100, 75, 50, 25}},
		{"negative", 4, []float64{-100, -75, -50, -25}, []float64{-25, -50, -75, -100}},
		{"one_element", 1, []float64{10}, []float64{10}},
		{"less_elements_than_quantiles", 3, []float64{1, 2, 3}, []float64{1, 2, 3}},
		{"less_elements_than_quantiles_negative", 3, []float64{-3, -2, -1}, []float64{-1, -3, -2}},
	}
	for _, s := range samples {
		t.Run(s.name, func(t *testing.T) {
			encoder := QuantileScaler{NQuantiles: s.n}
			encoder.Fit(s.vals)
			assert.Equal(t, QuantileScaler{NQuantiles: s.n, Quantiles: s.quantiles}, encoder)
		})
	}

	t.Run("nquantiles is larger than num input vals", func(t *testing.T) {
		encoder := QuantileScaler{NQuantiles: 10}
		encoder.Fit([]float64{1, 2, 3})
		assert.Equal(t, QuantileScaler{NQuantiles: 3, Quantiles: []float64{1, 2, 3}}, encoder)
	})

	t.Run("nquantiles is zero in beginning", func(t *testing.T) {
		encoder := QuantileScaler{}
		encoder.Fit(nil)
		assert.Equal(t, QuantileScaler{NQuantiles: 1000}, encoder)
	})
}
