package main

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEnvToService(t *testing.T) {
	val := int64(123456788)

	bd, err := getBlackDeath(map[string]string{
		annotation: strconv.Itoa(int(val)),
	})
	assert.Nil(t, err)

	assert.Equal(t, bd, val)

	_, err = getBlackDeath(map[string]string{
		"foo": "bar",
	})
	assert.NotNil(t, err)
}
