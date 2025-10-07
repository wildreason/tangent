# Contributing to Tangent

Thank you for your interest in contributing to Tangent! We welcome contributions from the community.

## Ways to Contribute

- 🐛 Report bugs
- 💡 Suggest features
- 📝 Improve documentation
- 🎨 Add new character designs to the library
- 🔧 Fix issues
- ✨ Add new features

## Getting Started

### 1. Fork the Repository

```bash
# Fork on GitHub, then clone
git clone https://github.com/YOUR_USERNAME/tangent.git
cd tangent
```

### 2. Set Up Development Environment

```bash
# Ensure Go 1.21+ is installed
go version

# Run tests
go test ./...

# Build tangent CLI
cd cmd/tangent
go build -o tangent .
./tangent
```

### 3. Make Your Changes

Create a feature branch:
```bash
git checkout -b feature/your-feature-name
```

## Development Guidelines

### Code Style

- Follow standard Go conventions (`gofmt`, `go vet`)
- Write clear, descriptive commit messages
- Add comments for complex logic
- Keep functions focused and testable

### Testing

```bash
# Run all tests
go test ./...

# Run specific package tests
go test ./pkg/characters

# Test the CLI
cd cmd/tangent && go run main.go session.go
```

### Examples

If adding features, include examples in `examples/`:
```bash
examples/
└── your_feature/
    ├── main.go
    └── README.md
```

## Submitting Changes

### 1. Commit Your Changes

```bash
git add .
git commit -m "feat: add awesome feature

- Detailed description
- What problem it solves
- Any breaking changes"
```

Use conventional commits:
- `feat:` - New feature
- `fix:` - Bug fix
- `docs:` - Documentation
- `refactor:` - Code refactoring
- `test:` - Adding tests
- `chore:` - Maintenance

### 2. Push to Your Fork

```bash
git push origin feature/your-feature-name
```

### 3. Create Pull Request

- Go to https://github.com/wildreason/tangent
- Click "New Pull Request"
- Select your branch
- Fill in the PR template

## Pull Request Guidelines

### Title
Use conventional commit format:
```
feat: add animation speed control
fix: handle EOF in pattern input
docs: improve quick start guide
```

### Description
Include:
- What changes were made
- Why the change is needed
- Screenshots/demos (if UI changes)
- Related issues (if any)

### Checklist
- [ ] Code follows project conventions
- [ ] Tests added/updated
- [ ] Documentation updated
- [ ] Examples work
- [ ] Commits are clean and descriptive

## Adding Library Characters

Want to add a character to the library?

### 1. Create Character File

```go
// pkg/characters/library/robot.go
package library

func Robot() LibraryCharacter {
    return LibraryCharacter{
        Name: "robot",
        Description: "Friendly robot with waving animation",
        Author: "Your Name",
        Frames: map[string][]string{
            "idle": {
                "__R6FFF6L__",
                "_T6FFFFF5T_",
                "___11_22___",
            },
            "wave": {
                "__R6FFF6L__",
                "__6FFFFF5T1",
                "___11_22___",
            },
        },
    }
}
```

### 2. Register in Library

```go
// pkg/characters/library/library.go
func init() {
    libraryChars["robot"] = Robot()
}
```

### 3. Add Documentation

Update `docs/LIBRARY.md` with your character details.

## Reporting Issues

### Bug Reports

Include:
- OS and Go version
- Steps to reproduce
- Expected vs actual behavior
- Error messages/screenshots

### Feature Requests

Describe:
- The problem you're trying to solve
- Your proposed solution
- Alternative solutions considered
- Examples of usage

## Code of Conduct

### Our Standards

- Be respectful and inclusive
- Welcome newcomers
- Accept constructive criticism
- Focus on what's best for the project

### Not Acceptable

- Harassment or discriminatory language
- Trolling or insulting comments
- Personal or political attacks
- Publishing private information

## Questions?

- 📧 Open an issue for questions
- 💬 Start a discussion on GitHub Discussions
- 🌐 Visit https://wildreason.com

## License

By contributing, you agree that your contributions will be licensed under the MIT License.

---

Thank you for contributing to Tangent! 🚀

© 2025 Wildreason, Inc.

