(ns bird-watcher)

(def last-week 
  [0,2,5,3,7,8,4])

(defn last-index [birds]
  (- (count birds) 1)
  )

(defn today [birds]
  (last birds)
  )

(defn inc-bird [birds]
  (conj (pop birds) (+ (int 1) (last birds)))  
  )

(defn day-without-birds? [birds]
  (not (empty? (filter zero? birds)))
  )

(defn n-days-count [birds n]
  (apply + (take n birds))
  )

(defn busy-days [birds]
  (count (filter (fn [x] (>= x 5)) birds))
  )

(defn odd-week? [birds]
  (= [1 0 1 0 1 0 1] birds)
  )
