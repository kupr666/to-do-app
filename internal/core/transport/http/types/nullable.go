package core_http_types

import (
	"encoding/json"

	"github.com/kupr666/to-do-app/internal/core/domain"
)

type Nullable[T any] struct {
	domain.Nullable[T]
}

func (n *Nullable[T]) UnmarshalJSON(b []byte) error {

	// if this method called - set = true
	n.Set = true

	// check th content of incoming json - if there are nothing - n.Value = nil 
	if string(b) == "null" {
		n.Value = nil

		return nil
	}

	// if incoming json is isn't "null" - we received phone_number -> try to unmarshal
	var value T
	if err := json.Unmarshal(b, &value); err != nil {
		return err
	}

	n.Value = &value

	return nil
}

func (n *Nullable[T]) ToDomain() domain.Nullable[T] {
	return domain.Nullable[T]{
		Value: n.Value,
		Set: n.Set,
	}
}