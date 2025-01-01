# File Watcher CLI

## Overview

**File Watcher CLI** is a command-line tool designed to monitor file changes and automatically restart processes based on these changes. It supports running scripts written in Go or JavaScript and is especially useful for development workflows that require frequent script re-executions during iteration.

---

## Features

- Watches a specific file for changes.
- Automatically restarts the associated process when the file is updated.
- Supports Go (`.go`) and JavaScript (`.js`) scripts.
- Configurable delay before restarting the process.
- Cross-platform compatibility.

---

## Installation

### Prerequisites

Ensure that you have the following installed on your system:

- [Go](https://golang.org/dl/) (for running the tool itself and `.go` scripts)
- [Node.js](https://nodejs.org/) (if you intend to run `.js` scripts)

### Build

1. Clone this repository:
   ```bash
   git clone https://github.com/abneribeiro/ptol-cli.git
   cd ptol-cli
   ```

2. Build the CLI:
   ```bash
   go build -o ptol
   ```

3. Move the binary to your system's PATH (optional):
   ```bash
   mv ptol /usr/local/bin/
   ```

---

## Usage

### Command

```bash
ptol [flags] <script>
```

### Flags

| Flag          | Description                        | Default         |
|---------------|------------------------------------|-----------------|
| `--delay`     | Delay before restarting the process | `2s` (2 seconds)|

### Example

#### Run a Go Script:
```bash
ptol my-script.go
```

#### Run a JavaScript File:
```bash
ptol --delay=5s app.js
```

---

## How It Works

1. The CLI initializes a file watcher using the [fsnotify](https://github.com/fsnotify/fsnotify) library.
2. When a modification is detected in the specified file:
   - A timer resets with the configured delay.
   - After the delay, the currently running process is terminated, and the script is restarted.
3. Output from the script is streamed directly to the terminal.

---

## Development

### Project Structure

```plaintext
.
├── cmd/
│   └── main.go       # Main entry point and command definition
├── go.mod            # Go module dependencies
└── README.md         # Project documentation
```

### Running Locally

To run the tool locally during development, use:

```bash
go run main.go <script>
```

---

## Limitations

- Only supports `.go` and `.js` files.
- Does not currently support recursive directory watching.
- Restarts processes but does not handle complex process dependency trees.

---

## Contributing

Contributions are welcome! Please follow these steps:

1. Fork the repository.
2. Create a new feature branch:
   ```bash
   git checkout -b feature-name
   ```
3. Commit your changes and push the branch:
   ```bash
   git commit -m "Add feature description"
   git push origin feature-name
   ```
4. Open a pull request.

---

## License

This project is licensed under the [MIT License](LICENSE).

---

## Acknowledgments

- Built using [Cobra](https://github.com/spf13/cobra) for CLI commands.
- File watching powered by [fsnotify](https://github.com/fsnotify/fsnotify).

For questions or support, feel free to reach out at [your.email@example.com].
