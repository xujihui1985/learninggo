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
	withFuture(<-fa(), <-fb())

	ca, cb := fa(), fb()
	withFuture(<-ca, <-cb)
}
