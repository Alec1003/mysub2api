package service

import (
	"context"
	"errors"
	infraerrors "github.com/Wei-Shaw/sub2api/internal/pkg/errors"
	"math"
	"strings"
	"time"
)

const SkillRechargeThreshold = 30000.0

var (
	ErrSkillNotFound        = infraerrors.NotFound("SKILL_NOT_FOUND", "skill not found")
	ErrSkillNotEligible     = infraerrors.Forbidden("SKILL_REDEMPTION_REQUIRED", "insufficient recharge points")
	ErrSkillAlreadyRedeemed = infraerrors.Conflict("SKILL_ALREADY_REDEEMED", "skill already redeemed")
)

type Skill struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	EffectImage string    `json:"effect_image"`
	DriveURL    string    `json:"drive_url,omitempty"`
	Status      string    `json:"status"`
	Redeemed    bool      `json:"redeemed"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
type SkillEntitlement struct {
	TotalRechargePoints  float64 `json:"total_recharge_points"`
	EarnedRedemptions    int64   `json:"earned_redemptions"`
	UsedRedemptions      int64   `json:"used_redemptions"`
	AvailableRedemptions int64   `json:"available_redemptions"`
}

type SkillRepository interface {
	List(ctx context.Context, includeInactive bool) ([]Skill, error)
	Create(ctx context.Context, s *Skill) error
	Update(ctx context.Context, s *Skill) error
	Delete(ctx context.Context, id int64) error
	Get(ctx context.Context, id int64) (*Skill, error)
	GetForUser(ctx context.Context, userID, skillID int64) (*Skill, bool, error)
	Entitlement(ctx context.Context, userID int64) (SkillEntitlement, error)
	Redeem(ctx context.Context, userID, skillID int64) (*Skill, error)
}

type SkillService struct{ repo SkillRepository }

func NewSkillService(repo SkillRepository) *SkillService { return &SkillService{repo: repo} }
func (s *SkillService) List(ctx context.Context, userID int64) ([]Skill, SkillEntitlement, error) {
	items, err := s.repo.List(ctx, true)
	if err != nil {
		return nil, SkillEntitlement{}, err
	}
	ent, err := s.repo.Entitlement(ctx, userID)
	if err != nil {
		return nil, ent, err
	}
	visible := make([]Skill, 0, len(items))
	for i := range items {
		_, redeemed, e := s.repo.GetForUser(ctx, userID, items[i].ID)
		if e != nil {
			return nil, ent, e
		}
		items[i].Redeemed = redeemed
		if !redeemed {
			if items[i].Status != "active" {
				continue
			}
			items[i].DriveURL = ""
		}
		visible = append(visible, items[i])
	}
	return visible, ent, nil
}
func (s *SkillService) AdminList(ctx context.Context) ([]Skill, error) { return s.repo.List(ctx, true) }
func (s *SkillService) Create(ctx context.Context, skill *Skill) error {
	if err := validateSkill(skill); err != nil {
		return err
	}
	return s.repo.Create(ctx, skill)
}
func validateSkill(s *Skill) error {
	s.Name = strings.TrimSpace(s.Name)
	s.Description = strings.TrimSpace(s.Description)
	s.EffectImage = strings.TrimSpace(s.EffectImage)
	s.DriveURL = strings.TrimSpace(s.DriveURL)
	if s.Name == "" || s.Description == "" || s.EffectImage == "" || s.DriveURL == "" {
		return errors.New("name, description, effect_image and drive_url are required")
	}
	if !strings.HasPrefix(s.EffectImage, "data:image/") && !strings.HasPrefix(s.EffectImage, "https://") && !strings.HasPrefix(s.EffectImage, "http://") {
		return errors.New("effect_image must be an image data URL or HTTP(S) URL")
	}
	if strings.HasPrefix(s.EffectImage, "data:image/") && len(s.EffectImage) > 8*1024*1024 {
		return errors.New("effect_image must be smaller than 6 MB")
	}
	if s.Status == "" {
		s.Status = "active"
	}
	if s.Status != "active" && s.Status != "offline" {
		return errors.New("status must be active or offline")
	}
	return nil
}
func (s *SkillService) Update(ctx context.Context, skill *Skill) error {
	if err := validateSkill(skill); err != nil {
		return err
	}
	return s.repo.Update(ctx, skill)
}
func (s *SkillService) Delete(ctx context.Context, id int64) error { return s.repo.Delete(ctx, id) }
func (s *SkillService) Redeem(ctx context.Context, userID, skillID int64) (*Skill, error) {
	return s.repo.Redeem(ctx, userID, skillID)
}

func entitlementFromRecharge(points float64, used int64) SkillEntitlement {
	earned := int64(math.Floor(points / SkillRechargeThreshold))
	available := earned - used
	if available < 0 {
		available = 0
	}
	return SkillEntitlement{TotalRechargePoints: points, EarnedRedemptions: earned, UsedRedemptions: used, AvailableRedemptions: available}
}

// SkillEntitlementFromRecharge keeps the qualification rule available to the
// SQL repository without duplicating the threshold calculation there.
func SkillEntitlementFromRecharge(points float64, used int64) SkillEntitlement {
	return entitlementFromRecharge(points, used)
}
