package repository

import (
	"context"
	"database/sql"
	"regexp"
	"testing"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"
)

func TestSkillRepositoryEntitlementIncludesManualAndAdminRecharge(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	rechargeQuery := `(?s)SELECT.*redeem_codes rc.*admin_balance.*NOT EXISTS.*payment_orders po`
	mock.ExpectQuery(rechargeQuery).
		WithArgs(int64(42)).
		WillReturnRows(sqlmock.NewRows([]string{"total_points"}).AddRow(72000.0))
	mock.ExpectQuery(regexp.QuoteMeta("SELECT COUNT(*) FROM user_skill_redemptions WHERE user_id=$1")).
		WithArgs(int64(42)).
		WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(int64(1)))

	entitlement, err := (&skillRepository{db: db}).Entitlement(context.Background(), 42)
	require.NoError(t, err)
	require.Equal(t, 72000.0, entitlement.TotalRechargePoints)
	require.Equal(t, int64(2), entitlement.EarnedRedemptions)
	require.Equal(t, int64(1), entitlement.UsedRedemptions)
	require.Equal(t, int64(1), entitlement.AvailableRedemptions)
	require.NoError(t, mock.ExpectationsWereMet())
}

func TestSkillRepositoryEntitlementPropagatesRechargeQueryError(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	mock.ExpectQuery("SELECT").WithArgs(int64(7)).WillReturnError(sql.ErrConnDone)
	_, err = (&skillRepository{db: db}).Entitlement(context.Background(), 7)
	require.ErrorIs(t, err, sql.ErrConnDone)
	require.NoError(t, mock.ExpectationsWereMet())
}
