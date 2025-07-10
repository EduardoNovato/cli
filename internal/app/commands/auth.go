package commands

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// CheckAuth verifica si el usuario está autenticado con GitHub CLI
func CheckAuth() bool {
	cmd := exec.Command("gh", "auth", "status")
	return cmd.Run() == nil
}

// Authenticate maneja el flujo de autenticación interactivo
func Authenticate() error {
	fmt.Println("Autenticación requerida para GitHub")
	fmt.Println("Por favor sigue los pasos:")

	cmd := exec.Command("gh", "auth", "login", "--hostname", "github.com", "--web")
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("error en autenticación: %v", err)
	}

	return nil
}

// GetGitHubUser obtiene el nombre de usuario autenticado
func GetGitHubUser() (string, error) {
	cmd := exec.Command("gh", "api", "user", "--jq", ".login")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("error obteniendo usuario: %v", err)
	}
	return strings.TrimSpace(string(output)), nil
}
