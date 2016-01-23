;;; anagram.el --- Anagram (exercism)

;;; Commentary:

;;; Code:

(require 'cl)


(defun string-to-ordered-list (string)
  (sort (string-to-list string ) '<))

(defun anagrams-p (word1 word2)
  (equal (string-to-ordered-list word1)
         (string-to-ordered-list word2)))

(defun anagrams-for (word anagrams)
  (cl-remove-if-not (lambda (x) (anagrams-p (downcase word) (downcase x)))
                    anagrams))


(provide 'anagram)
;;; anagram.el ends here
