(ns day1.core)

(def startweight 3232358)
(def totalfuel 3232358)

(defn calc_fuel [weight] (- (Math/floor (/ weight 3)) 2))

(loop [i startweight]
  (if (> i 0)
    ((def totalfuel (+ totalfuel i))
     recur (calc_fuel i))
    totalfuel)
  )
