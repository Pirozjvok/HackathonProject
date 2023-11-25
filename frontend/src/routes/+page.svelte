<div class="container">
    <div class="window">
        <h1 class="window__title">
            Вход
        </h1>
        <div class="window__alert">
            <input type="text" class="window__login" placeholder="Введите логин">
            <input type="text" class="window__password" placeholder="Введите пароль">
        </div>
        <button class="window__button" type="submit">Войти</button>
    </div>
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