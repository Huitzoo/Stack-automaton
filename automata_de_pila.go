package main

import (
	"fmt"
	"strings"
	"container/list"
)


func transition1(cadena string, p *list.List, c chan int){
	fmt.Println("HOLA")
	i:=1
	nd:=p.Front()
	if strings.Compare(string(cadena[i]),"a") == 0 {
		for strings.Compare(string(cadena[i]),"a")==0{
			//fmt.Println("HOLA")
			p.PushFront("A")
			i++
			if i == len(cadena) {
				//fmt.Println("No pertenece al automata de pila")
				c <- 0
				//fmt.Println("HOLA")
			}
		}
		/* BLOQUE DE B */
		if strings.Compare(string(cadena[i]),"b") ==0 {
			for strings.Compare(string(cadena[i]),"b")==0{
				p.PushFront("B")
				i++
				if i == len(cadena) {
					//fmt.Println("No pertenece al automata de pila")
					c <- 0
					//fmt.Println("HOLAB")
					return
				}
			}
		}else{
			//fmt.Println("No pertenece a la gramatica")
			c <- 0
			return
		}
		/*-----------------------------------------------*/
		/* BLOQUE DE C */
		if strings.Compare(string(cadena[i]),"c") == 0{
			for strings.Compare(string(cadena[i]),"c") == 0 && (strings.Compare(p.Front().Value.(string),"A") == 0 || strings.Compare(p.Front().Value.(string),"B") == 0) {
				p.Remove(p.Front())
				i++
				if i == len(cadena) {
					break
				}
			}
			if i != len(cadena){
				//fmt.Println("No pertenece al automata de pila")
				c <- 0
				return
			}else{
				c <- 1
				return
			}
		}
		/*---------------------------------------------------------*/
	}else if strings.Compare(string(cadena[i]),"b")==0 && strings.Compare(nd.Value.(string),"A")==0{
		if strings.Compare(string(cadena[i]),"b") ==0 {
			for strings.Compare(string(cadena[i]),"b")==0{
				p.PushFront("B")
				i++
				if i == len(cadena) {
					//fmt.Println("No pertenece al automata de pila")
					c <- 0
					return
				}
			}
		}else{
			c <- 0
			//fmt.Println("No pertenece a la gramatica")
			return
		}
		/*-----------------------------------------------*/
		/* BLOQUE DE C */
		if strings.Compare(string(cadena[i]),"c") == 0{
			for strings.Compare(string(cadena[i]),"c") == 0 && (strings.Compare(p.Front().Value.(string),"A") == 0 || strings.Compare(p.Front().Value.(string),"B") == 0) {
				p.Remove(p.Front())
				i++
				if i == len(cadena) {
					break
				}
			}
			if i == len(cadena){
				//fmt.Println("No pertenece al automata de pila")
				c <- 0
				return
			}else{
				c <- 1
				return
			}
		}
	}else{
		fmt.Println("goalal")
		c <- 0
		//fmt.Println("No pertenece al automata de pila")
		return
	}
}

func transition2(cadena string, p *list.List, c chan int){
	i:=1
	fmt.Println("HOLA1")
	if strings.Compare(string(cadena[i]),"a") == 0 {
		for strings.Compare(string(cadena[i]),"a")==0{
			//fmt.Println("HOLAs")
			if i%2 == 0{
				fmt.Println("HOLAs")
				p.PushFront("A")
			}
			i++
		}
		if i%2 != 0{
			//fmt.Println("dddddd")
			c <- 0
			return
		}
		/* BLOQUE DE B */
		if strings.Compare(string(cadena[i]),"b") == 0{
			for strings.Compare(string(cadena[i]),"b") == 0 && strings.Compare(p.Front().Value.(string),"A") == 0{
				fmt.Println("HOLAsB")
				p.Remove(p.Front())
				i++
				if i == len(cadena) {
					break
				}
			}
			if i != len(cadena){
				//fmt.Println("ffffffff")
				c <- 0
				return
			}else{
				c <- 1
				return
			}
		}
	}else{
		c <- 0
		return
	}
}

func main() {
	cadena := "_"
	pila := list.New()
	pila1 := list.New()
	c := make(chan int)
	fmt.Println("Acepta cadenas: (a^n b^m c^m+n) u (a^2n b^n)")
	fmt.Println("Ingresa su cadena")

	for strings.Compare(string(cadena[0]),"_")==0{
		fmt.Scanln(&cadena)
	}
	if strings.Compare(string(cadena[0]),"a") == 0{
		pila.PushFront("Z0")
		pila1.PushFront("Z0")
		pila.PushFront("A")
		pila1.PushFront("A")
		go transition1(cadena,pila,c)
		go transition2(cadena,pila1,c)
		t1, t2 := <- c, <- c
		if t1 == 1 || t2 == 1 {
			fmt.Println("Pertenece al automata")
		}else{
			fmt.Println("No pertenece al automata")
		}
	}else{
		fmt.Println("No pertenece al automata")
		for h:=pila.Front(); h!=nil ; h = h.Next() {
			fmt.Println(h.Value)
		}
		for h1:=pila1.Front(); h1!=nil ; h1 = h1.Next() {
			fmt.Println(h1.Value)
		}
	}
	if pila.Front() == nil || pila1.Front() == nil{
		fmt.Println("No se crearon las pilas")
	}else{
		if strings.Compare(pila.Front().Value.(string),"Z0") == 0{
			for h:=pila.Front(); h!=nil ; h = h.Next() {
				fmt.Println(h.Value)
			}
		}
		if strings.Compare(pila1.Front().Value.(string),"Z0") == 0{
			for h1:=pila1.Front(); h1!=nil ; h1 = h1.Next() {
				fmt.Println(h1.Value)
			}
		}
	}
}
