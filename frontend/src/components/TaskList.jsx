import React, { useEffect, useState } from "react";
import { taskService } from "../api/taskService";
import TaskItem from "./TaskItem";
import AlertMessage from "./AlertMessage";

export default function TaskList() {
  const [tasks, setTasks] = useState([]);
  const [loading, setLoading] = useState(true);
  const [alert, setAlert] = useState({ type: "", message: "" });

  useEffect(() => {
    async function fetchTasks() {
      try {
        setLoading(true);
        // const response = await taskService.getAll();
        // setTasks(response.data);

        //эти удалить
        const fakeTasks = [
          { id: 1, title: "Learn React", description: "Finish hooks", status: "in progress" },
          { id: 2, title: "Write report", description: "For project", status: "done" },
        ];
        setTasks(fakeTasks);
      } catch {
        setAlert({ type: "error", message: "Error loading tasks" });
      } finally {
        setLoading(false);
      }
    }
    fetchTasks();
  }, []);

  const handleDelete = async (id) => {
    try {
      setTasks(tasks.filter((t) => t.id !== id));
      setAlert({ type: "success", message: "The task was successfully deleted." });
    } catch {
      setAlert({ type: "error", message: "Error deleting task" });
    }
  };

  return (
    <div className="card shadow-sm border-0 mt-3">
      <div className="card-body">
        <h4 className="mb-3 text-dark">📋 List of tasks</h4>

        {alert.message && (
          <AlertMessage
            type={alert.type}
            message={alert.message}
            onClose={() => setAlert({ type: "", message: "" })}
          />
        )}

        {loading ? (
          <div className="text-center mt-4">
            <div className="spinner-border text-warning" role="status"></div>
            <p>Loading tasks...</p>
          </div>
        ) : tasks.length === 0 ? (
          <p>No tasks</p>
        ) : (
          <div className="list-group">
            {tasks.map((task) => (
              <TaskItem key={task.id} task={task} onDelete={handleDelete} />
            ))}
          </div>
        )}
      </div>
    </div>
  );
}
