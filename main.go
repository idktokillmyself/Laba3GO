package dbmsgo

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Application struct {
	db      *Database
	parser  *CommandParser
	scanner *bufio.Scanner
}

func NewApplication() *Application {
	db := NewDatabase()
	parser := NewCommandParser(db)
	return &Application{
		db:     db,
		parser: parser,
	}
}

func (app *Application) RunInteractive() {
	fmt.Println("=== Система управления базами данных ===")
	fmt.Println("Введите HELP для списка команд")
	
	app.scanner = bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("> ")
		if !app.scanner.Scan() {
			break
		}
		
		command := strings.TrimSpace(app.scanner.Text())
		if command == "EXIT" {
			break
		}
		
		app.parser.ProcessCommand(command)
	}
	
	fmt.Println("До свидания!")
}

func (app *Application) RunCommandLine(filename, query string) error {
	fileIO := NewFileIO()
	
	// Загружаем базу данных из файла
	if err := fileIO.LoadDatabaseFromFile(app.db, filename); err != nil {
		fmt.Printf("Создан новый файл: %s\n", filename)
	}
	
	// Выполняем команду
	app.parser.ProcessCommand(query)
	
	// Сохраняем изменения
	if err := fileIO.SaveDatabaseToFile(app.db, filename); err != nil {
		return fmt.Errorf("ошибка сохранения файла: %s", filename)
	}
	
	fmt.Printf("Изменения сохранены в файл: %s\n", filename)
	return nil
}

func (app *Application) GetDatabase() *Database {
	return app.db
}

func printUsage() {
	fmt.Println("Использование:")
	fmt.Println("  go run main.go --file <filename> --query <command>")
	fmt.Println("  go run main.go (для интерактивного режима)")
	fmt.Println("  go run main.go --help")
}

func main() {
	app := NewApplication()

	if len(os.Args) > 1 {
		// Режим командной строки
		var filename, query string
		
		for i := 1; i < len(os.Args); i++ {
			switch os.Args[i] {
			case "--file":
				if i+1 < len(os.Args) {
					filename = os.Args[i+1]
					i++
				}
			case "--query":
				if i+1 < len(os.Args) {
					query = os.Args[i+1]
					i++
				}
			case "--help":
				printUsage()
				return
			}
		}
		
		if filename != "" && query != "" {
			if err := app.RunCommandLine(filename, query); err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			return
		} else {
			printUsage()
			return
		}
	}

	// Интерактивный режим
	app.RunInteractive()
}
