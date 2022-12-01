(ns advent-of-code.day-04
  (:require [clojure.string :as str])
  (:require [clojure.edn :as edn])
  )

(def minprops '("byr:" "iyr:" "eyr:" "hgt:" "hcl:" "ecl:" "pid:"))

(defn part-1
  "Day 04 Part 1"
  [input]
  (let [passports (str/split (str/replace input #"(\r\n)" " ") #"  ")]
    (count 
     (filter 
      #(= % true) 
      (for [passport passports]
          (every? #(str/includes? passport %) minprops))))))

(defn data-checker
  [prop-map-vector]
  (for [prop-map prop-map-vector
        :let [byr (edn/read-string (get prop-map "byr"))
              iyr (edn/read-string (get prop-map "iyr"))
              eyr (edn/read-string (get prop-map "eyr"))
              hgt (get prop-map "hgt")
              hcl (get prop-map "hcl")
              pid (get prop-map "pid")
              ecl (get prop-map "ecl")]
        :when (not-any? nil? [byr iyr eyr hgt hcl ecl pid])]
    (every?
     #(= % true)
     [(>= byr 1920) (<= byr 2002)
      (>= iyr 2010) (<= iyr 2020)
      (>= eyr 2020) (<= eyr 2030)
      (if (re-find #"^\d{9}$" pid) true false)
      (if (some #(= % ecl) ["amb" "blu" "brn" "gry" "grn" "hzl" "oth"]) true false)
      (let [hgt-number (edn/read-string (re-find #"\d+" hgt))
            hgt-unit (re-find #"[a-z]+" hgt)]
        (if (= "cm" hgt-unit)
          (= true (>= hgt-number 150) (<= hgt-number 193))
          (if (= "in" hgt-unit)
            (= true (>= hgt-number 59) (<= hgt-number 76))
            false)))
      (if (re-find #"^#[a-f0-9]{6}$" hcl) true false)])))

(defn part-2
  "Day 04 Part 2"
  [input]
  (let 
   [passports (str/split (str/replace input #"(\r\n)" " ") #"  ")
    prop-map-vector (for 
                     [passport passports]
                      (into 
                       (sorted-map) 
                       (for
                        [[_, field, value] (re-seq #"(?:(\w\w\w):([a-z0-9#]+))" passport)]
                         [field, value]))
                    )
    ]
    (count
     (filter 
     #(= % true) 
     (data-checker prop-map-vector)))))

;;;; Part 2



(defn read-input
  [day]
  (slurp (clojure.java.io/resource day)))

(def input-example (read-input "day-04-example.txt"))
(def input (read-input "day-04.txt"))

(part-2 input)

