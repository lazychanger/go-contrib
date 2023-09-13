package multi

import (
	"github.com/lazychanger/go-contrib/app"
	"sort"
)

type SortApplication interface {
	app.Application

	Sort() int
}

type sortApplications struct {
	apps   []app.Application
	sorted bool
}

func (s *sortApplications) Len() int {
	return len(s.apps)
}

func (s *sortApplications) Less(i, j int) bool {
	return getSort(s.apps[i]) > getSort(s.apps[j])
}

func (s *sortApplications) Swap(i, j int) {
	s.apps[i], s.apps[j] = s.apps[j], s.apps[i]
}

func (s *sortApplications) Applications() []app.Application {
	if !s.sorted {
		sort.Sort(s)
		s.sorted = true
	}

	return s.apps
}

func getSort(app app.Application) int {
	if sortApp, ok := app.(SortApplication); ok {
		return sortApp.Sort()
	} else {
		return 0
	}
}

type TestSortApplication struct {
	app.TestApplication

	SortVal int
}

func (t *TestSortApplication) Sort() int {
	return t.SortVal
}
