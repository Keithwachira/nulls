package nulls

import (
	"database/sql"
	"encoding/json"
)

type NullString struct {
	sql.NullString
}

type NullInt64 struct {
	sql.NullInt64
}

type NullFloat64 struct {
	sql.NullFloat64
}

type NullBool struct {
	sql.NullBool
}

func (ni *NullString) MarshalJSON() ([]byte, error) {
	if !ni.Valid {
		return json.Marshal(ni.String)
	}
	return json.Marshal(ni.String)
}

func (ni *NullString) UnmarshalJSON(b []byte) error {
	err := json.Unmarshal(b, &ni.String)
	ni.Valid = (err == nil)
	return err
}

func (ni *NullInt64) MarshalJSON() ([]byte, error) {
	if !ni.Valid {
		return json.Marshal(ni.Int64)
	}
	return json.Marshal(ni.Int64)
}

func (ni *NullInt64) UnmarshalJSON(b []byte) error {
	err := json.Unmarshal(b, &ni.Int64)
	ni.Valid = (err == nil)
	return err
}

func (ni *NullFloat64) MarshalJSON() ([]byte, error) {
	if !ni.Valid {
		return json.Marshal(ni.Float64)
	}
	return json.Marshal(ni.Float64)
}

func (ni *NullFloat64) UnmarshalJSON(b []byte) error {
	err := json.Unmarshal(b, &ni.Float64)
	ni.Valid = (err == nil)
	return err
}

func (ni *NullBool) MarshalJSON() ([]byte, error) {
	if !ni.Valid {
		return json.Marshal(ni.Bool)
	}
	return json.Marshal(ni.Bool)
}

func (ni *NullBool) UnmarshalJSON(b []byte) error {
	err := json.Unmarshal(b, &ni.Bool)
	ni.Valid = (err == nil)
	return err
}
