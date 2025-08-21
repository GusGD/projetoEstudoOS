package repository

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/ordemservicos/backend/internal/os/domain"
)

// OrdemServicoRepository define as operações de persistência para OrdemServico
type OrdemServicoRepository interface {
	Create(ctx context.Context, os *domain.OrdemServico) error
	GetByID(ctx context.Context, id uuid.UUID) (*domain.OrdemServico, error)
	GetByPrefeitura(ctx context.Context, prefeituraID uuid.UUID, limit, offset int) ([]*domain.OrdemServico, error)
	GetByStatus(ctx context.Context, status domain.StatusOS, prefeituraID uuid.UUID) ([]*domain.OrdemServico, error)
	GetByResponsavel(ctx context.Context, responsavelID uuid.UUID) ([]*domain.OrdemServico, error)
	Update(ctx context.Context, os *domain.OrdemServico) error
	Delete(ctx context.Context, id uuid.UUID) error
	GetHistorico(ctx context.Context, osID uuid.UUID) ([]domain.Evento, error)
	AddEvento(ctx context.Context, evento domain.Evento) error
}

type ordemServicoRepository struct {
	db *sql.DB
}

func NewOrdemServicoRepository(db *sql.DB) OrdemServicoRepository {
	return &ordemServicoRepository{db: db}
}

func (r *ordemServicoRepository) Create(ctx context.Context, os *domain.OrdemServico) error {
	query := `
		INSERT INTO ordens_servico (
			id, titulo, descricao, status, prioridade, localizacao,
			responsavel_id, prefeitura_id, solicitante_id,
			data_criacao, data_atualizacao, data_conclusao, tags
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)
	`

	_, err := r.db.ExecContext(ctx, query,
		os.ID, os.Titulo, os.Descricao, os.Status, os.Prioridade, os.Localizacao,
		os.ResponsavelID, os.PrefeituraID, os.SolicitanteID,
		os.DataCriacao, os.DataAtualizacao, os.DataConclusao, os.Tags,
	)

	if err != nil {
		return err
	}

	for _, evento := range os.Historico {
		if err := r.AddEvento(ctx, evento); err != nil {
			return err
		}
	}

	return nil
}

func (r *ordemServicoRepository) GetByID(ctx context.Context, id uuid.UUID) (*domain.OrdemServico, error) {
	query := `
		SELECT id, titulo, descricao, status, prioridade, localizacao,
			   responsavel_id, prefeitura_id, solicitante_id,
			   data_criacao, data_atualizacao, data_conclusao, tags
		FROM ordens_servico
		WHERE id = $1
	`

	var os domain.OrdemServico
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&os.ID, &os.Titulo, &os.Descricao, &os.Status, &os.Prioridade, &os.Localizacao,
		&os.ResponsavelID, &os.PrefeituraID, &os.SolicitanteID,
		&os.DataCriacao, &os.DataAtualizacao, &os.DataConclusao, &os.Tags,
	)

	if err != nil {
		return nil, err
	}

	historico, err := r.GetHistorico(ctx, id)
	if err != nil {
		return nil, err
	}
	os.Historico = historico

	return &os, nil
}

func (r *ordemServicoRepository) GetByPrefeitura(ctx context.Context, prefeituraID uuid.UUID, limit, offset int) ([]*domain.OrdemServico, error) {
	query := `
		SELECT id, titulo, descricao, status, prioridade, localizacao,
			   responsavel_id, prefeitura_id, solicitante_id,
			   data_criacao, data_atualizacao, data_conclusao, tags
		FROM ordens_servico
		WHERE prefeitura_id = $1
		ORDER BY data_criacao DESC
		LIMIT $2 OFFSET $3
	`

	rows, err := r.db.QueryContext(ctx, query, prefeituraID, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var ordens []*domain.OrdemServico
	for rows.Next() {
		var os domain.OrdemServico
		err := rows.Scan(
			&os.ID, &os.Titulo, &os.Descricao, &os.Status, &os.Prioridade, &os.Localizacao,
			&os.ResponsavelID, &os.PrefeituraID, &os.SolicitanteID,
			&os.DataCriacao, &os.DataAtualizacao, &os.DataConclusao, &os.Tags,
		)
		if err != nil {
			return nil, err
		}
		ordens = append(ordens, &os)
	}

	return ordens, nil
}

func (r *ordemServicoRepository) GetByStatus(ctx context.Context, status domain.StatusOS, prefeituraID uuid.UUID) ([]*domain.OrdemServico, error) {
	query := `
		SELECT id, titulo, descricao, status, prioridade, localizacao,
			   responsavel_id, prefeitura_id, solicitante_id,
			   data_criacao, data_atualizacao, data_conclusao, tags
		FROM ordens_servico
		WHERE status = $1 AND prefeitura_id = $2
		ORDER BY prioridade DESC, data_criacao DESC
	`

	rows, err := r.db.QueryContext(ctx, query, status, prefeituraID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var ordens []*domain.OrdemServico
	for rows.Next() {
		var os domain.OrdemServico
		err := rows.Scan(
			&os.ID, &os.Titulo, &os.Descricao, &os.Status, &os.Prioridade, &os.Localizacao,
			&os.ResponsavelID, &os.PrefeituraID, &os.SolicitanteID,
			&os.DataCriacao, &os.DataAtualizacao, &os.DataConclusao, &os.Tags,
		)
		if err != nil {
			return nil, err
		}
		ordens = append(ordens, &os)
	}

	return ordens, nil
}

func (r *ordemServicoRepository) GetByResponsavel(ctx context.Context, responsavelID uuid.UUID) ([]*domain.OrdemServico, error) {
	query := `
		SELECT id, titulo, descricao, status, prioridade, localizacao,
			   responsavel_id, prefeitura_id, solicitante_id,
			   data_criacao, data_atualizacao, data_conclusao, tags
		FROM ordens_servico
		WHERE responsavel_id = $1
		ORDER BY prioridade DESC, data_criacao DESC
	`

	rows, err := r.db.QueryContext(ctx, query, responsavelID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var ordens []*domain.OrdemServico
	for rows.Next() {
		var os domain.OrdemServico
		err := rows.Scan(
			&os.ID, &os.Titulo, &os.Descricao, &os.Status, &os.Prioridade, &os.Localizacao,
			&os.ResponsavelID, &os.PrefeituraID, &os.SolicitanteID,
			&os.DataCriacao, &os.DataAtualizacao, &os.DataConclusao, &os.Tags,
		)
		if err != nil {
			return nil, err
		}
		ordens = append(ordens, &os)
	}

	return ordens, nil
}

func (r *ordemServicoRepository) Update(ctx context.Context, os *domain.OrdemServico) error {
	query := `
		UPDATE ordens_servico SET
			titulo = $1, descricao = $2, status = $3, prioridade = $4,
			localizacao = $5, responsavel_id = $6, data_atualizacao = $7,
			data_conclusao = $8, tags = $9
		WHERE id = $10
	`

	_, err := r.db.ExecContext(ctx, query,
		os.Titulo, os.Descricao, os.Status, os.Prioridade,
		os.Localizacao, os.ResponsavelID, os.DataAtualizacao,
		os.DataConclusao, os.Tags, os.ID,
	)

	return err
}

func (r *ordemServicoRepository) Delete(ctx context.Context, id uuid.UUID) error {
	query := `DELETE FROM ordens_servico WHERE id = $1`
	_, err := r.db.ExecContext(ctx, query, id)
	return err
}

func (r *ordemServicoRepository) GetHistorico(ctx context.Context, osID uuid.UUID) ([]domain.Evento, error) {
	query := `
		SELECT id, os_id, tipo, descricao, usuario_id, data_criacao, metadata
		FROM eventos_os
		WHERE os_id = $1
		ORDER BY data_criacao ASC
	`

	rows, err := r.db.QueryContext(ctx, query, osID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var eventos []domain.Evento
	for rows.Next() {
		var evento domain.Evento
		err := rows.Scan(
			&evento.ID, &evento.OSID, &evento.Tipo, &evento.Descricao,
			&evento.UsuarioID, &evento.DataCriacao, &evento.Metadata,
		)
		if err != nil {
			return nil, err
		}
		eventos = append(eventos, evento)
	}

	return eventos, nil
}

func (r *ordemServicoRepository) AddEvento(ctx context.Context, evento domain.Evento) error {
	query := `
		INSERT INTO eventos_os (id, os_id, tipo, descricao, usuario_id, data_criacao, metadata)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
	`

	_, err := r.db.ExecContext(ctx, query,
		evento.ID, evento.OSID, evento.Tipo, evento.Descricao,
		evento.UsuarioID, evento.DataCriacao, evento.Metadata,
	)

	return err
}
