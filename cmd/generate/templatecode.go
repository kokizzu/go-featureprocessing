package main

const templateCode = `
// Code generated by go-featureprocessing DO NOT EDIT

package {{$.PackageName}}

import (
	fp "github.com/nikolaydubina/go-featureprocessing/transformers"
)

// {{$.StructName}}FeatureTransformer is a feature processor for {{$.StructName}}.
// It was automatically generated by go-featureprocessing tool.
type {{$.StructName}}FeatureTransformer struct {
	{{range $i, $tr := $.Fields}}{{$tr.Name}} fp.{{$tr.Transformer}} ` + "`" + `json:"{{$tr.Name}}_{{$tr.TransformerTag}}"` + "`" + ` 
	{{end}}
}

// Fit fits transformer for each field
func (e *{{$.StructName}}FeatureTransformer) Fit(s []{{$.StructName}}) {
	if e == nil || len(s) == 0 {
		return
	}

	{{if $.HasNumericalTransformers}}dataNum := make([]float64, len(s)){{end}}
	{{if $.HasStringTransformers}}dataStr := make([]string, len(s)){{end}}

	{{range $i, $tr := $.Fields}}

	for i, v := range s {
		{{if $tr.NumericalInput }}dataNum[i] = float64(v.{{$tr.Name}}){{else}}dataStr[i] = string(v.{{$tr.Name}}){{end}}
	}

	e.{{$tr.Name}}.Fit({{if $tr.NumericalInput }}dataNum{{else}}dataStr{{end}})
	
	{{end}}
}

// Transform transforms struct into feature vector accordingly to transformers
func (e *{{$.StructName}}FeatureTransformer) Transform(s *{{$.StructName}}) []float64 {
	if s == nil || e == nil {
		return nil
	}

	features := make([]float64, e.NumFeatures())
	e.TransformInplace(features, s)
	return features
}

// TransformInplace transforms struct into feature vector accordingly to transformers, and does so inplace
func (e *{{$.StructName}}FeatureTransformer) TransformInplace(dst []float64, s *{{$.StructName}}) {
	if s == nil || e == nil || len(dst) != e.NumFeatures() {
		return
	}

	idx := 0
	{{range $i, $tr := $.Fields}}

	{{if $tr.Expanding }}
	e.{{$tr.Name}}.TransformInplace(dst[idx:idx + e.{{$tr.Name}}.NumFeatures()], s.{{$tr.Name}})
	idx += e.{{$tr.Name}}.NumFeatures()
	{{else}}
	dst[idx] = e.{{$tr.Name}}.Transform( {{if $tr.NumericalInput }}float64{{else}}string{{end}}( s.{{$tr.Name}} ))
	idx++
	{{end}}

	{{end}}
	return
}

// NumFeatures returns number of features in output feature vector
func (e *{{$.StructName}}FeatureTransformer) NumFeatures() int {
	if e == nil {
		return 0
	}

	count := {{$.NumFieldsFlat}}
	{{range $i, $tr := $.Fields}}{{if $tr.Expanding}}count += e.{{$tr.Name}}.NumFeatures(){{end}}
	{{end}}
	return count
}

// FeatureNames provides names of features that match output of transform
func (e *{{$.StructName}}FeatureTransformer) FeatureNames() []string {
	if e == nil {
		return nil
	}

	idx := 0
	names := make([]string, e.NumFeatures())

	{{range $i, $tr := $.Fields}}
	{{if $tr.Expanding }}
	for _, w := range e.{{$tr.Name}}.FeatureNames() {
		names[idx] = "{{$tr.Name}}_" + w
		idx++
	}
	{{else}}
	names[idx] = "{{$tr.Name}}"
	idx++
	{{end}}
	{{end}}

	return names
}
`
