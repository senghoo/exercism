;;; leap.el --- Leap exercise (exercism)

;;; Commentary:

;;; Code:



(defun leap-year-p (year)
  (defun year-divisible-p (a)
    (= (% year a) 0))
  (and (year-divisible-p 4) (or (year-divisible-p 400) (not (year-divisible-p 100)))))


(provide 'leap)
;;; leap.el ends here
