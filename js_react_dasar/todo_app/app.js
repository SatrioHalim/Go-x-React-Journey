const tasks = localStorage.getItem("tasks")
  ? JSON.parse(localStorage.getItem("tasks"))
  : [];

const taskList = document.getElementById("tasks-list");
const taskForm = document.getElementById("task-form");

taskForm.addEventListener("submit", function (event) {
  event.preventDefault();
  addTask();
});

function checkTask(index) {
  const taskElement = document.getElementById(`task${index}`);
  const checkboxElement = document.getElementById(`checkTask${index}`);

  if (checkboxElement.checked) {
    taskElement.classList.add("checked");
    tasks[index].checked = true;
    localStorage.setItem("tasks", JSON.stringify(tasks));
  } else {
    taskElement.classList.remove("checked");
    tasks[index].checked = false;
    localStorage.setItem("tasks", JSON.stringify(tasks));
  }
}

function deleteTask(index) {
  tasks.splice(index, 1);
  localStorage.setItem("tasks", JSON.stringify(tasks));
  renderTask();
}

function addTask() {
  const taskTitleValue = document.getElementById("task-title").value;

  if (taskTitleValue.trim() == "") {
    alert("Judul tugas tidak boleh kosong");
    return;
  }

  const newTask = {
    id: tasks.length + 1,
    title: taskTitleValue.trim(),
    checked: false,
  };

  tasks.push(newTask);
  document.getElementById("task-title").value = "";
  localStorage.setItem("tasks", JSON.stringify(tasks));

  renderTask();
}

function renderTask() {
  taskList.innerHTML = "";
  if (tasks.length == 0) {
    const li = document.createElement("li");
    li.classList.add("tasks-list-item");
    li.innerHTML = `
        <div class="task">
            <div class="task-title">
                No activity yet
            </div>
        </div>
        `;
    taskList.appendChild(li);
    return;
  }

  tasks.forEach((task, index) => {
    const li = document.createElement("li");
    li.classList.add("tasks-list-item");
    if (task.checked == true) {
      li.classList.add("checked");
    }
    li.innerHTML = `
            <div class = "task">
                <div>
                    <input type="checkbox" onchange="checkTask(${index})" id="checkTask${index}" ${task.checked ? 'checked' : ''}/>
                </div>
                <div class = "task-title">
                    ${task.title}
                </div>
                <button class ="button" onclick="deleteTask(${index})">Delete</button>
            </div>
        `;
    li.id = `task${index}`;
    taskList.appendChild(li);
  });
}

function main() {
  renderTask();
}

main();
