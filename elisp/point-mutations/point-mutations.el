;;; point-mutations.el --- Point Mutations (exercism)

;;; Commentary:

;;; Code:

(require 'cl)

(defun hamming-distance (a b)
  (unless (= (length a)
             (length b))
    (error "args should have same length"))
  (apply '+ (mapcar* (lambda (x y) (if (= x y) 0 1)) a b)))

(provide 'point-mutations)
;;; point-mutations.el ends here
