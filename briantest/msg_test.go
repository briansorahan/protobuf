package briantest

import (
	"encoding/json"
	"testing"

	"github.com/briansorahan/death"
	gogojson "github.com/gogo/protobuf/jsonpb"
	gjson "github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/ptypes/wrappers"
)

func TestMarshal(t *testing.T) {
	t.Run("with no gogoproto.jsontag", func(t *testing.T) {
		var (
			die       = death.By(t)
			marshaler = &gogojson.Marshaler{}
		)
		data, err := marshaler.MarshalToString(&Msg{
			AField: "foo",
			BField: "bar",
		})
		die(err)

		if expected, got := `{"aField":"foo","bField":"bar"}`, data; expected != got {
			t.Fatalf("expected %s, got %s", expected, got)
		}
	})

	t.Run("using gogoproto.jsontag", func(t *testing.T) {
		die := death.By(t)

		data, err := json.Marshal(&MsgTags{
			AField: "foo",
			BField: "bar",
		})
		die(err)

		if expected, got := `{"a":"foo","b":"bar","c":null}`, string(data); expected != got {
			t.Fatalf("expected %s, got %s", expected, got)
		}
	})

	t.Run("gogo jsonpb.Marshaler quotes around google.protobuf.Int64Value fields", func(t *testing.T) {
		var (
			die       = death.By(t)
			marshaler = &gogojson.Marshaler{}
		)
		data, err := marshaler.MarshalToString(&MsgTags{
			AField: "foo",
			BField: "bar",
			CField: &wrappers.Int64Value{Value: 1000000},
		})
		die(err)

		if expected, got := `{"a":"foo","b":"bar","c":1000000}`, data; expected != got {
			t.Fatalf("expected %s, got %s", expected, got)
		}
	})

	t.Run("google jsonpb.Marshaler quotes around google.protobuf.Int64Value fields", func(t *testing.T) {
		var (
			die       = death.By(t)
			marshaler = &gjson.Marshaler{}
		)
		data, err := marshaler.MarshalToString(&MsgTags{
			AField: "foo",
			BField: "bar",
			CField: &wrappers.Int64Value{Value: 1000000},
		})
		die(err)

		if expected, got := `{"a":"foo","b":"bar","c":1000000}`, data; expected != got {
			t.Fatalf("expected %s, got %s", expected, got)
		}
	})
}
