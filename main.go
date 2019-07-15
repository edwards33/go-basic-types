package main

type UserID int

func main() {
	idx := 1
	var uid UserID = 42

	// cannot use uid (type UserID) as type int64 in assignment
	// myID := idx

	myID := UserID(idx)
	
	println(uid, myID)
}
