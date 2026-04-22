
const MyTaskForm = (props) => {
    const {onSubmit,setTaskTitle,taskTitle} = props;
    return (
        <form className="form" id="task-form" onSubmit={(event) => onSubmit(event)}>
            <div className="form-group">
                <label for="task-title" className="form-label">Mau ngapain hari ini?</label>
                <input type="text" className="form-input form-input--full" id="task-title" name="task-title" placeholder="masukkan kegiatan" onChange={(event) => setTaskTitle(event.target.value)}
                value={taskTitle}/>
            </div>
            <div className="form-group form-group--button-right">
                <button className="button" type="submit">Save</button>
            </div>
        </form>
    );
};

const MyTaskList = (props) => {
    const {tasks,checkTask,deleteTask} = props
    return (
        <div className="tasks">
            <h2 className="tasks-title">My To-Do</h2>
            <ul className="tasks-list" id="tasks-list">
                {
                    tasks.length > 0 ? (
                        tasks.map((task,index) => (
                            <li key={task.id} className={`tasks-list-item ${task.checked ? "checked" : ""}`}>
                                <div className="task">
                                    <div>
                                        <input type="checkbox" checked={task.checked} onChange={(event) => checkTask(event,index)} />
                                    </div>
                                    <div className="task-title">{task.title}</div>
                                    <button className="button" onClick={() => deleteTask(index)}>Delete</button>
                                </div>
                            </li>
                        ))
                    ) : (
                        <div classNames="task">
                            <div classNames="task-title">
                                No activity yet
                            </div>
                        </div>
                    )
                }
            </ul>
            
        </div>
    )
};

const MyTaskContainer = (props) => {
    return (
        <div className="container">{props.children}</div>
    )
};

const App = () => {

    // state = cara menyimpan value di dalem komponen react, jadi kalo ada perubahan logicnya bisa terus berjalan
    const [taskTitle,setTaskTitle] = React.useState("")
    const [tasks,setTasks] = React.useState([]);

    React.useEffect(() => {
        const storedTasks = localStorage.getItem("tasks")
        if (storedTasks){
            setTasks(JSON.parse(storedTasks))
        }
    },[])

    const addTask = (event) => {
        event.preventDefault();
        if(taskTitle.trim() == ""){
            alert("Judul tugas tidak boleh kosong");
            return
        }

        const newTask = {
            id: tasks.length + 1,
            title: taskTitle
        }

        setTasks([...tasks, newTask]); // ... ini teknik spread
        setTaskTitle("");
        localStorage.setItem("tasks", JSON.stringify([...tasks,newTask]))
    }

    const deleteTask = (taskIndex) => {
        const updatedTasks = tasks.filter((_,index) => index != taskIndex)
        setTasks(updatedTasks)
        localStorage.setItem("tasks",JSON.stringify(updatedTasks));
    }

    const checkTask = (event, taskIndex) => {
        const isChecked = event.target.checked;
        const updatedTasks = tasks.map((task,index) => {
            if(taskIndex == index){
                return {
                    ...task,
                    checked:isChecked
                }
            }
            return task;
        });
        setTasks(updatedTasks);
        localStorage.setItem("tasks", JSON.stringify(updatedTasks))   
    }

    return (
        <MyTaskContainer>
            <MyTaskForm onSubmit={addTask} taskTitle={taskTitle} setTaskTitle={setTaskTitle}/>
            <MyTaskList tasks={tasks} checkTask={checkTask} deleteTask={deleteTask}/>
        </MyTaskContainer>
    );
};

const root = ReactDOM.createRoot(document.getElementById("app"));
root.render(<App/>);