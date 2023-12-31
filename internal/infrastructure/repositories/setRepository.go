package repositories

import (
	"fmt"
	"strings"
	"volleyapp/internal/core/domain"
	"volleyapp/internal/core/ports"

	"github.com/lib/pq"
)

type SetRepository struct {
	db ports.Database
}

var _ ports.SetRepository = (*SetRepository)(nil)

func NewSetRepository(db ports.Database) *SetRepository {
	return &SetRepository{
		db: db,
	}
}

func (s *SetRepository) SaveNewSet(newSet domain.SetMainInfo) (int, error) {
	query := `
		INSERT INTO set (game_id, started_at, is_active, last_update)
		VALUES($1, $2, $3, $4)
		RETURNING set_id
	`
	result := s.db.GetDB().QueryRow(
		query,
		newSet.GameId,
		newSet.StartedAt,
		newSet.IsActive,
		newSet.LastUpdate,
	)
	var newSetId int
	if err := result.Scan(&newSetId); err != nil {
		return 0, fmt.Errorf(
			"[DATABASE] Error in save new set: %s", err,
		)
	}
	return int(newSetId), nil
}

func (s *SetRepository) FinishSet(setId int, set domain.Set) (int, error) {
	// TODO check and set winner
	query := `
		UPDATE set
		SET
			is_active = $1,
			set_winner = $2,
			last_update = $3
		WHERE set_id = $4
	`
	result, err := s.db.GetDB().Exec(
		query,
		set.IsActive,
		set.SetWinner,
		set.LastUpdate,
		setId,
	)
	if err != nil {
		return 0, fmt.Errorf(
			"[DATABASE] Error in finish set: %s", err,
		)
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, fmt.Errorf(
			"[DATABASE] Error in finish set: %s", err,
		)
	}
	return int(rowsAffected), nil
}

func (s *SetRepository) GetSet(setId int) (domain.Set, error) {
	var set domain.Set
	query := `
		SELECT *
		FROM set
		WHERE set_id = $1
	`
	result := s.db.GetDB().QueryRow(query, setId)
	if err := result.Scan(
		&set.SetId,
		&set.GameId,
		&set.StartedAt,
		&set.IsActive,
		&set.TotalAttacks,
		&set.AttackPoints,
		&set.AttackNeutrals,
		&set.AttackErrors,
		&set.AttackEffectiveness,
		&set.TotalBlocks,
		&set.BlockPoints,
		&set.BlockNeutrals,
		&set.BlockErrors,
		&set.BlockEffectiveness,
		&set.TotalServes,
		&set.ServePoints,
		&set.ServeNeutrals,
		&set.ServeErrors,
		&set.ServeEffectiveness,
		&set.OpponentErrors,
		&set.TotalPoints,
		&set.TotalActions,
		&set.TotalEffectiveness,
		&set.OpponentAttacks,
		&set.OpponentBlocks,
		&set.OpponentServes,
		&set.TotalErrors,
		&set.OpponentPoints,
		&set.SetWinner,
		pq.Array(&set.GameActions),
		&set.SetCount,
		&set.LastUpdate,
	); err != nil {
		return set, fmt.Errorf("[DATABASE] Error in get set: %s", err)
	}
	return set, nil
}

func (s *SetRepository) SaveSet(set domain.Set) (int, error) {
	query := `
		UPDATE set
		SET
			total_attacks = $1,
			attack_points = $2,
			attack_neutrals = $3,
			attack_errors = $4,
			attack_effectiveness = $5,
			total_blocks = $6,
			block_points = $7,
			block_neutrals = $8,
			block_errors = $9,
			block_effectiveness = $10,
			total_serves = $11,
			serve_points = $12,
			serve_neutrals = $13,
			serve_errors = $14,
			serve_effectiveness = $15,
			opponent_errors = $16,
			total_points = $17,
			total_actions = $18,
			total_effectiveness = $19,
			opponent_attacks = $20,
			opponent_blocks = $21,
			opponent_serves = $22,
			total_errors = $23,
			opponent_points = $24,
			game_actions = $25,
			last_update = $26
		WHERE set_id = $27
	`
	result, err := s.db.GetDB().Exec(
		query,
		set.TotalAttacks,
		set.AttackPoints,
		set.AttackNeutrals,
		set.AttackErrors,
		set.AttackEffectiveness,
		set.TotalBlocks,
		set.BlockPoints,
		set.BlockNeutrals,
		set.BlockErrors,
		set.BlockEffectiveness,
		set.TotalServes,
		set.ServePoints,
		set.ServeNeutrals,
		set.ServeErrors,
		set.ServeEffectiveness,
		set.OpponentErrors,
		set.TotalPoints,
		set.TotalActions,
		set.TotalEffectiveness,
		set.OpponentAttacks,
		set.OpponentBlocks,
		set.OpponentServes,
		set.TotalErrors,
		set.OpponentPoints,
		fmt.Sprintf("{%s}", strings.Join(set.GameActions, ",")),
		set.LastUpdate,
		set.SetId,
	)
	if err != nil {
		return 0, fmt.Errorf(
			"[DATABASE] Error in save set: %s", err,
		)
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, fmt.Errorf(
			"[DATABASE] Error in save rally: %s", err,
		)
	}
	return int(rowsAffected), nil
}
