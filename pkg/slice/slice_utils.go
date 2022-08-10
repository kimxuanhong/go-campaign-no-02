package slice

func FirstOrDefault[T any](slice []T, filter func(*T) bool) (element *T) {

	for i := 0; i < len(slice); i++ {
		if filter(&slice[i]) {
			return &slice[i]
		}
	}

	return nil
}

func Where[T any](slice []T, filter func(*T) bool) []*T {

	var arr = make([]*T, 0)

	for i := 0; i < len(slice); i++ {
		if filter(&slice[i]) {
			arr = append(arr, &slice[i])
		}
	}

	return arr
}

func RemoveByIndex[T any](slice []T, index int) []T {
	arr := append(slice[:index], slice[index+1:]...)
	return arr
}

func Remove[T any](slice []T, filter func(*T) bool) []T {
	var indexList = make([]int, 0)
	for index, v := range slice {
		if filter(&v) {
			indexList = append(indexList, index)
		}
	}

	for _, v := range indexList {
		slice = RemoveByIndex(slice, v)
	}

	return slice
}
