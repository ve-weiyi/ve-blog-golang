# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

ve-blog-golang is a modern full-stack blog system built with Go + Go-Zero microservices architecture. The project includes:
- **blog-gozero/**: Main Go-Zero microservices implementation (primary)
- **blog-gin/**: Gin framework implementation (single architecture reference)
- **pkg/**: Reusable utility packages and plugins
- **tools/**: Code generation tools
- **stompws/**: WebSocket chat room using Stomp protocol

## Development Commands

### Working in blog-gozero/ (Main Implementation)

```bash
# Install dependencies and tools
cd blog-gozero
make deps

# Start infrastructure services (MySQL, Redis, RabbitMQ)
make docker-up

# Stop infrastructure services
make docker-down

# Run services (in separate terminals)
make run-rpc      # Start RPC service first (port 9999)
make run-blog     # Start blog frontend API (port 9090)
make run-admin    # Start admin backend API (port 9091)

# Build all services
make build        # Outputs to bin/ directory

# Clean build artifacts
make clean
```

### Direct Service Execution

```bash
# Using local config files (development)
go run service/rpc/blog/blog.go -f service/rpc/blog/etc/blog.yaml
go run service/api/blog/blog.go -f service/api/blog/etc/blog.yaml
go run service/api/admin/admin.go -f service/api/admin/etc/admin.yaml

# Using remote config (production with Nacos)
go run service/rpc/blog/blog.go
go run service/api/blog/blog.go
go run service/api/admin/admin.go
```

### Code Generation (tools/)

```bash
cd tools

# Generate Gin API code from .api files
make gen-api-gin-api

# Generate Gin API code from swagger.json
make gen-api-gin-swagger

# Generate database models from SQL DDL file
make gen-model-ddl

# Generate database models from database connection
make gen-model-dsn

# Generate TypeScript API code from .api files
make gen-web-ts-api

# Generate TypeScript API code from swagger.json
make gen-web-ts-swagger
```

## Architecture

### Service Layers (blog-gozero/)

**Three-tier microservices architecture:**

1. **API Gateway Layer** (`service/api/`)
   - `admin/`: Backend management API (port 9091)
   - `blog/`: Frontend blog API (port 9090)
   - Handles HTTP requests, authentication, and routes to RPC services
   - Each API service has:
     - `proto/`: API definition files (`.api` format)
     - `etc/`: Configuration files (`.yaml`)
     - `internal/`: Business logic (handler, logic, svc)
     - `docs/`: Swagger documentation

2. **RPC Service Layer** (`service/rpc/blog/`)
   - Core business logic service (port 9999)
   - Provides gRPC interfaces for API layer
   - Structure:
     - `proto/`: Protobuf definitions (`.proto` files)
     - `client/`: Generated RPC clients
     - `internal/`: Business logic implementation
     - `etc/`: Configuration files

3. **Model Layer** (`service/model/`)
   - GORM database models
   - Direct database access layer
   - Shared across all services

### Infrastructure Layer (infra/)

Provides cross-cutting concerns:
- `gormlogx/`: GORM logging integration
- `grpcerrors/`: gRPC error handling
- `interceptorx/`: gRPC interceptors (auth, logging, tracing)
- `middlewarex/`: HTTP middlewares (CORS, auth, rate limiting)
- `permissionx/`: RBAC permission management
- `queryx/`: Query builders and helpers
- `requestx/`: Request parsing utilities
- `responsex/`: Response formatting
- `staticx/`: Static file serving
- `tokenx/`: JWT token management
- `tracex/`: Distributed tracing

### Shared Packages (pkg/)

Reusable components organized by purpose:
- `infra/`: Business infrastructure (config, database connections)
- `kit/`: Technical components (HTTP client, cache, message queue)
- `plugins/`: Feature plugins (AI integration, API docs, music player)
- `utils/`: Pure utility functions (string, time, crypto)

### Service Communication

- **API → RPC**: API services call RPC via gRPC
- **RPC Connection Modes**:
  - Direct mode (development): Hardcoded endpoints in config
  - ETCD mode (production): Service discovery via ETCD
- **Configuration**: Supports local YAML files or Nacos config center

## API Design Standards

The project follows RESTful API conventions documented in `API设计规范.md`:

- Standard CRUD: `GET /api/v1/users`, `POST /api/v1/users`, etc.
- Batch operations: `POST /api/v1/users/batch`, `DELETE /api/v1/users/batch`
- Current user endpoints: `/api/v1/users/me/*` for authenticated user operations
- Complex queries: `POST /api/v1/users/query`

## Configuration Files

### API Services
- `service/api/blog/etc/blog.yaml` - Blog frontend API config
- `service/api/admin/etc/admin.yaml` - Admin backend API config

### RPC Services
- `service/rpc/blog/etc/blog.yaml` - Blog RPC service config

### Key Configuration Sections
- Database connection (MySQL)
- Redis configuration
- RPC connection mode (direct/ETCD)
- Upload service (Qiniu OSS)
- AI proxy settings
- Logging and health check

## Database

- **Database**: MySQL 8.0+
- **ORM**: GORM
- **Initialization**:
  - `blog-veweiyi-init.sql`: Table structure
  - `blog-veweiyi-data.sql`: Initial data
  - `blog-veweiyi.sql`: Full database dump

## Important Development Notes

1. **Service Startup Order**: Always start RPC service before API services
2. **Config Files**: Development uses local YAML files; production uses Nacos
3. **Code Generation**: Use tools/ for generating API handlers, models, and TypeScript clients
4. **API Definitions**:
   - API layer uses `.api` files (Go-Zero format)
   - RPC layer uses `.proto` files (Protobuf)
5. **Swagger Docs**:
   - Blog API: http://localhost:9090/blog-api/v1/swagger/index.html
   - Admin API: http://localhost:9091/admin-api/v1/swagger/index.html

## Testing

The project uses standard Go testing:
```bash
go test ./...
go test -v ./path/to/package
```

## Commit Convention

Follow Conventional Commits:
- `feat`: New features
- `fix`: Bug fixes
- `docs`: Documentation updates
- `style`: Code formatting
- `refactor`: Code refactoring
- `test`: Test-related changes
- `chore`: Build/toolchain changes
