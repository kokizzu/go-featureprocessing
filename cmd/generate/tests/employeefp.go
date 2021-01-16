// Code generated by go-featureprocessing DO NOT EDIT

package examplemodule

import (
	"sync"

	fp "github.com/nikolaydubina/go-featureprocessing/transformers"
)

// EmployeeFeatureTransformer is a feature processor for Employee.
// It was automatically generated by go-featureprocessing tool.
type EmployeeFeatureTransformer struct {
	Age         fp.Identity         `json:"Age_identity"`
	Salary      fp.MinMaxScaler     `json:"Salary_minmax"`
	Kids        fp.MaxAbsScaler     `json:"Kids_maxabs"`
	Weight      fp.StandardScaler   `json:"Weight_standard"`
	Height      fp.QuantileScaler   `json:"Height_quantile"`
	City        fp.OneHotEncoder    `json:"City_onehot"`
	Car         fp.OrdinalEncoder   `json:"Car_ordinal"`
	Income      fp.KBinsDiscretizer `json:"Income_kbins"`
	Description fp.TFIDFVectorizer  `json:"Description_tfidf"`
}

// Fit fits transformer for each field
func (e *EmployeeFeatureTransformer) Fit(s []Employee) {
	if e == nil || len(s) == 0 {
		return
	}

	dataNum := make([]float64, len(s))
	dataStr := make([]string, len(s))

	for i, v := range s {
		dataNum[i] = float64(v.Age)
	}

	e.Age.Fit(dataNum)

	for i, v := range s {
		dataNum[i] = float64(v.Salary)
	}

	e.Salary.Fit(dataNum)

	for i, v := range s {
		dataNum[i] = float64(v.Kids)
	}

	e.Kids.Fit(dataNum)

	for i, v := range s {
		dataNum[i] = float64(v.Weight)
	}

	e.Weight.Fit(dataNum)

	for i, v := range s {
		dataNum[i] = float64(v.Height)
	}

	e.Height.Fit(dataNum)

	for i, v := range s {
		dataStr[i] = v.City
	}

	e.City.Fit(dataStr)

	for i, v := range s {
		dataStr[i] = v.Car
	}

	e.Car.Fit(dataStr)

	for i, v := range s {
		dataNum[i] = float64(v.Income)
	}

	e.Income.Fit(dataNum)

	for i, v := range s {
		dataStr[i] = v.Description
	}

	e.Description.Fit(dataStr)

}

// Transform transforms struct into feature vector accordingly to transformers
func (e *EmployeeFeatureTransformer) Transform(s *Employee) []float64 {
	if s == nil || e == nil {
		return nil
	}
	features := make([]float64, e.NumFeatures())
	e.TransformInplace(features, s)
	return features
}

// TransformInplace transforms struct into feature vector accordingly to transformers, and does so inplace
func (e *EmployeeFeatureTransformer) TransformInplace(dst []float64, s *Employee) {
	if s == nil || e == nil || len(dst) != e.NumFeatures() {
		return
	}
	idx := 0

	dst[idx] = e.Age.Transform(float64(s.Age))
	idx++

	dst[idx] = e.Salary.Transform(float64(s.Salary))
	idx++

	dst[idx] = e.Kids.Transform(float64(s.Kids))
	idx++

	dst[idx] = e.Weight.Transform(float64(s.Weight))
	idx++

	dst[idx] = e.Height.Transform(float64(s.Height))
	idx++

	e.City.TransformInplace(dst[idx:idx+e.City.NumFeatures()], s.City)
	idx += e.City.NumFeatures()

	dst[idx] = e.Car.Transform((s.Car))
	idx++

	dst[idx] = e.Income.Transform(float64(s.Income))
	idx++

	e.Description.TransformInplace(dst[idx:idx+e.Description.NumFeatures()], s.Description)
	idx += e.Description.NumFeatures()

	return
}

// TransformAll transforms a slice of Employee
func (e *EmployeeFeatureTransformer) TransformAll(s []Employee) []float64 {
	if e == nil {
		return nil
	}
	features := make([]float64, len(s)*e.NumFeatures())
	e.TransformAllInplace(features, s)
	return features
}

// TransformAllInplace transforms a slice of Employee inplace
func (e *EmployeeFeatureTransformer) TransformAllInplace(dst []float64, s []Employee) {
	if e == nil {
		return
	}
	n := e.NumFeatures()
	if len(dst) != n*len(s) {
		return
	}
	for i, _ := range s {
		e.TransformInplace(dst[i*n:(i+1)*n], &s[i])
	}
}

// TransformAllParallel transforms a slice of Employee in parallel
func (e *EmployeeFeatureTransformer) TransformAllParallel(s []Employee, nworkers uint) []float64 {
	if e == nil {
		return nil
	}
	features := make([]float64, len(s)*e.NumFeatures())
	e.TransformAllInplaceParallel(features, s, nworkers)
	return features
}

// TransformAllInplaceParallel transforms a slice of Employee inplace parallel
// Useful for very large slices.
func (e *EmployeeFeatureTransformer) TransformAllInplaceParallel(dst []float64, s []Employee, nworkers uint) {
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
func (e *EmployeeFeatureTransformer) NumFeatures() int {
	if e == nil {
		return 0
	}

	count := 7

	count += e.City.NumFeatures()

	count += e.Description.NumFeatures()

	return count
}

// FeatureNames provides names of features that match output of transform
func (e *EmployeeFeatureTransformer) FeatureNames() []string {
	if e == nil {
		return nil
	}

	idx := 0
	names := make([]string, e.NumFeatures())

	names[idx] = "Age"
	idx++

	names[idx] = "Salary"
	idx++

	names[idx] = "Kids"
	idx++

	names[idx] = "Weight"
	idx++

	names[idx] = "Height"
	idx++

	for _, w := range e.City.FeatureNames() {
		names[idx] = "City_" + w
		idx++
	}

	names[idx] = "Car"
	idx++

	names[idx] = "Income"
	idx++

	for _, w := range e.Description.FeatureNames() {
		names[idx] = "Description_" + w
		idx++
	}

	return names
}
