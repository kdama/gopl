package console

import (
	"reflect"
	"testing"
)

func TestSprintTable(t *testing.T) {
	tests := []struct {
		rows [][]string
		want string
	}{
		{
			[][]string{},
			"",
		},
		{
			[][]string{[]string{}},
			"",
		},
		{
			[][]string{[]string{}, []string{}, []string{}},
			"",
		},
		{
			[][]string{
				[]string{"foo"},
			}, "" +
				"+-----+\n" +
				"| foo |\n" +
				"+-----+\n",
		},
		{
			[][]string{
				[]string{"foo", "12345"},
			}, "" +
				"+-----+-------+\n" +
				"| foo | 12345 |\n" +
				"+-----+-------+\n",
		},
		{
			[][]string{
				[]string{"A", "1", ""},
				[]string{"BBB", "1234"},
				[]string{"CC", "12", "", ""},
			}, "" +
				"+-----+------+--+\n" +
				"| A   | 1    |  |\n" +
				"+-----+------+--+\n" +
				"| BBB | 1234 |  |\n" +
				"+-----+------+--+\n" +
				"| CC  | 12   |  |\n" +
				"+-----+------+--+\n",
		},
	}
	for _, test := range tests {
		if got := SprintTable(test.rows); got != test.want {
			t.Errorf("SprintTable(%v) = %v, want %v", test.rows, got, test.want)
		}
	}
}

func TestColumnLens(t *testing.T) {
	tests := []struct {
		rows [][]string
		want []int
	}{
		{
			[][]string{},
			[]int{},
		},
		{
			[][]string{[]string{}},
			[]int{},
		},
		{
			[][]string{[]string{}, []string{}, []string{}},
			[]int{},
		},
		{
			[][]string{
				[]string{"foo"},
			},
			[]int{3},
		},
		{
			[][]string{
				[]string{"foo", "12345"},
			},
			[]int{3, 5},
		},
		{
			[][]string{
				[]string{"A", "1", ""},
				[]string{"BBB", "1234"},
				[]string{"CC", "12", "", ""},
			},
			[]int{3, 4, 0},
		},
	}
	for _, test := range tests {
		if got := columnLens(test.rows); !reflect.DeepEqual(got, test.want) {
			t.Errorf("columnLens(%v) = %v, want %v", test.rows, got, test.want)
		}
	}
}
