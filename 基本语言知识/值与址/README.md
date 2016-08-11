array值传递
```
var array  [5]float64=[5]float64{1,2,3,4,5}

	news:=array

	array[4]=5.1
	news[1]=1.1
	fmt.Println(array,news)
	
[1 2 3 4 5.1] [1 1.1 3 4 5]
```

slice址传递
```
var slice []float64=[]float64{1,2,3,4,5}

	news:=slice

	slice[4]=5.1
	news[1]=1.1
	fmt.Println(slice,news)
	
[1 1.1 3 4 5.1] [1 1.1 3 4 5.1]
```