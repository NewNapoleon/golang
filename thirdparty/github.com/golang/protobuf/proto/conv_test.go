package proto_test

import (
	"encoding/json"
	"github.com/searKing/golang/thirdparty/github.com/golang/protobuf/proto"
	"github.com/searKing/golang/thirdparty/github.com/golang/protobuf/ptypes/struct"
	"strings"
	"testing"
)

type Human struct {
	Name      string
	Friends   []string
	Strangers []Human
}

type ToGolangMapTests struct {
	input  Human
	output string
}

var (
	toGolangMapTests = []ToGolangMapTests{{
		input: Human{
			Name:    "Alice",
			Friends: []string{"Bob", "Carol", "Dave"},
			Strangers: []Human{{
				Name:    "Eve",
				Friends: []string{"Oscar"},
				Strangers: []Human{{
					Name:    "Isaac",
					Friends: []string{"Justin", "Trent", "Pat", "Victor", "Walter"},
				}},
			}},
		},
		output: `{
 "Friends": [
  "Bob",
  "Carol",
  "Dave"
 ],
 "Name": "Alice",
 "Strangers": [
  {
   "Friends": [
    "Oscar"
   ],
   "Name": "Eve",
   "Strangers": [
    {
     "Friends": [
      "Justin",
      "Trent",
      "Pat",
      "Victor",
      "Walter"
     ],
     "Name": "Isaac",
     "Strangers": null
    }
   ]
  }
 ]
}`,
	},
	}
)

func TestToGolangMap(t *testing.T) {
	for m, test := range toGolangMapTests {
		humanStructpb, err := struct_.ToProtoStruct(test.input)
		if err != nil {
			t.Errorf("#%d: ToGolangMap(%+v): got: _, %v exp: _, nil", m, test.input, err)
		}
		humanMap, err := proto.ToGolangMap(humanStructpb)
		if err != nil {
			t.Errorf("#%d: proto.ToGolangMap(%+v): got: _, %v exp: _, nil", m, test.input, err)
		}

		humanBytes, err := json.MarshalIndent(humanMap, "", " ")
		if err != nil {
			t.Errorf("#%d: json.Marshal(%+v): got: _, %v exp: _, nil", m, test.input, err)
		}

		if strings.Compare(string(humanBytes), test.output) != 0 {
			t.Errorf("#%d: json.Marshal(%+v): got: \n%v\n exp: \n%v", m, test.input, string(humanBytes), test.output)
		}
	}
}
