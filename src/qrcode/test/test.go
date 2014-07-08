package main

import (
        "image/png"
        "os"
)
import "github.com/qpliu/qrencode-go/qrencode"
func main(){
	
	grid, err := qrencode.Encode("Testing one two three four five six seven eight nine ten eleven twelve thirteen fourteen fifteen sixtee n seventeen eighteen nineteen twenty.", qrencode.ECLevelQ)
        if err != nil {
                return
        }
        f, err := os.Create("/tmp/qr.png")
        if err != nil {
                return
        }
        defer f.Close()
        png.Encode(f, grid.Image(8))
}
