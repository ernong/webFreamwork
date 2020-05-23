package util

var chanGo chan bool

func InitGo(size int) {
	chanGo = make(chan bool, size)

	for i := 0; i < size; i++ {
		chanGo <- true
	}
}

func GetGo() {
	<-chanGo
}

func ReleaseGo() {
	chanGo <- true
}
