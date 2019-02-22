package bamboo

import (
	"context"
	"testing"
)
var int_data  = []int {
	1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20,
	21,22,23,24,25,26,27,28,29,30,31,32,33,34,35,36,37,38,39,40,
	41,42,43,44,45,46,47,48,49,50,51,52,53,54,55,56,57,58,59,60,
	61,62,63,64,65,66,67,68,69,70,71,72,73,74,75,76,77,78,79,80,
	81,82,83,84,85,86,87,88,89,90,91,92,93,94,95,96,97,98,99,100,
}

func TestSeries_Lambda(t *testing.T) {
	tables := []struct{
		data_in []int
		data_out []int
		lambda func(ctx context.Context, column interface{})(column_out interface{}, override bool)
	} {
		{
			[]int{1,2,3},
			[]int{2,4,6},
			func(ctx context.Context, column_in interface{}) (column_out interface{}, override bool) {
				if val, ok := column_in.(int); ok {
					column_out = val * 2
				}

				return column_out, true
			},
		},

	}

	for _, table := range tables {
		var series = Series{}

		var err error
		if err = series.SetData(table.data_in); err == nil {
			series.Lambda(context.Background(), table.lambda)

			for index := range table.data_out {

				if value, ok := series.Get(index).(int); ok {
					if value != table.data_out[index] {
						t.Errorf("Got [%v], Expected [%v]\n", table.data_in[index], table.data_out[index])
					}
				}
			}
		} else {
			t.Error(err.Error())
		}
	}
}

func BenchmarkSeries_Lambda_IntAddition(b *testing.B) {

	var series = Series{}
	series.SetData(int_data)

	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		if _, err := series.Lambda(context.Background(), intAdditionLambda); err != nil {
			b.Error(err.Error())
		}
	}
}

func BenchmarkSeries_Lambda_IntMultiplier(b *testing.B) {
	var series = Series{}
	series.SetData(int_data)

	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		if _, err := series.Lambda(context.Background(), intMultiplicationLambda); err != nil {
			b.Error(err.Error())
		}
	}
}

func BenchmarkSeries_Lambda_IntDivision(b *testing.B) {

	var series = Series{}
	series.SetData(int_data)

	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		if _, err := series.Lambda(context.Background(), intDivisionLambda); err != nil {
			b.Error(err.Error())
		}
	}
}

func intAdditionLambda(ctx context.Context, column interface{}) (column_out interface{}, override bool)  {

	if val, ok := column.(int); ok {
		column_out = val + 2
	}

	return column_out, true
}

func intMultiplicationLambda(ctx context.Context, column interface{}) (column_out interface{}, override bool) {

	if val, ok := column.(int); ok {
		column_out = val * 2
	}

	return column_out, true
}

func intDivisionLambda(ctx context.Context, column interface{}) (column_out interface{}, override bool) {

	if val, ok := column.(int); ok {
		column_out = val / 2
	}

	return column_out, true
}