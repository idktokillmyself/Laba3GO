package dbmsgo

import (
	"bytes"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewApplication(t *testing.T) {
	app := NewApplication()
	
	assert.NotNil(t, app)
	assert.NotNil(t, app.db)
	assert.NotNil(t, app.parser)
	assert.Equal(t, app.db, app.parser.db)
}

func TestApplication_RunCommandLine_Success(t *testing.T) {
	app := NewApplication()
	
	// Создаем временный файл
	filename := "test_command_line.txt"
	defer os.Remove(filename)
	
	// Тестируем выполнение команды
	err := app.RunCommandLine(filename, "CREATE ARRAY test_array")
	assert.NoError(t, err)
	
	// Проверяем, что структура создалась
	assert.NotNil(t, app.db.FindArray("test_array"))
	
	// Проверяем, что файл создался
	_, err = os.Stat(filename)
	assert.NoError(t, err)
}

func TestApplication_RunCommandLine_FileCreation(t *testing.T) {
	app := NewApplication()
	
	filename := "nonexistent_file.txt"
	defer os.Remove(filename)
	
	// Запускаем с несуществующим файлом - должна создаться новая база
	err := app.RunCommandLine(filename, "CREATE ARRAY new_array")
	assert.NoError(t, err)
	
	assert.NotNil(t, app.db.FindArray("new_array"))
}

func TestApplication_RunCommandLine_SaveError(t *testing.T) {
	app := NewApplication()
	
	// Пытаемся сохранить в невалидный путь
	err := app.RunCommandLine("/invalid/path/file.txt", "CREATE ARRAY test")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "ошибка сохранения файла")
}

func TestApplication_GetDatabase(t *testing.T) {
	app := NewApplication()
	
	db := app.GetDatabase()
	assert.NotNil(t, db)
	assert.Equal(t, app.db, db)
}

func TestPrintUsage(t *testing.T) {
	// Перехватываем stdout
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	printUsage()

	w.Close()
	os.Stdout = oldStdout

	var buf bytes.Buffer
	buf.ReadFrom(r)
	output := buf.String()

	assert.Contains(t, output, "Использование:")
	assert.Contains(t, output, "--file")
	assert.Contains(t, output, "--query")
	assert.Contains(t, output, "--help")
}

func TestApplication_InteractiveModeLogic(t *testing.T) {
	app := NewApplication()
	
	// Тестируем обработку команд без реального сканера
	commands := []string{
		"CREATE ARRAY test_array",
		"MPUSH test_array value1",
		"CREATE SLL test_sll", 
		"FPUSH_BACK test_sll sll_value",
		"HELP",
		"EXIT",
	}
	
	for _, cmd := range commands {
		app.parser.ProcessCommand(cmd)
	}
	
	// Проверяем, что структуры создались
	assert.NotNil(t, app.db.FindArray("test_array"))
	assert.NotNil(t, app.db.FindSLL("test_sll"))
	
	// Проверяем данные
	arr := app.db.FindArray("test_array")
	assert.Equal(t, 1, arr.Length())
	
	sll := app.db.FindSLL("test_sll") 
	assert.False(t, sll.IsEmpty())
}

func TestApplication_CommandProcessing(t *testing.T) {
	app := NewApplication()
	
	testCases := []struct {
		name     string
		command  string
		validate func(t *testing.T, db *Database)
	}{
		{
			name:    "Create array",
			command: "CREATE ARRAY test_array",
			validate: func(t *testing.T, db *Database) {
				assert.NotNil(t, db.FindArray("test_array"))
			},
		},
		{
			name:    "Create stack",
			command: "CREATE STACK test_stack",
			validate: func(t *testing.T, db *Database) {
				assert.NotNil(t, db.FindStack("test_stack"))
			},
		},
		{
			name:    "Create queue",
			command: "CREATE QUEUE test_queue",
			validate: func(t *testing.T, db *Database) {
				assert.NotNil(t, db.FindQueue("test_queue"))
			},
		},
		{
			name:    "Create tree",
			command: "CREATE TREE test_tree",
			validate: func(t *testing.T, db *Database) {
				assert.NotNil(t, db.FindTree("test_tree"))
			},
		},
		{
			name:    "Create hash table",
			command: "CREATE HASH test_hash",
			validate: func(t *testing.T, db *Database) {
				assert.NotNil(t, db.FindHashTable("test_hash"))
			},
		},
	}
	
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			app.parser.ProcessCommand(tc.command)
			tc.validate(t, app.db)
		})
	}
}

func TestApplication_ComplexWorkflow(t *testing.T) {
	app := NewApplication()
	
	// Комплексный рабочий процесс
	workflow := []string{
		"CREATE ARRAY users",
		"MPUSH users Alice",
		"MPUSH users Bob",
		"CREATE STACK history", 
		"SPUSH history action1",
		"SPUSH history action2",
		"CREATE HASH config",
		"HINSERT config setting1 value1",
		"HINSERT config setting2 value2",
	}
	
	for _, cmd := range workflow {
		app.parser.ProcessCommand(cmd)
	}
	
	// Проверяем результаты
	users := app.db.FindArray("users")
	assert.Equal(t, 2, users.Length())
	
	history := app.db.FindStack("history")
	assert.Equal(t, 2, history.GetSize())
	
	config := app.db.FindHashTable("config")
	assert.Equal(t, 2, config.GetSize())
}

func TestApplication_ErrorCommands(t *testing.T) {
	app := NewApplication()
	
	// Команды, которые должны обрабатываться без паники
	errorCommands := []string{
		"",                      // пустая команда
		"   ",                   // пробелы
		"UNKNOWN_COMMAND",       // неизвестная команда
		"CREATE",                // недостаточно параметров
		"CREATE INVALID type",   // неверный тип
		"MPUSH nonexist value",  // несуществующая структура
	}
	
	for _, cmd := range errorCommands {
		assert.NotPanics(t, func() {
			app.parser.ProcessCommand(cmd)
		})
	}
}

func TestMain_Integration(t *testing.T) {
	// Сохраняем оригинальные аргументы и stdout
	oldArgs := os.Args
	oldStdout := os.Stdout
	defer func() {
		os.Args = oldArgs
		os.Stdout = oldStdout
	}()
	
	// Тест 1: help команда
	os.Args = []string{"main", "--help"}
	
	r, w, _ := os.Pipe()
	os.Stdout = w
	
	// Запускаем в горутине так как main может вызвать os.Exit
	done := make(chan bool)
	go func() {
		main()
		done <- true
	}()
	
	// Даем время на выполнение
	<-done
	
	w.Close()
	os.Stdout = oldStdout
	
	var buf bytes.Buffer
	buf.ReadFrom(r)
	output := buf.String()
	
	// Проверяем что вывод содержит справку
	assert.Contains(t, output, "Использование:")
}

func TestApplication_FileIOCreation(t *testing.T) {
	app := NewApplication()
	
	// Проверяем что FileIO создается в CommandParser
	assert.NotNil(t, app.parser.fileIO)
}

func TestApplication_ScannerInitialization(t *testing.T) {
	app := NewApplication()
	
	// До вызова RunInteractive сканер должен быть nil
	assert.Nil(t, app.scanner)
}

func TestApplication_MultipleInstances(t *testing.T) {
	// Тестируем создание нескольких экземпляров приложения
	app1 := NewApplication()
	app2 := NewApplication()
	
	assert.NotNil(t, app1)
	assert.NotNil(t, app2)
	assert.NotEqual(t, app1, app2)
	assert.NotEqual(t, app1.db, app2.db) // Базы должны быть разными
}

func TestApplication_EmptyDatabase(t *testing.T) {
	app := NewApplication()
	
	// Новая база должна быть пустой
	db := app.GetDatabase()
	assert.NotNil(t, db)
	
	// Проверяем что все коллекции пустые
	assert.Equal(t, 0, len(db.Arrays))
	assert.Equal(t, 0, len(db.SinglyLinkedLists))
	assert.Equal(t, 0, len(db.Stacks))
	assert.Equal(t, 0, len(db.Queues))
	assert.Equal(t, 0, len(db.Trees))
	assert.Equal(t, 0, len(db.HashTables))
}
