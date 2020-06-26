package datasource

import(
	"fmt"
	"time"
)

type Node struct{
	character *Character 
	next *Node
	prev *Node
}


type MarvelData struct{
	last_update time.Time
	capacity int
	head *Node
	tail *Node
	characterMap map[string]*Node
}

type Character struct{
	Name string
	Max_power int
}

var CharacterList [] *Character 
var md *MarvelData

func  Put(name string, power int) {

	_,present := md.characterMap[name]

	if present {
		    node := md.characterMap[name]
			node.character.Max_power = power
			remove(node)
			addToFront(node)
			
		}else{
			character := &Character{ Name : name, Max_power : power }
			newnode := &Node{ character :  character, next : nil , prev : nil }

			if md.capacity == 15 {

				delete(md.characterMap , md.tail.character.Name)
				remove(md.tail)				
				addToFront(newnode)

			} else {
			    addToFront(newnode)
			    md.capacity++
			}

			md.characterMap[character.Name] = newnode
					
		}

}

func Get(name string) int{

	  time_now := time.Now()

	  if(time_now.Sub(md.last_update) >= 10 * time.Second){
	  	fmt.Println("Old data found..")
	  	// function call to fetch new data 
	  	UpdateMarvelData("Update")
	  }

	  power := 0
	_,present := md.characterMap[name]

		if present {
			node := md.characterMap[name]
			power= node.character.Max_power;
			remove(node);
			addToFront(node);
		}
				
	return power;

}

func addToFront(node *Node){

	    node.next = md.head;
		node.prev = nil;
		
		if (md.head != nil){
			md.head.prev = node;
		}
		
		md.head = node;
		
		if (md.tail == nil){
			md.tail = md.head;
		}

}

func (character *Character) AddNodeToSortedMData(){


    // Initialising mdata
    if(md ==nil){
    	t_now := time.Now();
  	  	m := make(map[string]*Node)
    	s := 0
    	md = &MarvelData {last_update : t_now , capacity : s, characterMap : m }

    	fmt.Println(t_now)

    }

    
	// checking capacity
	if(md.capacity >= 15 ){
		    delete(md.characterMap, md.tail.character.Name)
			remove(md.tail)
			 // delete from map as well
			md.capacity--;
	}

	// sorted Marvel Data List for first Time addition

		var newnode *Node

        // first node to be added to the Marvel data
          if md.head == nil{          	
          	newnode = &Node{ character :  character, next : nil , prev : nil }
          	md.head = newnode

          }else{
        		

          	var currnode, prevnode *Node

          	currnode = md.head // point to head

          	// finding postion to insert new node ...
          	for currnode.next != nil && currnode.character.Max_power >= character.Max_power {
           			currnode = currnode.next          		
          	}


          	newnode = &Node{ character :  character, next : nil , prev : nil }


          	//Checking End position first

          	if(currnode.next == nil && currnode.prev == nil){
          		 	
          		if( currnode.character.Max_power >= character.Max_power ){
          		
          		    newnode.prev = currnode
          			currnode.next = newnode
                    md.tail = newnode
          			}else{
          				newnode.next = currnode
          				currnode.prev = newnode
          			    md.head = newnode
          			}

          	}else if(currnode.next == nil ){

          		if( currnode.character.Max_power >= character.Max_power ){

          			newnode.prev = currnode
          			currnode.next = newnode
                    md.tail = newnode
          			}else{
          				prevnode = currnode.prev
          				newnode.next = currnode
          				currnode.prev = newnode
          			    newnode.prev = prevnode
          			    prevnode.next = newnode
          			}

          	}else if( currnode.prev == nil ){

          			newnode.next = currnode
          			currnode.prev = newnode
          			md.head = newnode
          	}else{

          		        prevnode = currnode.prev
          				newnode.next = currnode
          				currnode.prev = newnode
          			    newnode.prev = prevnode
          			    prevnode.next = newnode
          	}

 			

          }

          md.characterMap[character.Name] = newnode
          md.capacity++


       
}

func remove(node *Node){

	if node.prev != nil {
			node.prev.next = node.next;
	}else{
			md.head = node.next;
	}
		
		
	if node.next != nil {
			node.next.prev = node.prev;
	}else{
			md.tail = node.prev;
	}

}


func PrintData(){

	currnode := md.head

	for(currnode != nil){

		fmt.Println(currnode.character.Name)
		fmt.Println(currnode.character.Max_power)

		currnode = currnode.next

	}

}


