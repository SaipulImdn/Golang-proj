package main

import (
	"fmt"
	"strconv"
)

func main() {
	{
		var str = "124"
		var num, err = strconv.Atoi(str)

		if err == nil {
			fmt.Println(num)
		}
	}

	{
		var num = 124
		var str = strconv.Itoa(num)

		fmt.Println(str)
	}

	{
		var str = "124"
		var num, err = strconv.ParseInt(str, 10, 64)

		if err == nil {
			fmt.Println(num)
		}
	}

	{
		var str = "1010"
		var num, err = strconv.ParseInt(str, 2, 8)

		if err == nil {
			fmt.Println(num)
		}
	}

	{
		var num = int64(24)
		var str = strconv.FormatInt(num, 10)

		fmt.Println(str)
	}

	{
		var str = "24.12"
		var num, err = strconv.ParseFloat(str, 32)

		if err == nil {
			fmt.Println(num)
		}
	}

	{
		var num = float64(24.12)
		var str = strconv.FormatFloat(num, 'g', 6, 64)

		fmt.Println(str)
	}

	{
		var str = "true"
		var bul, err = strconv.ParseBool(str)

		if err == nil {
			fmt.Println(bul)
		}
	}

	{
		var bul = true
		var str = strconv.FormatBool(bul)

		fmt.Println(str)
	}

	{
		var text1 = "halo"
		var b = []byte(text1)

		fmt.Printf("%d %d %d %d \n", b[0], b[1], b[2], b[3])
	}

	{
		var byte1 = []byte{104, 97, 108, 111}
		var s = string(byte1)

		fmt.Printf("%s \n", s)
	}

	{
		var data = map[string]interface{}{
			"nama":    "syaiful imanudin",
			"grade":   3,
			"height":  170.0,
			"isMale":  true,
			"hobbies": []string{"eating", "sleeping"},
		}

		fmt.Println(data["nama"].(string))
		fmt.Println(data["grade"].(int))
		fmt.Println(data["height"].(float64))
		fmt.Println(data["isMale"].(bool))
		fmt.Println(data["hobbies"].([]string))

		for _, val := range data {
			switch val.(type) {
			case string:
				fmt.Println(val.(string))
			case int:
				fmt.Println(val.(int))
			case float64:
				fmt.Println(val.(float64))
			case bool:
				fmt.Println(val.(bool))
			case []string:
				fmt.Println(val.([]string))
			default:
				fmt.Println(val.(int))
			}
		}
	}
}
