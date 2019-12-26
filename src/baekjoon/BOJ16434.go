/*
    baekjoon online judge
    problem number 16434
	https://www.acmicpc.net/problem/16434
	
	Binary Search
	get minimum and maximum value of result
	and prove the value can pass given input by using binary search
*/

package main

import (
	"os"
	"bufio"
	"strconv"
	"strings"
	"fmt"
)

func main(){
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)

	strTkn := getStrTkn(r)
	N,_ := strconv.Atoi(strTkn[0])
	attack,_ := strconv.Atoi(strTkn[1])

	rooms := make([]room,N,N)

	var sum int64 = 0
	for i:=0;i<N;i++ {
		strTkn = getStrTkn(r)

		rooms[i].roomType,_ = strconv.Atoi(strTkn[0])
		rooms[i].attack,_ = strconv.Atoi(strTkn[1])
		rooms[i].hp,_ = strconv.Atoi(strTkn[2])

		if rooms[i].roomType==1 {
			sum += int64(rooms[i].attack) * int64(rooms[i].hp)
		}
	}

	minHP := binarySearch(attack, sum, rooms)
	w.WriteString(strconv.FormatInt(minHP,10))
	w.Flush()
}

func binarySearch(attack int, maxHP int64, rooms []room) int64 {
	var left, right int64 = 1, maxHP
	var mid int64
	var h hero = hero{mid, mid, attack}

	var res int64 = maxHP
	for left <= right {
		// fmt.Println(mid)
		mid = (left+right)/2
		h.totalHP = mid
		h.currentHP = mid

		if h.meetRooms(rooms) {
			// fmt.Printf("hero alive at hp:%d\n",mid)
			right = mid-1
			if res > mid {
				res = mid
			}
		} else {
			// fmt.Printf("hero dead at hp:%d\n",mid)
			left = mid+1
		}
	}

	return res
}

type room struct{
	roomType int
	attack int
	hp int
}

func (r room) toString() string {
	return fmt.Sprintf("room type: %d\nattack: %d\nhp: %d\n",r.roomType,r.attack,r.hp)
}

type hero struct{
	totalHP int64
	currentHP int64
	attack int
}

func (h hero) toString() string {
	return fmt.Sprintf("total HP: %d\ncurrent HP: %d\nattack: %d\n",h.totalHP,h.currentHP,h.attack)
}

func (h hero) meetRooms(rooms []room) bool {
	var alive bool = true
	for i:=0;i<len(rooms);i++ {
		alive = h.meetRoom(rooms[i])
		if !alive {
			break
		}
	}

	return alive
}

func (h *hero) meetRoom(r room) bool {
	switch r.roomType {
	// monster
	case 1:
		var q int
		if r.hp % h.attack == 0 {
			q = r.hp / h.attack -1
		} else {
			q = r.hp / h.attack
		}
		h.currentHP -= int64(r.attack)*int64(q)
	
	// potion
	case 2:
		h.attack += r.attack
		h.currentHP += int64(r.hp)
		if h.currentHP > h.totalHP {
			h.currentHP = h.totalHP
		}
	default:
		panic("another type of room")
	}

	// fmt.Print(h.toString())
	if h.currentHP <= 0 {
		return false
	} else {
		return true
	}
}

func getString(r *bufio.Reader) string{
	str,_ := r.ReadString('\n')
	str = strings.TrimSpace(str)

	return str
}

func getStrTkn(r *bufio.Reader) []string{
	str,_ := r.ReadString('\n')
	str = strings.TrimSpace(str)
	strTkn := strings.Split(str," ")

	return strTkn
}

func printArr(arr []int, w *bufio.Writer) {
	for i:=0;i<len(arr);i++ {
		w.WriteString(strconv.Itoa(arr[i])+"\n")
	}
}