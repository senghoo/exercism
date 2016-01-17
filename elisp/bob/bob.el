;;; bob.el --- Bob exercise (exercism)

;;; Commentary:

;;; Code:

(defun response-for (str)
  (cond ((and (not (s-lowercase? str))
              (s-uppercase? str)) "Whoa, chill out!")
        ((s-ends-with? "?" str) "Sure.")
        ((s-blank? (s-trim str)) "Fine. Be that way!")
        (t "Whatever.")))


(provide 'bob)
;;; bob.el ends here
