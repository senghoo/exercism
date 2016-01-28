;;; raindrops.el --- Raindrops (exercism)

;;; Commentary:

;;; Code:

(require 'cl)

(setq raindrops '((3 . "Pling")
                  (5 . "Plang")
                  (7 . "Plong")))

(defun convert (n)
  "Convert integer N to its raindrops string."
  (let ((r (apply 'concat (mapcar (lambda (c)
                                    (if (divisble-by n (car c)) (cdr c) ""))
                                  raindrops))))
    (if (equal r "") (number-to-string n) r)))

(defun divisble-by (num d)
  (equal (% num d) 0))



(provide 'raindrops)
;;; raindrops.el ends here
