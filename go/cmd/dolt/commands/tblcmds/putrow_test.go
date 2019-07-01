package tblcmds

import (
	"context"
	"github.com/google/uuid"
	"github.com/liquidata-inc/ld/dolt/go/libraries/doltcore/dtestutils"
	"github.com/liquidata-inc/ld/dolt/go/libraries/doltcore/env"
	"github.com/liquidata-inc/ld/dolt/go/libraries/doltcore/row"
	"github.com/liquidata-inc/ld/dolt/go/libraries/doltcore/table"
	"github.com/liquidata-inc/ld/dolt/go/libraries/doltcore/table/typed/noms"
	"github.com/liquidata-inc/ld/dolt/go/store/types"
	"strings"
	"testing"
)

var expectedId = types.UUID(uuid.Must(uuid.Parse("11111111-1111-1111-1111-111111111111")))
var tableName = "people"

var expectedFieldVals = map[uint64]types.Value{
	dtestutils.IdTag:        expectedId,
	dtestutils.NameTag:      types.String("Eric Ericson"),
	dtestutils.AgeTag:       types.Uint(45),
	dtestutils.IsMarriedTag: types.Bool(true),
}

func TestPutRow(t *testing.T) {
	tests := []struct {
		args          []string
		expectedRes   int
		expectedTitle string
	}{
		{[]string{" "}, 1, ""},
		{[]string{tableName}, 1, ""},
		{[]string{tableName, "id:", "name:Eric Ericson", "age:45", "is_married:true"}, 1, ""},
		{[]string{tableName, "id:11111111-1111-1111-1111-111111111111", "name:Eric Ericson", "age:45", "is_married:true", "title:Dolt"}, 0, "Dolt"},
		{[]string{tableName, "id:11111111-1111-1111-1111-111111111111", "name:Eric Ericson", "age:45", "is_married:true", "title:"}, 0, ""},
		{[]string{tableName, "id:11111111-1111-1111-1111-111111111111", "name:Eric Ericson", "age:45", "is_married:true", "title"}, 1, ""},
		{[]string{tableName, "id:11111111-1111-1111-1111-111111111111", "name:Eric Ericson", "age:45", "is_married:true", ":Dolt"}, 1, ""},
		{[]string{tableName, "id:1", "name:Eric Ericson", "age:45", "is_married:true"}, 1, ""},
	}
	for _, test := range tests {
		dEnv := createEnvWithSeedData(t)

		commandStr := "dolt table put-row"
		result := PutRow(commandStr, test.args, dEnv)

		if result != test.expectedRes {
			commandLine := commandStr + " " + strings.Join(test.args, " ")
			t.Fatal("Unexpected failure. command", commandLine, "returned", result)
		}

		if result == 0 {
			root, _ := dEnv.WorkingRoot(context.Background())
			tbl, _ := root.GetTable(context.Background(), tableName)
			sch := tbl.GetSchema(context.Background())
			row, exists := tbl.GetRowByPKVals(context.Background(), row.TaggedValues{dtestutils.IdTag: expectedId}, sch)

			if !exists {
				t.Fatal("Could not find row")
			}

			for i, v := range expectedFieldVals {
				val, fld := row.GetColVal(i)

				if !val.Equals(types.Format_7_18, v) {
					t.Error("Unexpected value for", fld, "expected:", v, "actual:", val)
				}
			}

			titleVal, _ := row.GetColVal(dtestutils.TitleTag)

			if !titleVal.Equals(types.Format_7_18, types.String(test.expectedTitle)) {
				t.Error("Value of title was not the expected value. expected", test.expectedTitle, "actual", titleVal)
			}
		}
	}
}

func createEnvWithSeedData(t *testing.T) *env.DoltEnv {
	dEnv := dtestutils.CreateTestEnv()
	imt, sch := dtestutils.CreateTestDataTable(true)

	rd := table.NewInMemTableReader(imt)
	wr := noms.NewNomsMapCreator(context.Background(), dEnv.DoltDB.ValueReadWriter(), sch)

	_, _, err := table.PipeRows(context.Background(), rd, wr, false)
	rd.Close(context.Background())
	wr.Close(context.Background())

	if err != nil {
		t.Error("Failed to seed initial data", err)
	}

	err = dEnv.PutTableToWorking(context.Background(), *wr.GetMap(), wr.GetSchema(), tableName)

	if err != nil {
		t.Error("Unable to put initial value of table in in mem noms db", err)
	}

	return dEnv
}
