// components/AlertPopup.tsx
import { useEffect } from "react";

interface AlertPopupProps {
  message: string;
  type?: "success" | "error" | "warning" | "info";
  onClose: () => void;
  duration?: number;
}

const alertStyles = {
  success: "bg-green-100 text-green-700 border-green-400",
  error: "bg-red-100 text-red-700 border-red-400",
  warning: "bg-yellow-100 text-yellow-700 border-yellow-400",
  info: "bg-blue-100 text-blue-700 border-blue-400",
};

export default function AlertPopup({
  message,
  type = "info",
  onClose,
  duration = 3000,
}: AlertPopupProps) {
  useEffect(() => {
    const timer = setTimeout(() => {
      onClose();
    }, duration);
    return () => clearTimeout(timer);
  }, [duration, onClose]);

  return (
    <div
      className={`fixed top-5 right-5 z-50 px-4 py-3 rounded border shadow-lg transition-opacity ${
        alertStyles[type]
      }`}
    >
      <div className="flex justify-between items-center">
        <span>{message}</span>
        <button onClick={onClose} className="ml-4 font-bold">
          &times;
        </button>
      </div>
    </div>
  );
}
