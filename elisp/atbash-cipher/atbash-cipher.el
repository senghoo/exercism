;;; atbash-cipher.el --- Atbash-Cipher (exercism)

;;; Commentary:




;;; Code:

(defun encode-char (c)
  (if (<= ?a c ?z)
      (+ ?a (- ?z c))
    c))


(downcase "ABC")

(defun encode-prepare (s)
  (downcase (replace-regexp-in-string "[\s,.]+" "" s)))

(defun encode-add-space-per-five (l &optional counter res)
  (setq counter (or counter 0))
  (cond ((not res) (encode-add-space-per-five (cdr l) 1
                                              (cons (car l) nil) ))
        ((not l) res)
        ((>= counter 5)
         (encode-add-space-per-five l 0 (append res (cons ?  nil))))
        (t (encode-add-space-per-five (cdr l)
                                      (+ counter 1)
                                      (append res (cons (car l) nil))))))

(defun encode (plaintext)
  "Encode PLAINTEXT to atbash-cipher encoding."
  (apply 'string
         (encode-add-space-per-five (mapcar 'encode-char
                                            (string-to-list (encode-prepare plaintext))))))


(provide 'atbash-cipher)
;;; atbash-cipher.el ends here
