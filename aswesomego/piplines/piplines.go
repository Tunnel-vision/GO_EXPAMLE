package piplines

//receive values from upstream via inbound channels
//perform some function on that data, usually producing new values
//send values downstream via outbound channels

// receive values from upstream via inbound channels first stage
func gen(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

func sq(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			//println(">>>>>>>>>>>>", n)
			out <- n
		}
		close(out)

	}()
	return out
}
