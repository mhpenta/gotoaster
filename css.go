package gotoaser

// Embedded CSS for toasts
const toasterCSS = `
.toaster {
  position: fixed;
  display: flex;
  flex-direction: column;
  gap: 0.5rem;
  z-index: 9999;
  padding: 1rem;
}

.toaster-top-right {
  top: 0;
  right: 0;
}

.toaster-top-left {
  top: 0;
  left: 0;
}

.toaster-top-center {
  top: 0;
  left: 50%;
  transform: translateX(-50%);
}

.toaster-bottom-right {
  bottom: 0;
  right: 0;
}

.toaster-bottom-left {
  bottom: 0;
  left: 0;
}

.toaster-bottom-center {
  bottom: 0;
  left: 50%;
  transform: translateX(-50%);
}

.toast {
  position: relative;
  display: flex;
  flex-direction: column;
  min-width: 300px;
  max-width: 400px;
  background-color: white;
  border-radius: 0.375rem;
  box-shadow: 0 4px 6px -1px rgba(0, 0, 0, 0.1), 0 2px 4px -1px rgba(0, 0, 0, 0.06);
  overflow: hidden;
  animation: toast-in 0.3s ease forwards;
}

.toast-close {
  position: absolute;
  top: 0.5rem;
  right: 0.5rem;
  width: 1.5rem;
  height: 1.5rem;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 1.25rem;
  line-height: 1;
  border: none;
  background: transparent;
  color: rgba(0, 0, 0, 0.5);
  cursor: pointer;
  transition: color 0.2s;
}

.toast-close:hover {
  color: rgba(0, 0, 0, 0.8);
}

.toast-content {
  display: flex;
  padding: 1rem;
  gap: 0.75rem;
  align-items: center;
}

.toast-icon {
  flex-shrink: 0;
  color: currentColor;
}

.toast-message {
  flex-grow: 1;
}

.toast-progress {
  height: 4px;
  background-color: rgba(0, 0, 0, 0.1);
  width: 100%;
  position: absolute;
  bottom: 0;
  left: 0;
  transform-origin: left;
  animation: toast-progress linear forwards;
}

.toast-success {
  border-left: 4px solid #10b981;
  color: #065f46;
}

.toast-error {
  border-left: 4px solid #ef4444;
  color: #991b1b;
}

.toast-warning {
  border-left: 4px solid #f59e0b;
  color: #92400e;
}

.toast-info {
  border-left: 4px solid #3b82f6;
  color: #1e40af;
}

.toast-default {
  border-left: 4px solid #6b7280;
  color: #1f2937;
}

@keyframes toast-in {
  from {
    transform: translateY(100%);
    opacity: 0;
  }
  to {
    transform: translateY(0);
    opacity: 1;
  }
}

@keyframes toast-progress {
  0% {
    transform: scaleX(1);
  }
  100% {
    transform: scaleX(0);
  }
}
`
