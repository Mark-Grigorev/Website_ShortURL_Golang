 function shortenURL() {
            const inputUrl = document.getElementById("url").value;
            const apiUrl = "http://localhost:8080/v1/urlshort/urlshort"; //Адрес сервера

            // Отправляем запрос на сервер
            fetch(apiUrl, {
                method: 'POST',
                body: inputUrl,
            })
            .then(response => response.text())
            .then(shortUrl => {
                // Вставляем полученную короткую ссылку обратно в поле ввода
                document.getElementById("shortUrl").value = shortUrl;
            })
            .catch(error => console.error('Error:', error));
        }