<script>
  import { onMount } from "svelte";
    import AudioPlayer from "./audioPlayer.svelte";
    import Loading from "./loading.svelte"
    let loading = false
    let id;
    let chat;

    onMount(() => {
        id = document.location.href.split('/').slice(-1);
    });
</script>

{#if  loading}
    <Loading></Loading>
{/if}

<div class="container">
    <div class="arrow">
        <a href="/calls">
            <img src="/arrow.svg" alt="Назад">
        </a>
    </div>
    <div class="title">
        Звонок {id}
    </div>
    <div class="left">
        <div class="left__contacts">
            <div class="contact contact_contact">
                <div class="contact__ava">
                    <img src="/contact.svg" alt="">
                </div>
                <div class="contact__info">
                    <p>Контакт</p>
                    <p>Иванов Иван Иванович</p>
                    <p>89174563299</p>
                </div>
            </div>
            <div class="contact contact_operator">
                <div class="contact__ava">
                    <img src="/operator.svg" alt="">
                </div>
                <div class="contact__info">
                    <p>Оператор</p>
                    <p>Петр Петров Иванович</p>
                </div>
            </div>
        </div>
        <div class="left__footer">
            <div class="date">
                <p>Дата: 24.11.23</p>
                <p>Время: 10:47</p>
            </div>
            <button class="left__btn">
                Удалить карточку
            </button>
        </div>
    </div>
    <div class="audio-panel">
        <h2 class="audio-panel__title">Запись разговора</h2>
        <AudioPlayer class="player" src="/response.mp3"></AudioPlayer>

        <div class="audio-panel_btn-container">
            <button class="speak-transcription__btn" on:click={() => loading = true}>
                <p>Загрузить</p> 
                <p>Транскрибацию</p> 
            </button>
        </div>
    </div>     
    <div class="speak-transcription">
        <h2 class="speak-transcription__title">Транскрибация разговора</h2>
        

        <!-- <div class="speak-transcription__chat">
            <div class="speak-transcription__msg">
                <div class="speak-transcription__XZ"></div>
                <div class="speak-transcription__msg-bubble">
                    Сообщение
                </div>
                <div class="speak-transcription__ava">
                    <svg width="75" height="75" viewBox="0 0 75 75" fill="none" xmlns="http://www.w3.org/2000/svg">
                        <path d="M37.5 75C58.0107 75 75 58.047 75 37.5C75 16.9894 58.0107 0 37.4637 0C16.9531 0 0 16.9894 0 37.5C0 58.047 16.9894 75 37.5 75ZM37.5 43.4535C30.3848 43.4172 24.8669 37.4637 24.8669 29.6225C24.8306 22.2168 30.4211 16.0455 37.5 16.0455C44.5426 16.0455 50.0968 22.2168 50.0968 29.6225C50.0968 37.4637 44.5789 43.5261 37.5 43.4535ZM37.5 68.865C29.3683 68.865 20.7285 65.4889 15.2469 59.5716C19.3853 53.3277 27.6985 49.7338 37.5 49.7338C47.1926 49.7338 55.5421 53.2551 59.7168 59.5716C54.2352 65.4889 45.6317 68.865 37.5 68.865Z" fill="black"/>
                    </svg>      
                </div>
            </div>
            <div class="speak-transcription__msg speak-transcription__self">
                <div class="speak-transcription__XZ"></div>
                <div class="speak-transcription__msg-bubble">
                    Сообщение
                </div>
                <div class="speak-transcription__ava">
                    <svg xmlns="http://www.w3.org/2000/svg" width="75" height="75" viewBox="0 0 75 75" fill="none">
                        <circle cx="37.5" cy="37.5" r="37" fill="#FFFCFC" stroke="black"/>
                        <path d="M47.375 40.3437C47.375 40.8073 47.2376 41.2604 46.98 41.6459C46.7225 42.0313 46.3564 42.3317 45.9282 42.5091C45.4999 42.6865 45.0287 42.7329 44.574 42.6425C44.1194 42.552 43.7018 42.3288 43.374 42.001C43.0462 41.6732 42.823 41.2556 42.7325 40.801C42.6421 40.3463 42.6885 39.8751 42.8659 39.4468C43.0433 39.0186 43.3437 38.6525 43.7291 38.395C44.1146 38.1374 44.5677 38 45.0313 38C45.6529 38 46.249 38.2469 46.6885 38.6865C47.1281 39.126 47.375 39.7221 47.375 40.3437ZM30.9688 38C30.5052 38 30.0521 38.1374 29.6666 38.395C29.2812 38.6525 28.9808 39.0186 28.8034 39.4468C28.626 39.8751 28.5796 40.3463 28.67 40.801C28.7605 41.2556 28.9837 41.6732 29.3115 42.001C29.6393 42.3288 30.0569 42.552 30.5115 42.6425C30.9662 42.7329 31.4374 42.6865 31.8657 42.5091C32.2939 42.3317 32.66 42.0313 32.9175 41.6459C33.1751 41.2604 33.3125 40.8073 33.3125 40.3437C33.3125 39.7221 33.0656 39.126 32.626 38.6865C32.1865 38.2469 31.5904 38 30.9688 38ZM59.875 25.5V41.125C59.8709 44.8533 58.388 48.4278 55.7516 51.0641C53.1153 53.7004 49.5408 55.1834 45.8125 55.1875H30.1875C26.4592 55.1834 22.8847 53.7004 20.2484 51.0641C17.6121 48.4278 16.1291 44.8533 16.125 41.125V25.5C16.122 23.9766 16.6755 22.5046 17.6814 21.3606C18.6873 20.2166 20.0764 19.4794 21.5877 19.2875C23.0989 19.0956 24.6282 19.4622 25.8881 20.3185C27.148 21.1748 28.0519 22.4617 28.4297 23.9375H47.5703C47.9482 22.4617 48.852 21.1748 50.1119 20.3185C51.3718 19.4622 52.9011 19.0956 54.4124 19.2875C55.9236 19.4794 57.3127 20.2166 58.3186 21.3606C59.3245 22.5046 59.878 23.9766 59.875 25.5ZM52.0625 39.5625C52.0625 37.9049 51.404 36.3152 50.2319 35.1431C49.0598 33.971 47.4701 33.3125 45.8125 33.3125H30.1875C28.5299 33.3125 26.9402 33.971 25.7681 35.1431C24.596 36.3152 23.9375 37.9049 23.9375 39.5625V41.125C23.9375 42.7826 24.596 44.3723 25.7681 45.5444C26.9402 46.7165 28.5299 47.375 30.1875 47.375H45.8125C47.4701 47.375 49.0598 46.7165 50.2319 45.5444C51.404 44.3723 52.0625 42.7826 52.0625 41.125V39.5625Z" fill="#474545"/>
                    </svg>      
                </div>
            </div>
            <div class="speak-transcription__msg">             
                <div class="speak-transcription__XZ"></div>
                <div class="speak-transcription__msg-bubble">
                    Сообщение
                </div>
                <div class="speak-transcription__ava">
                    <svg width="75" height="75" viewBox="0 0 75 75" fill="none" xmlns="http://www.w3.org/2000/svg">
                        <path d="M37.5 75C58.0107 75 75 58.047 75 37.5C75 16.9894 58.0107 0 37.4637 0C16.9531 0 0 16.9894 0 37.5C0 58.047 16.9894 75 37.5 75ZM37.5 43.4535C30.3848 43.4172 24.8669 37.4637 24.8669 29.6225C24.8306 22.2168 30.4211 16.0455 37.5 16.0455C44.5426 16.0455 50.0968 22.2168 50.0968 29.6225C50.0968 37.4637 44.5789 43.5261 37.5 43.4535ZM37.5 68.865C29.3683 68.865 20.7285 65.4889 15.2469 59.5716C19.3853 53.3277 27.6985 49.7338 37.5 49.7338C47.1926 49.7338 55.5421 53.2551 59.7168 59.5716C54.2352 65.4889 45.6317 68.865 37.5 68.865Z" fill="black"/>
                    </svg>      
                </div>
            </div>
        </div> -->
    </div>   
</div>

<style>

h2 {
    font-size: 34px;
}

.container {
    display: grid;
    width: 100%;
    height: 100vh;
    grid-template-columns: 134px 0.6fr 1fr 132px;
    grid-template-rows: 17px 80px 360px 1fr 50px;
    padding-left: 40px;
}

.arrow {
    grid-row: 2;
    grid-column: 1;
}

.title {
    display: flex;
    align-items: center;
    font-size: 36px;
    padding: 0px 60px;
    grid-row: 2;
    grid-column: 2;
    width: fit-content;
    background-color: white;
    border-radius: 100px;
}

.left {
    display: flex;
    flex-direction: column;
    justify-content: space-between;
    grid-column: 2;
    grid-row: 3 / span 2;
}

.contact {
    display: flex;
    width: fit-content;
    background-color: white;
    border-radius: 1000px;
    padding: 40px 30px;
}

.contact__info {
    margin-left: 13px;
}

.contact__info>p {
    margin-top: 4px;
}

.contact_contact {
    margin-top: 71px;
}

.contact_operator {
    margin-top: 27px;
}

/* Date */

.date {
    background-color: white;
    border-radius: 40px;
    padding: 17px 28px;
    width: fit-content;
    margin-bottom: 35px;
}

.date p:first-child {
    margin-bottom: 12px;
}

.left__btn {
    font-size: 25px;
    padding: 20px 25px;
    border-radius: 100px;
    border: none;
    background-color: #D9D9D9;
    box-shadow: 0px 0px 3px 3px #D9D9D9;
    background-color: #D9D9D9;
}

/* Audio Panel */

.audio-panel {
    height: fit-content;
    width: 100%;
    background: white;
    padding: 30px 65px;
    border-radius: 105px;
    grid-row: 3;
    grid-column: 3;
    font-size: 20px;
}

.audio-panel__title {
    width: 100%;
    text-align: center;
    margin-bottom: 30px;
}

/* Speak Transcription */

.speak-transcription {
    margin-top: 48px;
    grid-column: 3;
    grid-row: 4;
    width: 100%;
    height: calc(100% - 48px) ;
    background: white;
    border-radius: 105px;
    padding: 40px;
    overflow-y: hidden;
}

.speak-transcription__title {
    margin-bottom: 40px;
    width: 100%;
    text-align: center;
    font-size: 34px;
}

.audio-panel_btn-container {
    display: flex;
    justify-content: center;
    margin-top: 20px;
}
.speak-transcription__btn {
    padding: 9px 20px;
    background-color: #272525;
    color: rgba(255, 255, 255, 0.90);
    text-align: center;
    font-family: Inter;
    font-size: 20px;
    font-style: normal;
    font-weight: 700;
    letter-spacing: -0.408px;
    border-radius: 200px;
    border: none;
    box-shadow: 0px 0px 3px 3px #272525;
}

.speak-transcription__chat {
    overflow-y: scroll;
    max-height: calc(100% - 48px - 50px);
    direction: rtl; 
    height: 100%;
    scrollbar-color: #565656 #D9D9D9;
}

.speak-transcription__chat::-webkit-scrollbar-track
{
	border-radius: 18px;
	background-color: #D9D9D9;
}

.speak-transcription__chat::-webkit-scrollbar
{
	width: 18px;
    border-radius: 18px;
}

.speak-transcription__chat::-webkit-scrollbar-thumb
{
	border-radius: 18px;
	background-color: #565656;
}

.speak-transcription__msg {
    direction: ltr; 
    display: flex;
    background-color: black;
    padding: 12px 6px;
    border-radius: 150px;
    margin-left: 16px;
    width: calc(100% - 16px) ;
}

.speak-transcription__msg:not(:first-child) {
    margin-top: 16px;
}

.speak-transcription__msg-bubble {
    background-color: white;
    padding: 12px;
    margin: 0px 24px;
    border-radius: 8px;
    flex: 1;
}

.speak-transcription__XZ {
    width: 75px;
}

.speak-transcription__ava {
    display: flex;
    justify-content: center;
    align-items: center;
    width: fit-content;
    height: 70px;
    border-radius: 50%;
    background-color: white;
}

.speak-transcription__self {
    flex-direction: row-reverse;
    background-color: #D7D7D7;
}
</style>