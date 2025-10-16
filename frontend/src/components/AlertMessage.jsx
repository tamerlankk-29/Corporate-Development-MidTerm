import React from "react";

export default function AlertMessage({ type = "info", message, onClose }) {
  if (!message) return null;

  const alertClass =
    type === "success"
      ? "alert alert-success alert-dismissible fade show"
      : type === "error"
      ? "alert alert-danger alert-dismissible fade show"
      : "alert alert-warning alert-dismissible fade show";

  return (
    <div className={alertClass} role="alert">
      {message}
      <button
        type="button"
        className="btn-close"
        onClick={onClose}
        aria-label="Close"
      ></button>
    </div>
  );
}
