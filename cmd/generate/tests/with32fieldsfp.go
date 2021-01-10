// Code generated by go-featureprocessing DO NOT EDIT

package examplemodule

import (
	fp "github.com/nikolaydubina/go-featureprocessing/transformers"
)

// With32FieldsFeatureTransformer is a feature processor for With32Fields.
// It was automatically generated by go-featureprocessing tool.
type With32FieldsFeatureTransformer struct {
	Name1  fp.MinMaxScaler
	Name2  fp.MinMaxScaler
	Name3  fp.MinMaxScaler
	Name4  fp.MinMaxScaler
	Name5  fp.MinMaxScaler
	Name6  fp.MinMaxScaler
	Name7  fp.MinMaxScaler
	Name8  fp.MinMaxScaler
	Name9  fp.MinMaxScaler
	Name10 fp.MinMaxScaler
	Name11 fp.MinMaxScaler
	Name12 fp.MinMaxScaler
	Name13 fp.MinMaxScaler
	Name14 fp.MinMaxScaler
	Name15 fp.MinMaxScaler
	Name16 fp.MinMaxScaler
	Name17 fp.MinMaxScaler
	Name18 fp.MinMaxScaler
	Name19 fp.MinMaxScaler
	Name21 fp.MinMaxScaler
	Name22 fp.MinMaxScaler
	Name23 fp.MinMaxScaler
	Name24 fp.MinMaxScaler
	Name25 fp.MinMaxScaler
	Name26 fp.MinMaxScaler
	Name27 fp.MinMaxScaler
	Name28 fp.MinMaxScaler
	Name29 fp.MinMaxScaler
	Name30 fp.MinMaxScaler
	Name31 fp.MinMaxScaler
	Name32 fp.MinMaxScaler
}

// Fit fits transformer for each field
func (e *With32FieldsFeatureTransformer) Fit(s []With32Fields) {
	if e == nil || len(s) == 0 {
		return
	}

	dataNum := make([]float64, len(s))

	for i, v := range s {
		dataNum[i] = float64(v.Name1)
	}

	e.Name1.Fit(dataNum)

	for i, v := range s {
		dataNum[i] = float64(v.Name2)
	}

	e.Name2.Fit(dataNum)

	for i, v := range s {
		dataNum[i] = float64(v.Name3)
	}

	e.Name3.Fit(dataNum)

	for i, v := range s {
		dataNum[i] = float64(v.Name4)
	}

	e.Name4.Fit(dataNum)

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

	for i, v := range s {
		dataNum[i] = float64(v.Name9)
	}

	e.Name9.Fit(dataNum)

	for i, v := range s {
		dataNum[i] = float64(v.Name10)
	}

	e.Name10.Fit(dataNum)

	for i, v := range s {
		dataNum[i] = float64(v.Name11)
	}

	e.Name11.Fit(dataNum)

	for i, v := range s {
		dataNum[i] = float64(v.Name12)
	}

	e.Name12.Fit(dataNum)

	for i, v := range s {
		dataNum[i] = float64(v.Name13)
	}

	e.Name13.Fit(dataNum)

	for i, v := range s {
		dataNum[i] = float64(v.Name14)
	}

	e.Name14.Fit(dataNum)

	for i, v := range s {
		dataNum[i] = float64(v.Name15)
	}

	e.Name15.Fit(dataNum)

	for i, v := range s {
		dataNum[i] = float64(v.Name16)
	}

	e.Name16.Fit(dataNum)

	for i, v := range s {
		dataNum[i] = float64(v.Name17)
	}

	e.Name17.Fit(dataNum)

	for i, v := range s {
		dataNum[i] = float64(v.Name18)
	}

	e.Name18.Fit(dataNum)

	for i, v := range s {
		dataNum[i] = float64(v.Name19)
	}

	e.Name19.Fit(dataNum)

	for i, v := range s {
		dataNum[i] = float64(v.Name21)
	}

	e.Name21.Fit(dataNum)

	for i, v := range s {
		dataNum[i] = float64(v.Name22)
	}

	e.Name22.Fit(dataNum)

	for i, v := range s {
		dataNum[i] = float64(v.Name23)
	}

	e.Name23.Fit(dataNum)

	for i, v := range s {
		dataNum[i] = float64(v.Name24)
	}

	e.Name24.Fit(dataNum)

	for i, v := range s {
		dataNum[i] = float64(v.Name25)
	}

	e.Name25.Fit(dataNum)

	for i, v := range s {
		dataNum[i] = float64(v.Name26)
	}

	e.Name26.Fit(dataNum)

	for i, v := range s {
		dataNum[i] = float64(v.Name27)
	}

	e.Name27.Fit(dataNum)

	for i, v := range s {
		dataNum[i] = float64(v.Name28)
	}

	e.Name28.Fit(dataNum)

	for i, v := range s {
		dataNum[i] = float64(v.Name29)
	}

	e.Name29.Fit(dataNum)

	for i, v := range s {
		dataNum[i] = float64(v.Name30)
	}

	e.Name30.Fit(dataNum)

	for i, v := range s {
		dataNum[i] = float64(v.Name31)
	}

	e.Name31.Fit(dataNum)

	for i, v := range s {
		dataNum[i] = float64(v.Name32)
	}

	e.Name32.Fit(dataNum)

}

// Transform transforms struct into feature vector accordingly to transformers
func (e *With32FieldsFeatureTransformer) Transform(s *With32Fields) []float64 {
	if s == nil || e == nil {
		return nil
	}

	features := make([]float64, 0, e.GetNumFeatures())

	return e.transform(features, s)
}

// transform is utilizing mid-stack inliner, the idea is that publicly exported function will be inlined,
// meaning final features slice will not escape to heap
func (e *With32FieldsFeatureTransformer) transform(features []float64, s *With32Fields) []float64 {
	features = append(features, e.Name1.Transform(float64(s.Name1)))
	features = append(features, e.Name2.Transform(float64(s.Name2)))
	features = append(features, e.Name3.Transform(float64(s.Name3)))
	features = append(features, e.Name4.Transform(float64(s.Name4)))
	features = append(features, e.Name5.Transform(float64(s.Name5)))
	features = append(features, e.Name6.Transform(float64(s.Name6)))
	features = append(features, e.Name7.Transform(float64(s.Name7)))
	features = append(features, e.Name8.Transform(float64(s.Name8)))
	features = append(features, e.Name9.Transform(float64(s.Name9)))
	features = append(features, e.Name10.Transform(float64(s.Name10)))
	features = append(features, e.Name11.Transform(float64(s.Name11)))
	features = append(features, e.Name12.Transform(float64(s.Name12)))
	features = append(features, e.Name13.Transform(float64(s.Name13)))
	features = append(features, e.Name14.Transform(float64(s.Name14)))
	features = append(features, e.Name15.Transform(float64(s.Name15)))
	features = append(features, e.Name16.Transform(float64(s.Name16)))
	features = append(features, e.Name17.Transform(float64(s.Name17)))
	features = append(features, e.Name18.Transform(float64(s.Name18)))
	features = append(features, e.Name19.Transform(float64(s.Name19)))
	features = append(features, e.Name21.Transform(float64(s.Name21)))
	features = append(features, e.Name22.Transform(float64(s.Name22)))
	features = append(features, e.Name23.Transform(float64(s.Name23)))
	features = append(features, e.Name24.Transform(float64(s.Name24)))
	features = append(features, e.Name25.Transform(float64(s.Name25)))
	features = append(features, e.Name26.Transform(float64(s.Name26)))
	features = append(features, e.Name27.Transform(float64(s.Name27)))
	features = append(features, e.Name28.Transform(float64(s.Name28)))
	features = append(features, e.Name29.Transform(float64(s.Name29)))
	features = append(features, e.Name30.Transform(float64(s.Name30)))
	features = append(features, e.Name31.Transform(float64(s.Name31)))
	features = append(features, e.Name32.Transform(float64(s.Name32)))

	return features
}

// GetNumFeatures returns number of features in output feature vector
func (e *With32FieldsFeatureTransformer) GetNumFeatures() int {
	if e == nil {
		return 0
	}

	count := 31

	return count
}
