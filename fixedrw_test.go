package multilock

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFixedRW_Get(t *testing.T) {
	s := []int{1}
	fixed := NewFixedRW(10)
	m := fixed.Get(s)
	m.Lock()
	m.Unlock()
	assert.Equal(t, addr(fixed.Get(s)), addr(m))
}

func TestFixedRW_Get_One(t *testing.T) {
	s := []int{1}
	s2 := []int{1, 2}
	fixed := NewFixed(1)
	m := fixed.Get(s)
	m.Lock()
	m.Unlock()
	assert.Equal(t, addr(fixed.Get(s2)), addr(m))
}

func TestFixedRW_Get_CustomDistribution(t *testing.T) {
	s := []int{1}
	s2 := []int{1, 2}
	fixed := NewFixed(100, WithCustomDistribution(func(_ string, _ int) int {
		return 0
	}))
	m := fixed.Get(s)
	m.Lock()
	m.Unlock()
	assert.Equal(t, addr(fixed.Get(s2)), addr(m))
}
