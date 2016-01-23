;;; binary.el --- Binary exercise (exercism)

;;; Commentary:

;;; Code:

(defun to-decimal (string)
  (to-decimal--inner (string-to-list string) 0))

(defun to-decimal--inner (lst val)
  (if (not lst)
    val
    (to-decimal--inner
     (cdr lst)
     (+ (lsh val 1)
       (if (equal (car lst) ?1) 1 0)))))




(provide 'binary)
;;; binary.el ends here
