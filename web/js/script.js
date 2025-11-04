document.getElementById('shorten-form').addEventListener('submit', async function (event) {
    event.preventDefault();
    const longUrlInput = document.getElementById('long-url');
    const resultDiv = document.getElementById('result');
    const errorDiv = document.getElementById('error-message');
    const originalInput = longUrlInput.value.trim();

    let longUrl = originalInput;

    resultDiv.style.display = 'none';
    errorDiv.style.display = 'none';
    errorDiv.innerHTML = '';

    if (!longUrl) {
        errorDiv.style.display = 'flex';
        errorDiv.innerHTML = '❌ Пожалуйста, введите ссылку.';
        return;
    }

    const hasProtocol = /^(https?|ftp|s?ftp|file|data):\/\//i.test(longUrl);

    if (!hasProtocol) {
        longUrl = 'https://' + longUrl;
    }

    const MAX_URL_LENGTH = 2000;
    if (longUrl.length > MAX_URL_LENGTH) {
        errorDiv.style.display = 'flex';
        errorDiv.innerHTML = '❌ Слишком длинная ссылка. Максимум ' + MAX_URL_LENGTH + ' символов.';
        return;
    }

    if (typeof punycode === 'undefined' || !punycode.toUnicode) {
        errorDiv.style.display = 'flex';
        errorDiv.innerHTML = '❌ Ошибка: Библиотека punycode.js не найдена. Пожалуйста, подключите её в HTML.';
        return;
    }

    try {
        const response = await fetch('/api/shorten', {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify({ long_url: longUrl })
        });

        const data = await response.json();

        if (response.ok) {
            const shortCode = data.short_url;

            const technicalHost = window.location.hostname;

            const displayHost = punycode.toUnicode(technicalHost);

            const displayLink = displayHost + '/' + shortCode;

            const copyLink = 'https://' + displayLink;

            const technicalLink = 'https://' + technicalHost + '/' + shortCode;

            resultDiv.style.display = 'flex';

            resultDiv.innerHTML = `
        <div style="flex-grow: 1;">
            Успех! ✨ Сокращенная ссылка: <a href="${technicalLink}">${displayLink}</a>
        </div>
        <button id="copy-button" class="copy-button"> Копировать</button>`;

            longUrlInput.value = originalInput;

            document.getElementById('copy-button').addEventListener('click', () => {

                navigator.clipboard.writeText(copyLink).then(() => {
                    alert('Ссылка скопирована!');
                }).catch(err => {
                    console.error('Ошибка копирования:', err);
                    alert('Не удалось скопировать. Попробуйте вручную.');
                });
            });

        } else {
            errorDiv.style.display = 'flex';
            errorDiv.innerHTML = '❌ Ошибка: ' + (data.error || 'Неизвестная ошибка сервера');
            longUrlInput.value = originalInput;
        }
    } catch (e) {
        errorDiv.style.display = 'flex';
        errorDiv.innerHTML = '❌ Ошибка подключения. Убедитесь, что сервер запущен.';
        longUrlInput.value = originalInput;
    }
});