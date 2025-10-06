// web/js/script.js

document.getElementById('shorten-form').addEventListener('submit', async function (event) {
    event.preventDefault();
    const longUrlInput = document.getElementById('long-url');
    const longUrl = longUrlInput.value;
    const resultDiv = document.getElementById('result');
    const errorDiv = document.getElementById('error-message');

    // Сброс сообщений
    resultDiv.style.display = 'none';
    errorDiv.style.display = 'none';
    errorDiv.innerHTML = '';

    if (!longUrl) {
        errorDiv.style.display = 'flex';
        errorDiv.innerHTML = '❌ Пожалуйста, введите ссылку.';
        return;
    }

    // Сохраняем введенную ссылку перед отправкой
    const initialUrl = longUrl;

    try {
        const response = await fetch('/api/shorten', {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify({ long_url: longUrl })
        });

        const data = await response.json();

        if (response.ok) {
            const shortLink = data.short_url;

            resultDiv.style.display = 'flex';

            // --- ИЗМЕНЕНИЕ: УДАЛЕНА ХЛОПУШКА, ОСТАВЛЕНО ТОЛЬКО "Успех! ✨" ---
            resultDiv.innerHTML = `
                <div style="flex-grow: 1;">
                    Успех! ✨ Сокращенная ссылка: <a href="${shortLink}">${shortLink}</a>
                </div>
                <button id="copy-button" class="copy-button">📄 Копировать</button>
            `;

            // Ссылку оставляем в поле ввода
            longUrlInput.value = initialUrl;

            document.getElementById('copy-button').addEventListener('click', () => {
                navigator.clipboard.writeText(shortLink).then(() => {
                    alert('Ссылка скопирована!');
                }).catch(err => {
                    console.error('Ошибка копирования: ', err);
                    alert('Не удалось скопировать. Попробуйте вручную.');
                });
            });

        } else {
            errorDiv.style.display = 'flex';
            errorDiv.innerHTML = '❌ Ошибка: ' + (data.error || 'Неизвестная ошибка сервера');
            longUrlInput.value = initialUrl; // Оставляем ссылку при ошибке
        }
    } catch (e) {
        errorDiv.style.display = 'flex';
        errorDiv.innerHTML = '❌ Ошибка подключения. Убедитесь, что сервер запущен.';
        longUrlInput.value = initialUrl; // Оставляем ссылку при ошибке
    }
});