# Tangent Architecture

## Overview

Tangent implements a layered architecture with clear separation of concerns, following domain-driven design principles. The architecture is designed to be maintainable, testable, and extensible while preserving zero external dependencies and single binary distribution.

## Architecture Layers

### 1. Domain Layer (`pkg/characters/domain/`)

The domain layer contains the core business logic and entities:

- **`character.go`**: Core business entities (`Character`, `Frame`, `CharacterSpec`, `FrameSpec`)
- **`errors.go`**: Domain-specific errors
- **`interfaces.go`**: Repository, compiler, and animation interfaces

**Key Principles:**
- Pure business logic with no external dependencies
- Domain entities are immutable where possible
- Clear interfaces define contracts for infrastructure

### 2. Service Layer (`pkg/characters/service/`)

The service layer orchestrates domain objects and handles application logic:

- **`character_service.go`**: Main service with dependency injection
- Validation logic
- Business rule enforcement
- Error propagation with context

**Key Principles:**
- Depends only on domain interfaces
- Handles complex business workflows
- Provides clean API for CLI and other consumers

### 3. Infrastructure Layer (`pkg/characters/infrastructure/`)

The infrastructure layer provides concrete implementations:

- **`pattern_compiler.go`**: Pattern compilation and validation
- **`file_repository.go`**: File-based character persistence
- **`animation_engine.go`**: Terminal animation implementation

**Key Principles:**
- Implements domain interfaces
- Handles external concerns (file I/O, terminal output)
- Can be easily swapped for testing or different implementations

### 4. Package Interface (`pkg/characters/`)

The main package provides backward-compatible API:

- **`characters.go`**: Convenience functions and compatibility layer
- **`builder.go`**: Legacy builder pattern (deprecated but maintained)
- **`registry.go`**: Global character registry
- **`animator.go`**: Legacy animation functions

**Key Principles:**
- Maintains backward compatibility
- Provides easy migration path
- Wraps new architecture for existing users

## Dependency Flow

```
CLI/Examples → Package Interface → Service Layer → Domain Layer
                    ↓                    ↓
              Infrastructure Layer ← Domain Interfaces
```

## Design Patterns Used

### 1. Dependency Injection
Services receive their dependencies through constructor injection, making them testable and flexible.

### 2. Interface Segregation
Small, focused interfaces (`CharacterRepository`, `PatternCompiler`, `AnimationEngine`) allow for easy testing and implementation swapping.

### 3. Repository Pattern
Abstracts data persistence behind a clean interface, allowing different storage implementations.

### 4. Service Layer Pattern
Encapsulates business logic and orchestrates domain objects.

### 5. Builder Pattern (Legacy)
Maintained for backward compatibility, though new code should use the service layer.

## Error Handling Strategy

### Domain Errors
- Defined in `pkg/characters/domain/errors.go`
- Provide clear, actionable error messages
- Include context about what went wrong

### Service Layer Errors
- Wrap domain errors with additional context
- Use `fmt.Errorf("...: %w", err)` for error chaining
- Provide helpful suggestions in CLI

### Infrastructure Errors
- Handle external concerns (file I/O, terminal issues)
- Convert to domain errors where appropriate
- Provide detailed error information for debugging

## Testing Strategy

### Unit Tests
- Each layer has comprehensive unit tests
- Mock dependencies for isolated testing
- Test both success and error cases

### Integration Tests
- Test service layer with real infrastructure
- Verify end-to-end workflows
- Test backward compatibility

### Test Structure
```
pkg/characters/
├── domain/
│   └── *_test.go
├── service/
│   └── *_test.go
└── infrastructure/
    └── *_test.go
```

## Migration Guide

### For New Projects
Use the service layer directly:

```go
import (
    "github.com/wildreason/tangent/pkg/characters/domain"
    "github.com/wildreason/tangent/pkg/characters/service"
    "github.com/wildreason/tangent/pkg/characters/infrastructure"
)

// Create service with dependencies
repo := infrastructure.NewFileCharacterRepository("/path/to/storage")
compiler := infrastructure.NewPatternCompiler()
animationEngine := infrastructure.NewAnimationEngine()
service := service.NewCharacterService(repo, compiler, animationEngine)

// Use service
spec := domain.NewCharacterSpec("my-char", 5, 3)
spec.AddFrame("idle", []string{"FRF", "LRL", "FRF"})
character, err := service.CreateCharacter(*spec)
```

### For Existing Projects
Continue using the legacy API - it's fully backward compatible:

```go
import "github.com/wildreason/tangent/pkg/characters"

// This still works exactly as before
char, err := characters.NewCharacterSpec("my-char", 5, 3).
    AddFrame("idle", []string{"FRF", "LRL", "FRF"}).
    Build()
```

## Benefits of New Architecture

### 1. Testability
- Each layer can be tested independently
- Easy to mock dependencies
- Clear separation of concerns

### 2. Maintainability
- Business logic is isolated in domain layer
- Infrastructure concerns are separated
- Easy to understand and modify

### 3. Extensibility
- New implementations can be added easily
- Different storage backends possible
- Animation engines can be swapped

### 4. Performance
- No unnecessary dependencies
- Efficient error handling
- Clean resource management

## Future Enhancements

### 1. Additional Infrastructure Implementations
- Database repository
- Memory repository
- Network repository

### 2. Advanced Animation Engines
- Web-based animation
- File-based animation export
- Custom animation formats

### 3. Enhanced Validation
- Pattern syntax validation
- Character dimension validation
- Animation parameter validation

### 4. Plugin System
- Custom pattern compilers
- Custom animation engines
- Custom repositories

## Conclusion

The new architecture provides a solid foundation for Tangent's future development while maintaining full backward compatibility. The layered design makes the codebase more maintainable, testable, and extensible, addressing the concerns raised in the architectural review.
