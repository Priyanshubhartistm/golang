# Contributing to Golang Learning Repository

Thank you for your interest in contributing to this Golang learning repository! This document provides guidelines for contributing to the project.

## 📋 Table of Contents

- [Code of Conduct](#code-of-conduct)
- [How Can I Contribute?](#how-can-i-contribute)
- [Getting Started](#getting-started)
- [Development Guidelines](#development-guidelines)
- [Submitting Changes](#submitting-changes)
- [Style Guidelines](#style-guidelines)

## Code of Conduct

This project and everyone participating in it is governed by our Code of Conduct. By participating, you are expected to uphold this code.

## How Can I Contribute?

### Reporting Bugs

- Check if the bug has already been reported in the Issues section
- Use a clear and descriptive title
- Provide detailed steps to reproduce the issue
- Include code samples and error messages if applicable

### Suggesting Enhancements

- Use a clear and descriptive title
- Provide a detailed description of the suggested enhancement
- Explain why this enhancement would be useful
- Include examples if applicable

### Adding New Examples

We welcome contributions of new Go examples! To add a new example:

1. Create a new directory with a descriptive name (follow the existing numbering scheme)
2. Add your Go code with clear comments explaining the concepts
3. Ensure the code follows Go best practices
4. Test your code to make sure it compiles and runs correctly
5. Update the main README.md to include your new example

### Improving Documentation

- Fix typos, grammar, or unclear explanations
- Add more detailed comments to existing code
- Improve the README or other documentation files
- Translate documentation (if applicable)

## Getting Started

1. **Fork the repository** on GitHub
2. **Clone your fork** locally:
   ```bash
   git clone https://github.com/YOUR-USERNAME/golang.git
   cd golang
   ```
3. **Create a new branch** for your contribution:
   ```bash
   git checkout -b feature/your-feature-name
   ```
4. **Make your changes** and commit them with clear commit messages
5. **Push to your fork**:
   ```bash
   git push origin feature/your-feature-name
   ```
6. **Open a Pull Request** from your fork to the main repository

## Development Guidelines

### Prerequisites

- Go 1.19 or higher installed
- Basic understanding of Git and GitHub
- A text editor or IDE (VS Code, GoLand, etc.)

### Code Structure

Each topic should be in its own directory with:
- Descriptive directory name
- One or more `.go` files with examples
- Clear comments explaining the concepts
- Working, runnable code

### Testing Your Code

Before submitting, make sure:
1. Your code compiles without errors:
   ```bash
   go build ./...
   ```
2. Your code runs as expected:
   ```bash
   go run your-file.go
   ```
3. Your code follows Go conventions:
   ```bash
   go fmt ./...
   go vet ./...
   ```

## Submitting Changes

### Pull Request Process

1. Ensure your code follows the style guidelines below
2. Update the README.md if you're adding new content
3. Write a clear pull request description explaining:
   - What you changed and why
   - Any relevant issue numbers (e.g., "Fixes #123")
   - Screenshots or examples if applicable
4. Wait for review from maintainers
5. Address any requested changes
6. Once approved, your PR will be merged!

### Commit Messages

- Use clear, descriptive commit messages
- Start with a verb (Add, Fix, Update, Remove, etc.)
- Keep the first line under 50 characters
- Add more details in the body if needed

Examples:
```
Add example for goroutines and channels
Fix typo in pointers documentation
Update README with new examples section
```

## Style Guidelines

### Go Code Style

- Follow standard Go formatting (use `go fmt`)
- Use meaningful variable and function names
- Add comments for complex logic
- Keep functions small and focused
- Follow Go naming conventions (camelCase for private, PascalCase for public)

### Comment Style

- Write comments in English
- Explain *why*, not just *what* the code does
- Use examples in comments when helpful
- Keep comments up-to-date with code changes

### Documentation Style

- Use Markdown for documentation files
- Keep lines under 100 characters when possible
- Use proper headings and formatting
- Include code examples in documentation

## Questions?

If you have questions about contributing, feel free to:
- Open an issue with the "question" label
- Reach out to the maintainers

Thank you for contributing to help others learn Go! 🎉
