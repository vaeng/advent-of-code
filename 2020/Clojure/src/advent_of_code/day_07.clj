(ns advent-of-code.day-07
  (:require [clojure.string :as str])
  (:require [clojure.set :as set])
   (:require [clojure.edn :as edn])
  )

(defn color-containers
  "returns all the bag colors that can include *color*ed bag"
  [input, color]
  (set (for
        [[_, bag] (re-seq (re-pattern (str "(?:(\\w+ \\w+) bags contain .*" color ")")) input)]
         bag)))

(defn recursive-bag-finder
  "finds bags in bags recusirvly"
  [input searchcolorset]
  (let [newset (set/union searchcolorset (apply set/union (for [bag searchcolorset] (color-containers input bag))))]
    (if (= (count searchcolorset) (count newset))
      searchcolorset
      (recursive-bag-finder input newset))))

(defn part-1
  "Day 07 Part 1"
  [input]
  (count (disj (recursive-bag-finder input #{"shiny gold"}) "shiny gold")))

(defn bags-in-a-bag
  "returns all the bag colors and amounts that are in a specific *color*ed bag"
  [input color]
  (let [rule (second (re-find (re-pattern (str color " bags contain ([\\w ,]+)\\.")) input))]
    (for [[_, amount bagcolor] (re-seq #"(\d+) (\w+ \w+) bags?" rule)]
      [bagcolor (edn/read-string amount)])))

(defn recursive-bag-counter
  "counts all possible bags in a bag"
  [input color]
  (let [bagcontent (bags-in-a-bag input color)]
    (if (empty? bagcontent)
      0
      (apply + (for
                [[innercolor amount] bagcontent]
                 (+ amount (* amount (recursive-bag-counter input innercolor))))))))

(defn part-2
  "Day 07 Part 2"
  [input]
  (recursive-bag-counter input "shiny gold"))
