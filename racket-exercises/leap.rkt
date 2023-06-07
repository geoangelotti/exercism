#lang racket

(provide leap-year?)

(provide leap-year?)
(define (leap-year? year)
  (define (divisible-by n) (zero? (remainder year n)))
  (cond
    [(divisible-by 400) #t]
    [(divisible-by 100) #f]
    [(divisible-by   4) #t]
    [else               #f]))
