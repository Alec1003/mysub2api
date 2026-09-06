package repository

import (
	"context"
	"database/sql"
	"errors"
	"github.com/Wei-Shaw/sub2api/internal/service"
)

type skillRepository struct{ db *sql.DB }

func NewSkillRepository(db *sql.DB) service.SkillRepository { return &skillRepository{db: db} }
func scanSkill(row interface{ Scan(...any) error }) (*service.Skill, error) {
	var s service.Skill
	err := row.Scan(&s.ID, &s.Name, &s.Description, &s.EffectImage, &s.DriveURL, &s.Status, &s.CreatedAt, &s.UpdatedAt)
	return &s, err
}
func (r *skillRepository) List(ctx context.Context, all bool) ([]service.Skill, error) {
	q := `SELECT id,name,description,effect_image,drive_url,status,created_at,updated_at FROM skills`
	if !all {
		q += ` WHERE status='active'`
	}
	q += ` ORDER BY created_at DESC`
	rows, err := r.db.QueryContext(ctx, q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	out := []service.Skill{}
	for rows.Next() {
		s, err := scanSkill(rows)
		if err != nil {
			return nil, err
		}
		out = append(out, *s)
	}
	return out, rows.Err()
}
func (r *skillRepository) Get(ctx context.Context, id int64) (*service.Skill, error) {
	s, err := scanSkill(r.db.QueryRowContext(ctx, `SELECT id,name,description,effect_image,drive_url,status,created_at,updated_at FROM skills WHERE id=$1`, id))
	if errors.Is(err, sql.ErrNoRows) {
		return nil, service.ErrSkillNotFound
	}
	return s, err
}
func (r *skillRepository) Create(ctx context.Context, s *service.Skill) error {
	return r.db.QueryRowContext(ctx, `INSERT INTO skills(name,description,effects,effect_image,drive_url,status) VALUES($1,$2,'',$3,$4,$5) RETURNING id,created_at,updated_at`, s.Name, s.Description, s.EffectImage, s.DriveURL, s.Status).Scan(&s.ID, &s.CreatedAt, &s.UpdatedAt)
}
func (r *skillRepository) Update(ctx context.Context, s *service.Skill) error {
	res, err := r.db.ExecContext(ctx, `UPDATE skills SET name=$1,description=$2,effects='',effect_image=$3,drive_url=$4,status=$5,updated_at=NOW() WHERE id=$6`, s.Name, s.Description, s.EffectImage, s.DriveURL, s.Status, s.ID)
	if err != nil {
		return err
	}
	n, _ := res.RowsAffected()
	if n == 0 {
		return service.ErrSkillNotFound
	}
	return nil
}
func (r *skillRepository) Delete(ctx context.Context, id int64) error {
	res, err := r.db.ExecContext(ctx, `DELETE FROM skills WHERE id=$1`, id)
	if err != nil {
		return err
	}
	n, _ := res.RowsAffected()
	if n == 0 {
		return service.ErrSkillNotFound
	}
	return nil
}
func (r *skillRepository) GetForUser(ctx context.Context, uid, sid int64) (*service.Skill, bool, error) {
	s, err := r.Get(ctx, sid)
	if err != nil {
		return nil, false, err
	}
	var exists bool
	err = r.db.QueryRowContext(ctx, `SELECT EXISTS(SELECT 1 FROM user_skill_redemptions WHERE user_id=$1 AND skill_id=$2)`, uid, sid).Scan(&exists)
	return s, exists, err
}
func (r *skillRepository) Entitlement(ctx context.Context, uid int64) (service.SkillEntitlement, error) {
	var points float64
	err := r.db.QueryRowContext(ctx, `
		SELECT
			COALESCE((
				SELECT SUM(amount) * 1000
				FROM payment_orders
				WHERE user_id = $1
				  AND order_type = 'balance'
				  AND status IN ('PAID', 'COMPLETED', 'RECHARGING', 'paid', 'completed', 'recharging')
			), 0)
			+ COALESCE((
				SELECT SUM(rc.value) * 1000
				FROM redeem_codes rc
				WHERE rc.used_by = $1
				  AND rc.value > 0
				  AND rc.type IN ('balance', 'admin_balance')
				  AND NOT EXISTS (
					  SELECT 1
					  FROM payment_orders po
					  WHERE po.recharge_code = rc.code
				  )
			), 0)
	`, uid).Scan(&points)
	if err != nil {
		return service.SkillEntitlement{}, err
	}
	var used int64
	err = r.db.QueryRowContext(ctx, `SELECT COUNT(*) FROM user_skill_redemptions WHERE user_id=$1`, uid).Scan(&used)
	if err != nil {
		return service.SkillEntitlement{}, err
	}
	return service.SkillEntitlementFromRecharge(points, used), nil
}
func (r *skillRepository) Redeem(ctx context.Context, uid, sid int64) (*service.Skill, error) {
	tx, err := r.db.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()
	var s service.Skill
	err = tx.QueryRowContext(ctx, `SELECT id,name,description,effect_image,drive_url,status,created_at,updated_at FROM skills WHERE id=$1 FOR UPDATE`, sid).Scan(&s.ID, &s.Name, &s.Description, &s.EffectImage, &s.DriveURL, &s.Status, &s.CreatedAt, &s.UpdatedAt)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, service.ErrSkillNotFound
	}
	if err != nil {
		return nil, err
	}
	if s.Status != "active" {
		return nil, service.ErrSkillNotFound
	}
	var exists bool
	_ = tx.QueryRowContext(ctx, `SELECT EXISTS(SELECT 1 FROM user_skill_redemptions WHERE user_id=$1 AND skill_id=$2)`, uid, sid).Scan(&exists)
	if exists {
		return nil, service.ErrSkillAlreadyRedeemed
	}
	var points float64
	err = tx.QueryRowContext(ctx, `
		SELECT
			COALESCE((
				SELECT SUM(amount) * 1000
				FROM payment_orders
				WHERE user_id = $1
				  AND order_type = 'balance'
				  AND status IN ('PAID', 'COMPLETED', 'RECHARGING', 'paid', 'completed', 'recharging')
			), 0)
			+ COALESCE((
				SELECT SUM(rc.value) * 1000
				FROM redeem_codes rc
				WHERE rc.used_by = $1
				  AND rc.value > 0
				  AND rc.type IN ('balance', 'admin_balance')
				  AND NOT EXISTS (
					  SELECT 1
					  FROM payment_orders po
					  WHERE po.recharge_code = rc.code
				  )
			), 0)
	`, uid).Scan(&points)
	if err != nil {
		return nil, err
	}
	var used int64
	err = tx.QueryRowContext(ctx, `SELECT COUNT(*) FROM user_skill_redemptions WHERE user_id=$1`, uid).Scan(&used)
	if err != nil {
		return nil, err
	}
	if service.SkillEntitlementFromRecharge(points, used).AvailableRedemptions < 1 {
		return nil, service.ErrSkillNotEligible
	}
	if _, err = tx.ExecContext(ctx, `INSERT INTO user_skill_redemptions(user_id,skill_id) VALUES($1,$2)`, uid, sid); err != nil {
		return nil, err
	}
	return &s, tx.Commit()
}
