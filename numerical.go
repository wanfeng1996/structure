package structure

import (
	"slices"
)

// Sum 求和
func Sum[T Numerical](nums []T) *T {
	var sum T
	for num := range nums {
		sum += nums[num]
	}
	return &sum
}

func SumIndex[T Numerical](nums []*T) *T {
	var sum *T
	for num := range nums {
		if nums[num] == nil {
			continue
		}
		if sum == nil {
			sum = new(T)
		}
		*sum += *nums[num]
	}
	return sum
}

// Avg 求平均
func Avg[T Numerical](nums []T) *T {
	if len(nums) < 1 {
		return nil
	}

	return Address(*Sum(nums) / T(len(nums)))
}

func AvgIndex[T Numerical](nums []*T) *T {
	sum := SumIndex(nums)

	if sum == nil {
		return nil
	}

	return Address(*sum / T(len(nums)))
}

// Max 获取最大值
func Max[T Numerical](nums []T) *T {
	if len(nums) == 0 {
		return nil
	}

	return Address(slices.Max(nums))
}

func MaxIndex[T Numerical](nums []*T) *T {
	var mx *T
	found := false // 标记是否存在有效元素

	for num := range nums {
		if nums[num] == nil {
			continue
		}

		if !found { // 第一个有效元素
			mx = nums[num]
			found = true
			continue
		}

		if *nums[num] > *mx { // 后续元素比较
			mx = nums[num]
		}

	}

	return mx // 自动处理全nil或空切片的情况
}

// Min 获取最小值
func Min[T Numerical](nums []T) *T {
	if len(nums) < 1 {
		return nil
	}
	return Address(slices.Min(nums))
}

func MinIndex[T Numerical](nums []*T) *T {
	var mn *T
	found := false
	for num := range nums {
		if nums[num] == nil {
			continue

		}
		if !found {
			mn = nums[num]
			found = true
			continue
		}

		if *nums[num] < *mn {
			mn = nums[num]
		}
	}
	return mn
}

func Address[T any](data T) *T {
	return &data
}

func Default[T any](data *T) T {
	if data != nil {
		return *data
	}

	var zero T
	return zero
}
