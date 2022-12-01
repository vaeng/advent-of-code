(ns advent-of-code.day-12
  (:require [clojure.string :as str])
  (:require [clojure.edn :as edn]))

(defn update-position
  "returns the coordinate changes from an instruction"
  [instruction northsouth eastwest facing]
  (let [code (first (re-find #"[NSEWLRF]" instruction))
        number (edn/read-string (re-find #"\d+" instruction))]
    (cond
      (= code \N) [(+ northsouth number) eastwest facing]
      (= code \S) [(- northsouth number) eastwest facing]
      (= code \E) [northsouth (+ eastwest number) facing]
      (= code \W) [northsouth (- eastwest number) facing]
      (= code \L) [northsouth eastwest (mod (+ 360 (- facing number)) 360)]
      (= code \R) [northsouth eastwest (mod (+ 360 (+ facing number)) 360)]
      (= code \F) (cond
                    (= facing 270) [(+ northsouth number) eastwest facing]
                    (= facing 90) [(- northsouth number) eastwest facing]
                    (= facing 0) [northsouth (+ eastwest number) facing]
                    (= facing 180) [northsouth (- eastwest number) facing]))))

(defn part-1
  "Day 12 Part 1"
  [input]
  (loop
   [instructions input
    northsouth 0
    eastwest 0
    facing 0]
    (let [instruction (first instructions)
          otherinstructions (rest instructions)
          [newns, newew, newface] (update-position instruction northsouth eastwest facing)]

      (if (empty? otherinstructions)
        (+ (Math/abs newns) (Math/abs newew))
        (recur otherinstructions newns newew newface)))))

(defn update-position2
  "returns the coordinate changes from an instruction"
  [instruction northsouth eastwest shipNS shipEW]
  (let [code (first (re-find #"[NSEWLRF]" instruction))
        number (edn/read-string (re-find #"\d+" instruction))]
    (cond
      (= code \N) [(+ northsouth number) eastwest shipNS shipEW]
      (= code \S) [(- northsouth number) eastwest shipNS shipEW]
      (= code \E) [northsouth (+ eastwest number) shipNS shipEW]
      (= code \W) [northsouth (- eastwest number) shipNS shipEW]
      (= code \F) [northsouth eastwest (+ shipNS (* number northsouth)) (+ shipEW (* number eastwest))]
      (= code \L) (cond
                    (= number 90) [eastwest (- northsouth) shipNS shipEW] ;; n=a e=b => n=b e=-a) 
                    (= number 270) [(- eastwest) northsouth shipNS shipEW]
                    (= number 180) [(- northsouth) (- eastwest) shipNS shipEW])
      (= code \R) (cond
                    (= number 270) [eastwest (- northsouth) shipNS shipEW] ;; n=a e=b => n=b e=-a) 
                    (= number 90) [(- eastwest) northsouth shipNS shipEW]
                    (= number 180) [(- northsouth) (- eastwest) shipNS shipEW]))))

(defn part-2
  "Day 12 Part 2"
  [input]
  (loop
   [instructions input
    northsouth 1
    eastwest 10
    shipNS 0
    shipEW 0]
    (let [instruction (first instructions)
          otherinstructions (rest instructions)
          [newns, newew, newShipNS, newShipEW] (update-position2 instruction northsouth eastwest shipNS shipEW)]

      (if (empty? otherinstructions)
        (+ (Math/abs newShipNS) (Math/abs newShipEW))
        (recur otherinstructions newns newew newShipNS newShipEW)))))