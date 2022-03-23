package util

import (
	"testing"
)

func TestNewSnowflake(t *testing.T) {
	t.Log(nodeBits, sequenceBits,
		nodeMax, sequenceMax,
		timestampShift, nodeShift,
		epoch)
	snowflake0, _ := NewSnowflake(0)
	snowflake1, _ := NewSnowflake(1)
	snowflake2, _ := NewSnowflake(2)
	for i := 0; i < 100; i++ {
		t.Log(snowflake0.NextId(), snowflake1.NextId(), snowflake2.NextId())
	}
}

func BenchmarkSnowflake_NextId(b *testing.B) {
	snowflake, _ := NewSnowflake(0)
	for i := 0; i < b.N; i++ {
		snowflake.NextId()
	}
}
