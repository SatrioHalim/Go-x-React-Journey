const tasks = localStorage.getItem("tasks") ? JSON.parse(localStorage.getItem("tasks")) : [];

const taskList = document.getElementById("tasks-list");
const taskForm = document.getElementById("task-form");

taskForm.addEventListener("submit", function (event) {
    event.preventDefault();
    addTask();
})

function addTask(){
    const taskTitleValue = document.getElementById("task-title").value;

    if(taskTitleValue.trim() == ""){
        alert("Judul tugas tidak boleh kosong");
        return;
    };

    const newTask = {
        id : tasks.length + 1,
        title : taskTitleValue.trim()
    };

    tasks.push(newTask);
    document.getElementById("task-title").value = ""
    localStorage.setItem("tasks",JSON.stringify(tasks))
}