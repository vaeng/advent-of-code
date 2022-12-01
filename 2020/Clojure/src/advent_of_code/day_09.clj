(ns advent-of-code.day-09
  (:require [clojure.string :as str])
  (:require [clojure.edn :as edn])
 )

(defn find-invalid
  "returns the first invalid number in an XMAS sequence"
  [col preamble-size]
  (let [preamble (take preamble-size col)
        candidate (last (take (inc preamble-size) col))]
    (if (contains?
         (set (for [x preamble y preamble :when (> x y)] (+ x y)))
         candidate)
      (find-invalid (drop 1 col) preamble-size)
      candidate)))


(defn part-1
  "Day 09 Part 1"
  [input]
  (find-invalid (for [num (str/split-lines input)] (edn/read-string num)) 25))


(defn summed-sequence
  "returns the subsequence that adds to the sum"
  [col start end target]
  (let [subcol (drop start (take end col))
        current-sum (apply + subcol)]
    (if (= current-sum target)
      subcol
      (if (< current-sum target)
        (summed-sequence col start (inc end) target)
        (summed-sequence col (inc start) end target)))))

(defn part-2
  "Day 09 Part 2"
  [input]
  (let [number-vec (for [num (str/split-lines input)] (edn/read-string num))
        faulty-number (part-1 input)
        solution-seq (summed-sequence number-vec 0 1 faulty-number)
        ]
    (+ (apply min solution-seq) (apply max solution-seq))
    ))
