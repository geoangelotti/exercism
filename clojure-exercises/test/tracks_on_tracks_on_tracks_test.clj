(ns tracks_on_tracks_on_tracks_test
  (:require [clojure.test :refer [deftest is]] tracks_on_tracks_on_tracks))

(deftest list-empty-test
  (is (= (quote ()) (tracks-on-tracks-on-tracks/new-list))))

(deftest list-add-test
  (is (= (quote ("JavaScript" "Java" "Lisp" "Clojure"))
    (->
      (tracks-on-tracks-on-tracks/new-list)
      (tracks-on-tracks-on-tracks/add-language "Clojure")
      (tracks-on-tracks-on-tracks/add-language "Lisp")
      (tracks-on-tracks-on-tracks/add-language "Java")
      (tracks-on-tracks-on-tracks/add-language "JavaScript")))))

(deftest first-test
  (is (= "Lisp" (tracks-on-tracks-on-tracks/first-language (quote ("Lisp" "Clojure"))))))

(deftest list-remove-test
  (is (= (quote ("Clojure")) (tracks-on-tracks-on-tracks/remove-language (quote ("Lisp" "Clojure"))))))
  
(deftest list-count-test
  (is (= 3 (tracks-on-tracks-on-tracks/count-languages (quote ("JavaScript" "Java" "Clojure"))))))

(deftest learning-list-test
  (is (= 3 (tracks-on-tracks-on-tracks/learning-list))))
