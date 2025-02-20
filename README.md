# Process Orders

## 📌 Sobre o Projeto

O **Process Orders** é um sistema desenvolvido em **Golang** que processa pedidos e determina qual **Centro de Distribuição (CD)** deve ser utilizado para o envio de cada item.

### 🛠 Problema que resolve

Atualmente, a empresa de e-commerce possui apenas um centro de distribuição, mas está em processo de expansão. Nos próximos 6 meses, mais 5 CDs serão adicionados e, em 2 anos, mais 20 centros estarão operacionais. O sistema anterior não possuía uma inteligência para determinar a origem de envio dos itens de um pedido, tornando necessária uma solução otimizada e escalável.

## 🚀 Funcionalidades

- 🔄 Todas as operações são realizadas **em memória**.
- ⚡ Uso de **workers** para aumentar a performance na determinação dos CDs.
- 📊 Monitoramento com **Prometheus** e **Grafana**.

## 📥 Instalação

```bash
# Clone o repositório
git clone https://github.com/LordRadamanthys/teste_meli.git
cd teste_meli
```

## ▶️ Rodando o Projeto

```bash
# Iniciar os serviços com Docker
docker-compose up --build
```

## 📡 Uso da API

### 🔹 Processar um novo pedido

- **Endpoint:** `/orders`
- **Método:** `POST`
- **Descrição:** Processa um novo pedido.

#### 📩 Exemplo de Request Body:

```json
{
  "items": [
    {
      "id": <idItem>,
      "quantity": 2
    }
  ]
}
```

#### 📤 Exemplo de Resposta:

**HTTP Status:** `201 Created`

```json
{
  "order_id": "c0d8d243-4680-4f67-b980-afa02bf5cd1f"
}
```

### 🔹 Obter informações de um pedido

- **Endpoint:** `/orders/{idOrder}`
- **Método:** `GET`
- **Descrição:** Retorna os detalhes de um pedido pelo seu ID.

#### 📤 Exemplo de Resposta:

```json
{
  "id": "c0d8d243-4680-4f67-b980-afa02bf5cd1f",
  "items": {
    "processed_items": [
      {
        "id": "12345",
        "primary_distribution_center": "CD2",
        "distribution_centers": ["CD1", "CD2", "CD3"]
      }
    ],
    "not_processed_items": [
      {
        "id": "67890"
      }
    ]
  }
}
```

## 📊 Monitoramento

### 🔹 Métricas da aplicação

- **Endpoint:** `/metrics`
- **Método:** `GET`
- **Descrição:** Retorna as métricas da aplicação para monitoramento via Prometheus.

### 🔹 Configurando o Monitoramento

A aplicação utiliza o Prometheus para expor e coletar as metricas e Grafana como dash de visualização dessas metricas.
Para iniciar o **Prometheus** e o **Grafana**, execute:

```bash
docker-compose up -d
```

Verifique se todos os serviços estão em execução.

### 🔹 Acessando o Prometheus

- Abra um navegador e acesse: [http://localhost:9090](http://localhost:9090)

### 🔹 Acessando o Grafana

- Abra um navegador e acesse: [http://localhost:3000](http://localhost:3000)
- Faça login no **Grafana**:
  - **Usuário:** `admin`
  - **Senha:** `admin`
- Vá até **Dashboards** onde ja tem um dash chamado **Process orders Dashboard** que já esta configurado.

---

Feito com ❤️ por [LordRadamanthys](https://github.com/LordRadamanthys) 🚀

