(ns advent-of-code.day-11
  (:require [clojure.string :as str]))

(defn count-adjacent-occupied
  [input row column]
  (letfn [(look-up [seat-row seat-column]
            (try
              (nth (nth input seat-row) seat-column)
              (catch Exception _ nil)))]
    (count (filter
            #(= \# %)
            [(look-up  (dec row) (dec column))
             (look-up  (dec row) column)
             (look-up  (dec row) (inc column))
             (look-up  row (dec column))
             (look-up  row (inc column))
             (look-up  (inc row) (dec column))
             (look-up  (inc row) column)
             (look-up  (inc row) (inc column))]))))

(defn count-visible-occupied
  [input row column]
  (letfn [(incvalue [value]
                    (cond
                      (= value 0) 0
                      (< value 0) (dec value)
                      (> value 0) (inc value)))
          
          (look-up [horiz vertical]
           (let [nextseat (try
                            (nth (nth input (+ row horiz)) (+ column vertical))
                            (catch Exception _ nil))]
             (if (= nextseat \.)
               (look-up (incvalue horiz) (incvalue vertical))
               nextseat
               )
             ))]
    (count (filter
            #(= \# %)
            [(look-up  -1 -1)
             (look-up  -1 0)
             (look-up  -1 1)
             (look-up  0 -1)
             (look-up  0 1)
             (look-up  1 -1)
             (look-up  1 0)
             (look-up  1 1)]))))

(defn cycleseats [input row column]
  (let [status (nth (nth input row) column)
        occupied (count-adjacent-occupied input row column)]
    (cond
      (and (= status \#) (>= occupied 4)) \L
      (and (= status \L) (= occupied 0)) \#
      :else status)))

(defn cycleseats2 [input row column]
  (let [status (nth (nth input row) column)
        occupied (count-visible-occupied input row column)]
    (cond
      (and (= status \#) (>= occupied 5)) \L
      (and (= status \L) (= occupied 0)) \#
      :else status)))

(defn cycle-map
  [input]
  (let [rows (count input)
        columns (count (first input))]
    (for [row (range rows)]
      (for [column (range columns)]
        (cycleseats input row column)))))


(defn cycle-map2
  [input]
  (let [rows (count input)
        columns (count (first input))]
    (for [row (range rows)]
      (for [column (range columns)]
        (cycleseats2 input row column)))))

(defn part-1
  "Day 11 Part 1"
  [input]
  (let [seatmap (str/split-lines input)
        finalmap (loop
                  [loopinput seatmap]
                   (let [nextmap (cycle-map loopinput)]
                     (if (= nextmap loopinput)
                       nextmap
                       (recur nextmap))))]
    (apply + (for [row finalmap]
               (count
                (filter #(= \# %) row))))))



(defn part-2
  "Day 11 Part 2"
  [input]
  (let [seatmap (str/split-lines input)
        finalmap (loop
                  [loopinput seatmap]
                   (let [nextmap (cycle-map2 loopinput)]
                     (if (= nextmap loopinput)
                       nextmap
                       (recur nextmap))))]
    (apply + (for [row finalmap]
               (count
                (filter #(= \# %) row))))))




