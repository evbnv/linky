document.getElementById('shorten-form').addEventListener('submit', async function (event) {
    event.preventDefault();

    const longUrlInput = document.getElementById('long-url');
    const longUrl = longUrlInput.value;
    const resultDiv = document.getElementById('result');
    const errorDiv = document.getElementById('error-message');

    // Сброс предыдущих сообщений
    resultDiv.style.display = 'none';
    errorDiv.style.display = 'none';
    errorDiv.innerHTML = '';

    try {
        const response = await fetch('/api/shorten', {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            // Отправляем JSON с ключом long_url
            body: JSON.stringify({ long_url: longUrl })
        });

        // ВСЕГДА пытаемся получить JSON-тело, даже при ошибке, 
        // т.к. бэкенд отправляет структурированный JSON-ответ с ошибкой
        const data = await response.json();

        if (response.ok) {
            // --- УСПЕХ (HTTP Status 201) ---
            const shortLink = data.short_url;

            resultDiv.style.display = 'block';
            resultDiv.innerHTML = '✅ Сокращенная ссылка: <a href="' + shortLink + '">' + shortLink + '</a>';

            // Очистка поля ввода
            longUrlInput.value = '';

        } else {
            // --- ОШИБКА (HTTP Status 400, 500 и т.д.) ---
            errorDiv.style.display = 'block';
            // Показываем сообщение об ошибке, пришедшее из Go-бэкенда
            errorDiv.innerHTML = '❌ Ошибка: ' + (data.error || 'Неизвестная ошибка сервера');
        }
    } catch (e) {
        // Ошибка сети (сервер не запущен, CORS, и т.п.)
        errorDiv.style.display = 'block';
        errorDiv.innerHTML = '❌ Ошибка подключения. Убедитесь, что сервер запущен.';
    }
});