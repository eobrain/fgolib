package os
import type (
	java.io.FileInputStream
)

func Open(path String){
	try {
		new FileInputStream(path)
	} catch Exception e {
		e
	}
}
