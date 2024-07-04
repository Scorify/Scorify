package model

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/99designs/gqlgen/graphql"
	"github.com/google/uuid"
)

type UUID uuid.UUID

func MarshalUUID(val uuid.UUID) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		err := json.NewEncoder(w).Encode(val)
		if err != nil {
			panic(err)
		}
	})
}

func UnmarshalUUID(v interface{}) (uuid.UUID, error) {
	if m, ok := v.(string); ok {
		return uuid.MustParse(m), nil
	}

	return uuid.UUID{}, fmt.Errorf("%T is not a UUID", v)
}
