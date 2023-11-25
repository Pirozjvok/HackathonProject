<script>
    export let src = ''
    let ext_class;
    export { ext_class as class };
    import { onMount } from 'svelte';

    let audio;
    let duration = '00:00'
    let current = '00:00'
    let progress;
    let play = false

    function getTime(time) {
        let m = Math.floor(time / 60)
        let s = Math.floor(time % 60)

        let m1 = m % 10
        let m2 = Math.floor(m / 10)

        let s1 = s % 10
        let s2 = Math.floor(s / 10)
        
        return `${m2}${m1}:${s2}${s1}`
    }

    function playPause() { 
        if (play) {
            audio.pause()
        } else {
            audio.play()
        }
    }

    onMount(() => {
        audio.addEventListener('loadedmetadata', () => {
            duration = getTime(audio.duration)
            progress.value = 0
        }); 
        audio.addEventListener('timeupdate', () => {
            current = getTime(audio.currentTime)
            progress.value = audio.currentTime / audio.duration * 1000;
        })
        audio.addEventListener('play', () => {
            play = true
        })
        audio.addEventListener('pause', () => {
            play = false
        })
        progress.addEventListener('input', () => {
            audio.currentTime = progress.value / 1000 * audio.duration;
        });
        audio.load(src)
    }); 
</script>

<div class="audio-player {ext_class || '.'}">
    <audio src="{src}" bind:this={audio}></audio>
    <!-- <div class="audio-player__progress">
        <div class="audio-player__progress-bar"  bind:this={progress}></div>
    </div> -->
    <input type="range" id="range1" min="0" max="1000" value="0" bind:this={progress}/>
    <div class="audio-player__flex">
        <div class="audio-player__time">
            {current}
        </div>
        <div class="audio-player__tools">
            <button class="audio-player__btn" on:click={() => audio.currentTime -= 10}>
                <svg xmlns="http://www.w3.org/2000/svg" width="48" height="48" viewBox="0 0 48 48" fill="none">
                    <path d="M24 9.99997V1.99997L14 12L24 22V14C30.6 14 36 19.4 36 26C36 32.6 30.6 38 24 38C17.4 38 12 32.6 12 26H8C8 34.8 15.2 42 24 42C32.8 42 40 34.8 40 26C40 17.2 32.8 9.99997 24 9.99997ZM21.8 32H20.1V25.5L18.1 26.1V24.7L21.6 23.4H21.8V32ZM30.3 28.5C30.3 29.1 30.2 29.7 30.1 30.1C30 30.5 29.8 30.9 29.5 31.2C29.2 31.5 28.9 31.7 28.6 31.9C28.3 32.1 27.9 32.1 27.4 32.1C26.9 32.1 26.6 32 26.2 31.9C25.8 31.8 25.5 31.5 25.3 31.2C25.1 30.9 24.8 30.5 24.7 30.1C24.6 29.7 24.5 29.1 24.5 28.5V27C24.5 26.4 24.6 25.8 24.7 25.4C24.8 25 25 24.6 25.3 24.3C25.6 24 25.9 23.8 26.2 23.6C26.5 23.4 26.9 23.4 27.4 23.4C27.9 23.4 28.2 23.5 28.6 23.6C29 23.7 29.3 24 29.5 24.3C29.7 24.6 30 25 30.1 25.4C30.2 25.8 30.3 26.4 30.3 27V28.5ZM28.7 26.8C28.7 26.4 28.7 26.1 28.6 25.8C28.5 25.5 28.5 25.3 28.4 25.2C28.3 25.1 28.2 24.9 28 24.9C27.8 24.9 27.7 24.8 27.5 24.8C27.3 24.8 27.1 24.8 27 24.9C26.9 25 26.7 25.1 26.6 25.2C26.5 25.3 26.4 25.6 26.4 25.8C26.4 26 26.3 26.4 26.3 26.8V28.7C26.3 29.1 26.3 29.4 26.4 29.7C26.5 30 26.5 30.2 26.6 30.3C26.7 30.4 26.8 30.6 27 30.6C27.2 30.6 27.3 30.7 27.5 30.7C27.7 30.7 27.9 30.7 28 30.6C28.1 30.5 28.3 30.4 28.4 30.3C28.5 30.2 28.6 29.9 28.6 29.7C28.6 29.5 28.7 29.1 28.7 28.7V26.8Z" fill="white"/>
                </svg>
            </button>

            <button class="audio-player__btn" on:click={playPause}>
                {#if play}
                    <svg xmlns="http://www.w3.org/2000/svg" width="40" height="40" viewBox="0 0 40 40" fill="none">
                        <path d="M6.10576 34.3858C4.19556 32.5408 2.67192 30.3339 1.62374 27.8939C0.575556 25.4538 0.0238315 22.8294 0.000755135 20.1738C-0.0223212 17.5182 0.483713 14.8846 1.48933 12.4267C2.49495 9.96876 3.98001 7.73572 5.85786 5.85786C7.73572 3.98001 9.96876 2.49495 12.4267 1.48933C14.8846 0.483713 17.5182 -0.0223212 20.1738 0.000755135C22.8294 0.0238315 25.4538 0.575556 27.8939 1.62374C30.3339 2.67192 32.5408 4.19556 34.3858 6.10576C38.0289 9.87781 40.0448 14.9299 39.9992 20.1738C39.9537 25.4177 37.8503 30.434 34.1421 34.1421C30.434 37.8503 25.4177 39.9537 20.1738 39.9992C14.9299 40.0448 9.87781 38.0289 6.10576 34.3858ZM14.2458 12.2458V28.2458H18.2458V12.2458H14.2458ZM22.2458 12.2458V28.2458H26.2458V12.2458H22.2458Z" fill="white"/>
                    </svg>
                {:else}
                    <svg xmlns="http://www.w3.org/2000/svg" width="48" height="48" viewBox="0 0 48 48" fill="none">
                        <path d="M24 3.99997C12.95 3.99997 4 12.95 4 24C4 35.05 12.95 44 24 44C35.05 44 44 35.05 44 24C44 12.95 35.05 3.99997 24 3.99997ZM20 33V15L32 24L20 33Z" fill="white"/>
                    </svg>
                {/if}

            </button>

            <button class="audio-player__btn" on:click={() => audio.currentTime += 10}>
                <svg xmlns="http://www.w3.org/2000/svg" width="48" height="48" viewBox="0 0 48 48" fill="none">
                    <path d="M8 26C8 34.8 15.2 42 24 42C32.8 42 40 34.8 40 26H36C36 32.6 30.6 38 24 38C17.4 38 12 32.6 12 26C12 19.4 17.4 14 24 14V22L34 12L24 1.99997V9.99997C15.2 9.99997 8 17.2 8 26ZM21.7 32H20V25.5L18 26.1V24.7L21.5 23.4H21.7V32ZM30.2 28.5C30.2 29.1 30.1 29.7 30 30.1C29.9 30.5 29.7 30.9 29.4 31.2C29.1 31.5 28.8 31.7 28.5 31.9C28.2 32.1 27.8 32.1 27.3 32.1C26.8 32.1 26.5 32 26.1 31.9C25.7 31.8 25.4 31.5 25.2 31.2C25 30.9 24.7 30.5 24.6 30.1C24.5 29.7 24.4 29.1 24.4 28.5V27C24.4 26.4 24.5 25.8 24.6 25.4C24.7 25 24.9 24.6 25.2 24.3C25.5 24 25.8 23.8 26.1 23.6C26.4 23.4 26.8 23.4 27.3 23.4C27.8 23.4 28.1 23.5 28.5 23.6C28.9 23.7 29.2 24 29.4 24.3C29.6 24.6 29.9 25 30 25.4C30.1 25.8 30.2 26.4 30.2 27V28.5ZM28.5 26.8C28.5 26.4 28.5 26.1 28.4 25.8C28.3 25.5 28.3 25.3 28.2 25.2C28.1 25.1 28 24.9 27.8 24.9C27.6 24.9 27.5 24.8 27.3 24.8C27.1 24.8 26.9 24.8 26.8 24.9C26.7 25 26.5 25.1 26.4 25.2C26.3 25.3 26.2 25.6 26.2 25.8C26.2 26 26.1 26.4 26.1 26.8V28.7C26.1 29.1 26.1 29.4 26.2 29.7C26.3 30 26.3 30.2 26.4 30.3C26.5 30.4 26.6 30.6 26.8 30.6C27 30.6 27.1 30.7 27.3 30.7C27.5 30.7 27.7 30.7 27.8 30.6C27.9 30.5 28.1 30.4 28.2 30.3C28.3 30.2 28.4 29.9 28.4 29.7C28.4 29.5 28.5 29.1 28.5 28.7V26.8Z" fill="white"/>
                </svg>
            </button>
        </div>
        <div class="audio-player__time audio-player__time_duration">
            {duration}
        </div>
    </div>
</div>

<style>
    .audio-player {
        background-color: #D7D7D7;
        border-radius: 20px;
        padding: 35px;
    }

    .audio-player__progress {
        width: 100%;
        height: 8px;
        background-color: white;
        border-radius: 8px;
    }

    .audio-player__progress-bar {
        border-radius: 8px;
        height: 8px;
        background-color: black;
        width: 0%;
    }

    #range1 {
        -webkit-appearance: none;
        appearance: none; 
        width: 100%;
        height: 8px;
        cursor: pointer;
        outline: none;
        overflow: hidden;
        border-radius: 8px;
    }

    #range1::-webkit-slider-runnable-track {
        height: 15px;
        background: #FAFAFA;
        border-radius: 16px;
    }

    #range1::-moz-range-track {
        height: 15px;
        background: #FAFAFA;
        border-radius: 16px;
    }

    #range1::-webkit-slider-thumb {
        -webkit-appearance: none;
        appearance: none; 
        height: 15px;
        width: 1px;
        border-radius: 50%;
        box-shadow: -10000px 0 0 10000px black;
        background-color: black;
    }

    #range1::-moz-range-thumb {
        height: 15px;
        width: 1px;
        border-radius: 50%;
        box-shadow: -10000px 0 0 10000px black;
        background-color: black;
    }

    .audio-player__flex {
        display: flex;
        justify-content: space-between;
        margin-top: 8px;
    }

    .audio-player__time {
        color: white;
        width: 50px;
        text-align: left;
    }

    .audio-player__time_duration {
        text-align: right;
    }

    .audio-player__tools {
        display: flex;
        gap: 21px;
    }

    .audio-player__btn {
        display: flex;
        justify-content: center;
        align-items: center;
        outline: none;
        border: none;
        background: transparent;
        width: 48px;
        height: 48px;
    }

    .audio-player__btn:hover {
        cursor: pointer;
    }
</style>