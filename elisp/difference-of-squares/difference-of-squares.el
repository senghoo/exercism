;;; difference-of-squares.el --- Difference of Squares (exercism)

;;; Commentary:

;;; Code:

(defun square-of-sums (num)
  (let ((sum (/ (* (+ 1 num)
                   num)
                2)))
    (* sum sum)))

(defun sum-of-squares (num)
  (/ (* num
        (+ num 1)
        (+ (* num 2) 1))
     6))

(defun difference (num)
  (- (square-of-sums num)
     (sum-of-squares num)))

(provide 'difference-of-squares)
;;; difference-of-squares.el ends here
