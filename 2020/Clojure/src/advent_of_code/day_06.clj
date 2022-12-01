(ns advent-of-code.day-06
   (:require [clojure.string :as str])
  (:require [clojure.set :as set])
  )

(defn part-1
  "Day 06 Part 1"
  [input]
  (apply + (for
            [group (str/split (str/replace input #"\r\n" " ") #"  ")]
             (count (set (re-seq #"[a-z]" group))))))

(defn part-2
  "Day 06 Part 2"
  [input]
  (let [groups (str/split input #"\r\n\r\n")]
    (apply + (for
            [group groups]
             (let [choices (str/split group #"\r\n")]
               (count (apply set/intersection (for
                                               [choice choices]
                                                (set choice))))))))
)