(ns advent-of-code.day-02
  (:require [clojure.string :as str])
  (:require [clojure.edn :as edn])
 )

(defn password-vars [text]
  (let [[_, min, max, letter, password] (re-matches #"(\d+)-(\d+) (\w): (\w*)" text)]
    [(edn/read-string min) (edn/read-string max) letter password])) 

(defn count_letters [letter text] (count (filter #(= % true) (for [x text] (= letter x)))))

(defn check-correct
  [min max letter password]
  (let [amount (count_letters (first (char-array letter)) password)]
    (= true (>= amount min) (<= amount max))))

(defn part-1
  "Day 02 Part 1"
  [input]
  (let 
   [database (mapv password-vars (str/split input #"\r\n"))]
    (count (for [[min max letter password] database
                 :let [status (check-correct min max letter password)]
                 :when (= true status)] status))
   )
  )

(defn convert-string-char
  [letter]
  (first (char-array letter)))

(defn check-correct2
  [min max letter password]
  (let [newletter (convert-string-char letter)]
    (not=
     (= newletter (nth password (- min 1)))
     (= newletter (nth password (- max 1))))))

(defn part-2
  "Day 02 Part 2"
  [input]
  (let 
   [database (mapv password-vars (str/split input #"\r\n"))]
    (count (for [[min max letter password] database
                 :let [status (check-correct2 min max letter password)]
                 :when (= true status)] status))
   )
  )