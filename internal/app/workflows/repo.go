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

	// Validación de repositorios protegidos
	protected := []string{"production", "main"} // Se puede cargar dinámicamente de configs/default.yaml
	for _, p := range protected {
		if strings.EqualFold(repo, p) {
			fmt.Printf("No se puede eliminar el repositorio protegido: '%s'\n", repo)
			return
		}
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
