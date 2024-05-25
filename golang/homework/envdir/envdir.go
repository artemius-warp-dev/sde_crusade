// Домашнее задание
// Утилита envdir
// Цель: Реализовать утилиту envdir на Go. Эта утилита позволяет запускать программы получая переменные окружения из определенной директории. См man envdir Пример go-envdir /path/to/evndir command arg1 arg2
// Завести в репозитории отдельный пакет (модуль) для этого ДЗ
// Реализовать функцию вида ReadDir(dir string) (map[string]string, error), которая сканирует указанный каталог и возвращает все переменные окружения, определенные в нем.
// Реализовать функцию вида RunCmd(cmd []string, env map[string]string) int , которая запускает программу с аргументами (cmd) c переопределнным окружением.
// Реализовать функцию main, анализирующую аргументы командной строки и вызывающую ReadDir и RunCmd

// Протестировать утилиту.
// Тестировать можно утилиту целиком с помощью shell скрипта, а можно написать unit тесты на отдельные функции.
// Критерии оценки: Стандартные потоки ввода/вывода/ошибок должны пробрасываться в вызываемую программу.
// Код выхода утилиты envdir должен совпадать с кодом выхода программы.
// Код должен проходить проверки go vet и golint
// У преподавателя должна быть возможность скачать и установить пакет с помощью go get / go install

package main

import (
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func ReadDir(dir string) (map[string]string, error) {
	files, err := os.ReadDir(dir)
	envs := map[string]string{}
	if err != nil {
		log.Fatalln("Couldn't read directory", err)
		return envs, err
	}

	for _, file := range files {
		filename := file.Name()
		fullPath := filepath.Join(dir, filename)

		content, err := os.ReadFile(fullPath)
		if err != nil {
			log.Fatalln("Couldn't read data from file", err)
			continue
		}
		str_parts := strings.SplitN(string(content), "=", 2)
		if len(str_parts) != 1 {
			log.Fatalln("ENV file has wrong string format")
			continue
		}
		key := filename
		envs[key] = str_parts[0]

	}

	return envs, nil
}
func RunCmd(cmd []string, env map[string]string) int {
	for k, v := range env {
		os.Setenv(k, v)
	}
	var err error
	for _, c := range cmd {
		var cmd_string strings.Builder
		cmd_string.WriteString("./")
		cmd_string.WriteString(c)
		cc := exec.Command(cmd_string.String())
		cc.Stdout = os.Stdout
		cc.Stderr = os.Stderr
		cc.Stdin = os.Stdin
		err = cc.Start()
		if err != nil {
			log.Println("Fail during start command", err)
			return 1

		}
		err = cc.Wait()
		if err != nil {
			log.Println("Fail during wait end of the command", err)
			return 1
		}
	}
	return 0

}



func main() {
	args := os.Args
	envs, err := ReadDir(args[1])
	if err != nil {
		log.Fatalln(err)
	}

	RunCmd(args[2:], envs)

}
