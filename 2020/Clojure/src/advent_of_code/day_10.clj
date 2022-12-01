(ns advent-of-code.day-10
    (:require [clojure.string :as str])
   (:require [clojure.set :as set])
(:require [clojure.edn :as edn])
)

(defn part-1
  "Day 10 Part 1"
  [input]
  (let [input-vector (for [num (str/split-lines input)] (edn/read-string num))]
    (loop [count1 0
           count3 0
           col (sort (conj input-vector 0))]
      (if (= 1 (count col))
        (* count1 (inc count3))
        (let [left (first col) right (second col)]
          (if (= 1 (- right left))
            (recur (inc count1) count3 (drop 1 col))
            (if (= 3 (- right left))
              (recur count1 (inc count3) (drop 1 col))
              (recur count1 count3 (drop 1 col)))))))))

(defn part-2
  "Day 10 Part 2"
  [input]
  input)


(defn read-input
  [day]
  (slurp (clojure.java.io/resource day)))

(def input-example (read-input "day-10-example.txt"))
(def input (read-input "day-10.txt"))

(part-1 input)



(def numb-seq (sort (conj (for [num (str/split-lines input-example)] (edn/read-string num)) 0)))
numb-seq

(defn adapter-needed? 
  [before next]
  (> (- next before ) 3)
  )

(adapter-needed? 1 5)

(def short-test (sort [16
10
15
5
1
11
7
19
6
12
4
                 0
                       22]))

short-test

0 1 2 3 4 ; neeeded
0 2 3 4   ; not needed
0 3 4    ; not needed second neither
0 2 4    ; not needed third neither

0 1 2 3 _4_ 5 6 7 8

(defn drop-variations
  "returns the number of variations possible with the drop of the second member"
  [col]
  (let [x0 (first col)
        x1 (second col)
        x2 (nth col 2)
        x3 (nth col 3)
        x3 (nth col 3)
        ]
   (count 
   (filter 
    true? [true
           (<= (- x2 x0) 3)
           (<= (- x3 x0) 3)
           (<= (- x3 x0) 3)])))
  )

(loop
 [col numb-seq
  variations 1]
  (if (= 2 (count col))
    (if (adapter-needed? (first col) (+ 3 (second col)))
      variations
      (* 2 variations))
    (if (adapter-needed? (first col) (nth col 2))
      (recur (drop 1 col) variations)
      (recur (drop 1 col) (+ () variations)))))

(defn drop-nth [n coll]
  (keep-indexed #(if (not= %1 n) %2) coll))

(nth [1 2 3 4 5] 3)


(defn check-valid? 
  "checks if sequence is still valid if nth element were removed"
  [col n]
  (<= (- (nth col (inc n)) (nth col (dec n))) 3)
  )

(check-valid? [1 2 4] 1)

(defn total-posibilites
  "checks all possible scenarios"
  [col results]
  (let
   [lentgh (count col)
    subresult (conj #{} col)
    completeresult (set/union results subresult)]
    (if (= 2 lentgh)
      (if (<= (- (last col) (first col)) 3)
        completeresult
        results)
      (apply set/union (for 
                        [n (range 1 (dec lentgh))
                         droppedcol (drop-nth n col)
                         nextset (set/union completeresult droppedcol)
                         :when (and (not (contains? completeresult droppedcol)) (check-valid? col n))]
                        (total-posibilites droppedcol nextset))))))

(count (total-posibilites short-test #{}))

(total-posibilites '(0 1 3 4) #{})

(set [0 1])
(def avector [0 1])
(conj #{} avector)

(set/union (conj #{} [1 2 3]) #{[1 3]})

(not true)

(filter true? [false true])