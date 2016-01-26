;;; robot-name.el --- Robot Name (exercism)

;;; Commentary:
;;
;; Build a robot with a name like AA000, that can be reset
;; to a new name. Instructions are in README.md
;;

;;; Code:

(defun random-from (s n)
  (let ((len (length s))
        (res))
    (apply 'string (dotimes (i n res)
                     (setq res (cons (elt s (random len)) res))))))

(defun new-robot-name ()
  (concat (random-from "ABCDEFGHIJKLMNOPQRSTUVWXYZ" 2)
          (random-from "0123456789" 3)))

(defun build-robot ()
  (cons (cons (new-robot-name) nil) nil))

(defun new-robot-name-for (robot)
  (let ((new-name (new-robot-name)))
    (while (member new-name (car robot))
      (setq new-name (new-robot-name)))
    new-name))

(defun reset-robot (robot)
    (setcar robot (cons (new-robot-name-for robot) (car robot))))

(defun robot-name (robot)
  (caar robot))

(provide 'robot-name)
;;; robot-name.el ends here
