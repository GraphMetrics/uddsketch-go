// Unless explicitly stated otherwise all files in this repository are licensed
// under the BSD-3-Clause License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2018 Datadog, Inc.

package dataset

import (
	"math"
	"sort"
)

type Dataset struct {
	Values []float64
	Count  int64
	sorted bool
}

func NewDataset() *Dataset { return &Dataset{} }

func (d *Dataset) Add(v float64) {
	d.Values = append(d.Values, v)
	d.Count++
	d.sorted = false
}

// Quantile returns the lower quantile of the dataset
func (d *Dataset) Quantile(q float64) float64 {
	if q < 0 || q > 1 || d.Count == 0 {
		return math.NaN()
	}

	d.Sort()
	rank := q * float64(d.Count-1)
	return d.Values[int64(rank)]
}

func (d *Dataset) Rank(v float64) int64 {
	d.Sort()
	i := int64(0)
	for ; i < d.Count; i++ {
		if d.Values[i] >= v {
			return i + 1
		}
	}
	return d.Count
}

func (d *Dataset) Min() float64 {
	d.Sort()
	return d.Values[0]
}

func (d *Dataset) Max() float64 {
	d.Sort()
	return d.Values[len(d.Values)-1]
}

func (d *Dataset) Sum() float64 {
	s := float64(0)
	for _, v := range d.Values {
		s += v
	}
	return s
}

func (d *Dataset) Avg() float64 {
	return d.Sum() / float64(d.Count)
}

func (d *Dataset) Sort() {
	if d.sorted {
		return
	}
	sort.Float64s(d.Values)
	d.sorted = true
}
