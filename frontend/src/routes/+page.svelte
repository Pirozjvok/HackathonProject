<script>
    const login = 'admin'
    const password = 'admin'

    let login_input = ''
    let password_input = ''
    
    function logIn() {
        if (login_input === login && password === password_input) {
            document.location.href = "/calls"
        } else {
            alert('Неправильный логин или пароль')
        }
    }
</script>
<div class="container">
    <form class="window" on:submit={logIn}>
        <h1 class="window__title">
            Вход
        </h1>
        <div class="window__alert">
            <input type="text" class="window__login" placeholder="Введите логин" bind:value={login_input}>
            <input type="password" class="window__password" placeholder="Введите пароль" bind:value={password_input}>
        </div>
        <button class="window__button" type="submit">Войти</button>
    </form>
</div>

<style>
    * {
    box-sizing: border-box;
}

.container {
    max-width: 750px;
    height: 500px;
    margin: 0 auto;
    margin-top: 220px;
}
.window {
    height: 500px;
    border-radius: 60px;
    background: rgba(255, 255, 255, 0.95);
    text-align: center;
    margin: 0 auto;
}
.window__title {
    color: #000;
    font-size: 50px;
    font-weight: 400;
    padding-top: 30px;
}
.window__alert {
    text-decoration: none;
}
.window__login, .window__password {
    display: flex;
    width: 590px;
    padding: 13px 418px 12px 0px;
    align-items: center;
    border: none;
    border-bottom: 2px solid #000;
    margin: 0 auto;
    margin-top: 41px;
    background-color: transparent;
}
.window__button {
    margin-top: 82px;
    width: 260px;
    height: 60px;
    border-radius: 40px;
    background: #272525;
    box-shadow: 0px 10px 10px 0px rgba(0, 0, 0, 0.50);
    color: #FFF;
    font-size: 32px;
    font-weight: 400;
}

</style>
