package utills

import (
	"errors"
	"testing"

	"github.com/scipie28/note-service-api/internal/app/api"
	"github.com/stretchr/testify/require"
)

func TestSwapKeyAndValue(t *testing.T) {
	t.Run("input value to nil", func(t *testing.T) {
		output, err := SwapKeyAndValue(nil)
		expected := make(map[string]int32)
		require.Nil(t, err)
		require.Equal(t, expected, output)
	})

	t.Run("input map to zero value", func(t *testing.T) {
		input := make(map[int32]string)
		output, err := SwapKeyAndValue(input)

		expected := map[string]int32{}

		require.Nil(t, err)
		require.Equal(t, expected, output)
	})

	t.Run("values match the condition", func(t *testing.T) {
		input := map[int32]string{1: "one", 2: "two"}
		output, err := SwapKeyAndValue(input)

		expected := map[string]int32{"one": 1, "two": 2}

		require.Nil(t, err)
		require.Equal(t, expected, output)
	})

	t.Run("identical keys", func(t *testing.T) {
		input := map[int32]string{1: "one", 2: "two", 3: "one", 4: "two"}
		output, err := SwapKeyAndValue(input)
		expected := map[string]int32{"one": 1, "two": 2}

		actual := equalMap(output, expected)

		require.Nil(t, err)
		require.Equal(t, true, actual)
	})
}

func equalMap(output map[string]int32, expected map[string]int32) bool {
	if len(expected) != len(output) {
		return false
	}

	for key := range output {
		if _, found := expected[key]; !found {

			return false
		}
	}

	return true
}

func TestFilterSlice(t *testing.T) {
	t.Run("input to nil value", func(t *testing.T) {
		var expected []string

		output, err := FilterSlice(nil)

		require.Nil(t, err)
		require.Equal(t, expected, output)
	})

	t.Run("input slice to zero values", func(t *testing.T) {
		var expected []string
		input := []string{""}

		output, err := FilterSlice(input)

		require.Nil(t, err)
		require.Equal(t, expected, output)
	})

	t.Run("not values conforming for filter", func(t *testing.T) {
		var expected []string
		input := []string{"r", "t", "z", "1"}
		output, err := FilterSlice(input)

		require.Nil(t, err)
		require.Equal(t, expected, output)
	})

	t.Run("values match the condition", func(t *testing.T) {
		input := []string{"a", "t", "d", "1"}
		expected := []string{"a", "d"}

		output, err := FilterSlice(input)

		require.Nil(t, err)
		require.Equal(t, expected, output)
	})

	t.Run("only matching values for filter", func(t *testing.T) {
		expected := []string{"a", "b", "c", "d"}
		input := []string{"a", "b", "c", "d"}
		output, err := FilterSlice(input)

		require.Nil(t, err)
		require.Equal(t, expected, output)
	})
}

func TestConvertSliceToMap(t *testing.T) {

	t.Run("input to nil value", func(t *testing.T) {
		expected := map[uint64]api.Note{}

		output, err := ConvertSliceToMap(nil)

		require.Nil(t, err)
		require.Equal(t, expected, output)
	})

	t.Run("input map to zero values", func(t *testing.T) {
		expected := map[uint64]api.Note{}
		input := make([]api.Note, 0)

		output, err := ConvertSliceToMap(input)

		require.Nil(t, err)
		require.Equal(t, expected, output)
	})

	t.Run("values match the condition", func(t *testing.T) {
		input := []api.Note{
			{Id: 1, UserId: 1, ClassroomId: 23, DocumentId: 6},
			{Id: 2, UserId: 2, ClassroomId: 24, DocumentId: 7},
		}

		expected := map[uint64]api.Note{
			1: {Id: 1, UserId: 1, ClassroomId: 23, DocumentId: 6},
			2: {Id: 2, UserId: 2, ClassroomId: 24, DocumentId: 7},
		}

		output, err := ConvertSliceToMap(input)

		require.Nil(t, err)
		require.Equal(t, expected, output)
	})
}

func TestSplitSlice(t *testing.T) {
	t.Run("input to nil value", func(t *testing.T) {
		var expectedValue [][]api.Note
		expectedError := errors.New("ErrorInputValues")
		output, err := SplitSlice(nil, 1)

		require.NotNil(t, err)
		require.Equal(t, expectedError, err)
		require.Equal(t, expectedValue, output)
	})

	t.Run("input map to zero values", func(t *testing.T) {
		var expectedValue [][]api.Note
		expectedError := errors.New("ErrorInputValues")
		input := make([]api.Note, 0)

		output, err := SplitSlice(input, 0)

		require.NotNil(t, err)
		require.Equal(t, expectedError, err)
		require.Equal(t, expectedValue, output)
	})

	t.Run("input values match the condition", func(t *testing.T) {
		input := []api.Note{
			{Id: 1, UserId: 1, ClassroomId: 23, DocumentId: 6},
			{Id: 2, UserId: 2, ClassroomId: 24, DocumentId: 7},
			{Id: 3, UserId: 3, ClassroomId: 23, DocumentId: 6},
			{Id: 4, UserId: 4, ClassroomId: 24, DocumentId: 7},
		}

		t.Run("map with butch size equal to 1", func(t *testing.T) {
			expected := [][]api.Note{
				{{Id: 1, UserId: 1, ClassroomId: 23, DocumentId: 6}},
				{{Id: 2, UserId: 2, ClassroomId: 24, DocumentId: 7}},
				{{Id: 3, UserId: 3, ClassroomId: 23, DocumentId: 6}},
				{{Id: 4, UserId: 4, ClassroomId: 24, DocumentId: 7}},
			}

			output, err := SplitSlice(input, 1)

			require.Nil(t, err)
			require.Equal(t, expected, output)
		})

		t.Run("the number value is a multiple of butch size", func(t *testing.T) {
			expected := [][]api.Note{
				{{Id: 1, UserId: 1, ClassroomId: 23, DocumentId: 6},
					{Id: 2, UserId: 2, ClassroomId: 24, DocumentId: 7}},
				{{Id: 3, UserId: 3, ClassroomId: 23, DocumentId: 6},
					{Id: 4, UserId: 4, ClassroomId: 24, DocumentId: 7}},
			}

			output, err := SplitSlice(input, 2)

			require.Nil(t, err)
			require.Equal(t, expected, output)
		})

		t.Run("the number value is not a multiple of butch size", func(t *testing.T) {
			expected := [][]api.Note{
				{{Id: 1, UserId: 1, ClassroomId: 23, DocumentId: 6},
					{Id: 2, UserId: 2, ClassroomId: 24, DocumentId: 7},
					{Id: 3, UserId: 3, ClassroomId: 23, DocumentId: 6}},
				{{Id: 4, UserId: 4, ClassroomId: 24, DocumentId: 7}},
			}

			actual, err := SplitSlice(input, 3)

			require.Nil(t, err)
			require.Equal(t, expected, actual)
		})

		t.Run("numeric value is larger than the batch size", func(t *testing.T) {
			var expectedValue [][]api.Note
			expectedError := errors.New("ErrorInputValues")

			output, err := SplitSlice(input, 6)

			require.NotNil(t, err)
			require.Equal(t, expectedError, err)
			require.Equal(t, expectedValue, output)
		})

		t.Run("butch size less or equal to zero", func(t *testing.T) {
			var expectedValue [][]api.Note
			expectedError := errors.New("ErrorInputValues")

			output, err := SplitSlice(input, 0)

			require.NotNil(t, err)
			require.Equal(t, expectedError, err)
			require.Equal(t, expectedValue, output)
		})
	})
}
