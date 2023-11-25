<script>
import { onMount } from "svelte";

let list = [];
let draggedElement = null;

onMount(async () => {
    document.addEventListener('mousemove', onMouseMove);
    document.addEventListener('mouseup', onMouseUp);
    document.addEventListener('mousedown', onMouseDown);
    await getAll();
});


function onMouseDown(event) {
    let aaaa = document.elementsFromPoint(event.clientX, event.clientY);
    
    for (let element of aaaa) {
        if (element.id == 'call_data') {
            draggedElement = element;
            draggedElement.style.position = 'absolute';
            return
        }
    }
}

function onMouseMove(event) {
    if (draggedElement) {
        draggedElement.style.left = event.clientX + 'px';
        draggedElement.style.top = event.clientY + 'px';
    }
}

function onMouseUp(event) {
    if (draggedElement) {
        draggedElement.style.position = '';
        draggedElement.style.left = '';
        draggedElement.style.top = '';

        const targetTable = draggedElement.closest('.call_data_table');

        let target;

        document.elementsFromPoint(event.clientX, event.clientY).forEach((element) => {
            if (element.classList.contains('call_data_table')) {
                target = element;
            }
        });
        console.log(event)
        console.log(targetTable)
        if (target) {
            let state = draggedElement.dataset.status;
            let id = target.id;

            console.log(state)
            if (id == "new"){
              draggedElement = null;
              return;
            }
            if (id == "process" && state == "ok"){
              draggedElement = null;
              return;
            }

            draggedElement.dataset.status = target.id
            target.appendChild(draggedElement);
        }

        draggedElement = null;
    }
}

async function getAll() {
    let response = await fetch("http://185.192.247.23/get/all");

    if (response.ok) {
        list = await response.json();
    } else {
        alert("Ошибка HTTP: " + response.status);
    }
}

</script>
<div class="container">
        <div class="container_column">

            <div class="container_left">
                <div class="status_new">Новые</div>
                <div class="call_data_table" id="new">
                    {#each list as listitem}
                    {#if listitem.status_bell == "new"}
                    <div class="call_data" id="call_data" data-status="{listitem.status_bell}">
                        <div class="call_num_1">Звонок {listitem.id}</div>
                        <div class="call_date">Дата: {listitem.call_date.split(' ')[0]}</div>
                        <div class="call_time">Время: {listitem.call_date.split(' ')[1]}</div>
                        <div class="call_operator">Оператор: {listitem.operator_fio}</div>
                    </div>
                    {/if}
	                {/each}
                </div>
            </div>

            <div class="container_center">
                <div class="status_process">В работе</div>
                <div class="call_data_table" id="process">
                    {#each list as listitem}
                    {#if listitem.status_bell == "process"}
                    <div class="call_data" id="call_data" data-status="{listitem.status_bell}">
                        <div class="call_num_1">Звонок {listitem.id}</div>
                        <div class="call_date">Дата: {listitem.call_date.split(' ')[0]}</div>
                        <div class="call_time">Время: {listitem.call_date.split(' ')[1]}</div>
                        <div class="call_operator">Оператор: {listitem.operator_fio}</div>
                    </div>
                    {/if}
	                {/each}             
                </div>
            </div>

            <div class="container_right">
                <div class="status_done">Обработано</div>
                <div class="call_data_table" id="ok">
                    {#each list as listitem}
                    {#if listitem.status_bell == "ok"}
                    <div class="call_data" id="call_data" data-status="{listitem.status_bell}">
                        <div class="call_num_1">Звонок {listitem.id}</div>
                        <div class="call_date">Дата: {listitem.call_date.split(' ')[0]}</div>
                        <div class="call_time">Время: {listitem.call_date.split(' ')[1]}</div>
                        <div class="call_operator">Оператор: {listitem.operator_fio}</div>
                    </div>
                    {/if}
	                {/each}
                </div>
            </div>

        </div>
        <button on:mousedown={(e) => e.stopPropagation()} class="button button_bottom_right">Выгрузить с сервера</button>
    </div>

    <style>

    
  
  .container {
    display: grid;
    gap: 20px;
    padding: 20px;
  }

  .container {
  -webkit-touch-callout: none; /* iOS Safari */
    -webkit-user-select: none; /* Safari */
     -khtml-user-select: none; /* Konqueror HTML */
       -moz-user-select: none; /* Old versions of Firefox */
        -ms-user-select: none; /* Internet Explorer/Edge */
            user-select: none; /* Non-prefixed version, currently
                                  supported by Chrome, Edge, Opera and Firefox */
}
  
  .container_column {
    display: grid;
    grid-template-columns: repeat(3, 1fr);
    grid-template-rows: auto;
    gap: 20px;
    align-content: start;
    padding: 20px;
  }
  
  .container_left,
  .container_center,
  .container_right {
    margin: 10%;
    border-radius: 8px;
    font-size: 16px;
  }
  
  .container_left {
    grid-column: 1;
  }
  
  .container_center {
    grid-column: 2;
  }
  
  .container_right {
    grid-column: 3;
  }
  
  .call_data {
    border: 1px solid #eee;
    padding: 10px;
    background-color: white;
    border-radius: 8px;
    margin: 10px;
    cursor: grab;
    height: fit-content;
  }
  
.call_data_table {
    min-height: 100vh;
}

  .status_new,
  .status_process,
  .status_done {
    border-radius: 10px 10px 10px 10px;
    background: rgba(161, 169, 234, 0.80);
    text-align: center;
    height: 67px;
    display: flex;
    justify-content: center;
    align-items: center;
    font-size: 40px;
  }
  
  .status_process {
    border-radius: 10px 10px 10px 10px;
    background: rgba(255, 227, 126, 0.80);
  }
  
  .status_done {
    border-radius: 10px 10px 10px 10px;
    background: rgba(138, 241, 148, 0.80);
  }
  
  .call_date,
  .call_time,
  .call_operator {
    margin-bottom: 20px;
    font-size: 20px;
  }
  
  [class^="call_num_"] {
    font-size: 32px;
  }
  
  .button {
    border-radius: 50px;
    border: 3px solid #272525;
    background: rgba(255, 255, 255, 0.95);
    color: #000;
    font-family: Open Sans;
    font-size: 24px;
    font-style: normal;
    font-weight: 400;
    line-height: normal;
    cursor: pointer;
    
}

.button_bottom_right {
    position: fixed;
    bottom: 20px; 
    right: 20px; 
    width: 276px;
    height: 64px;
}


    </style>