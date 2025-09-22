// これ作ったけどほんまに後続で使っているのか使っていなければ消していいと思う

package common

// Specification は仕様パターンのインターフェース
type Specification[T any] interface {
	IsSatisfiedBy(candidate T) bool
	And(other Specification[T]) Specification[T]
	Or(other Specification[T]) Specification[T]
	Not() Specification[T]
}

// BaseSpecification は仕様パターンの基底実装
type BaseSpecification[T any] struct {
	predicate func(T) bool
}

// NewSpecification は新しい仕様を作成
func NewSpecification[T any](predicate func(T) bool) Specification[T] {
	return &BaseSpecification[T]{predicate: predicate}
}

// IsSatisfiedBy は候補が仕様を満たすかチェック
func (s *BaseSpecification[T]) IsSatisfiedBy(candidate T) bool {
	return s.predicate(candidate)
}

// And は AND条件の複合仕様を作成
func (s *BaseSpecification[T]) And(other Specification[T]) Specification[T] {
	return NewSpecification(func(candidate T) bool {
		return s.IsSatisfiedBy(candidate) && other.IsSatisfiedBy(candidate)
	})
}

// Or は OR条件の複合仕様を作成
func (s *BaseSpecification[T]) Or(other Specification[T]) Specification[T] {
	return NewSpecification(func(candidate T) bool {
		return s.IsSatisfiedBy(candidate) || other.IsSatisfiedBy(candidate)
	})
}

// Not は NOT条件の仕様を作成
func (s *BaseSpecification[T]) Not() Specification[T] {
	return NewSpecification(func(candidate T) bool {
		return !s.IsSatisfiedBy(candidate)
	})
}

// CompositeSpecification は複合仕様の基底型
type CompositeSpecification[T any] struct {
	left  Specification[T]
	right Specification[T]
	op    string
}

// AndSpecification は AND複合仕様
type AndSpecification[T any] struct {
	CompositeSpecification[T]
}

// NewAndSpecification は AND複合仕様を作成
func NewAndSpecification[T any](left, right Specification[T]) *AndSpecification[T] {
	return &AndSpecification[T]{
		CompositeSpecification: CompositeSpecification[T]{
			left:  left,
			right: right,
			op:    "AND",
		},
	}
}

// IsSatisfiedBy は候補が仕様を満たすかチェック
func (s *AndSpecification[T]) IsSatisfiedBy(candidate T) bool {
	return s.left.IsSatisfiedBy(candidate) && s.right.IsSatisfiedBy(candidate)
}

// And は AND条件の複合仕様を作成
func (s *AndSpecification[T]) And(other Specification[T]) Specification[T] {
	return NewAndSpecification(s, other)
}

// Or は OR条件の複合仕様を作成
func (s *AndSpecification[T]) Or(other Specification[T]) Specification[T] {
	return NewOrSpecification(s, other)
}

// Not は NOT条件の仕様を作成
func (s *AndSpecification[T]) Not() Specification[T] {
	return NewNotSpecification(s)
}

// OrSpecification は OR複合仕様
type OrSpecification[T any] struct {
	CompositeSpecification[T]
}

// NewOrSpecification は OR複合仕様を作成
func NewOrSpecification[T any](left, right Specification[T]) *OrSpecification[T] {
	return &OrSpecification[T]{
		CompositeSpecification: CompositeSpecification[T]{
			left:  left,
			right: right,
			op:    "OR",
		},
	}
}

// IsSatisfiedBy は候補が仕様を満たすかチェック
func (s *OrSpecification[T]) IsSatisfiedBy(candidate T) bool {
	return s.left.IsSatisfiedBy(candidate) || s.right.IsSatisfiedBy(candidate)
}

// And は AND条件の複合仕様を作成
func (s *OrSpecification[T]) And(other Specification[T]) Specification[T] {
	return NewAndSpecification(s, other)
}

// Or は OR条件の複合仕様を作成
func (s *OrSpecification[T]) Or(other Specification[T]) Specification[T] {
	return NewOrSpecification(s, other)
}

// Not は NOT条件の仕様を作成
func (s *OrSpecification[T]) Not() Specification[T] {
	return NewNotSpecification(s)
}

// NotSpecification は NOT仕様
type NotSpecification[T any] struct {
	spec Specification[T]
}

// NewNotSpecification は NOT仕様を作成
func NewNotSpecification[T any](spec Specification[T]) *NotSpecification[T] {
	return &NotSpecification[T]{spec: spec}
}

// IsSatisfiedBy は候補が仕様を満たすかチェック
func (s *NotSpecification[T]) IsSatisfiedBy(candidate T) bool {
	return !s.spec.IsSatisfiedBy(candidate)
}

// And は AND条件の複合仕様を作成
func (s *NotSpecification[T]) And(other Specification[T]) Specification[T] {
	return NewAndSpecification(s, other)
}

// Or は OR条件の複合仕様を作成
func (s *NotSpecification[T]) Or(other Specification[T]) Specification[T] {
	return NewOrSpecification(s, other)
}

// Not は NOT条件の仕様を作成
func (s *NotSpecification[T]) Not() Specification[T] {
	return s.spec
}
