document.addEventListener('DOMContentLoaded', function () {
    const loginForm = document.querySelector('#loginForm');

    loginForm.addEventListener('submit', function (e) {
        e.preventDefault(); // Предотвращаем отправку формы по умолчанию

        // Получаем значения полей ввода
        const login = document.querySelector('#login').value;
        const password = document.querySelector('#password').value;

        // Проверяем наличие данных в полях ввода
        if (!login || !password) {
            console.error('Заполните все поля   !');
            return; // Прекращаем выполнение функции, если данные отсутствуют
        }

        // Создаем объект с данными для отправки на сервер
        const formData = {
            login: login,
            password: password
        };

        // Отправляем данные на сервер для авторизации
        fetch('http://localhost:8080/authorization', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(formData)
        })
            .then(response => {
                // Обрабатываем ответ от сервера
                if (response.ok) {
                    
                } else {
                    throw new Error('Ошибка сервера');
                }
            })
            .then(data => {
                // Действия при успешной авторизации
                alert('Успешная авторизация:');
                window.location.href = "urlinfo.html";
            })
            .catch(error => {
                // Ошибка сети или другая ошибка
                console.error('Ошибка:', error);
            });
    });
});
