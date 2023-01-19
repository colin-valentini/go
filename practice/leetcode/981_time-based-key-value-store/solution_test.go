package timebasedkeyvaluestore

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTimeMap(t *testing.T) {
	timeMap := Constructor()
	timeMap.Set("foo", "bar", 1)
	assert.Equal(t, "bar", timeMap.Get("foo", 1))
	assert.Equal(t, "bar", timeMap.Get("foo", 3))
	timeMap.Set("foo", "bar2", 4)
	assert.Equal(t, "bar2", timeMap.Get("foo", 4))
	assert.Equal(t, "bar2", timeMap.Get("foo", 5))
	timeMap.Set("foo", "baz", 5)
	assert.Equal(t, "", timeMap.Get("foo", 0))
	assert.Equal(t, "baz", timeMap.Get("foo", 5))
	assert.Equal(t, "bar2", timeMap.Get("foo", 4))
	assert.Equal(t, "bar", timeMap.Get("foo", 3))
	assert.Equal(t, "", timeMap.Get("foo2", 10))
}
