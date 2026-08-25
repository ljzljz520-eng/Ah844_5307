package parser

import "testing"

func TestParseJSONSchema(t *testing.T) {
	_, e := ParseJSON([]byte(`{"database":"mysql","tables":[{"Name":"vehicles","Columns":["id"],"Types":["INT"]}]}`))
	if e != nil {
		t.Fatal(e)
	}
}
