package multi

import (
	"github.com/lazychanger/go-contrib/app"
	"reflect"
	"sort"
	"testing"
)

func Test_getSort(t *testing.T) {

	tests := []struct {
		name string
		args app.Application
		want int
	}{
		{
			name: "application",
			args: &app.TestApplication{},
			want: 0,
		},
		{
			name: "sortApplication",
			args: &TestSortApplication{
				SortVal: 3,
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getSort(tt.args); got != tt.want {
				t.Errorf("getSort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sortApplications_Applications(t *testing.T) {

	var (
		app3, app1, app10, app0 = &TestSortApplication{SortVal: 3},
			&TestSortApplication{SortVal: 1},
			&TestSortApplication{SortVal: 10},
			&app.TestApplication{}
	)

	tests := []struct {
		name   string
		fields []app.Application
		want   []app.Application
	}{
		{
			name:   "sort applications-(auto sort)",
			fields: []app.Application{app3, app0, app1, app10},
			want:   []app.Application{app10, app3, app1, app0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &sortApplications{
				apps: tt.fields,
			}
			if got := s.Applications(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Applications() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkMulti_Sort(b *testing.B) {

	var (
		apps = sortApplications{apps: []app.Application{
			&TestSortApplication{SortVal: 3},
			&TestSortApplication{SortVal: 1},
			&TestSortApplication{SortVal: 4},
			&TestSortApplication{SortVal: 9},
			&app.TestApplication{},
			&TestSortApplication{SortVal: 2},
			&TestSortApplication{SortVal: 3},
			&TestSortApplication{SortVal: 20},
			&TestSortApplication{SortVal: 12},
			&TestSortApplication{SortVal: 93},
			&TestSortApplication{SortVal: 14},
			&TestSortApplication{SortVal: 25},
			&TestSortApplication{SortVal: 21},
			&TestSortApplication{SortVal: 11},
			&app.TestApplication{},
		}}
		apps2 = make([]app.Application, apps.Len())
	)

	copy(apps2, apps.apps)

	b.Run("sort.Sort()", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			apps.Applications()
		}
	})

	b.Run("sort.Slice()", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			sort.Slice(apps2, func(i, j int) bool {
				return getSort(apps2[i]) > getSort(apps2[j])
			})
		}
	})
}
