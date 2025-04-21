package dto_test

import (
	"testing"
	pkg_data "yamanmnur/simple-dashboard/pkg/data"

	"github.com/stretchr/testify/assert"
)

func TestGetSort(t *testing.T) {
	t.Run("should return default sort when Sort is empty", func(t *testing.T) {
		pageData := &pkg_data.PageData{}
		expectedSort := "Id desc"

		actualSort := pageData.GetSort()

		assert.Equal(t, expectedSort, actualSort)
		assert.Equal(t, expectedSort, pageData.Sort)
	})

	t.Run("should return existing Sort when it is not empty", func(t *testing.T) {
		pageData := &pkg_data.PageData{Sort: "Name asc"}
		expectedSort := "Name asc"

		actualSort := pageData.GetSort()

		assert.Equal(t, expectedSort, actualSort)
		assert.Equal(t, expectedSort, pageData.Sort)
	})
}
func TestGetOffset(t *testing.T) {
	t.Run("should calculate offset correctly when Page and Limit are set", func(t *testing.T) {
		pageData := &pkg_data.PageData{Page: 3, Limit: 20}
		expectedOffset := 40

		actualOffset := pageData.GetOffset()

		assert.Equal(t, expectedOffset, actualOffset)
	})

	t.Run("should calculate offset correctly when Page is default and Limit is set", func(t *testing.T) {
		pageData := &pkg_data.PageData{Limit: 15}
		expectedOffset := 0

		actualOffset := pageData.GetOffset()

		assert.Equal(t, expectedOffset, actualOffset)
	})

	t.Run("should calculate offset correctly when Limit is default and Page is set", func(t *testing.T) {
		pageData := &pkg_data.PageData{Page: 2}
		expectedOffset := 10

		actualOffset := pageData.GetOffset()

		assert.Equal(t, expectedOffset, actualOffset)
	})

	t.Run("should calculate offset correctly when both Page and Limit are default", func(t *testing.T) {
		pageData := &pkg_data.PageData{}
		expectedOffset := 0

		actualOffset := pageData.GetOffset()

		assert.Equal(t, expectedOffset, actualOffset)
	})
}

func TestGetLimit(t *testing.T) {
	t.Run("should return default Limit when Limit is not set", func(t *testing.T) {
		pageData := &pkg_data.PageData{}
		expectedLimit := 10

		actualLimit := pageData.GetLimit()

		assert.Equal(t, expectedLimit, actualLimit)
		assert.Equal(t, expectedLimit, pageData.Limit)
	})

	t.Run("should return existing Limit when it is set", func(t *testing.T) {
		pageData := &pkg_data.PageData{Limit: 25}
		expectedLimit := 25

		actualLimit := pageData.GetLimit()

		assert.Equal(t, expectedLimit, actualLimit)
		assert.Equal(t, expectedLimit, pageData.Limit)
	})
}

func TestGetPage(t *testing.T) {
	t.Run("should return default Page when Page is not set", func(t *testing.T) {
		pageData := &pkg_data.PageData{}
		expectedPage := 1

		actualPage := pageData.GetPage()

		assert.Equal(t, expectedPage, actualPage)
		assert.Equal(t, expectedPage, pageData.Page)
	})

	t.Run("should return existing Page when it is set", func(t *testing.T) {
		pageData := &pkg_data.PageData{Page: 5}
		expectedPage := 5

		actualPage := pageData.GetPage()

		assert.Equal(t, expectedPage, actualPage)
		assert.Equal(t, expectedPage, pageData.Page)
	})
}
