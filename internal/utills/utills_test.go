package utills

import (
	"testing"

	"github.com/scipie28/note-service-api/internal/app/api"
	"github.com/stretchr/testify/require"
)

func TestSwapKeyAndValue(t *testing.T) {
	t.Run("input value equal nil", func(t *testing.T) {
		res, err := SwapKeyAndValue(nil)
		expectedRes := make(map[string]int32)

		require.Nil(t, err)
		require.Equal(t, expectedRes, res)
	})

	t.Run("input value equal zero", func(t *testing.T) {
		req := make(map[int32]string)
		res, err := SwapKeyAndValue(req)

		expectedRes := map[string]int32{}

		require.Nil(t, err)
		require.Equal(t, expectedRes, res)
	})

	t.Run("success case", func(t *testing.T) {
		req := map[int32]string{1: "one", 2: "two"}
		res, err := SwapKeyAndValue(req)

		expectedRes := map[string]int32{"one": 1, "two": 2}

		require.Nil(t, err)
		require.Equal(t, expectedRes, res)
	})

	t.Run("input identical keys", func(t *testing.T) {
		req := map[int32]string{1: "one", 2: "two", 3: "one", 4: "two"}
		res, err := SwapKeyAndValue(req)

		expectedRes := map[string]int32{"one": 1, "two": 2}

		actual := equalMap(res, expectedRes)

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
	t.Run("input value equal nil", func(t *testing.T) {
		res, err := FilterSlice(nil)
		var expectedRes []string

		require.Nil(t, err)
		require.Equal(t, expectedRes, res)
	})

	t.Run("input values equal zero", func(t *testing.T) {
		req := []string{""}
		res, err := FilterSlice(req)
		var expectedRes []string

		require.Nil(t, err)
		require.Equal(t, expectedRes, res)
	})

	t.Run("not success case", func(t *testing.T) {
		req := []string{"r", "t", "z", "1"}
		res, err := FilterSlice(req)

		var expectedRes []string

		require.Nil(t, err)
		require.Equal(t, expectedRes, res)
	})

	t.Run("nothing passed the filter", func(t *testing.T) {
		req := []string{"a", "t", "d", "1"}
		res, err := FilterSlice(req)

		expectedRes := []string{"a", "d"}

		require.Nil(t, err)
		require.Equal(t, expectedRes, res)
	})

	t.Run("everything passed the filter", func(t *testing.T) {
		req := []string{"a", "b", "c", "d"}
		res, err := FilterSlice(req)

		expectedRes := []string{"a", "b", "c", "d"}

		require.Nil(t, err)
		require.Equal(t, expectedRes, res)
	})
}

func TestConvertSliceToMap(t *testing.T) {
	t.Run("input value equal nil", func(t *testing.T) {
		res, err := ConvertSliceToMap(nil)
		expectedRes := make(map[uint64]api.Note)

		require.Nil(t, err)
		require.Equal(t, expectedRes, res)
	})

	t.Run("input values equal zero", func(t *testing.T) {
		req := make([]api.Note, 0)
		res, err := ConvertSliceToMap(req)

		expectedRes := map[uint64]api.Note{}

		require.Nil(t, err)
		require.Equal(t, expectedRes, res)
	})

	t.Run("success case", func(t *testing.T) {
		req := []api.Note{
			{Id: 1, UserId: 1, ClassroomId: 23, DocumentId: 6},
			{Id: 2, UserId: 2, ClassroomId: 24, DocumentId: 7},
		}
		res, err := ConvertSliceToMap(req)

		expectedRes := map[uint64]api.Note{
			1: {Id: 1, UserId: 1, ClassroomId: 23, DocumentId: 6},
			2: {Id: 2, UserId: 2, ClassroomId: 24, DocumentId: 7},
		}

		require.Nil(t, err)
		require.Equal(t, expectedRes, res)
	})
}

func TestSplitSlice(t *testing.T) {
	var (
		req = []api.Note{
			{Id: 1, UserId: 1, ClassroomId: 23, DocumentId: 6},
			{Id: 2, UserId: 2, ClassroomId: 24, DocumentId: 7},
			{Id: 3, UserId: 3, ClassroomId: 23, DocumentId: 6},
			{Id: 4, UserId: 4, ClassroomId: 24, DocumentId: 7},
		}
	)

	t.Run("req value equal nil", func(t *testing.T) {
		_, err := SplitSlice(nil, 1)

		expectedError := "error input values"

		require.Error(t, err)
		require.Equal(t, expectedError, err.Error())

	})

	t.Run("req values equal zero ", func(t *testing.T) {
		req := make([]api.Note, 0)
		_, err := SplitSlice(req, 0)

		expectedError := "error input values"

		require.Error(t, err)
		require.Equal(t, expectedError, err.Error())

	})

	t.Run("butch size equal to 1", func(t *testing.T) {
		res, err := SplitSlice(req, 1)

		expectedRes := [][]api.Note{
			{{Id: 1, UserId: 1, ClassroomId: 23, DocumentId: 6}},
			{{Id: 2, UserId: 2, ClassroomId: 24, DocumentId: 7}},
			{{Id: 3, UserId: 3, ClassroomId: 23, DocumentId: 6}},
			{{Id: 4, UserId: 4, ClassroomId: 24, DocumentId: 7}},
		}

		require.Nil(t, err)
		require.Equal(t, expectedRes, res)
	})

	t.Run("the number value is a multiple of butch size", func(t *testing.T) {
		res, err := SplitSlice(req, 2)

		expectedRes := [][]api.Note{
			{{Id: 1, UserId: 1, ClassroomId: 23, DocumentId: 6},
				{Id: 2, UserId: 2, ClassroomId: 24, DocumentId: 7}},
			{{Id: 3, UserId: 3, ClassroomId: 23, DocumentId: 6},
				{Id: 4, UserId: 4, ClassroomId: 24, DocumentId: 7}},
		}

		require.Nil(t, err)
		require.Equal(t, expectedRes, res)
	})

	t.Run("the number value is not a multiple of butch size", func(t *testing.T) {
		res, err := SplitSlice(req, 3)

		expectedRes := [][]api.Note{
			{{Id: 1, UserId: 1, ClassroomId: 23, DocumentId: 6},
				{Id: 2, UserId: 2, ClassroomId: 24, DocumentId: 7},
				{Id: 3, UserId: 3, ClassroomId: 23, DocumentId: 6}},
			{{Id: 4, UserId: 4, ClassroomId: 24, DocumentId: 7}},
		}

		require.Nil(t, err)
		require.Equal(t, expectedRes, res)
	})

	t.Run("numeric value is larger than the batch size", func(t *testing.T) {
		res, err := SplitSlice(req, 6)

		expectedRes := [][]api.Note{
			{{Id: 1, UserId: 1, ClassroomId: 23, DocumentId: 6},
				{Id: 2, UserId: 2, ClassroomId: 24, DocumentId: 7},
				{Id: 3, UserId: 3, ClassroomId: 23, DocumentId: 6},
				{Id: 4, UserId: 4, ClassroomId: 24, DocumentId: 7}},
		}

		require.Nil(t, err)
		require.Equal(t, expectedRes, res)
	})

	t.Run("butch size less or equal to zero", func(t *testing.T) {
		_, err := SplitSlice(req, 0)
		expectedError := "error input values"

		require.NotNil(t, err)
		require.Equal(t, expectedError, err.Error())

	})
}

func TestTwoToOneDimensionalSlice(t *testing.T) {
	t.Run("input value equal nil", func(t *testing.T) {
		expectedError := "error input values"
		_, err := TwoToOneDimensionalSlice(nil)

		require.Error(t, err)
		require.Equal(t, expectedError, err.Error())
	})

	t.Run("input value equal zero", func(t *testing.T) {
		req := make([][]api.Note, 0)
		_, err := TwoToOneDimensionalSlice(req)

		expectedRes := "error input values"

		require.Error(t, err)
		require.Equal(t, expectedRes, err.Error())
	})

	t.Run("input slice with one value", func(t *testing.T) {
		req := [][]api.Note{
			{{Id: 1, UserId: 1, ClassroomId: 23, DocumentId: 6}},
		}
		res, err := TwoToOneDimensionalSlice(req)

		expectedRes := []api.Note{
			{Id: 1, UserId: 1, ClassroomId: 23, DocumentId: 6},
		}

		require.Nil(t, err)
		require.Equal(t, expectedRes, res)
	})

	t.Run("success case", func(t *testing.T) {
		req := [][]api.Note{
			{{Id: 1, UserId: 1, ClassroomId: 23, DocumentId: 6},
				{Id: 2, UserId: 2, ClassroomId: 24, DocumentId: 7},
				{Id: 3, UserId: 3, ClassroomId: 23, DocumentId: 6}},
			{{Id: 4, UserId: 4, ClassroomId: 24, DocumentId: 7}},
		}
		res, err := TwoToOneDimensionalSlice(req)

		expectedRes := []api.Note{
			{Id: 1, UserId: 1, ClassroomId: 23, DocumentId: 6},
			{Id: 2, UserId: 2, ClassroomId: 24, DocumentId: 7},
			{Id: 3, UserId: 3, ClassroomId: 23, DocumentId: 6},
			{Id: 4, UserId: 4, ClassroomId: 24, DocumentId: 7},
		}

		require.Nil(t, err)
		require.Equal(t, expectedRes, res)
	})
}
