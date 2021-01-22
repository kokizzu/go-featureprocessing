// Code generated by go-featureprocessing DO NOT EDIT

package examplemodule

import (
	"sync"

	fp "github.com/nikolaydubina/go-featureprocessing/transformers"
)

// With32FieldsFeatureTransformer is a feature processor for With32Fields.
// It was automatically generated by go-featureprocessing tool.
type With32FieldsFeatureTransformer struct {
	Name1  fp.MinMaxScaler `json:"Name1_minmax"`
	Name2  fp.MinMaxScaler `json:"Name2_minmax"`
	Name3  fp.MinMaxScaler `json:"Name3_minmax"`
	Name4  fp.MinMaxScaler `json:"Name4_minmax"`
	Name5  fp.MinMaxScaler `json:"Name5_minmax"`
	Name6  fp.MinMaxScaler `json:"Name6_minmax"`
	Name7  fp.MinMaxScaler `json:"Name7_minmax"`
	Name8  fp.MinMaxScaler `json:"Name8_minmax"`
	Name9  fp.MinMaxScaler `json:"Name9_minmax"`
	Name10 fp.MinMaxScaler `json:"Name10_minmax"`
	Name11 fp.MinMaxScaler `json:"Name11_minmax"`
	Name12 fp.MinMaxScaler `json:"Name12_minmax"`
	Name13 fp.MinMaxScaler `json:"Name13_minmax"`
	Name14 fp.MinMaxScaler `json:"Name14_minmax"`
	Name15 fp.MinMaxScaler `json:"Name15_minmax"`
	Name16 fp.MinMaxScaler `json:"Name16_minmax"`
	Name17 fp.MinMaxScaler `json:"Name17_minmax"`
	Name18 fp.MinMaxScaler `json:"Name18_minmax"`
	Name19 fp.MinMaxScaler `json:"Name19_minmax"`
	Name21 fp.MinMaxScaler `json:"Name21_minmax"`
	Name22 fp.MinMaxScaler `json:"Name22_minmax"`
	Name23 fp.MinMaxScaler `json:"Name23_minmax"`
	Name24 fp.MinMaxScaler `json:"Name24_minmax"`
	Name25 fp.MinMaxScaler `json:"Name25_minmax"`
	Name26 fp.MinMaxScaler `json:"Name26_minmax"`
	Name27 fp.MinMaxScaler `json:"Name27_minmax"`
	Name28 fp.MinMaxScaler `json:"Name28_minmax"`
	Name29 fp.MinMaxScaler `json:"Name29_minmax"`
	Name30 fp.MinMaxScaler `json:"Name30_minmax"`
	Name31 fp.MinMaxScaler `json:"Name31_minmax"`
	Name32 fp.MinMaxScaler `json:"Name32_minmax"`
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
	features := make([]float64, e.NumFeatures())
	e.TransformInplace(features, s)
	return features
}

// TransformInplace transforms struct into feature vector accordingly to transformers, and does so inplace
func (e *With32FieldsFeatureTransformer) TransformInplace(dst []float64, s *With32Fields) {
	if s == nil || e == nil || len(dst) != e.NumFeatures() {
		return
	}
	idx := 0

	dst[idx] = e.Name1.Transform(float64(s.Name1))
	idx++

	dst[idx] = e.Name2.Transform(float64(s.Name2))
	idx++

	dst[idx] = e.Name3.Transform(float64(s.Name3))
	idx++

	dst[idx] = e.Name4.Transform(float64(s.Name4))
	idx++

	dst[idx] = e.Name5.Transform(float64(s.Name5))
	idx++

	dst[idx] = e.Name6.Transform(float64(s.Name6))
	idx++

	dst[idx] = e.Name7.Transform(float64(s.Name7))
	idx++

	dst[idx] = e.Name8.Transform(float64(s.Name8))
	idx++

	dst[idx] = e.Name9.Transform(float64(s.Name9))
	idx++

	dst[idx] = e.Name10.Transform(float64(s.Name10))
	idx++

	dst[idx] = e.Name11.Transform(float64(s.Name11))
	idx++

	dst[idx] = e.Name12.Transform(float64(s.Name12))
	idx++

	dst[idx] = e.Name13.Transform(float64(s.Name13))
	idx++

	dst[idx] = e.Name14.Transform(float64(s.Name14))
	idx++

	dst[idx] = e.Name15.Transform(float64(s.Name15))
	idx++

	dst[idx] = e.Name16.Transform(float64(s.Name16))
	idx++

	dst[idx] = e.Name17.Transform(float64(s.Name17))
	idx++

	dst[idx] = e.Name18.Transform(float64(s.Name18))
	idx++

	dst[idx] = e.Name19.Transform(float64(s.Name19))
	idx++

	dst[idx] = e.Name21.Transform(float64(s.Name21))
	idx++

	dst[idx] = e.Name22.Transform(float64(s.Name22))
	idx++

	dst[idx] = e.Name23.Transform(float64(s.Name23))
	idx++

	dst[idx] = e.Name24.Transform(float64(s.Name24))
	idx++

	dst[idx] = e.Name25.Transform(float64(s.Name25))
	idx++

	dst[idx] = e.Name26.Transform(float64(s.Name26))
	idx++

	dst[idx] = e.Name27.Transform(float64(s.Name27))
	idx++

	dst[idx] = e.Name28.Transform(float64(s.Name28))
	idx++

	dst[idx] = e.Name29.Transform(float64(s.Name29))
	idx++

	dst[idx] = e.Name30.Transform(float64(s.Name30))
	idx++

	dst[idx] = e.Name31.Transform(float64(s.Name31))
	idx++

	dst[idx] = e.Name32.Transform(float64(s.Name32))
	idx++

}

// TransformAll transforms a slice of With32Fields
func (e *With32FieldsFeatureTransformer) TransformAll(s []With32Fields) []float64 {
	if e == nil {
		return nil
	}
	features := make([]float64, len(s)*e.NumFeatures())
	e.TransformAllInplace(features, s)
	return features
}

// TransformAllInplace transforms a slice of With32Fields inplace
func (e *With32FieldsFeatureTransformer) TransformAllInplace(dst []float64, s []With32Fields) {
	if e == nil {
		return
	}
	n := e.NumFeatures()
	if len(dst) != n*len(s) {
		return
	}
	for i := range s {
		e.TransformInplace(dst[i*n:(i+1)*n], &s[i])
	}
}

// TransformAllParallel transforms a slice of With32Fields in parallel
func (e *With32FieldsFeatureTransformer) TransformAllParallel(s []With32Fields, nworkers uint) []float64 {
	if e == nil {
		return nil
	}
	features := make([]float64, len(s)*e.NumFeatures())
	e.TransformAllInplaceParallel(features, s, nworkers)
	return features
}

// TransformAllInplaceParallel transforms a slice of With32Fields inplace parallel
// Useful for very large slices.
func (e *With32FieldsFeatureTransformer) TransformAllInplaceParallel(dst []float64, s []With32Fields, nworkers uint) {
	if e == nil || nworkers == 0 {
		return
	}
	ns := uint(len(s))
	nf := uint(e.NumFeatures())
	if uint(len(dst)) != nf*ns {
		return
	}

	nbatch := ns / nworkers
	var wg sync.WaitGroup

	for i := uint(0); i < nworkers; i++ {
		wg.Add(1)
		go func(i uint) {
			defer wg.Done()
			iStart := nbatch * i
			iEnd := nbatch * (i + 1)
			if i == (nworkers - 1) {
				iEnd = ns
			}
			e.TransformAllInplace(dst[iStart*nf:iEnd*nf], s[iStart:iEnd])
		}(i)
	}

	wg.Wait()
}

// NumFeatures returns number of features in output feature vector
func (e *With32FieldsFeatureTransformer) NumFeatures() int {
	if e == nil {
		return 0
	}

	count := 31

	return count
}

// FeatureNames provides names of features that match output of transform
func (e *With32FieldsFeatureTransformer) FeatureNames() []string {
	if e == nil {
		return nil
	}

	idx := 0
	names := make([]string, e.NumFeatures())

	names[idx] = "Name1"
	idx++

	names[idx] = "Name2"
	idx++

	names[idx] = "Name3"
	idx++

	names[idx] = "Name4"
	idx++

	names[idx] = "Name5"
	idx++

	names[idx] = "Name6"
	idx++

	names[idx] = "Name7"
	idx++

	names[idx] = "Name8"
	idx++

	names[idx] = "Name9"
	idx++

	names[idx] = "Name10"
	idx++

	names[idx] = "Name11"
	idx++

	names[idx] = "Name12"
	idx++

	names[idx] = "Name13"
	idx++

	names[idx] = "Name14"
	idx++

	names[idx] = "Name15"
	idx++

	names[idx] = "Name16"
	idx++

	names[idx] = "Name17"
	idx++

	names[idx] = "Name18"
	idx++

	names[idx] = "Name19"
	idx++

	names[idx] = "Name21"
	idx++

	names[idx] = "Name22"
	idx++

	names[idx] = "Name23"
	idx++

	names[idx] = "Name24"
	idx++

	names[idx] = "Name25"
	idx++

	names[idx] = "Name26"
	idx++

	names[idx] = "Name27"
	idx++

	names[idx] = "Name28"
	idx++

	names[idx] = "Name29"
	idx++

	names[idx] = "Name30"
	idx++

	names[idx] = "Name31"
	idx++

	names[idx] = "Name32"
	idx++

	return names
}
