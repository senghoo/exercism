;;; gigasecond.el --- Gigasecond exercise (exercism)

;;; Commentary:
;; Calculate the date one gigasecond (10^9 seconds) from the
;; given date.
;;
;; NB: Pay attention to  Emacs' handling of time zones and dst
;; in the encode-time and decode-time functions.

;;; Code:

(defun from (&rest vars)
  (butlast (decode-time (time-add (apply 'encode-time vars)
                                  (seconds-to-time 1e9)))
           3))

(provide 'gigasecond)
;;; gigasecond.el ends here
