;;; allergies.el --- Allergies Exercise (exercism)

;;; Commentary:

;;; Code:

(setq allergens '(("eggs" . 1)
                  ("peanuts" . 2)
                  ("shellfish" . 4)
                  ("strawberries" . 8)
                  ("tomatoes" . 16)
                  ("chocolate". 32)
                  ("pollen" . 64)
                  ("cats" . 128)))

(defun allergen-list (score)
  (mapcar 'car
          (remove-if (lambda (c)
                       (message (car c) (cdr c))
                       (equal (logand (cdr c)
                                      score)
                              0))
                     allergens)))

(defun allergic-to-p (score allergen)
  (member allergen (allergen-list score)))


(provide 'allergies)
;;; allergies.el ends here
