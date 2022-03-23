package model

import (
	"encoding/json"
	"testing"
)

func Test_Marshal(t *testing.T) {
	var desc Description = []string{"1", "2"}

	b, err := json.Marshal(&desc)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log("success", string(b))
}
