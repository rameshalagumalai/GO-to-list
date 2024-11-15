import 'bootstrap/dist/css/bootstrap.css';
import './App.css';
import Task from './Task';
import { useState } from 'react';

function App() {

  const [tasks, setTasks] = useState([]);
  const [newTaskName, setNewTaskName] = useState("");

  function handleNewTask(e) {
    e.preventDefault();
    setTasks([...tasks, newTaskName]);
    setNewTaskName("");
  }

  function handleDelete(i) {
    setTasks([...tasks.slice(0, i), ...tasks.slice(i + 1)]);
  }

  function handleEdit(e, i, editedTaskName, setIsEdit) {
    e.preventDefault();
    setTasks([...tasks.slice(0, i), editedTaskName, ...tasks.slice(i + 1)]);
    setIsEdit(false);
  }

  return (
    <div className="App vh-100 d-flex flex-column align-items-center justify-content-center">
      <div style={{ width: "26rem" }} className="h-100 d-flex flex-column shadow-lg my-3 rounded-5">
        <h1 className="text-info f-700 mx-auto mt-2">GO - to list</h1>
        <div className="h-100 d-flex flex-column overflow-y-auto border-top border-bottom">
          {
            tasks.map((task, i) => {
              return (
                <Task task={task} key={i} i={i} handleDelete={handleDelete} handleEdit={handleEdit} />
              )
            })
          }
        </div>
        <form onSubmit={e => handleNewTask(e)} className="p-3">
          <input value={newTaskName} onChange={e => setNewTaskName(e.target.value)} type="text" className="form-control" placeholder="Enter a new item" />
        </form>
      </div>
    </div>
  );
}

export default App;
