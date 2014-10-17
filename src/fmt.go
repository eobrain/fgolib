package fmt
import "clojure/string"
import type (
	clojure.lang.{Ratio, IPersistentVector, APersistentMap}
)

var (
	Printf = printf
	Print = print
	Sprint = str
)


func stripZeroAfterDecimal(s) { string.replace(str(s), /^([1-9][0-9]*)\.0$/, "$1") }


// Create a string in the same way Go does.
func toStringLikeGo(x) {
	switch x.(type) {
	case String:           x
	case Ratio:            str(int(x))
	case Number:           stripZeroAfterDecimal(x)
	case IPersistentVector: str(
		"[", 
		" "  string.join  (toStringLikeGo  map  x),
		"]"
	)
        case APersistentMap:
		str("map[",
			" "  string.join  (for [key, val] := lazy x {
				str(toStringLikeGo(key), ":", toStringLikeGo(val))
			})
		, "]")
	default:               str(x)
	//default:               str(x->getClass(), ":", x)
	}
}

func Println(args...) {
	const as = func{toStringLikeGo($1)}  map  args
	println(" "  string.join  as)
}
