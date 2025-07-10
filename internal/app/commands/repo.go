package commands

import (
	"fmt"
	"os/exec"
	"strings"
)

func CreateRepo(name, visibility string, initReadme bool) error {

	if !CheckAuth() {
		if err := Authenticate(); err != nil {
			return fmt.Errorf("error autenticando con GitHub: %v", err)
		}
	}

	args := []string{"repo", "create", name, "--" + visibility}

	if initReadme {
		args = append(args, "--add-readme")
	}

	cmd := exec.Command("gh", args...)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("error creando repositorio: %v\n%s", err, string(output))
	}

	fmt.Printf("Repositorio creado exitosamente: https://github.com/%s/%s\n", getGitHubUser(), name)
	return nil
}

func getGitHubUser() string {
	cmd := exec.Command("gh", "api", "user", "--jq", ".login")
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error obteniendo usuario de GitHub:", err)
		return "usuario-desconocido"
	}
	return strings.TrimSpace(string(output))
}
