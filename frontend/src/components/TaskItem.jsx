import React from "react";
import { Link } from "react-router-dom";

export default function TaskItem({ task, onDelete }) {
  return (
    <div className="card mb-3 shadow-sm border-0 rounded-4 btn-form-primary" >
      <div className="card-body">
        <div className="d-flex justify-content-between align-items-start">
          <div>
            <h5 className="card-title mb-1 text-dark">{task.title}</h5>
            <p className="card-text text-muted mb-2">{task.description}</p>

            <span
              className={`badge ${
                task.status === "done"
                  ? "bg-success"
                  : task.status === "in progress"
                  ? "bg-warning text-dark"
                  : "bg-secondary"
              }`}
            >
              {task.status}
            </span>
          </div>

          <div className="d-flex flex-column align-items-end">
            <Link
              to={`/edit/${task.id}`}
              className="btn btn-sm btn-outline-primary mb-2"
              
            >
              Edit
            </Link>

            <button
              className="btn btn-sm btn-outline-danger"
              onClick={() => onDelete(task.id)}
            >
              Delete
            </button>
          </div>
        </div>
      </div>
    </div>
  );
}
