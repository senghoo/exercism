;;; hello-world.el --- Hello World Exercise (exercism)

;;; Commentary:

;;; Code:

(defun hello (&optional name)
  (concat "Hello, " (or name "World") "!"))

(provide 'hello-world)
;;; hello-world.el ends here
