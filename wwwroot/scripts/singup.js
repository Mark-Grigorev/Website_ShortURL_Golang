document.addEventListener('DOMContentLoaded', function () {
    const signupForm = document.querySelector('#signupForm');

    signupForm.addEventListener('submit', function (e) {
        e.preventDefault(); // Предотвращаем отправку формы по умолчанию

        // Получаем значения полей ввода
        const login = document.querySelector('#login').value;
        const password = document.querySelector('#password').value;
        const name = document.querySelector('#name').value;

        // Проверяем наличие данных в полях ввода
        if (!login || !password || !name) {
            console.error('Заполните все поля!');
            return; // Прекращаем выполнение функции, если данные отсутствуют
        }

        // Создаем объект с данными для отправки на сервер
        const formData = {
            login: login,
            password: password,
            name: name
        };

        // Отправляем данные на сервер для регистрации
        fetch('http://localhost:8080/v1/auth/registration', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(formData)
        })
            .then(response => {
                // Обрабатываем ответ от сервера
                if (response.ok) {
                    return response.json(); // Парсим JSON из ответа
                } else {
                    throw new Error('Ошибка сервера');
                }
            })
            .then(data => {
                // Действия при успешной регистрации
                alert('Успешная регистрация!');
                window.location.href = "login.html";
            })
            .catch(error => {
                // Ошибка сети или другая ошибка
                console.error('Ошибка:', error);
            });
    });
});
