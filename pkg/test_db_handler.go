package pkg

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/uptrace/bun"
)

type TestDBHandler[Schema any] struct {
	dbConn *sql.Tx
}

func NewTestDBHandler[Schema any](dbConn *sql.Tx) *TestDBHandler[Schema] {
	return &TestDBHandler[Schema]{
		dbConn: dbConn,
	}
}

func (h *TestDBHandler[Schema]) SeedTable(ctx context.Context, schemaWithData Schema) {
	//if _, err := h.dbConn.NewInsert().Model(&schemaWithData).Exec(ctx); err != nil {
	//	panic(err)
	//}

	schemaElem := reflect.ValueOf(schemaWithData).Elem()
	schemaType := reflect.TypeOf(schemaWithData)

	numfields := schemaType.NumField()

	fields := make([]any, 0, numfields)
	fieldValues := make([]any, 0, numfields)

	for i := 0; i < numfields; i++ {
		fields = append(fields, schemaType.Field(i).Name)
		fieldValues = append(fields, schemaElem.Field(i).Interface())
	}

	// TODO: parse the fields and fiedlValues into the query
	query := "INSERT INTO table_name (field1, field2) VALUES ($1, $2)"
	_, err := h.dbConn.ExecContext(ctx, query, fields...)

	if err != nil {
		panic(err)
	}
}

func (h *TestDBHandler[Schema]) AssertCountInTable(t *testing.T, size int, field map[string]interface{}) []Schema {
	t.Helper()

	var entries []Schema
	err := h.buildWhere(field).Scan(context.Background(), &entries)

	if size == 0 {
		assert.Error(t, err, sql.ErrNoRows)
	} else {
		assert.NoError(t, err)
	}

	assert.Equal(t, size, len(entries))

	return entries
}

func (h *TestDBHandler[Schema]) AssertInTable(t *testing.T, expectedFields map[string]interface{}) Schema {
	t.Helper()

	var schema Schema
	err := h.buildWhere(expectedFields).Scan(context.Background(), &schema)

	assert.NoError(t, err)
	assert.NotNil(t, schema)

	return schema
}

func (h *TestDBHandler[Schema]) buildWhere(fields map[string]interface{}) *bun.SelectQuery {
	query := h.dbConn.NewSelect().Model((*Schema)(nil)).QueryBuilder()

	for field, value := range fields {
		where := fmt.Sprintf("%s = ?", field)
		query.Where(where, value)
	}

	selectQuery, ok := query.Unwrap().(*bun.SelectQuery)
	if !ok {
		panic("unable to build query")
	}

	return selectQuery
}
