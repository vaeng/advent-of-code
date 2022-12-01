(ns advent-of-code.day-05
  (:require [clojure.string :as str])
)

(defn resolve-code
  "Transforms binary space partition into ID"
  [code]
  (. Integer parseInt (str/replace (str/replace code #"[BR]" "1") #"[FL]" "0") 2))

(defn part-1
  "Day 05 Part 1"
  [input]
  (apply max (for [code (str/split input #"\r\n")] (resolve-code code)))
  )

(defn part-2
  "Day 05 Part 2"
  [input]
  (let [
        alltickets (for [code (str/split input #"\r\n")] (resolve-code code))
        firstseat (apply min alltickets)
        lastseat (apply max alltickets)
        ]
   (for
   [seat (range firstseat lastseat)
    :when (nil? (some #(= % seat) alltickets))]
    seat))
  )