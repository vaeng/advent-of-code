(ns advent-of-code.day-08
   (:require [clojure.string :as str])
   (:require [clojure.edn :as edn])
  )


(defn buffercheck
  "checks the buffer before loop"
  [input-vector currentline visited-lines acc]
  (if (contains? visited-lines currentline)
    acc
    (let [[_, command, numberstring] (re-find #"([a-z]{3}) ([\+\-]\d+)" (nth input-vector currentline))
          number (edn/read-string numberstring)]
      (case command
        "nop" (buffercheck input-vector (inc currentline) (conj visited-lines currentline) acc)
        "acc" (buffercheck input-vector (inc currentline) (conj visited-lines currentline) (+ number acc))
        "jmp" (buffercheck input-vector (+ currentline number) (conj visited-lines currentline) acc)))))

(defn part-1
  "Day 08 Part 1"
  [input]
  (buffercheck (str/split-lines input) 0 #{} 0))

(defn corrected-boot-up-test
  "returns acc for uncorrupted code, no changes"
  [input-vector currentline visited-lines acc]
  (if (contains? visited-lines currentline)
    false
    (if (= currentline (count input-vector))
      acc
      (let [[_, command, numberstring] (re-find #"([a-z]{3}) ([\+\-]\d+)" (nth input-vector currentline))
            number (edn/read-string numberstring)]
        (case command
          "nop" (corrected-boot-up-test input-vector (inc currentline) (conj visited-lines currentline) acc)
          "acc" (corrected-boot-up-test input-vector (inc currentline) (conj visited-lines currentline) (+ number acc))
          "jmp" (corrected-boot-up-test input-vector (+ currentline number) (conj visited-lines currentline) acc))))))

(defn boot-up-test
  "returns acc for uncorrupted code"
  [input-vector currentline visited-lines acc]
  (if (contains? visited-lines currentline)
    false
    (if (= currentline (count input-vector))
      acc
      (let [[_, command, numberstring] (re-find #"([a-z]{3}) ([\+\-]\d+)" (nth input-vector currentline))
            number (edn/read-string numberstring)]
        (case command
          "nop" (let [nop (boot-up-test input-vector (inc currentline) (conj visited-lines currentline) acc)]
                  (if nop
                    nop
                    (corrected-boot-up-test input-vector (+ currentline number) (conj visited-lines currentline) acc)))
          "acc" (boot-up-test input-vector (inc currentline) (conj visited-lines currentline) (+ number acc))
          "jmp" (let [jmp (boot-up-test input-vector (+ currentline number) (conj visited-lines currentline) acc)]
                  (if jmp
                    jmp
                    (corrected-boot-up-test input-vector (inc currentline) (conj visited-lines currentline) acc))))))))

(defn part-2
  "Day 08 Part 2"
  [input]
  (boot-up-test (str/split-lines input) 0 #{} 0))






