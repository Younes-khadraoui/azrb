# azrb

**azrb** is a cross-platform CLI cheat sheet manager written in Go. With azrb you can easily create, view, edit, delete, and list cheat sheets—simple Markdown files that store useful commands, tips, and reference material. Built with [Cobra](https://github.com/spf13/cobra) and [Viper](https://github.com/spf13/viper), azrb is designed to work on Linux, Windows, and macOS with minimal configuration.

## Table of Contents

- [Features](#features)
- [Installation](#installation)
  - [From Source](#from-source)
  - [Pre-built Binaries](#pre-built-binaries)
- [Configuration](#configuration)
- [Usage](#usage)
- [Release Process](#release-process)
- [Automated Builds with GoReleaser](#automated-builds-with-goreleaser)
- [Contributing](#contributing)
- [License](#license)
- [Contact](#contact)

## Features

- **View Cheat Sheets**: Run `azrb <cheatsheet-name>` to read and display a cheat sheet.
- **Create Cheat Sheets**: Use `azrb create <cheatsheet-name>` to create a new cheat sheet (opens your editor automatically).
- **Edit Cheat Sheets**: Run `azrb edit <cheatsheet-name>` to modify an existing cheat sheet.
- **Delete Cheat Sheets**: Remove a cheat sheet with `azrb delete <cheatsheet-name>`.
- **List Cheat Sheets**: See all available cheat sheets using `azrb list`.
- **Show Configuration**: Run `azrb config` to display your current settings (storage directory and default editor).

## Installation

### From Source

Make sure you have [Go](https://golang.org/) (version 1.18 or later) installed. Then clone the repository and build the binary:

```bash
git clone https://github.com/Younes-khadraoui/azrb.git
cd azrb
go build -o azrb .
```

Alternatively, install via:

```bash
go install github.com/Younes-khadraoui/azrb@latest
```

Ensure your `$GOPATH/bin` is in your system PATH so that you can run `azrb` from anywhere.

### Pre-built Binaries

Pre-built binaries for Linux, Windows, and macOS are available on the [GitHub Releases](https://github.com/Younes-khadraoui/azrb/releases) page. We use [GoReleaser](https://goreleaser.com/) to automate the build and release process.

## Configuration

By default, azrb uses the following settings:
- **Storage Directory**: `~/.azrb` (where cheat sheets are saved)
- **Default Editor**:
  - Unix-like systems: `nvim` (or the value of the `$EDITOR` environment variable)
  - Windows: `notepad`

You can override these settings in two ways:

1. **Using a Configuration File**:  
   Create a file named `~/.azrb.yaml` with contents like:

   ```yaml
   storageDir: "/path/to/your/cheatsheets"
   editor: "your_preferred_editor"
   ```

2. **Using Environment Variables**:

   ```bash
   export AZRB_STORAGE_DIR="/path/to/your/cheatsheets"
   export AZRB_EDITOR="your_preferred_editor"
   ```

## Usage

Once installed, you can use azrb with the following commands:

- **View a Cheat Sheet**:  
  ```bash
  azrb <cheatsheet-name>
  ```
  Example:
  ```bash
  azrb pandas
  ```

- **Create a Cheat Sheet**:  
  ```bash
  azrb create <cheatsheet-name>
  ```
  This command creates a new Markdown file (e.g., `~/.azrb/pandas.md`) and opens your default editor.

- **Edit a Cheat Sheet**:  
  ```bash
  azrb edit <cheatsheet-name>
  ```

- **Delete a Cheat Sheet**:  
  ```bash
  azrb delete <cheatsheet-name>
  ```

- **List All Cheat Sheets**:  
  ```bash
  azrb list
  ```

- **Show Current Configuration**:  
  ```bash
  azrb config
  ```

- **Help**:  
  Display help text:
  ```bash
  azrb --help
  ```

## Release Process

We follow semantic versioning. **To release version v1.0.0:**

1. **Tag your release**:

   ```bash
   git tag -a v1.0.0 -m "Release version v1.0.0"
   git push origin v1.0.0
   ```

2. **Run GoReleaser**:  
   With your tag in place and your `.goreleaser.yml` configured, run:

   ```bash
   goreleaser release --rm-dist
   ```

   GoReleaser will build your CLI tool for Linux, Windows, and macOS, create archives, and publish them to GitHub Releases under version **v1.0.0**.

## Automated Builds with GoReleaser

We use [GoReleaser](https://goreleaser.com/) to automate the building, archiving, and releasing process.

## Contributing

Contributions are welcome! To contribute:
1. Fork the repository.
2. Create a branch for your feature or bugfix.
3. Write your code and tests.
4. Open a pull request with a clear description of your changes.

Please follow our coding guidelines and include tests where applicable.

## License

azrb is released under the [MIT License](LICENSE).

## Contact

For questions, suggestions, or support, please open an issue on GitHub or contact [younes.khadraoui.pro@gmail.com](mailto:younes.khadraoui.pro@gmail.com).

---

Happy cheat sheet managing with azrb!