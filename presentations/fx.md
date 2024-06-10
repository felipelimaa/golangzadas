---
marp: true
theme: gaia
_class: lead
paginate: true
backgroundColor: #fff
backgroundImage: url('https://marp.app/assets/hero-background.svg')
---

# **Golang FX**

https://pkg.go.dev/go.uber.org/fx
https://uber-go.github.io/fx/

---
# **O que é**

É uma biblioteca que facilita a construção de aplicações que dependem de injeções de dependências e gerenciamento de componentes, oferecendo de maneira estruturada e modular a possibilidade de organização de código.

---
# **Principais Funcionalidades**

***Injeção de Dependências***

Declaração de dependências de forma declarativa, evitando a criação manual de instâncias de componentes e a gerenciar suas dependências de maneira mais eficiente.

---

# **Principais Funcionalidades**
***Ciclo de Vida de Componentes***

Ciclo de vida gerenciado para componentes da aplicação, permitindo que você defina fases de inicialização e finalização para seus serviços. Isso é útil para garantir que os recursos sejam corretamente inicializados e liberados.

---

# **Principais Funcionalidades**
***Configuração Simplificada***

Configuração centralizada e simplificada de componentes, ajudando a evitar a propagação de parâmetros de configuração por todo o código.

---

# **Principais Funcionalidades**
***Integração com Bibliotecas Populares***

Compatível com diversas bibliotecas populares no ecossistema Go, como `Uber's zap` para logging e prometheus para monitoramento.

---

# **Principais Funcionalidades**
***Testabilidade***

A injeção de dependências e a modularidade facilitam a escrita de testes unitários e de integração, uma vez que as dependências podem ser facilmente mockadas.

---
<style scoped>
p {
    font-size: 0.7rem;
}
</style>
# **Exemplo simples**

Instância de Logger

```go
// NewLogger cria uma nova instância de logger
func NewLogger() (*zap.Logger, error) {
    return zap.NewProduction()
}
```

---

# **Exemplo simples**
<style scoped>
pre {
   font-size: 0.8rem;
}
p {
    font-size: 0.7rem;
}
</style>
Instância do Mux

```go
// NewMux cria um novo http.ServeMux
func NewMux(lc fx.Lifecycle, logger *zap.Logger) *http.ServeMux {
    mux := http.NewServeMux()
    lc.Append(fx.Hook{
        OnStart: func() error {
            logger.Info("Starting HTTP server")
            go http.ListenAndServe(":8080", mux)
            return nil
        },
        OnStop: func() error {
            logger.Info("Stopping HTTP server")
            return nil
        },
    })
    return mux
}
```

---

# **Exemplo simples**
<style scoped>
pre {
   font-size: 1rem;
}
p {
    font-size: 0.7rem;
}
</style>
Registro de Handlers
```go
// RegisterHandlers registra handlers HTTP
func RegisterHandlers(mux *http.ServeMux, logger *zap.Logger) {
    mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        logger.Info("Handling request", zap.String("path", r.URL.Path))
        w.Write([]byte("Hello, FX!"))
    })
}
```

---

# **Exemplo simples**
<style scoped>
pre {
   font-size: 1rem;
}
p {
    font-size: 0.7rem;
}
</style>
Inicialização da aplicação
```go
func main() {
    app := fx.New(
        fx.Provide(
            NewLogger,
            NewMux,
        ),
        fx.Invoke(RegisterHandlers),
    )
    app.Run()
}
```

---

# **Comparativo - Sem uso do FX**
<style scoped>
pre {
   font-size: 0.53rem;
}
</style>
```go
func main() {
    // Configurar o logger
    logger := log.New(os.Stdout, "INFO: ", log.LstdFlags)

    // Configurar o servidor HTTP
    mux := http.NewServeMux()
    mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        logger.Println("Handling request:", r.URL.Path)
        w.Write([]byte("Hello, World!"))
    })

    // Iniciar o servidor HTTP
    server := &http.Server{
        Addr:    ":8080",
        Handler: mux,
    }

    go func() {
        logger.Println("Starting HTTP server on :8080")
        if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
            logger.Fatalf("Could not listen on :8080: %v\n", err)
        }
    }()

    if err := server.Close(); err != nil {
        logger.Fatalf("Server Close: %v", err)
    }

    logger.Println("Server gracefully stopped")
}
```

---

# **Comparativo - Com uso do FX**
<style scoped>
pre {
   font-size: 0.53rem;
   margin-top: -5px;
}
</style>
```go
func NewLogger() (*zap.Logger, error) {
    return zap.NewProduction()
}

func NewMux(lc fx.Lifecycle, logger *zap.Logger) *http.ServeMux {
    mux := http.NewServeMux()
    lc.Append(fx.Hook{
        OnStart: func() error {
            go http.ListenAndServe(":8080", mux)
            return nil
        },
        OnStop: func() error {
            return nil
        },
    })
    return mux
}

func RegisterHandlers(mux *http.ServeMux, logger *zap.Logger) {
    mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("Hello, FX!"))
    })
}

func main() {
    app := fx.New(
        fx.Provide(NewLogger, NewMux),
        fx.Invoke(RegisterHandlers),
    )
    app.Run()
}
```

---
<style scoped>
p, ul {
    font-size: 0.7rem;
}
</style>
# **Comparativo**
***Gerenciamento de Dependências***
- **Sem FX:** As dependências são criadas e gerenciadas manualmente. Isso pode levar a um código mais acoplado e menos modular.
- **Com FX:** As dependências são declaradas e injetadas automaticamente, resultando em um código mais modular e desacoplado.

---
<style scoped>
p, ul {
    font-size: 0.7rem;
}
</style>
# **Comparativo**

***Ciclo de Vida dos Componentes***
- **Sem FX:** O ciclo de vida (inicialização e finalização) dos componentes deve ser gerenciado manualmente, o que pode ser propenso a erros.
- **Com FX:** O ciclo de vida dos componentes é gerenciado automaticamente pelo FX, garantindo uma inicialização e finalização ordenadas e seguras.

---

<style scoped>
p, ul {
    font-size: 0.7rem;
}
</style>
# **Comparativo**

***Configuração e Manutenção***
- **Sem FX:** A configuração é espalhada pelo código, o que pode dificultar a manutenção e a leitura.
- **Com FX:** A configuração é centralizada, facilitando a manutenção e a leitura do código.

---

<style scoped>
p, ul {
    font-size: 0.7rem;
}
</style>
# **Comparativo**

***Escalabilidade***
- **Sem FX:** Adicionar novos componentes ou serviços pode requerer alterações significativas no código existente.
- **Com FX:** Adicionar novos componentes ou serviços é mais simples, pois o FX facilita a integração e a configuração de novos módulos.

---

<style scoped>
p {
    font-size: 0.7rem;
}
ul {
    font-size: 0.6rem;
}
</style>
# **"Concorrentes"**

**Wire (by Google)**
- Gerador de código que cria automaticamente injeções de dependência para o seu aplicativo Go. Ele usa diretivas declarativas em seu código para gerar o código de injeção de dependências.

**blabla**


---

# **Exemplo de Estruturação**
<style scoped>
pre {
   font-size: 0.8rem;
}
</style>

```
📦your-api
 ┣ 📂 cmd
 ┃ ┗ 📂 api
 ┃ ┃ ┗ 📜 main.go
 ┗ 📂 internal
   ┗ 📂 domain
     ┗ 📂 handler
     ┃ ┣ 📜 client_handler.go
     ┃ ┣ 📜 tenant_handler.go
     ┃ ┗ 📜 module.go
     ┗ 📂 repository
     ┃ ┣ 📜 client_repository.go
     ┃ ┣ 📜 tenant_repository.go
     ┃ ┗ 📜 module.go
     ┗ 📂 service
       ┣ 📜 client_service.go
       ┣ 📜 tenant_service.go
       ┗ 📜 module.go
```