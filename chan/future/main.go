package main

func fa() <-chan string {
	return nil
}

func fb() <-chan string {
	return nil
}

func withFuture(ca, cb string) {

}

func main(){

	// this is ok, but not ideal, because fa and fb will be invoke sequencely
	withFuture(<-fa(), <-fb())

	// this is good because fa and fb are invoke concurrently
	ca, cb := fa(), fb()
	withFuture(<-ca, <-cb)
}
