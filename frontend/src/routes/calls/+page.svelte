<script>
import { onMount, afterUpdate } from "svelte";
import { BASE_API_URL } from "$lib/config.js"
import Loading from "./loading.svelte"

let list = [];
let draggedElement;
let loading;

onMount(async () => {
    document.addEventListener('mousemove', onMouseMove);
    document.addEventListener('mouseup', onMouseUp);
    document.addEventListener('mousedown', onMouseDown);
    await getAll();
});


afterUpdate(() => {
    let elements = document.getElementsByClassName("call_data")
    for (let element of elements) {
        element.addEventListener('dblclick', (event) => {
            if (event.target.dataset.id)
                document.location.href = `/call/${event.target.dataset.id}`
        });
    }
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

async function updateStatus(id, status){
    try {
        let response = await fetch(BASE_API_URL + "/update/" + id, {
            method: "POST",
            headers: {
                "Content-Type": "application/json"
            },
            body: JSON.stringify({
                id: Number(id),
                status_bell: status
            })
        });
        if (response.ok) {
            return true;
        } else {
            alert("Произошла ошибка при попытке обновить статус: " + response.status + " " + response.statusText)
        }
    } catch (err) {
        alert("Произошла ошибка при попытке обновить статус: " + err)
    }
}

async function onMouseUp(event) {
    if (draggedElement) {
        
        event.stopPropagation();

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
            let call_id = draggedElement.dataset.id
            console.log(state)

            if (id == state) {
                if (event.button == 0){
                    document.location.href = `/call/${call_id}`
                }
                return;
            }

            if (event.button != 2) {
                if (id == "new") {
                    draggedElement = null;
                    return;
                }
                if (id == "process" && state == "ok"){
                    draggedElement = null;
                    return;
                }
            }

            loading = "Обновление..."
            if (await updateStatus(call_id, id)){
                draggedElement.dataset.status = target.id
                target.appendChild(draggedElement);
            }
            loading = false
        }

        draggedElement = null;
    }
}

async function getAll() {

    try {
        loading = true
        let response = await fetch(BASE_API_URL + "/get/all");

        if (response.ok) {
            list = await response.json();
        } else {
            alert("Ошибка HTTP: " + response.status);
        }
    } catch (e) {
        alert("Ошибка HTTP: " + e);
    }
    finally {
        loading = false
    }


}
</script>

{#if loading}
    <Loading text={loading === true ? "Загрузка" : loading}></Loading>
{/if}

<div class="container">
        <div class="container_column">
            <div class="container_left">
                <div class="status_new">Новые</div>
                <div class="call_data_table" id="new">
                    {#each list as listitem}
                    {#if listitem.status_bell == "new"}
                    <div class="call_data" id="call_data" data-status="{listitem.status_bell}" data-id="{listitem.id}">
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
                    <div class="call_data" id="call_data" data-status="{listitem.status_bell}" data-id="{listitem.id}">
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
                    <div class="call_data" id="call_data" data-status="{listitem.status_bell}" data-id="{listitem.id}">
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
        <button on:mousedown={(e) => e.stopPropagation()} on:click={getAll} class="button button_bottom_right">Выгрузить с сервера</button>
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