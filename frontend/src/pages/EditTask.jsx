import React from "react";
import TaskForm from "../components/TaskForm";

export default function EditTask() {
  return (
    <div>
      <h2>Edit Task</h2>
      <TaskForm isEdit />
    </div>
  );
}
