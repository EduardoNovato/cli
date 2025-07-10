# GitFlow CLI

GitFlow is a command-line tool designed to simplify GitHub workflows through interactive commands.

## Features

- **Create GitHub repositories interactively**: Answer prompts to set up new repositories with your preferred visibility and optional README initialization.
- **Authentication with GitHub**: Handles authentication using the GitHub CLI (`gh`).
- **User-friendly prompts**: Uses interactive questions for repository details.

## Requirements

- [Go](https://golang.org/) 1.18+
- [GitHub CLI (`gh`)](https://cli.github.com/) installed and authenticated

## Installation

Clone the repository and build the binary:

```sh
git clone https://github.com/EduardoNovato/gitflow.git
cd gitflow
go build -o gitflow
```

## Usage

Run the CLI:

```sh
./gitflow repo
```

You will be prompted for:
- Repository name
- Visibility (public/private)
- Whether to initialize with a README

If you are not authenticated, the tool will guide you through the GitHub authentication process.

## Troubleshooting

- If you encounter authentication issues, try running `gh auth login` manually.
- Ensure you have a stable internet connection.
