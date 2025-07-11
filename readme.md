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

### Interactive Mode

Run the CLI and select an action from the interactive menu:

```sh
./gitflow
```

You will see:

```
? What do you want to do?  [Use arrows to move, type to filter]
> Create Repository
  Delete Repository
  Exit
```

#### Create Repository

If you select **Create Repository**, you will be prompted for:
- **Repository name**
- **Visibility** (`Public` / `Private`)
- **Initialize with README?** (confirmation)

#### Delete Repository

If you select **Delete Repository**, you will be prompted for:
- **Repository name**
- **Confirmation** before deletion

> **Warning:** You cannot delete protected repositories such as `production` or `main` (see `configs/default.yaml`).

### Direct Commands

Create a repository directly:

```sh
./gitflow repo
```

Delete a repository directly:

```sh
./gitflow delete <repo-name>
```

If you are not authenticated, the tool will guide you through the GitHub authentication process.

## Contributing

Contributions are welcome! Please open issues or pull requests. For major changes, open an issue first to discuss what you would like to change.

## Protected repositories

Repositorios listados en `configs/default.yaml` bajo `protected_repos` no pueden ser eliminados por seguridad.

## Troubleshooting

- If you encounter authentication issues, try running `gh auth login` manually.
- Ensure you have a stable internet connection.
