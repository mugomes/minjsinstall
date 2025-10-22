// Copyright (C) 2025 Murilo Gomes Julio
// SPDX-License-Identifier: GPL-2.0-only

// Site: https://www.mugomes.com.br

package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println()
		fmt.Println("MiNJSInstall")
		fmt.Println("Version: 2.0.0")
		fmt.Println()
		fmt.Println("Copyright (C) 2025 Murilo Gomes Julio")
		fmt.Println("License: GPL-2.0-only")
		fmt.Println()
		fmt.Println("Site: https://www.mugomes.com.br")
		fmt.Println()
		fmt.Println("------------- MiNJSInstall -------------")
		fmt.Println("Selecione uma opção:")
		fmt.Println("1. Instalar NVM")
		fmt.Println("2. Instalar Node")
		fmt.Println("3. Remover o NVM, NPM e Node")
		fmt.Println("4. Sair")
		fmt.Println("------------- /MiNJSInstall -------------")
		fmt.Println()

		fmt.Print("Opções: ")
		option, _ := reader.ReadString('\n')
		option = strings.TrimSpace(option)

		switch option {
		case "1":
			fmt.Print("Digite a versão do NVM: ")
			version, _ := reader.ReadString('\n')
			version = strings.TrimSpace(version)

			// Baixa e instala o NVM
			cmd := exec.Command("bash", "-c",
				fmt.Sprintf("wget -qO- https://raw.githubusercontent.com/nvm-sh/nvm/v%s/install.sh | bash", version))
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			if err := cmd.Run(); err != nil {
				fmt.Println("Erro ao instalar NVM:", err)
				continue
			}

			// Exporta o NVM_DIR
			nvmDir := os.ExpandEnv("${HOME}/.config/nvm")
			os.Setenv("NVM_DIR", nvmDir)

			// Carrega o NVM
			loadNVM := exec.Command("bash", "-c", fmt.Sprintf("source %s/nvm.sh", nvmDir))
			loadNVM.Stdout = os.Stdout
			loadNVM.Stderr = os.Stderr
			loadNVM.Run()

			fmt.Println("NVM instalado e configurado com sucesso!")

		case "2":
			// Instala o Node
			cmd := exec.Command("bash", "-c", "source ~/.config/nvm/nvm.sh && nvm install node")
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			if err := cmd.Run(); err != nil {
				fmt.Println("Erro ao instalar Node:", err)
			} else {
				fmt.Println("Node instalado com sucesso!")
			}

		case "3":
			// Remover NVM, NPM e Node
			nvmDir := os.ExpandEnv("${HOME}/.config/nvm")
			cmd := exec.Command("bash", "-c",
				fmt.Sprintf("rm -rf %s ~/.config/nvm ~/.npm ~/.bower", nvmDir))
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			
			bashrc := os.ExpandEnv("${HOME}/.bashrc")
			data, err := os.ReadFile(bashrc)
			if err != nil {
				fmt.Println("Erro ao ler .bashrc:", err)
				break
			}
		
			content := string(data)
			linesToRemove := []string{
				`export NVM_DIR="$HOME/.config/nvm"`,
				`[ -s "$NVM_DIR/nvm.sh" ] && \. "$NVM_DIR/nvm.sh"  # This loads nvm`,
				`[ -s "$NVM_DIR/bash_completion" ] && \. "$NVM_DIR/bash_completion"  # This loads nvm bash_completion`,
			}
		
			var newLines []string
			for _, line := range strings.Split(content, "\n") {
				trimmed := strings.TrimSpace(line)
				remove := false
				for _, bad := range linesToRemove {
					if trimmed == bad {
						remove = true
						break
					}
				}
				if !remove {
					newLines = append(newLines, line)
				}
			}
		
			newContent := strings.Join(newLines, "\n")
		
			if newContent != content {
				err = os.WriteFile(bashrc, []byte(newContent), 0644)
				if err != nil {
					fmt.Println("Erro ao atualizar .bashrc:", err)
				} else {
					fmt.Println("Linhas do NVM removidas do .bashrc com sucesso!")
				}
			} else {
				fmt.Println("Nenhuma linha do NVM encontrada no .bashrc.")
			}

		
			if err := cmd.Run(); err != nil {
				fmt.Println("Erro ao remover:", err)
			} else {
				fmt.Println("NVM, NPM e Node removidos com sucesso.")
			}

		case "4":
			fmt.Println("Saindo...")
			os.Exit(0)

		default:
			fmt.Println("Opção inválida. Por favor, escolha uma opção válida.")
		}
	}
}
