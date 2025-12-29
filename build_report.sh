#!/bin/bash

echo "🚀 Компиляция и генерация отчета DBMS Go..."

# 1. Установка зависимостей
echo "📦 Установка зависимостей..."
go mod tidy

# 2. Компиляция проекта
echo "🔧 Компиляция проекта..."
go build -o dbms
if [ $? -eq 0 ]; then
    echo "✅ Компиляция успешна!"
else
    echo "❌ Ошибка компиляции!"
    exit 1
fi

# 3. Запуск тестов из всех пакетов
echo "🧪 Запуск тестов..."
go test -v -coverprofile=coverage.out ./...
TEST_RESULT=$?

if [ $TEST_RESULT -eq 0 ]; then
    echo "✅ Тесты прошли успешно!"
else
    echo "❌ Некоторые тесты не прошли!"
fi

# 4. Фильтрация coverage.out - исключаем только main.go
echo "🔍 Фильтрация отчета покрытия (исключая main.go)..."
if [ -f "coverage.out" ]; then
    # Создаем временный файл без main.go
    grep -v "main.go" coverage.out > coverage_filtered.out 2>/dev/null || true
    
    # Проверяем, не пустой ли отфильтрованный файл
    if [ -s "coverage_filtered.out" ]; then
        mv coverage_filtered.out coverage.out
        echo "✅ main.go исключен из отчета покрытия"
    else
        # Если отфильтрованный файл пуст, используем оригинальный
        rm -f coverage_filtered.out
        echo "⚠️  Фильтрация не удалась, используем оригинальный coverage"
    fi
else
    echo "⚠️  Файл coverage.out не найден"
fi

# 5. Генерация HTML отчета покрытия
echo "📈 Генерация отчета покрытия..."
if [ -f "coverage.out" ] && [ -s "coverage.out" ]; then
    go tool cover -html=coverage.out -o coverage.html
    echo "✅ Отчет покрытия создан"
else
    echo "⚠️  Не удалось создать отчет покрытия (файл coverage.out отсутствует или пуст)"
    # Создаем пустой HTML файл чтобы избежать ошибок
    cat > coverage.html << 'EOF'
<!DOCTYPE html>
<html>
<head>
    <title>Coverage Report</title>
</head>
<body>
    <h1>No coverage data available</h1>
    <p>Coverage report could not be generated.</p>
</body>
</html>
EOF
fi

# 6. Создание главного отчета
echo "📄 Создание главного отчета..."
cat > index.html << 'EOF'
<!DOCTYPE html>
<html>
<head>
    <title>DBMS Go - Отчет</title>
    <style>
        body { 
            font-family: Arial, sans-serif; 
            margin: 40px; 
            background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
            min-height: 100vh;
        }
        .container {
            max-width: 1200px;
            margin: 0 auto;
            background: white;
            padding: 40px;
            border-radius: 15px;
            box-shadow: 0 10px 30px rgba(0,0,0,0.2);
        }
        .header {
            text-align: center;
            margin-bottom: 40px;
        }
        .status {
            padding: 20px;
            border-radius: 10px;
            margin: 20px 0;
            text-align: center;
            font-size: 18px;
            font-weight: bold;
        }
        .success {
            background: #d4edda;
            color: #155724;
            border: 2px solid #c3e6cb;
        }
        .error {
            background: #f8d7da;
            color: #721c24;
            border: 2px solid #f5c6cb;
        }
        .warning {
            background: #fff3cd;
            color: #856404;
            border: 2px solid #ffeaa7;
        }
        .btn {
            display: inline-block;
            padding: 15px 30px;
            background: #007bff;
            color: white;
            text-decoration: none;
            border-radius: 8px;
            margin: 10px;
            font-size: 16px;
            transition: all 0.3s;
        }
        .btn:hover {
            background: #0056b3;
            transform: translateY(-2px);
        }
        .commands {
            background: #f8f9fa;
            padding: 20px;
            border-radius: 8px;
            margin: 20px 0;
        }
        pre {
            background: #1e1e1e;
            color: #00ff00;
            padding: 20px;
            border-radius: 8px;
            overflow-x: auto;
        }
        .note {
            background: #e7f3ff;
            color: #004085;
            padding: 15px;
            border-radius: 8px;
            margin: 20px 0;
            border: 1px solid #b8daff;
        }
    </style>
</head>
<body>
    <div class="container">
        <div class="header">
            <h1>🧪 DBMS Go - Отчет о компиляции и тестах</h1>
            <p>Полный отчет о состоянии проекта</p>
        </div>

        <div class="status success">
            ✅ Проект успешно скомпилирован!
        </div>

        <div class="note">
            <strong>ℹ️ Примечание:</strong> Отчет покрытия не включает main.go (точку входа)
        </div>

        <div style="text-align: center; margin: 30px 0;">
            <a href="coverage.html" class="btn">📊 Отчет покрытия кода</a>
        </div>

        <div class="commands">
            <h2>🚀 Команды для запуска:</h2>
            <pre>./dbms                    # Запуск в интерактивном режиме
./dbms --help            # Показать справку
go run main.go           # Запуск без компиляции</pre>
        </div>

        <div style="background: #e7f3ff; padding: 20px; border-radius: 8px; margin: 20px 0;">
            <h2>📋 Структура проекта:</h2>
            <ul>
                <li><strong>structures/</strong> - все структуры данных</li>
                <li><strong>serialization/</strong> - сериализация и файловый ввод-вывод</li>
                <li><strong>command/</strong> - парсер команд</li>
                <li><strong>tests/</strong> - модульные тесты</li>
                <li><strong>main.go</strong> - точка входа (исключено из покрытия)</li>
            </ul>
        </div>
    </div>
</body>
</html>
EOF

# 7. Открываем отчет
echo "✅ Все отчеты созданы!"
echo "📊 Открываю отчет в браузере..."

if command -v xdg-open > /dev/null; then
    xdg-open index.html
elif command -v open > /dev/null; then
    open index.html
else
    echo "📋 Отчеты созданы:"
    echo "   - Главный: file://$(pwd)/index.html"
    echo "   - Покрытие: file://$(pwd)/coverage.html"
fi

echo ""
echo "🎯 КОМАНДЫ ДЛЯ ЗАПУСКА:"
echo "   ./dbms                    # Запуск программы"
echo "   ./dbms --help             # Справка"
echo "   go run main.go           # Запуск без компиляции"
