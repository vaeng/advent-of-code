(ns advent-of-code.day-13
  (:require [clojure.string :as str])
  (:require [clojure.set :as set])
  (:require [clojure.edn :as edn]))

(defn part-1
  "Day 13 Part 1"
  [input]
  (let
   [[timestampTxt, linesTxt] (str/split-lines input)
    timestamp (edn/read-string timestampTxt)
    lines (for [x (str/split linesTxt  #",") :when (not= "x" x)] (edn/read-string x))
    [waittime lineID] (first (sort (for [line lines] [(- line (mod timestamp line)) line])))]
    (* waittime lineID)))



(defn part-2
  "Day 13 Part 2"
  [input]
  (letfn
   [(euclid-q-s-t-calc
     ;"calculates q, s and t for a and b so that 'q = s*a + t*b'"
      [a b]
      (if (= 0 b)
        [a 1 0]
        (let [[d s t] (euclid-q-s-t-calc b (mod a b))]
          [d, t, (- s (* t (quot a b)))])))]
    (let [raw_lines (for [x (str/split (second (str/split-lines input))  #",")] (edn/read-string x))
          M (apply * (for [linecode raw_lines :when (not= linecode 'x)] linecode))
          bus-data (sort-by first > (for
                                     [[offset line] (map-indexed (fn [idx itm] [idx itm]) raw_lines) :when (not= line 'x)]
                                      [line (mod (- line offset) line) (/ M line)]))
          kongruent (apply + (for
                              [[divisor rest Mi] bus-data]
                               (let [[q s t] (euclid-q-s-t-calc divisor Mi)]
                                 (* rest Mi t))))]
      (loop
       [result kongruent]
        (if (> result 0)
          result
          (recur (+ result M)))))))