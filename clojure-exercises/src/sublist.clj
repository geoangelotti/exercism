(ns sublist)

(defn sublist? [a b]
  (some #{a} (partition (count a) 1 b) ))

(defn classify [coll1 coll2]
  (cond (= coll1 coll2) :equal
        (sublist? coll1 coll2) :sublist
        (sublist? coll2 coll1) :superlist
        :else :unequal))