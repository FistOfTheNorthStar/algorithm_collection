package shopping

import "fmt"

// GetNumberOfOptions calculates how many shopping combinations are possible within budget
func GetNumberOfOptions(priceOfJeans, priceOfShoes, priceOfSkirts, priceOfTops []int, dollars int) (int, error) {
	if dollars < 1 || dollars > 1000000000 {
		return 0, fmt.Errorf("wrong value for budget")
	}

	// Validate all price arrays
	if err := validate(priceOfJeans, "jeans"); err != nil {
		return 0, err
	}
	if err := validate(priceOfShoes, "shoes"); err != nil {
		return 0, err
	}
	if err := validate(priceOfSkirts, "skirts"); err != nil {
		return 0, err
	}
	if err := validate(priceOfTops, "tops"); err != nil {
		return 0, err
	}

	numberOfOptions := 0
	for _, priceOfJean := range priceOfJeans {
		if priceOfJean >= dollars {
			continue
		}

		for _, priceOfShoe := range priceOfShoes {
			if priceOfJean+priceOfShoe >= dollars {
				continue
			}

			for _, priceOfSkirt := range priceOfSkirts {
				if priceOfJean+priceOfShoe+priceOfSkirt >= dollars {
					continue
				}

				for _, priceOfTop := range priceOfTops {
					if priceOfJean+priceOfShoe+priceOfSkirt+priceOfTop <= dollars {
						numberOfOptions++
					}
				}
			}
		}
	}

	return numberOfOptions, nil
}

func validate(array []int, arrayName string) error {
	if len(array) < 1 || len(array) > 1000 {
		return fmt.Errorf("wrong size in array %s", arrayName)
	}

	for _, price := range array {
		if price < 1 || price > 1000000000 {
			return fmt.Errorf("wrong value in array %s", arrayName)
		}
	}
	return nil
}
