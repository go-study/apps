package main
import("fmt"
	"time"
)

func main(){
	var i=1;
	go func(i int){
		var j=0;
		for{
			fmt.Println(i,"\n");
			time.Sleep(1);
			j=j+1;
			if (j<10){
				break;
}
		}

	}(i)
	go func(i int){
		var j=0;
		for{
			fmt.Println(i,"\n");
			time.Sleep(1);
			j=j+1;
			if (j<10){
				break;
}
		}
	}(i)
 time.Sleep(60 * time.Second)
}
