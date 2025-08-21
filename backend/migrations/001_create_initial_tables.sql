-- Migration: 001_create_initial_tables.sql
-- Descrição: Criação das tabelas principais do sistema

-- Tabela de prefeituras
CREATE TABLE IF NOT EXISTS prefeituras (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    nome VARCHAR(255) NOT NULL,
    cnpj VARCHAR(18) UNIQUE NOT NULL,
    endereco TEXT,
    telefone VARCHAR(20),
    email VARCHAR(255),
    ativo BOOLEAN DEFAULT true,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

-- Tabela de usuários
CREATE TABLE IF NOT EXISTS usuarios (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    nome VARCHAR(255) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    senha_hash VARCHAR(255) NOT NULL,
    cpf VARCHAR(14) UNIQUE,
    telefone VARCHAR(20),
    cargo VARCHAR(100),
    prefeitura_id UUID NOT NULL REFERENCES prefeituras(id),
    ativo BOOLEAN DEFAULT true,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

-- Tabela de perfis de usuário (RBAC)
CREATE TABLE IF NOT EXISTS perfis (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    nome VARCHAR(100) NOT NULL,
    descricao TEXT,
    prefeitura_id UUID NOT NULL REFERENCES prefeituras(id),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

-- Tabela de relacionamento usuário-perfil
CREATE TABLE IF NOT EXISTS usuario_perfis (
    usuario_id UUID REFERENCES usuarios(id) ON DELETE CASCADE,
    perfil_id UUID REFERENCES perfis(id) ON DELETE CASCADE,
    PRIMARY KEY (usuario_id, perfil_id)
);

-- Tabela de permissões
CREATE TABLE IF NOT EXISTS permissoes (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    nome VARCHAR(100) NOT NULL,
    descricao TEXT,
    recurso VARCHAR(100) NOT NULL,
    acao VARCHAR(50) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

-- Tabela de relacionamento perfil-permissão
CREATE TABLE IF NOT EXISTS perfil_permissoes (
    perfil_id UUID REFERENCES perfis(id) ON DELETE CASCADE,
    permissao_id UUID REFERENCES permissoes(id) ON DELETE CASCADE,
    PRIMARY KEY (perfil_id, permissao_id)
);

-- Tabela principal de ordens de serviço
CREATE TABLE IF NOT EXISTS ordens_servico (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    titulo VARCHAR(255) NOT NULL,
    descricao TEXT NOT NULL,
    status VARCHAR(50) NOT NULL DEFAULT 'aberta',
    prioridade INTEGER NOT NULL DEFAULT 3 CHECK (prioridade >= 1 AND prioridade <= 5),
    localizacao TEXT,
    responsavel_id UUID REFERENCES usuarios(id),
    prefeitura_id UUID NOT NULL REFERENCES prefeituras(id),
    solicitante_id UUID NOT NULL REFERENCES usuarios(id),
    data_criacao TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    data_atualizacao TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    data_conclusao TIMESTAMP WITH TIME ZONE,
    tags TEXT[] DEFAULT '{}',
    sla_id UUID,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

-- Tabela de eventos para Event Sourcing
CREATE TABLE IF NOT EXISTS eventos_os (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    os_id UUID NOT NULL REFERENCES ordens_servico(id) ON DELETE CASCADE,
    tipo VARCHAR(100) NOT NULL,
    descricao TEXT NOT NULL,
    usuario_id UUID NOT NULL REFERENCES usuarios(id),
    data_criacao TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    metadata JSONB
);

-- Tabela de SLA (Service Level Agreement)
CREATE TABLE IF NOT EXISTS slas (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    nome VARCHAR(255) NOT NULL,
    descricao TEXT,
    tempo_limite_horas INTEGER NOT NULL,
    tempo_alerta_horas INTEGER,
    prioridade INTEGER NOT NULL DEFAULT 3,
    prefeitura_id UUID NOT NULL REFERENCES prefeituras(id),
    ativo BOOLEAN DEFAULT true,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

-- Tabela de notificações
CREATE TABLE IF NOT EXISTS notificacoes (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    usuario_id UUID NOT NULL REFERENCES usuarios(id),
    titulo VARCHAR(255) NOT NULL,
    mensagem TEXT NOT NULL,
    tipo VARCHAR(50) NOT NULL DEFAULT 'info',
    lida BOOLEAN DEFAULT false,
    data_envio TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    data_leitura TIMESTAMP WITH TIME ZONE,
    metadata JSONB
);

-- Índices para performance
CREATE INDEX IF NOT EXISTS idx_ordens_servico_prefeitura ON ordens_servico(prefeitura_id);
CREATE INDEX IF NOT EXISTS idx_ordens_servico_status ON ordens_servico(status);
CREATE INDEX IF NOT EXISTS idx_ordens_servico_responsavel ON ordens_servico(responsavel_id);
CREATE INDEX IF NOT EXISTS idx_ordens_servico_solicitante ON ordens_servico(solicitante_id);
CREATE INDEX IF NOT EXISTS idx_ordens_servico_data_criacao ON ordens_servico(data_criacao);
CREATE INDEX IF NOT EXISTS idx_ordens_servico_prioridade ON ordens_servico(prioridade);

CREATE INDEX IF NOT EXISTS idx_eventos_os_os_id ON eventos_os(os_id);
CREATE INDEX IF NOT EXISTS idx_eventos_os_tipo ON eventos_os(tipo);
CREATE INDEX IF NOT EXISTS idx_eventos_os_data_criacao ON eventos_os(data_criacao);

CREATE INDEX IF NOT EXISTS idx_usuarios_prefeitura ON usuarios(prefeitura_id);
CREATE INDEX IF NOT EXISTS idx_usuarios_email ON usuarios(email);

CREATE INDEX IF NOT EXISTS idx_notificacoes_usuario ON notificacoes(usuario_id);
CREATE INDEX IF NOT EXISTS idx_notificacoes_lida ON notificacoes(lida);

-- Trigger para atualizar updated_at automaticamente
CREATE OR REPLACE FUNCTION update_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = NOW();
    RETURN NEW;
END;
$$ language 'plpgsql';

CREATE TRIGGER update_ordens_servico_updated_at BEFORE UPDATE ON ordens_servico
    FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();

CREATE TRIGGER update_usuarios_updated_at BEFORE UPDATE ON usuarios
    FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();

CREATE TRIGGER update_prefeituras_updated_at BEFORE UPDATE ON prefeituras
    FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();

CREATE TRIGGER update_slas_updated_at BEFORE UPDATE ON slas
    FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();
