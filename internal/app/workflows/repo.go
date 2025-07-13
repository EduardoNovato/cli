package workflows

import (
	"fmt"
	"strings"

	"github.com/EduardoNovato/gitflow/internal/app/utils"
)

func CreateRepoFlow() {
	fmt.Println("Creando nuevo repositorio")

	name := utils.Input("Nombre del repositorio:")
	if name == "" {
		fmt.Println("El nombre del repositorio no puede estar vacío.")
		return
	}
	// Validación de caracteres inválidos (solo ejemplo básico)
	if strings.ContainsAny(name, " !@#$%^&*()[]{};:'\"|\\<>,/?") {
		fmt.Println("El nombre contiene caracteres inválidos.")
		return
	}
	visibility := utils.Select("Visibilidad:", []string{"Público", "Privado"})
	visMap := map[string]string{"Público": "public", "Privado": "private"}
	addReadme := utils.Confirm("¿Agregar README?", true)

	err := utils.RepoCreate(name, visMap[visibility], addReadme)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Repositorio '%s' creado exitosamente!\n", name)
}

func DeleteRepoFlow() {
	fmt.Println("Eliminando repositorio")

	repo := utils.Input("Nombre del repositorio a eliminar:")
	if repo == "" {
		fmt.Println("El nombre del repositorio no puede estar vacío.")
		return
	}

	confirm := utils.Confirm(fmt.Sprintf("¿Estás seguro de eliminar '%s'?", repo), false)
	if !confirm {
		fmt.Println("Operación cancelada")
		return
	}

	err := utils.RepoDelete(repo)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Repositorio '%s' eliminado exitosamente.\n", repo)
}

func CloneRepoFlow() {
	fmt.Println("Clonando repositorio")

	repo := utils.Input("URL del repositorio a clonar:")
	if repo == "" {
		fmt.Println("La URL del repositorio no puede estar vacía.")
		return
	}

	err := utils.RepoClone(repo)
	if err != nil {
		fmt.Printf("Error al clonar el repositorio: %v\n", err)
		return
	}

	fmt.Printf("Repositorio '%s' clonado exitosamente.\n", repo)
}

func ListReposFlow() {
	fmt.Println("Listado de repositorios")

	defaultOwner := utils.GetCurrentUser()
	owner := utils.Input(fmt.Sprintf("Usuario u organización [%s]:", defaultOwner))
	if owner == "" {
		owner = defaultOwner
	}

	visibility := utils.Select("Filtrar por visibilidad:", []string{"Todos", "Público", "Privado", "Interno"})
	visMap := map[string]string{"Todos": "", "Público": "public", "Privado": "private", "Interno": "internal"}
	vis := visMap[visibility]

	filterType := utils.Select("¿Qué tipo de repos mostrar?", []string{"Todos", "Solo forks", "Solo originales"})
	isFork := false
	isSource := false
	if filterType == "Solo forks" {
		isFork = true
	}
	if filterType == "Solo originales" {
		isSource = true
	}

	lang := utils.Input("Filtrar por lenguaje (deja vacío para todos):")

	limitStr := utils.Input("¿Cuántos repositorios deseas mostrar? (por defecto 30):")
	limit := 30
	if strings.TrimSpace(limitStr) != "" {
		fmt.Sscanf(limitStr, "%d", &limit)
	}

	err := utils.RepoList(owner, limit, vis, isFork, isSource, lang)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}
