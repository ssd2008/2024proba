package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAddToSchedule(t *testing.T) {
	t.Run("wrong day", func(t *testing.T) {
		res := AddToSchedule("asd", "as", "asd")
		assert.False(t, res)
	})

	t.Run("limit 3", func(t *testing.T) {
		for i := 0; i < 3; i++ {
			res := AddToSchedule("Mon", "as", "asd")
			assert.True(t, res)
		}

		res := AddToSchedule("Mon", "as", "asd")
		assert.False(t, res)

		for i := 0; i < 3; i++ {
			res = AddToSchedule("Tue", "as", "asd")
			assert.True(t, res)
		}

		res = AddToSchedule("Tue", "as", "asd")
		assert.False(t, res)
	})

}

func TestGetNameSchedule(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		sch := GetNameSchedule("asd")
		t.Log(sch)
	})
	t.Run("sort by day", func(t *testing.T) {
		res := AddToSchedule("Fri", "as", "asd")
		assert.True(t, res)
		res = AddToSchedule("Tue", "as", "asd")
		assert.True(t, res)
		res = AddToSchedule("Thu", "as", "asd")
		assert.True(t, res)
		res = AddToSchedule("Wed", "as", "asd")
		assert.True(t, res)

		sch := GetNameSchedule("asd")
		t.Log(sch)
	})

	t.Run("sort by subject", func(t *testing.T) {
		res := AddToSchedule("Mon", "AAc", "asd")
		assert.True(t, res)
		res = AddToSchedule("Mon", "AAa", "asd")
		assert.True(t, res)
		res = AddToSchedule("Mon", "AAb", "asd")
		assert.True(t, res)

		sch := GetNameSchedule("asd")
		t.Log(sch)
	})

	t.Run("dupl", func(t *testing.T) {
		for i := 0; i < 4; i++ {
			res := AddToSchedule("Wed", "as", "asd")
			assert.True(t, res)
		}

		sch := GetNameSchedule("asd")
		t.Log(sch)
	})
}
