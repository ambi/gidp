package sqlgateway

import (
	"errors"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/ambi/gidp/model"
	"github.com/stretchr/testify/assert"
)

func eqRegexp(s string) string {
	return "^" + regexp.QuoteMeta(s) + "$"
}

func TestTenantGateway(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	tenants := []*model.Tenant{
		{ID: "ID1", Status: "active"},
		{ID: "ID2", Status: "active"},
	}

	gw := NewTenantRepo(db)

	t.Run("NewTenantRepo", func(t *testing.T) {
		assert.NotNil(t, gw)
	})

	t.Run("List", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"id", "status"}).
			AddRow(tenants[0].ID, tenants[0].Status).
			AddRow(tenants[1].ID, tenants[1].Status)
		mock.ExpectQuery(eqRegexp(sqlTenantList)).WillReturnRows(rows)

		result, err := gw.List()

		assert.Nil(t, err)
		expected := make([]*model.Tenant, 2)
		for i := 0; i < len(tenants); i++ {
			expected[i] = tenants[i]
		}
		assert.Equal(t, expected, result)
		assert.Nil(t, mock.ExpectationsWereMet())
	})

	t.Run("Get", func(t *testing.T) {
		t.Run("tenant ID is valid", func(t *testing.T) {
			rows := sqlmock.NewRows([]string{"id", "status"}).
				AddRow(tenants[0].ID, tenants[0].Status)
			mock.ExpectQuery(eqRegexp(sqlTenantGet)).WillReturnRows(rows)

			result, err := gw.Get(tenants[0].ID)

			assert.Nil(t, err)
			assert.Equal(t, tenants[0], result)
			assert.Nil(t, mock.ExpectationsWereMet())
		})

		t.Run("tenant ID is invalid", func(t *testing.T) {
			e := errors.New("not found")
			mock.ExpectQuery(eqRegexp(sqlTenantGet)).WillReturnError(e)

			result, err := gw.Get(tenants[0].ID)

			assert.NotNil(t, err)
			assert.Nil(t, result)
			assert.Nil(t, mock.ExpectationsWereMet())
		})
	})

	t.Run("Create", func(t *testing.T) {
		result := sqlmock.NewResult(1, 1)
		mock.ExpectExec(eqRegexp(sqlTenantCreate)).WillReturnResult(result)

		tenant := model.Tenant{
			Status: "inactive",
		}
		err := gw.Create(&tenant)

		assert.Nil(t, err)
		assert.Equal(t, "inactive", tenant.Status)
		assert.Regexp(t, "^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$", tenant.ID)
		assert.Nil(t, mock.ExpectationsWereMet())
	})

	t.Run("Delete", func(t *testing.T) {
		result := sqlmock.NewResult(1, 1)
		mock.ExpectExec(eqRegexp(sqlTenantDelete)).WillReturnResult(result)

		err := gw.Delete(tenants[0])

		assert.Nil(t, err)
		assert.Nil(t, mock.ExpectationsWereMet())
	})
}
