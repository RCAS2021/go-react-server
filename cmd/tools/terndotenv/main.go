package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/joho/godotenv"
)

func main() {
	// Carregar variáveis de ambiente
	if err := godotenv.Load(); err != nil {
		fmt.Fprintf(os.Stderr, "Erro ao carregar .env: %v\n", err)
		os.Exit(1)
	}

	// Executar comando de migração
	cmd := exec.Command("tern", "migrate", "--migrations", "./internal/store/pgstore/migrations", "--config", "./internal/store/pgstore/migrations/tern.conf")
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Erro ao executar comando: %v\nSaída: %s\n", err, string(output))
		os.Exit(1)
	}

	fmt.Printf("Comando executado com sucesso.")
}
