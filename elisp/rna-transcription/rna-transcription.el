;;; rna-transcription.el -- RNA Transcription (exercism)

;;; Commentary:



;;; Code:

(setq dna-to-rna '((?C . ?G)
                   (?G . ?C)
                   (?A . ?U)
                   (?T . ?A)))

(defun to-rna-nucleotides (nucleotide)
  (cdr (assoc nucleotide dna-to-rna)))


(defun to-rna (dna)
  (apply 'string (mapcar 'to-rna-nucleotides
                  (string-to-list dna))))

(provide 'rna-transcription)
;;; rna-transcription.el ends here
