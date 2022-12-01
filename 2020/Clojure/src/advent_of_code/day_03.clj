(ns advent-of-code.day-03
  (:require [clojure.string :as str])
  )

(defn hit-position [line right width down] (rem (/ (* line right) down) width))

(defn hit-checker 
  [input right down]
    (let [treemap (str/split input #"\r\n")
          width (count (first treemap))]
      (count
       (filter
        #(= % \#)
        (for [[linecount treeline] (map-indexed vector treemap)
              :when (= 0 (rem linecount down))]
          (nth treeline (hit-position linecount right width down)))))))

(defn part-1
  "Day 03 Part 1"
  [input]
  (hit-checker input 3 1)
  )

(defn part-2
  "Day 03 Part 2"
  [input]
  (*
   (hit-checker input 1 1)
   (hit-checker input 3 1)
   (hit-checker input 5 1)
   (hit-checker input 7 1)
   (hit-checker input 1 2)
   ))


