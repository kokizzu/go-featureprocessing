// Code generated by go-featureprocessing DO NOT EDIT

package examplemodule

import (
	fp "github.com/nikolaydubina/go-featureprocessing/transformers"
)

// LargeMemoryTransformerFeatureTransformer is a feature processor for LargeMemoryTransformer.
// It was automatically generated by go-featureprocessing tool.
type LargeMemoryTransformerFeatureTransformer struct {
	Name1 fp.OneHotEncoder
	Name2 fp.OneHotEncoder
	Name3 fp.OrdinalEncoder
	Name4 fp.OrdinalEncoder
	Name5 fp.QuantileScaler
	Name6 fp.QuantileScaler
	Name7 fp.KBinsDiscretizer
	Name8 fp.KBinsDiscretizer
}

// Fit fits transformer for each field
func (e *LargeMemoryTransformerFeatureTransformer) Fit(s []LargeMemoryTransformer) {
	if e == nil || len(s) == 0 {
		return
	}

	dataNum := make([]float64, len(s))
	dataStr := make([]string, len(s))

	for i, v := range s {
		dataStr[i] = string(v.Name1)
	}

	e.Name1.Fit(dataStr)

	for i, v := range s {
		dataStr[i] = string(v.Name2)
	}

	e.Name2.Fit(dataStr)

	for i, v := range s {
		dataStr[i] = string(v.Name3)
	}

	e.Name3.Fit(dataStr)

	for i, v := range s {
		dataStr[i] = string(v.Name4)
	}

	e.Name4.Fit(dataStr)

	for i, v := range s {
		dataNum[i] = float64(v.Name5)
	}

	e.Name5.Fit(dataNum)

	for i, v := range s {
		dataNum[i] = float64(v.Name6)
	}

	e.Name6.Fit(dataNum)

	for i, v := range s {
		dataNum[i] = float64(v.Name7)
	}

	e.Name7.Fit(dataNum)

	for i, v := range s {
		dataNum[i] = float64(v.Name8)
	}

	e.Name8.Fit(dataNum)

}

// Transform transforms struct into feature vector accordingly to transformers
func (e *LargeMemoryTransformerFeatureTransformer) Transform(s *LargeMemoryTransformer) []float64 {
	if s == nil || e == nil {
		return nil
	}

	features := make([]float64, 0, e.GetNumFeatures())

	return e.transform(features, s)
}

// transform is utilizing mid-stack inliner, the idea is that publicly exported function will be inlined,
// meaning final features slice will not escape to heap
func (e *LargeMemoryTransformerFeatureTransformer) transform(features []float64, s *LargeMemoryTransformer) []float64 {
	features = append(features, e.Name1.Transform(string(s.Name1))...)
	features = append(features, e.Name2.Transform(string(s.Name2))...)
	features = append(features, e.Name3.Transform(string(s.Name3)))
	features = append(features, e.Name4.Transform(string(s.Name4)))
	features = append(features, e.Name5.Transform(float64(s.Name5)))
	features = append(features, e.Name6.Transform(float64(s.Name6)))
	features = append(features, e.Name7.Transform(float64(s.Name7)))
	features = append(features, e.Name8.Transform(float64(s.Name8)))

	return features
}

// GetNumFeatures returns number of features in output feature vector
func (e *LargeMemoryTransformerFeatureTransformer) GetNumFeatures() int {
	if e == nil {
		return 0
	}

	count := 6
	count += e.Name1.NumFeatures()
	count += e.Name2.NumFeatures()

	return count
}
