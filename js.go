package gotoaser

// JavaScript for handling toasts
const toasterJS = `
function removeToast(id) {
  const toast = document.getElementById(id);
  if (toast) {
    toast.style.animation = 'toast-out 0.3s forwards';
    setTimeout(() => {
      toast.remove();
    }, 300);
  }
}

// Set up auto-dismiss for toasts
document.addEventListener('DOMContentLoaded', () => {
  const toasts = document.querySelectorAll('.toast');
  
  toasts.forEach(toast => {
    const duration = parseInt(toast.dataset.duration || 3, 10);
    const progress = toast.querySelector('.toast-progress');
    
    if (progress) {
      progress.style.animationDuration = duration + 's';
    }
    
    setTimeout(() => {
      removeToast(toast.id);
    }, duration * 1000);
  });
});

// Add toast-out animation
const style = document.createElement('style');
style.textContent = ` + "`" + `
@keyframes toast-out {
	from {
		transform: translateY(0);
		opacity: 1;
	}
	to {
		transform: translateY(-20px);
		opacity: 0;
	}
}
` + "`" + `;
document.head.appendChild(style);
`
