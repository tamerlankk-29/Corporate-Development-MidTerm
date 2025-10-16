import React from "react";
import TaskList from "../components/TaskList";
import { Link } from "react-router-dom";

export default function Home() {
  return (
    <div>
      <Link to="/create" className="btn btn-primary mb-3">+ Add Task</Link>
      <TaskList />
    </div>
  );
}
