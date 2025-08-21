package domain

import (
	"time"

	"github.com/google/uuid"
)

// StatusOS representa os possíveis estados de uma Ordem de Serviço
type StatusOS string

const (
	StatusAberta      StatusOS = "aberta"
	StatusEmAndamento StatusOS = "em_andamento"
	StatusConcluida   StatusOS = "concluida"
	StatusTransferida StatusOS = "transferida"
	StatusCancelada   StatusOS = "cancelada"
	StatusEmAnalise   StatusOS = "em_analise"
)

// OrdemServico representa a entidade principal do sistema
type OrdemServico struct {
	ID              uuid.UUID  `json:"id"`
	Titulo          string     `json:"titulo"`
	Descricao       string     `json:"descricao"`
	Status          StatusOS   `json:"status"`
	Prioridade      int        `json:"prioridade"` // 1-5, onde 5 é mais alta
	Localizacao     string     `json:"localizacao"`
	ResponsavelID   *uuid.UUID `json:"responsavel_id,omitempty"`
	PrefeituraID    uuid.UUID  `json:"prefeitura_id"`
	SolicitanteID   uuid.UUID  `json:"solicitante_id"`
	DataCriacao     time.Time  `json:"data_criacao"`
	DataAtualizacao time.Time  `json:"data_atualizacao"`
	DataConclusao   *time.Time `json:"data_conclusao,omitempty"`
	SLA             *SLA       `json:"sla,omitempty"`
	Tags            []string   `json:"tags"`
	Historico       []Evento   `json:"historico"`
}

type SLA struct {
	ID          uuid.UUID `json:"id"`
	TempoLimite time.Time `json:"tempo_limite"`
	TempoAlerta time.Time `json:"tempo_alerta"`
	Prioridade  int       `json:"prioridade"`
	Descricao   string    `json:"descricao"`
}

// Evento representa um evento no histórico da OS (Event Sourcing)
type Evento struct {
	ID          uuid.UUID              `json:"id"`
	OSID        uuid.UUID              `json:"os_id"`
	Tipo        string                 `json:"tipo"`
	Descricao   string                 `json:"descricao"`
	UsuarioID   uuid.UUID              `json:"usuario_id"`
	DataCriacao time.Time              `json:"data_criacao"`
	Metadata    map[string]interface{} `json:"metadata,omitempty"`
}

func NewOrdemServico(titulo, descricao, localizacao string, prefeituraID, solicitanteID uuid.UUID, prioridade int) *OrdemServico {
	now := time.Now()
	return &OrdemServico{
		ID:              uuid.New(),
		Titulo:          titulo,
		Descricao:       descricao,
		Status:          StatusAberta,
		Prioridade:      prioridade,
		Localizacao:     localizacao,
		PrefeituraID:    prefeituraID,
		SolicitanteID:   solicitanteID,
		DataCriacao:     now,
		DataAtualizacao: now,
		Tags:            []string{},
		Historico:       []Evento{},
	}
}

// AdicionarEvento adiciona um evento ao histórico (Observer Pattern)
func (os *OrdemServico) AdicionarEvento(tipo, descricao string, usuarioID uuid.UUID, metadata map[string]interface{}) {
	evento := Evento{
		ID:          uuid.New(),
		OSID:        os.ID,
		Tipo:        tipo,
		Descricao:   descricao,
		UsuarioID:   usuarioID,
		DataCriacao: time.Now(),
		Metadata:    metadata,
	}

	os.Historico = append(os.Historico, evento)
	os.DataAtualizacao = time.Now()
}

func (os *OrdemServico) AlterarStatus(novoStatus StatusOS, usuarioID uuid.UUID, motivo string) error {
	statusAnterior := os.Status
	os.Status = novoStatus
	os.DataAtualizacao = time.Now()

	if novoStatus == StatusConcluida {
		now := time.Now()
		os.DataConclusao = &now
	}

	os.AdicionarEvento(
		"status_alterado",
		"Status alterado de "+string(statusAnterior)+" para "+string(novoStatus)+": "+motivo,
		usuarioID,
		map[string]interface{}{
			"status_anterior": statusAnterior,
			"status_novo":     novoStatus,
			"motivo":          motivo,
		},
	)

	return nil
}

func (os *OrdemServico) AtribuirResponsavel(responsavelID uuid.UUID, usuarioID uuid.UUID) {
	os.ResponsavelID = &responsavelID
	os.AdicionarEvento(
		"responsavel_atribuido",
		"Responsável atribuído à OS",
		usuarioID,
		map[string]interface{}{
			"responsavel_id": responsavelID,
		},
	)
}

func (os *OrdemServico) AdicionarTag(tag string) {
	for _, t := range os.Tags {
		if t == tag {
			return
		}
	}
	os.Tags = append(os.Tags, tag)
}

func (os *OrdemServico) RemoverTag(tag string) {
	for i, t := range os.Tags {
		if t == tag {
			os.Tags = append(os.Tags[:i], os.Tags[i+1:]...)
			break
		}
	}
}
