package utils

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/AlecAivazis/survey/v2"
)

// RepoCreate crea un nuevo repositorio en GitHub
func RepoCreate(name, visibility string, initReadme bool) error {
	args := []string{"repo", "create", name, "--" + strings.ToLower(visibility)}
	if initReadme {
		args = append(args, "--add-readme")
	}
	cmd := exec.Command("gh", args...)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("error al crear repositorio: %v\n%s", err, string(output))
	}
	return nil
}
func AskRepoDetails() (name string, visibility string, initReadme bool, err error) {
	qs := []*survey.Question{
		{
			Name: "name",
			Prompt: &survey.Input{
				Message: "Nombre del repositorio:",
			},
			Validate: survey.Required,
		},
		{
			Name: "visibility",
			Prompt: &survey.Select{
				Message: "Visibilidad:",
				Options: []string{"Público", "Privado"},
				Default: "Público",
			},
		},
		{
			Name: "initReadme",
			Prompt: &survey.Confirm{
				Message: "¿Inicializar con README?",
				Default: true,
			},
		},
	}

	answers := struct {
		Name       string
		Visibility string
		InitReadme bool
	}{}

	err = survey.Ask(qs, &answers)
	visMap := map[string]string{"Público": "public", "Privado": "private"}
	return answers.Name, visMap[answers.Visibility], answers.InitReadme, err
}

// RepoDelete elimina un repositorio de GitHub
func RepoDelete(name string) error {
	args := []string{"repo", "delete", name, "--yes"}
	cmd := exec.Command("gh", args...)
	output, err := cmd.CombinedOutput()
	outputStr := string(output)
	if err != nil {
		if strings.Contains(outputStr, "delete_repo") {
			// Intentar refresh automático
			refreshCmd := exec.Command("gh", "auth", "refresh", "-h", "github.com", "-s", "delete_repo")
			refreshCmd.Stdin = os.Stdin
			refreshCmd.Stdout = os.Stdout
			refreshCmd.Stderr = os.Stderr
			refreshErr := refreshCmd.Run()
			if refreshErr != nil {
				return fmt.Errorf("permisos insuficientes y no se pudo refrescar el scope. Ejecuta manualmente: gh auth refresh -h github.com -s delete_repo")
			}
			// Reintentar eliminación
			cmd2 := exec.Command("gh", args...)
			output2, err2 := cmd2.CombinedOutput()
			if err2 != nil {
				return fmt.Errorf("error al eliminar repositorio tras refresh: %v\n%s", err2, string(output2))
			}
			return nil
		}
		return fmt.Errorf("error al eliminar repositorio: %v\n%s", err, outputStr)
	}
	return nil
}

func GetCurrentUser() string {
	cmd := exec.Command("gh", "api", "user", "--jq", ".login")
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error obteniendo usuario de GitHub:", err)
		return "usuario-desconocido"
	}
	return strings.TrimSpace(string(output))
}

// IsAuthenticated verifica si el usuario está autenticado
func IsAuthenticated() bool {
	cmd := exec.Command("gh", "auth", "status")
	return cmd.Run() == nil
}

// AuthLogin inicia el flujo de autenticación
func AuthLogin() error {
	cmd := exec.Command("gh", "auth", "login", "--web")
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func checkDeletePermissions() error {
	cmd := exec.Command("gh", "auth", "refresh", "-h", "github.com", "-s", "delete_repo")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("error al verificar permisos: %v\n%s", err, string(output))
	}
	return nil
}

func ConfirmDelete(repoName string) bool {
	var confirm bool
	prompt := &survey.Confirm{
		Message: fmt.Sprintf("¿Estás seguro de eliminar '%s'?", repoName),
		Default: false,
	}
	survey.AskOne(prompt, &confirm)
	return confirm
}

// RepoClone clona un repositorio de GitHub
func RepoClone(repoURL string) error {
	cmd := exec.Command("gh", "repo", "clone", repoURL)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

// AskRepoURL solicita al usuario la URL del repositorio a clonar
func AskRepoURL() (string, error) {
	var repoURL string
	prompt := &survey.Input{
		Message: "Introduce la URL del repositorio a clonar:",
	}
	err := survey.AskOne(prompt, &repoURL, survey.WithValidator(survey.Required))
	if err != nil {
		return "", fmt.Errorf("error al solicitar URL del repositorio: %v", err)
	}
	return repoURL, nil
}

// RepoList lista repositorios de un usuario u organización usando `gh repo list`
func RepoList(owner string, limit int, visibility string, isFork, isSource bool, language string) error {
	args := []string{"repo", "list"}

	// Owner (usuario u organización)
	if owner != "" {
		args = append(args, owner)
	}

	// Filtros de flags
	if limit > 0 {
		args = append(args, "--limit", fmt.Sprintf("%d", limit))
	}
	if visibility != "" {
		args = append(args, "--visibility", visibility) // public | private | internal
	}
	if isFork {
		args = append(args, "--fork")
	}
	if isSource {
		args = append(args, "--source")
	}
	if language != "" {
		args = append(args, "--language", language)
	}

	// Campos JSON válidos y JQ expression robusta
	args = append(args,
		"--json", "name,isPrivate,isFork,primaryLanguage,url",
		"--jq", `.[] | "\(.name) - \((if .isPrivate then "private" else "public" end)) - \((.primaryLanguage.name // "")) - \(.url)"`,
	)

	cmd := exec.Command("gh", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("error al listar repositorios: %v", err)
	}
	return nil
}
