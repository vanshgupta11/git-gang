# LeetCode Auto Submission Tool (Go Implementation)

This is the Go implementation of the LeetCode Auto Submission tool, designed to produce a small, efficient executable.

## Features

- Minimal dependencies (uses standard library)
- Small binary size (no VM/runtime required)
- Cross-platform support
- Memory-efficient operation
- Same functionality as Node.js version

## Building from Source

1. Install Go 1.21 or later
2. Clone this repository
3. Build the binary:

```bash
cd leetcode-auto
go build ./cmd/leetcode-auto
```

This creates a single executable that you can run on your system.

## Usage

```bash
./leetcode-auto -session "your_session_token" -csrf "your_csrf_token"
```

Optional flags:
- `-max-questions int`: Maximum question number (default 3691)
- `-target int`: Target successful submissions (default 1)

## Development

Project structure:
```
leetcode-auto/
├── cmd/
│   └── leetcode-auto/      # Main entry point
├── internal/
│   ├── api/               # LeetCode API client
│   ├── config/            # Configuration
│   ├── models/            # Data structures
│   └── submitter/         # Core logic
└── go.mod
```

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Run tests: `go test ./...`
5. Submit a pull request

## License

Same as the original project.