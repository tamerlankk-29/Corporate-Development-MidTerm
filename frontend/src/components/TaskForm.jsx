import React, { useState, useEffect } from "react";
import { useNavigate, useParams } from "react-router-dom";
import { taskService } from "../api/taskService";

export default function TaskForm({ isEdit = false }) {
  const navigate = useNavigate();
  const { id } = useParams(); 
  const [task, setTask] = useState({
    title: "",
    description: "",
    status: "in progress",
  });
  const [error, setError] = useState("");

  useEffect(() => {
    if (isEdit && id) {
      async function fetchTask() {
        try {
          const data = await taskService.getById(id);
          setTask({
            title: data.title || "",
            description: data.description || "",
            status: data.status || "in progress",
          });
        } catch (err) {
          console.error("Error loading task:", err);
          setError("Failed to load task");
        }
      }
      fetchTask();
    }
  }, [id, isEdit]);

  const handleChange = (e) => {
    const { name, value } = e.target;
    setTask((prev) => ({ ...prev, [name]: value }));
  };

  const validate = () => {
    if (!task.title.trim()) return "The 'Title' field is required.";
    if (!task.description.trim()) return "The 'Description' field is required.";
    return "";
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    const validationError = validate();
    if (validationError) {
      setError(validationError);
      return;
    }

    try {
      if (isEdit) {
        await taskService.update(id, task);
      } else {
        await taskService.create(task);
      }

      navigate("/"); 
    } catch (err) {
      console.error("Error while saving:", err);
      setError("Failed to save task. Check your connection to the server.");
    }
  };

  return (
    <div className="card shadow-sm btn-form-primary text-black " >
      <div className="card-body">
        <h4 className="mb-3">{isEdit ? "Edit task" : "Create a task"}</h4>

        {error && (
          <div className="alert alert-danger py-2" role="alert">
            {error}
          </div>
        )}

        <form onSubmit={handleSubmit}>
          <div className="mb-3">
            <label className="form-label">Task name</label>
            <input
              type="text"
              name="title"
              value={task.title}
              onChange={handleChange}
              className="form-control"
              placeholder="Enter task name..."
            />
          </div>

          <div className="mb-3">
            <label className="form-label">Description</label>
            <textarea
              name="description"
              value={task.description}
              onChange={handleChange}
              className="form-control"
              placeholder="Describe the task..."
            />
          </div>

          <div className="mb-3">
            <label className="form-label">Status</label>
            <select
              name="status"
              value={task.status}
              onChange={handleChange}
              className="form-select"
            >
              <option value="in progress">In progress</option>
              <option value="done">Done</option>
            </select>
          </div>

          <button type="submit" className="btn btn-primary " >
            {isEdit ? "Save changes" : "Create a task"}
          </button>
          <button
            type="button"
            className="btn btn-secondary ms-2"
            onClick={() => navigate("/")}
          >
            Cancel
          </button>
        </form>
      </div>
    </div>
  );
}
