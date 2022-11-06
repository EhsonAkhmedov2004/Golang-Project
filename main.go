package main
import (

	"fmt"
	"io/ioutil"
	"encoding/json"
	"base/Models"
	"base/Shop"

)


// ==================================START=======================================




var mainRange = []int{3,6,9,12,18,24};





func main(){

	filess,_ := ioutil.ReadFile("./demoDb/database.json")
	var Database Models.Db;
	var message int  = 0;
	var price   int  = 0;
	var Range   int  = 0;
	json.Unmarshal(filess,&Database)

	
	fmt.Println("=======================PROGRAM HAS BEEN STARTED !=======================")




	processFindRange := []int{}
	fmt.Println("\nПожалуйста выберите продукт ( введите 1 или 2 или 3) ")
	fmt.Println("1 - Телефон")
	fmt.Println("2 - Компьютер")
	fmt.Println("3 - Телевизор")
	
	fmt.Scan(&message)
	
	if message <= 0{
		fmt.Println("Введите 1 или 2 или 3 {}")
		return
	}

	if message > len(Database.Db){
		fmt.Println("Пожалуйтста выберите от 1 до 3")
		return
	} 
			
	
	fmt.Print("Введите цену : ")
	fmt.Scan(&price)
	if price <= 0 {
		fmt.Println("Введите сумму больше 0 или уберите символы!")
		return;
	}
	fmt.Println("Диапазон начинается с 3 месяцев до какого месяца \n хотите расширить диапазон выберите из -> [6,12,18,24] ")
			
	fmt.Print(" > ")
			
	fmt.Scan(&Range) 
	
			
	if Range <= 0{
		fmt.Printf("Диапазон указан некорректно по умолчанию диапазон [ %v ] от 3 до %v \n",
														Database.Db[message - 1].Name,
														Database.Db[message - 1].Range[1])
		Range = Database.Db[message - 1].Range[1]
	}
		
			//Little logic !....
	
			for _,v := range mainRange{
				

				if Range == v {
					break;
				}

				if Range > v {
					processFindRange = append(processFindRange,v)
				} 
				if Range < v {
					processFindRange = append(processFindRange,v)				
				}


				if len(processFindRange) == 2 {
					if Range > processFindRange[0] && Range < processFindRange[1]{
						Range = processFindRange[1]
						break;

					} else{
						processFindRange = processFindRange[:len(processFindRange) - 1]
				
						continue;
					}
					
				}
			}
			
			
			product := Database.Db[message - 1]
			product.Price = price;
			
			Shop.Shop(&product,[]int{3,Range})
		

	



	
	
}