package utills

import (
	"testing"

	"github.com/scipie28/note-service-api/internal/app/api"
	"github.com/stretchr/testify/require"
)

func TestSwapKeyAndValue(t *testing.T) {
	t.Run("nil", func(t *testing.T) {
		actual, _ := SwapKeyAndValue(make(map[int32]string))
		expected := make(map[string]int32)

		require.Equal(t, expected, actual)
	})

	t.Run("zeroValue", func(t *testing.T) {
		data := make(map[int32]string)
		actual, _ := SwapKeyAndValue(data)

		expected := map[string]int32{}

		require.Equal(t, expected, actual)
	})

	t.Run("correct filled", func(t *testing.T) {
		actual, _ := SwapKeyAndValue(map[int32]string{1: "one", 2: "two"})
		realRes := map[string]int32{"one": 1, "two": 2}

		require.Equal(t, realRes, actual)
	})

	t.Run("identical keys", func(t *testing.T) {
		actual, _ := SwapKeyAndValue(map[int32]string{1: "one", 2: "two", 3: "one"})
		expected := map[string]int32{"one": 1, "two": 2}

		require.Equal(t, expected, actual)
	})
}

func TestFilterSlice(t *testing.T) {
	t.Run("nil", func(t *testing.T) {
		var a []string
		var expected []string

		actual, _ := FilterSlice(a)

		require.Equal(t, expected, actual)
	})

	t.Run("zeroValue", func(t *testing.T) {
		var expected []string

		actual, _ := FilterSlice([]string{""})

		require.Equal(t, expected, actual)
	})

	t.Run("not needed value", func(t *testing.T) {
		var expected []string

		actual, _ := FilterSlice([]string{"r", "t", "z", "1"})

		require.Equal(t, expected, actual)
	})

	t.Run("correct filled", func(t *testing.T) {
		var expected = []string{"a", "d"}

		res, _ := FilterSlice([]string{"a", "t", "d", "1"})

		require.Equal(t, expected, res)
	})

	t.Run("only needed value", func(t *testing.T) {
		var expected = []string{"a", "b", "c", "d"}

		actual, _ := FilterSlice([]string{"a", "b", "c", "d"})

		require.Equal(t, expected, actual)
	})
}

func TestConvertSliceToMap(t *testing.T) {

	t.Run("nil", func(t *testing.T) {
		var data []api.Note
		var expected = map[uint64]api.Note{}

		actual, _ := ConvertSliceToMap(data)

		require.Equal(t, expected, actual)
	})

	t.Run("zeroValue", func(t *testing.T) {
		var expected = map[uint64]api.Note{}
		data := make([]api.Note, 0)

		actual, _ := ConvertSliceToMap(data)

		require.Equal(t, expected, actual)
	})

	t.Run("correct filled", func(t *testing.T) {
		data := []api.Note{
			{Id: 1, UserId: 1, ClassroomId: 23, DocumentId: 6},
			{Id: 2, UserId: 2, ClassroomId: 24, DocumentId: 7},
		}

		expected := map[uint64]api.Note{
			1: {Id: 1, UserId: 1, ClassroomId: 23, DocumentId: 6},
			2: {Id: 2, UserId: 2, ClassroomId: 24, DocumentId: 7},
		}

		actual, _ := ConvertSliceToMap(data)

		require.Equal(t, expected, actual)
	})
}

func TestSplitSlice(t *testing.T) {
	t.Run("nil", func(t *testing.T) {
		var data []api.Note
		expected := make([][]api.Note, 0)

		actual, _ := SplitSlice(data, 5)

		require.Equal(t, expected, actual)
	})

	t.Run("zeroValue", func(t *testing.T) {
		expected := make([][]api.Note, 0)
		data := make([]api.Note, 0)

		actual, _ := SplitSlice(data, 1)

		require.Equal(t, expected, actual)
	})

	t.Run("correct filled", func(t *testing.T) {
		data := []api.Note{
			{Id: 1, UserId: 1, ClassroomId: 23, DocumentId: 6},
			{Id: 2, UserId: 2, ClassroomId: 24, DocumentId: 7},
			{Id: 3, UserId: 3, ClassroomId: 23, DocumentId: 6},
			{Id: 4, UserId: 4, ClassroomId: 24, DocumentId: 7},
		}

		t.Run("butchSize=1", func(t *testing.T) {
			expected := [][]api.Note{
				{{Id: 1, UserId: 1, ClassroomId: 23, DocumentId: 6}},
				{{Id: 2, UserId: 2, ClassroomId: 24, DocumentId: 7}},
				{{Id: 3, UserId: 3, ClassroomId: 23, DocumentId: 6}},
				{{Id: 4, UserId: 4, ClassroomId: 24, DocumentId: 7}},
			}

			actual, _ := SplitSlice(data, 1)

			require.Equal(t, expected, actual)
		})

		t.Run("len(data) % butchSize = 0", func(t *testing.T) {
			expected := [][]api.Note{
				{{Id: 1, UserId: 1, ClassroomId: 23, DocumentId: 6},
					{Id: 2, UserId: 2, ClassroomId: 24, DocumentId: 7}},
				{{Id: 3, UserId: 3, ClassroomId: 23, DocumentId: 6},
					{Id: 4, UserId: 4, ClassroomId: 24, DocumentId: 7}},
			}

			actual, _ := SplitSlice(data, 2)

			require.Equal(t, expected, actual)
		})

		t.Run("len(data) % butchSize != 0", func(t *testing.T) {
			expected := [][]api.Note{
				{{Id: 1, UserId: 1, ClassroomId: 23, DocumentId: 6},
					{Id: 2, UserId: 2, ClassroomId: 24, DocumentId: 7},
					{Id: 3, UserId: 3, ClassroomId: 23, DocumentId: 6}},
				{{Id: 4, UserId: 4, ClassroomId: 24, DocumentId: 7}},
			}

			actual, _ := SplitSlice(data, 3)

			require.Equal(t, expected, actual)
		})
	})
}
