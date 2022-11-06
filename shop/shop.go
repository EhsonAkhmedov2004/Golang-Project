package Shop
import(
	"base/Models"
	"fmt"
)
var mainRange = []int{3,6,9,12,18,24};
func Shop(product *Models.Product,Range []int) {
	
	count := 0;

	for _,v := range mainRange {
		if v > product.Range[1] && v <= Range[1] {
			count++
			continue
		}
			
	}
	

	percent := product.Percent * count
	income := percent * product.Price / 100
	result := product.Price + income
	fmt.Printf("\nДобавленный процент ( %v )\n",percent)
	fmt.Printf("Цена ( %v ) составляет %v сомони",product.Name,result)


}