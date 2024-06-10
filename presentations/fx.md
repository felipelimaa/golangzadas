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
<style>
section {
  
}
</style>
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

# **Exemplo simples de Estruturação**
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