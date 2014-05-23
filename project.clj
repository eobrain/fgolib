(defproject org.eamonn.funcgo/fgolib "0.2.1"
  :description "Library that mimics the standard Go library"
  :url "http://funcgo.org"
  :license {:name "Eclipse Public License"
            :url "http://www.eclipse.org/legal/epl-v10.html"}
  :dependencies [[org.clojure/clojure "1.5.1"]
                 [org.eamonn.funcgo/funcgo-lein-plugin "0.2.3"]
                 [midje "1.5.1" :scope "test"]]
  :plugins [[lein-midje "3.1.1"]
            [org.eamonn.funcgo/funcgo-lein-plugin "0.2.3"]])
