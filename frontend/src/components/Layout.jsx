import React from "react";
import { Link, Outlet } from "react-router-dom";

export default function Layout() {
  return (
    <div className="container mt-4">
      <nav className="navbar navbar-expand-lg navbar-light bg-light rounded mb-4 p-3 shadow-sm">
        <Link className="navbar-brand fw-bold" to="/">
          🗂️ Task Manager
        </Link>
        <div>
          <Link to="/" className="btn btn-outline-primary me-2">
            Home
          </Link>
          <Link to="/create" className="btn btn-primary">
            + Create Task
          </Link>
        </div>
      </nav>

      <Outlet />
    </div>
  );
}
