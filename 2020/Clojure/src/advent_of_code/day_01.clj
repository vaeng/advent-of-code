(ns advent-of-code.day-01)

(require '[clojure.edn :as edn])
(require '[clojure.string :as str])

(defn read-numbers
  [raw-input]
  (mapv edn/read-string (str/split raw-input #"\r\n")))

(defn part-1
  "Day 01 Part 1"
  [input]
  (let [numbers (read-numbers input)]
    (first (for [in1 numbers in2 numbers
          :let [x (+ in1 in2)]
          :when (= x 2020)]
      (* in1 in2)))))


(defn part-2
  "Day 01 Part 2"
  [input]
  (let [numbers (read-numbers input)]
    (first (for [in1 numbers in2 numbers in3 numbers
        :let [x (+ in1 in2 in3)] 
        :when (= x 2020)] 
      (* in1 in2 in3))))
  )
