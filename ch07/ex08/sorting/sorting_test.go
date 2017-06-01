package sorting

import (
	"bytes"
	"sort"
	"testing"
)

func TestMultiSort(t *testing.T) {
	var tests = []struct {
		tracks  []*Track
		columns []string
		want    []*Track // sort by artist, then sort by year
	}{
		{
			[]*Track{
				&Track{"A", "Artist C", "A", 2, Length("1m")},
				&Track{"B", "Artist A", "B", 1, Length("1m")},
				&Track{"C", "Artist C", "C", 2, Length("1m")},
				&Track{"D", "Artist B", "D", 4, Length("1m")},
				&Track{"E", "Artist C", "D", 1, Length("1m")},
				&Track{"F", "Artist D", "D", 3, Length("1m")},
				&Track{"G", "Artist C", "D", 1, Length("1m")},
				&Track{"H", "Artist E", "D", 3, Length("1m")},
			},
			[]string{
				"Artist",
				"Year",
			},
			[]*Track{
				&Track{"B", "Artist A", "B", 1, Length("1m")},
				&Track{"E", "Artist C", "D", 1, Length("1m")},
				&Track{"G", "Artist C", "D", 1, Length("1m")},
				&Track{"A", "Artist C", "A", 2, Length("1m")},
				&Track{"C", "Artist C", "C", 2, Length("1m")},
				&Track{"F", "Artist D", "D", 3, Length("1m")},
				&Track{"H", "Artist E", "D", 3, Length("1m")},
				&Track{"D", "Artist B", "D", 4, Length("1m")},
			},
		},
	}

	for _, test := range tests {
		sort.Sort(MultiSort(test.tracks, test.columns))
		if !equals(test.tracks, test.want) {
			t.Errorf("MultiSort result is:\n%s\nBut expected is:\n%s", SprintTracks(test.tracks), SprintTracks(test.want))
		}
	}
}

// BenchmarkSortStable が正しく多段ソートすることを確認します。
func TestSortStable(t *testing.T) {
	var tests = []struct {
		tracks []*Track
		want   []*Track // sort by artist, then sort by year
	}{
		{
			[]*Track{
				&Track{"A", "Artist C", "A", 2, Length("1m")},
				&Track{"B", "Artist A", "B", 1, Length("1m")},
				&Track{"C", "Artist C", "C", 2, Length("1m")},
				&Track{"D", "Artist B", "D", 4, Length("1m")},
				&Track{"E", "Artist C", "D", 1, Length("1m")},
				&Track{"F", "Artist D", "D", 3, Length("1m")},
				&Track{"G", "Artist C", "D", 1, Length("1m")},
				&Track{"H", "Artist E", "D", 3, Length("1m")},
			},
			[]*Track{
				&Track{"B", "Artist A", "B", 1, Length("1m")},
				&Track{"E", "Artist C", "D", 1, Length("1m")},
				&Track{"G", "Artist C", "D", 1, Length("1m")},
				&Track{"A", "Artist C", "A", 2, Length("1m")},
				&Track{"C", "Artist C", "C", 2, Length("1m")},
				&Track{"F", "Artist D", "D", 3, Length("1m")},
				&Track{"H", "Artist E", "D", 3, Length("1m")},
				&Track{"D", "Artist B", "D", 4, Length("1m")},
			},
		},
	}

	for _, test := range tests {
		sort.Stable(byArtist(test.tracks))
		sort.Stable(byYear(test.tracks))
		if !equals(test.tracks, test.want) {
			t.Errorf("sort.Stable result is:\n%s\nBut expected is:\n%s", SprintTracks(test.tracks), SprintTracks(test.want))
		}
	}
}

func BenchmarkMultiSort(b *testing.B) {
	tracks := []*Track{
		&Track{"A", "Artist C", "A", 2, Length("1m")},
		&Track{"B", "Artist A", "B", 1, Length("1m")},
		&Track{"C", "Artist C", "C", 2, Length("1m")},
		&Track{"D", "Artist B", "D", 4, Length("1m")},
		&Track{"E", "Artist C", "D", 1, Length("1m")},
		&Track{"F", "Artist D", "D", 3, Length("1m")},
		&Track{"G", "Artist C", "D", 1, Length("1m")},
		&Track{"H", "Artist E", "D", 3, Length("1m")},
	}
	for i := 0; i < b.N; i++ {
		sort.Sort(MultiSort(tracks, []string{"Artist", "Year"}))
		sort.Sort(sort.Reverse(MultiSort(tracks, []string{"Artist", "Year"})))
	}
}

func BenchmarkSortStable(b *testing.B) {
	tracks := []*Track{
		&Track{"A", "Artist C", "A", 2, Length("1m")},
		&Track{"B", "Artist A", "B", 1, Length("1m")},
		&Track{"C", "Artist C", "C", 2, Length("1m")},
		&Track{"D", "Artist B", "D", 4, Length("1m")},
		&Track{"E", "Artist C", "D", 1, Length("1m")},
		&Track{"F", "Artist D", "D", 3, Length("1m")},
		&Track{"G", "Artist C", "D", 1, Length("1m")},
		&Track{"H", "Artist E", "D", 3, Length("1m")},
	}
	for i := 0; i < b.N; i++ {
		sort.Stable(byArtist(tracks))
		sort.Stable(byYear(tracks))
		sort.Stable(sort.Reverse(byArtist(tracks)))
		sort.Stable(sort.Reverse(byYear(tracks)))
	}
}

func equals(x, y []*Track) bool {
	if len(x) != len(y) {
		return false
	}
	for i := range x {
		if x[i].Album != y[i].Album {
			return false
		} else if x[i].Artist != y[i].Artist {
			return false
		} else if x[i].Length != y[i].Length {
			return false
		} else if x[i].Title != y[i].Title {
			return false
		} else if x[i].Year != y[i].Year {
			return false
		}
	}
	return true
}

func SprintTracks(tracks []*Track) string {
	var b bytes.Buffer
	FprintTracks(&b, tracks)
	return b.String()
}
