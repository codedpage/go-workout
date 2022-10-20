package main
func main() {
	
s:= []string{"a","b","c","d","e","f","g","h","i","j"} 
var b = s[2:5]		  
	
for i, e := range b {
		print(i)
		print(".................")
		println(e)
	}	
println("---------------------------")
for _, e := range b {		
		print(".................")
		println(e)
	}	

println("---------------------------")
for i := range b {
		print(i)
		println(".................")		
	}	
}