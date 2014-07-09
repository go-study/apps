package prints

//#include "prints.h"
// // some comment
import "C"

func Prints(s string) {
  p := C.CString(s);
  C.prints(p);
}
